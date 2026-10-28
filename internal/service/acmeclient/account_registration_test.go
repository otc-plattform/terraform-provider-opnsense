package acmeclient

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	client "github.com/browningluke/opnsense-go/pkg/acmeclient"
	"github.com/browningluke/opnsense-go/pkg/api"
)

func TestRegisterAccountChecksStatus(t *testing.T) {
	for _, tc := range []struct {
		status    string
		wantError bool
	}{{"200", false}, {"400", true}, {"100", true}} {
		t.Run(tc.status, func(t *testing.T) {
			calls := []string{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls = append(calls, r.Method+" "+r.URL.Path)
				switch r.URL.Path {
				case "/api/acmeclient/accounts/register/id":
					fmt.Fprint(w, `{"response":"OK\n\n"}`)
				case "/api/acmeclient/accounts/get/id":
					fmt.Fprintf(w, `{"account":{"name":"account","statusCode":%q}}`, tc.status)
				default:
					http.NotFound(w, r)
				}
			}))
			defer server.Close()
			controller := client.Controller{Api: api.NewClient(api.Options{Uri: server.URL, Logger: log.New(io.Discard, "", 0)})}
			_, err := registerAccount(context.Background(), &controller, "id")
			if (err != nil) != tc.wantError {
				t.Fatalf("error = %v", err)
			}
			if len(calls) != 2 || !strings.HasPrefix(calls[0], "POST ") || !strings.HasPrefix(calls[1], "GET ") {
				t.Fatalf("calls = %v", calls)
			}
		})
	}
}
