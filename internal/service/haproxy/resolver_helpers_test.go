package haproxy

import (
	"context"
	"reflect"
	"testing"

	client "github.com/browningluke/opnsense-go/pkg/haproxy"
)

func TestResolverResponseNameservers(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  []string
	}{
		{"empty object", map[string]any{}, nil},
		{"empty placeholder", map[string]any{"": map[string]any{"selected": float64(1), "value": ""}}, nil},
		{"selected options", map[string]any{"1.1.1.1:53": map[string]any{"selected": float64(1)}, "8.8.8.8:53": map[string]any{"selected": float64(0)}}, []string{"1.1.1.1:53"}},
		{"array", []any{"1.1.1.1:53", "8.8.8.8:53"}, []string{"1.1.1.1:53", "8.8.8.8:53"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := resolverResponseToModel("id", client.HAProxyObject{"nameservers": tt.value})
			if tt.want == nil {
				if !m.Nameservers.IsNull() {
					t.Fatalf("expected null, got %v", m.Nameservers)
				}
				return
			}
			var got []string
			if diags := m.Nameservers.ElementsAs(context.Background(), &got, false); diags.HasError() {
				t.Fatal(diags)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
