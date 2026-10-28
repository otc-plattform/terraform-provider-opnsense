package haproxy

import (
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type healthMonitorResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	Description           types.String `tfsdk:"description"`
	Type                  types.String `tfsdk:"type"`
	Interval              types.String `tfsdk:"interval"`
	SSL                   types.String `tfsdk:"ssl"`
	SSLSNI                types.String `tfsdk:"ssl_sni"`
	CheckPort             types.String `tfsdk:"check_port"`
	HTTPMethod            types.String `tfsdk:"http_method"`
	HTTPURI               types.String `tfsdk:"http_uri"`
	HTTPVersion           types.String `tfsdk:"http_version"`
	HTTPHost              types.String `tfsdk:"http_host"`
	HTTPExpressionEnabled types.Bool   `tfsdk:"http_expression_enabled"`
	HTTPExpression        types.String `tfsdk:"http_expression"`
	HTTPNegate            types.Bool   `tfsdk:"http_negate"`
	HTTPValue             types.String `tfsdk:"http_value"`
	TCPEnabled            types.Bool   `tfsdk:"tcp_enabled"`
	TCPSendValue          types.String `tfsdk:"tcp_send_value"`
	TCPMatchType          types.String `tfsdk:"tcp_match_type"`
	TCPNegate             types.Bool   `tfsdk:"tcp_negate"`
	TCPMatchValue         types.String `tfsdk:"tcp_match_value"`
	AgentPort             types.String `tfsdk:"agent_port"`
	MySQLUser             types.String `tfsdk:"mysql_user"`
	MySQLPost41           types.Bool   `tfsdk:"mysql_post41"`
	PostgresUser          types.String `tfsdk:"postgres_user"`
	SMTPDomain            types.String `tfsdk:"smtp_domain"`
	ESMTPDomain           types.String `tfsdk:"esmtp_domain"`
}

func healthMonitorResourceSchema() schema.Schema {
	str := func(description, value string) schema.StringAttribute {
		return schema.StringAttribute{MarkdownDescription: description, Optional: true, Computed: true, Default: stringdefault.StaticString(value)}
	}
	flag := func(description string) schema.BoolAttribute {
		return schema.BoolAttribute{MarkdownDescription: description, Optional: true, Computed: true, Default: booldefault.StaticBool(false)}
	}
	return schema.Schema{MarkdownDescription: "Manage an OPNsense HAProxy health monitor.", Attributes: map[string]schema.Attribute{
		"id":                      schema.StringAttribute{Computed: true, MarkdownDescription: "Health monitor UUID.", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
		"name":                    schema.StringAttribute{Required: true, MarkdownDescription: "Health monitor name."},
		"description":             str("Description.", ""),
		"type":                    str("Check type, such as http, tcp, agent, ldap, mysql, pgsql, redis, smtp, esmtp or ssl.", "http"),
		"interval":                str("Check interval.", "2s"),
		"ssl":                     str("TLS mode: nopref, ssl, sslsni or nossl.", "nopref"),
		"ssl_sni":                 str("SNI name for TLS checks.", ""),
		"check_port":              str("Override check port.", ""),
		"http_method":             str("HTTP check method.", "options"),
		"http_uri":                str("HTTP check URI.", "/"),
		"http_version":            str("HTTP version: http10, http11 or http2.", "http10"),
		"http_host":               str("HTTP host header.", "localhost"),
		"http_expression_enabled": flag("Enable an HTTP response expression."),
		"http_expression":         str("HTTP response expression type.", ""),
		"http_negate":             flag("Negate the HTTP response match."),
		"http_value":              str("HTTP response match value.", ""),
		"tcp_enabled":             flag("Enable TCP send and response matching."),
		"tcp_send_value":          str("TCP payload to send.", ""),
		"tcp_match_type":          str("TCP response match type.", "string"),
		"tcp_negate":              flag("Negate the TCP response match."),
		"tcp_match_value":         str("TCP response match value.", ""),
		"agent_port":              str("Agent check port.", ""),
		"mysql_user":              str("MySQL check username.", ""),
		"mysql_post41":            flag("Use MySQL 4.1 or newer checks."),
		"postgres_user":           str("PostgreSQL check username.", ""),
		"smtp_domain":             str("SMTP check domain.", ""),
		"esmtp_domain":            str("ESMTP check domain.", ""),
	}}
}

func healthMonitorDataSourceSchema() dschema.Schema {
	attrs := map[string]dschema.Attribute{"id": dschema.StringAttribute{Required: true, MarkdownDescription: "Health monitor UUID."}}
	for _, name := range []string{"name", "description", "type", "interval", "ssl", "ssl_sni", "check_port", "http_method", "http_uri", "http_version", "http_host", "http_expression", "http_value", "tcp_send_value", "tcp_match_type", "tcp_match_value", "agent_port", "mysql_user", "postgres_user", "smtp_domain", "esmtp_domain"} {
		attrs[name] = dschema.StringAttribute{Computed: true}
	}
	for _, name := range []string{"http_expression_enabled", "http_negate", "tcp_enabled", "tcp_negate", "mysql_post41"} {
		attrs[name] = dschema.BoolAttribute{Computed: true}
	}
	return dschema.Schema{MarkdownDescription: "Read an OPNsense HAProxy health monitor.", Attributes: attrs}
}
