package haproxy

import (
	"testing"

	client "github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestHealthMonitorResponseAndPayload(t *testing.T) {
	option := func(selected string) map[string]interface{} {
		return map[string]interface{}{selected: map[string]interface{}{"selected": float64(1), "value": selected}, "unused": map[string]interface{}{"selected": float64(0)}}
	}
	model := healthMonitorResponseToModel("test-id", client.HAProxyObject{
		"name": "health", "type": option("http"), "interval": "2s", "ssl": option("nopref"), "sslSNI": "example.com", "checkport": "443",
		"http_method": option("get"), "http_uri": "/ready", "http_version": option("http11"), "http_host": "example.com",
		"http_expressionEnabled": "1", "http_expression": option("status"), "http_negate": "0", "http_value": "200",
		"tcp_enabled": "1", "tcp_sendValue": "PING", "tcp_matchType": option("string"), "tcp_negate": "1", "tcp_matchValue": "PONG",
		"agent_port": "123", "mysql_user": "mysql", "mysql_post41": "1", "pgsql_user": "postgres", "smtp_domain": "mail.example.com", "esmtp_domain": "example.com",
	})
	if model.Id.ValueString() != "test-id" || model.Type.ValueString() != "http" || model.HTTPMethod.ValueString() != "get" || !model.HTTPExpressionEnabled.ValueBool() || !model.TCPNegate.ValueBool() {
		t.Fatalf("incorrect model: %#v", model)
	}
	model.Description = types.StringValue("example")
	payload := model.toHealthMonitorObject()
	for key, expected := range map[string]string{"name": "health", "sslSNI": "example.com", "checkport": "443", "http_method": "get", "http_expressionEnabled": "1", "http_expression": "status", "http_value": "200", "tcp_sendValue": "PING", "tcp_matchType": "string", "tcp_negate": "1", "tcp_matchValue": "PONG", "mysql_post41": "1", "pgsql_user": "postgres"} {
		if payload[key] != expected {
			t.Errorf("payload[%s] = %v, want %q", key, payload[key], expected)
		}
	}
}
