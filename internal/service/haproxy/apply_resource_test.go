package haproxy

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/api"
	client "github.com/browningluke/opnsense-go/pkg/haproxy"
)

func TestApplyHAProxy(t *testing.T) {
	tests := []struct {
		name                string
		output              string
		reconfigureResponse string
		wantReconfigure     bool
		wantError           string
		wantWarning         bool
	}{
		{"valid", "Configuration file is valid\n", `{"status":"ok"}`, true, "", false},
		{"warning", "[WARNING] deprecated option\nConfiguration file is valid\n", `{"status":"ok"}`, true, "", true},
		{"alert", "[ALERT] unknown keyword\nFatal errors found", `{"status":"ok"}`, false, "unknown keyword", false},
		{"empty", "\n", `{"status":"ok"}`, true, "", false},
		{"unknown", "unexpected response", `{"status":"ok"}`, false, "unrecognized output", false},
		{"reconfigure failed", "Configuration file is valid", `{"status":"failed"}`, true, "failed", false},
		{"missing status", "Configuration file is valid", `{}`, true, "status", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var calls []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls = append(calls, r.Method+" "+r.URL.Path)
				switch r.URL.Path {
				case "/api/haproxy/service/configtest":
					json.NewEncoder(w).Encode(map[string]string{"result": tt.output})
				case "/api/haproxy/service/reconfigure":
					io.WriteString(w, tt.reconfigureResponse)
				default:
					http.NotFound(w, r)
				}
			}))
			defer server.Close()
			c := client.Controller{Api: api.NewClient(api.Options{Uri: server.URL, Logger: log.New(io.Discard, "", 0)})}
			warning, err := applyHAProxy(context.Background(), &c)
			if tt.wantError == "" && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tt.wantError != "" && (err == nil || !strings.Contains(err.Error(), tt.wantError)) {
				t.Fatalf("got error %v, want %q", err, tt.wantError)
			}
			if (warning != "") != tt.wantWarning {
				t.Fatalf("warning = %q", warning)
			}
			wantCalls := 1
			if tt.wantReconfigure {
				wantCalls = 2
			}
			if len(calls) != wantCalls || calls[0] != "POST /api/haproxy/service/configtest" || (tt.wantReconfigure && calls[1] != "POST /api/haproxy/service/reconfigure") {
				t.Fatalf("unexpected calls: %v", calls)
			}
		})
	}
}
