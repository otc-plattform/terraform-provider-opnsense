package haproxy

import (
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type actionResourceModel struct {
	Id types.String `tfsdk:"id"`
	Enabled types.Bool `tfsdk:"enabled"`
	Name types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	TestType types.String `tfsdk:"test_type"`
	LinkedAcls types.Set `tfsdk:"linked_acls"`
	Operator types.String `tfsdk:"operator"`
	Type_ types.String `tfsdk:"type"`
	UseBackend types.String `tfsdk:"use_backend"`
	UseServer types.String `tfsdk:"use_server"`
	FcgiPassHeader types.String `tfsdk:"fcgi_pass_header"`
	FcgiSetParam types.String `tfsdk:"fcgi_set_param"`
	MonitorFailUri types.String `tfsdk:"monitor_fail_uri"`
	Custom types.String `tfsdk:"custom"`
	HttpAfterResponseAction types.String `tfsdk:"http_after_response_action"`
	HttpAfterResponseOption types.String `tfsdk:"http_after_response_option"`
	HttpRequestAction types.String `tfsdk:"http_request_action"`
	HttpRequestOption types.String `tfsdk:"http_request_option"`
	HttpResponseAction types.String `tfsdk:"http_response_action"`
	HttpResponseOption types.String `tfsdk:"http_response_option"`
	TcpRequestAction types.String `tfsdk:"tcp_request_action"`
	TcpRequestOption types.String `tfsdk:"tcp_request_option"`
	TcpResponseAction types.String `tfsdk:"tcp_response_action"`
	TcpResponseOption types.String `tfsdk:"tcp_response_option"`
	HttpRequestAuth types.String `tfsdk:"http_request_auth"`
	HttpRequestDenyStatus types.String `tfsdk:"http_request_deny_status"`
	HttpRequestRedirect types.String `tfsdk:"http_request_redirect"`
	HttpRequestLua types.String `tfsdk:"http_request_lua"`
	HttpRequestUseService types.String `tfsdk:"http_request_use_service"`
	HttpRequestAddHeaderName types.String `tfsdk:"http_request_add_header_name"`
	HttpRequestAddHeaderContent types.String `tfsdk:"http_request_add_header_content"`
	HttpRequestSetHeaderName types.String `tfsdk:"http_request_set_header_name"`
	HttpRequestSetHeaderContent types.String `tfsdk:"http_request_set_header_content"`
	HttpRequestDelHeaderName types.String `tfsdk:"http_request_del_header_name"`
	HttpRequestReplaceHeaderName types.String `tfsdk:"http_request_replace_header_name"`
	HttpRequestReplaceHeaderRegex types.String `tfsdk:"http_request_replace_header_regex"`
	HttpRequestReplaceValueName types.String `tfsdk:"http_request_replace_value_name"`
	HttpRequestReplaceValueRegex types.String `tfsdk:"http_request_replace_value_regex"`
	HttpRequestSetPath types.String `tfsdk:"http_request_set_path"`
	HttpRequestSetVarScope types.String `tfsdk:"http_request_set_var_scope"`
	HttpRequestSetVarName types.String `tfsdk:"http_request_set_var_name"`
	HttpRequestSetVarExpr types.String `tfsdk:"http_request_set_var_expr"`
	HttpResponseLua types.String `tfsdk:"http_response_lua"`
	HttpResponseAddHeaderName types.String `tfsdk:"http_response_add_header_name"`
	HttpResponseAddHeaderContent types.String `tfsdk:"http_response_add_header_content"`
	HttpResponseSetHeaderName types.String `tfsdk:"http_response_set_header_name"`
	HttpResponseSetHeaderContent types.String `tfsdk:"http_response_set_header_content"`
	HttpResponseDelHeaderName types.String `tfsdk:"http_response_del_header_name"`
	HttpResponseReplaceHeaderName types.String `tfsdk:"http_response_replace_header_name"`
	HttpResponseReplaceHeaderRegex types.String `tfsdk:"http_response_replace_header_regex"`
	HttpResponseReplaceValueName types.String `tfsdk:"http_response_replace_value_name"`
	HttpResponseReplaceValueRegex types.String `tfsdk:"http_response_replace_value_regex"`
	HttpResponseSetStatusCode types.String `tfsdk:"http_response_set_status_code"`
	HttpResponseSetStatusReason types.String `tfsdk:"http_response_set_status_reason"`
	HttpResponseSetVarScope types.String `tfsdk:"http_response_set_var_scope"`
	HttpResponseSetVarName types.String `tfsdk:"http_response_set_var_name"`
	HttpResponseSetVarExpr types.String `tfsdk:"http_response_set_var_expr"`
	TcpRequestContentLua types.String `tfsdk:"tcp_request_content_lua"`
	TcpRequestContentUseService types.String `tfsdk:"tcp_request_content_use_service"`
	TcpRequestInspectDelay types.String `tfsdk:"tcp_request_inspect_delay"`
	TcpResponseContentLua types.String `tfsdk:"tcp_response_content_lua"`
	TcpResponseInspectDelay types.String `tfsdk:"tcp_response_inspect_delay"`
	MapDataUseBackendFile types.String `tfsdk:"map_data_use_backend_file"`
	MapDataUseBackendDefault types.String `tfsdk:"map_data_use_backend_default"`
	MapDataUseBackendInput types.String `tfsdk:"map_data_use_backend_input"`
	MapUseBackendFile types.String `tfsdk:"map_use_backend_file"`
	MapUseBackendDefault types.String `tfsdk:"map_use_backend_default"`
	CompressionAlgoRes types.Set `tfsdk:"compression_algo_res"`
	CompressionAlgoReq types.Set `tfsdk:"compression_algo_req"`
	CompressionMimeRes types.Set `tfsdk:"compression_mime_res"`
	CompressionMimeReq types.Set `tfsdk:"compression_mime_req"`
	CompressionOffloading types.Bool `tfsdk:"compression_offloading"`
	CompressionMinsizeRes types.String `tfsdk:"compression_minsize_res"`
	CompressionMinsizeReq types.String `tfsdk:"compression_minsize_req"`
	CompressionDirection types.String `tfsdk:"compression_direction"`
	GpcNumber types.String `tfsdk:"gpc_number"`
	GptNumber types.String `tfsdk:"gpt_number"`
	ScNumber types.String `tfsdk:"sc_number"`
	Mapfile types.String `tfsdk:"mapfile"`
	MapDefault types.String `tfsdk:"map_default"`
	SampleFetch types.String `tfsdk:"sample_fetch"`
}

func actionResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense HAProxy rule/action.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID of the haproxy action.",
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the enabled option for this action.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Action name value.",
				Required: true,
				Validators: []validator.String{stringvalidator.LengthAtLeast(1)},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Action description value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"test_type": schema.StringAttribute{
				MarkdownDescription: "Action test_type option. One of: if, unless.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("if", "unless")},
			},
			"linked_acls": schema.SetAttribute{
				MarkdownDescription: "List of linked_acls values for this action.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"operator": schema.StringAttribute{
				MarkdownDescription: "Action operator option. One of: , and, or.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "and", "or")},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Action type option. One of: compression, fcgi_pass_header, fcgi_set_param, http-after-response, http-request, http-response, map_data_use_backend, map_use_backend, monitor_fail, tcp-request, tcp-response, use_backend, use_server, custom.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("compression", "fcgi_pass_header", "fcgi_set_param", "http-after-response", "http-request", "http-response", "map_data_use_backend", "map_use_backend", "monitor_fail", "tcp-request", "tcp-response", "use_backend", "use_server", "custom")},
			},
			"use_backend": schema.StringAttribute{
				MarkdownDescription: "Action use_backend (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"use_server": schema.StringAttribute{
				MarkdownDescription: "Action use_server (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"fcgi_pass_header": schema.StringAttribute{
				MarkdownDescription: "Action fcgi_pass_header value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"fcgi_set_param": schema.StringAttribute{
				MarkdownDescription: "Action fcgi_set_param value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"monitor_fail_uri": schema.StringAttribute{
				MarkdownDescription: "Action monitor_fail_uri value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"custom": schema.StringAttribute{
				MarkdownDescription: "Action custom value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_after_response_action": schema.StringAttribute{
				MarkdownDescription: "Action http_after_response_action option. One of: , add-header, allow, capture, del-header, del-map, do-log, replace-header, replace-value, sc-add-gpc, sc-inc-gpc, sc-inc-gpc0, sc-inc-gpc1, sc-set-gpt, sc-set-gpt0, set-header, set-log-level, set-map, set-status, set-var, set-var-fmt, strict-mode, unset-var.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "add-header", "allow", "capture", "del-header", "del-map", "do-log", "replace-header", "replace-value", "sc-add-gpc", "sc-inc-gpc", "sc-inc-gpc0", "sc-inc-gpc1", "sc-set-gpt", "sc-set-gpt0", "set-header", "set-log-level", "set-map", "set-status", "set-var", "set-var-fmt", "strict-mode", "unset-var")},
			},
			"http_after_response_option": schema.StringAttribute{
				MarkdownDescription: "Action http_after_response_option value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_action": schema.StringAttribute{
				MarkdownDescription: "Action http_request_action option. One of: , add-acl, add-header, allow, auth, cache-use, capture, del-acl, del-header, del-map, deny, disable-l7-retry, do-log, do-resolve, early-hint, lua, normalize-uri, redirect, reject, replace-header, replace-path, replace-pathq, replace-uri, replace-value, return, sc-add-gpc, sc-inc-gpc, sc-inc-gpc0, sc-inc-gpc1, sc-set-gpt, sc-set-gpt0, send-spoe-group, set-dst, set-dst-port, set-fc-mark, set-fc-tos, set-header, set-log-level, set-map, set-method, set-nice, set-path, set-pathq, set-priority-class, set-priority-offset, set-query, set-src, set-src-port, set-timeout, set-uri, set-var, set-var-fmt, silent-drop, strict-mode, tarpit, track-sc0, track-sc1, track-sc2, unset-var, use-service, wait-for-body, wait-for-handshake.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "add-acl", "add-header", "allow", "auth", "cache-use", "capture", "del-acl", "del-header", "del-map", "deny", "disable-l7-retry", "do-log", "do-resolve", "early-hint", "lua", "normalize-uri", "redirect", "reject", "replace-header", "replace-path", "replace-pathq", "replace-uri", "replace-value", "return", "sc-add-gpc", "sc-inc-gpc", "sc-inc-gpc0", "sc-inc-gpc1", "sc-set-gpt", "sc-set-gpt0", "send-spoe-group", "set-dst", "set-dst-port", "set-fc-mark", "set-fc-tos", "set-header", "set-log-level", "set-map", "set-method", "set-nice", "set-path", "set-pathq", "set-priority-class", "set-priority-offset", "set-query", "set-src", "set-src-port", "set-timeout", "set-uri", "set-var", "set-var-fmt", "silent-drop", "strict-mode", "tarpit", "track-sc0", "track-sc1", "track-sc2", "unset-var", "use-service", "wait-for-body", "wait-for-handshake")},
			},
			"http_request_option": schema.StringAttribute{
				MarkdownDescription: "Action http_request_option value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_action": schema.StringAttribute{
				MarkdownDescription: "Action http_response_action option. One of: , add-acl, add-header, allow, cache-store, capture, del-acl, del-header, del-map, deny, do-log, lua, redirect, replace-header, replace-value, return, sc-add-gpc, sc-inc-gpc, sc-inc-gpc0, sc-inc-gpc1, sc-set-gpt, sc-set-gpt0, send-spoe-group, set-fc-mark, set-fc-tos, set-header, set-log-level, set-map, set-nice, set-status, set-timeout, set-var, set-var-fmt, silent-drop, strict-mode, track-sc0, track-sc1, track-sc2, unset-var, wait-for-body.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "add-acl", "add-header", "allow", "cache-store", "capture", "del-acl", "del-header", "del-map", "deny", "do-log", "lua", "redirect", "replace-header", "replace-value", "return", "sc-add-gpc", "sc-inc-gpc", "sc-inc-gpc0", "sc-inc-gpc1", "sc-set-gpt", "sc-set-gpt0", "send-spoe-group", "set-fc-mark", "set-fc-tos", "set-header", "set-log-level", "set-map", "set-nice", "set-status", "set-timeout", "set-var", "set-var-fmt", "silent-drop", "strict-mode", "track-sc0", "track-sc1", "track-sc2", "unset-var", "wait-for-body")},
			},
			"http_response_option": schema.StringAttribute{
				MarkdownDescription: "Action http_response_option value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tcp_request_action": schema.StringAttribute{
				MarkdownDescription: "Action tcp_request_action option. One of: , connection_accept, connection_expect-netscaler-cip, connection_expect-proxy, connection_fc-silent-drop, connection_reject, connection_sc-add-gpc, connection_sc-inc-gpc, connection_sc-inc-gpc0, connection_sc-inc-gpc1, connection_sc-set-gpt, connection_sc-set-gpt0, connection_send-spoe-group, connection_set-dst, connection_set-dst-port, connection_set-fc-mark, connection_set-fc-tos, connection_set-log-level, connection_set-src, connection_set-src-port, connection_set-var, connection_set-var-fmt, connection_silent-drop, connection_track-sc0, connection_track-sc1, connection_track-sc2, connection_unset-var, content_accept, content_capture, content_do-resolve, content_lua, content_reject, content_sc-add-gpc, content_sc-inc-gpc, content_sc-inc-gpc0, content_sc-inc-gpc1, content_sc-set-gpt, content_sc-set-gpt0, content_send-spoe-group, content_set-dst, content_set-dst-port, content_set-fc-mark, content_set-fc-tos, content_set-log-level, content_set-nice, content_set-priority-class, content_set-priority-offset, content_set-src, content_set-src-port, content_set-var, content_set-var-fmt, content_silent-drop, content_switch-mode, content_track-sc0, content_track-sc1, content_track-sc2, content_unset-var, content_use-service, inspect-delay, session_accept, session_attach-srv, session_reject, session_sc-add-gpc, session_sc-inc-gpc, session_sc-inc-gpc0, session_sc-inc-gpc1, session_sc-set-gpt, session_sc-set-gpt0, session_send-spoe-group, session_set-dst, session_set-dst-port, session_set-fc-mark, session_set-fc-tos, session_set-log-level, session_set-src, session_set-src-port, session_set-var, session_set-var-fmt, session_silent-drop, session_track-sc0, session_track-sc1, session_track-sc2, session_unset-var.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "connection_accept", "connection_expect-netscaler-cip", "connection_expect-proxy", "connection_fc-silent-drop", "connection_reject", "connection_sc-add-gpc", "connection_sc-inc-gpc", "connection_sc-inc-gpc0", "connection_sc-inc-gpc1", "connection_sc-set-gpt", "connection_sc-set-gpt0", "connection_send-spoe-group", "connection_set-dst", "connection_set-dst-port", "connection_set-fc-mark", "connection_set-fc-tos", "connection_set-log-level", "connection_set-src", "connection_set-src-port", "connection_set-var", "connection_set-var-fmt", "connection_silent-drop", "connection_track-sc0", "connection_track-sc1", "connection_track-sc2", "connection_unset-var", "content_accept", "content_capture", "content_do-resolve", "content_lua", "content_reject", "content_sc-add-gpc", "content_sc-inc-gpc", "content_sc-inc-gpc0", "content_sc-inc-gpc1", "content_sc-set-gpt", "content_sc-set-gpt0", "content_send-spoe-group", "content_set-dst", "content_set-dst-port", "content_set-fc-mark", "content_set-fc-tos", "content_set-log-level", "content_set-nice", "content_set-priority-class", "content_set-priority-offset", "content_set-src", "content_set-src-port", "content_set-var", "content_set-var-fmt", "content_silent-drop", "content_switch-mode", "content_track-sc0", "content_track-sc1", "content_track-sc2", "content_unset-var", "content_use-service", "inspect-delay", "session_accept", "session_attach-srv", "session_reject", "session_sc-add-gpc", "session_sc-inc-gpc", "session_sc-inc-gpc0", "session_sc-inc-gpc1", "session_sc-set-gpt", "session_sc-set-gpt0", "session_send-spoe-group", "session_set-dst", "session_set-dst-port", "session_set-fc-mark", "session_set-fc-tos", "session_set-log-level", "session_set-src", "session_set-src-port", "session_set-var", "session_set-var-fmt", "session_silent-drop", "session_track-sc0", "session_track-sc1", "session_track-sc2", "session_unset-var")},
			},
			"tcp_request_option": schema.StringAttribute{
				MarkdownDescription: "Action tcp_request_option value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tcp_response_action": schema.StringAttribute{
				MarkdownDescription: "Action tcp_response_action option. One of: , content_accept, content_close, content_lua, content_reject, content_sc-add-gpc, content_sc-inc-gpc, content_sc-inc-gpc0, content_sc-inc-gpc1, content_sc-set-gpt, content_sc-set-gpt0, content_send-spoe-group, content_set-fc-mark, content_set-fc-tos, content_set-log-level, content_set-nice, content_set-var, content_set-var-fmt, content_silent-drop, content_unset-var, inspect-delay.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "content_accept", "content_close", "content_lua", "content_reject", "content_sc-add-gpc", "content_sc-inc-gpc", "content_sc-inc-gpc0", "content_sc-inc-gpc1", "content_sc-set-gpt", "content_sc-set-gpt0", "content_send-spoe-group", "content_set-fc-mark", "content_set-fc-tos", "content_set-log-level", "content_set-nice", "content_set-var", "content_set-var-fmt", "content_silent-drop", "content_unset-var", "inspect-delay")},
			},
			"tcp_response_option": schema.StringAttribute{
				MarkdownDescription: "Action tcp_response_option value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_auth": schema.StringAttribute{
				MarkdownDescription: "Action http_request_auth value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_deny_status": schema.StringAttribute{
				MarkdownDescription: "Action http_request_deny_status value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_redirect": schema.StringAttribute{
				MarkdownDescription: "Action http_request_redirect value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_lua": schema.StringAttribute{
				MarkdownDescription: "Action http_request_lua value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_use_service": schema.StringAttribute{
				MarkdownDescription: "Action http_request_use_service value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_add_header_name": schema.StringAttribute{
				MarkdownDescription: "Action http_request_add_header_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_add_header_content": schema.StringAttribute{
				MarkdownDescription: "Action http_request_add_header_content value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_set_header_name": schema.StringAttribute{
				MarkdownDescription: "Action http_request_set_header_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_set_header_content": schema.StringAttribute{
				MarkdownDescription: "Action http_request_set_header_content value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_del_header_name": schema.StringAttribute{
				MarkdownDescription: "Action http_request_del_header_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_replace_header_name": schema.StringAttribute{
				MarkdownDescription: "Action http_request_replace_header_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_replace_header_regex": schema.StringAttribute{
				MarkdownDescription: "Action http_request_replace_header_regex value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_replace_value_name": schema.StringAttribute{
				MarkdownDescription: "Action http_request_replace_value_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_replace_value_regex": schema.StringAttribute{
				MarkdownDescription: "Action http_request_replace_value_regex value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_set_path": schema.StringAttribute{
				MarkdownDescription: "Action http_request_set_path value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_set_var_scope": schema.StringAttribute{
				MarkdownDescription: "Action http_request_set_var_scope option. One of: , proc, sess, txn, req, res.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "proc", "sess", "txn", "req", "res")},
			},
			"http_request_set_var_name": schema.StringAttribute{
				MarkdownDescription: "Action http_request_set_var_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_request_set_var_expr": schema.StringAttribute{
				MarkdownDescription: "Action http_request_set_var_expr value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_lua": schema.StringAttribute{
				MarkdownDescription: "Action http_response_lua value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_add_header_name": schema.StringAttribute{
				MarkdownDescription: "Action http_response_add_header_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_add_header_content": schema.StringAttribute{
				MarkdownDescription: "Action http_response_add_header_content value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_set_header_name": schema.StringAttribute{
				MarkdownDescription: "Action http_response_set_header_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_set_header_content": schema.StringAttribute{
				MarkdownDescription: "Action http_response_set_header_content value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_del_header_name": schema.StringAttribute{
				MarkdownDescription: "Action http_response_del_header_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_replace_header_name": schema.StringAttribute{
				MarkdownDescription: "Action http_response_replace_header_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_replace_header_regex": schema.StringAttribute{
				MarkdownDescription: "Action http_response_replace_header_regex value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_replace_value_name": schema.StringAttribute{
				MarkdownDescription: "Action http_response_replace_value_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_replace_value_regex": schema.StringAttribute{
				MarkdownDescription: "Action http_response_replace_value_regex value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_set_status_code": schema.StringAttribute{
				MarkdownDescription: "Action http_response_set_status_code value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_set_status_reason": schema.StringAttribute{
				MarkdownDescription: "Action http_response_set_status_reason value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_set_var_scope": schema.StringAttribute{
				MarkdownDescription: "Action http_response_set_var_scope option. One of: , proc, sess, txn, req, res.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "proc", "sess", "txn", "req", "res")},
			},
			"http_response_set_var_name": schema.StringAttribute{
				MarkdownDescription: "Action http_response_set_var_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http_response_set_var_expr": schema.StringAttribute{
				MarkdownDescription: "Action http_response_set_var_expr value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tcp_request_content_lua": schema.StringAttribute{
				MarkdownDescription: "Action tcp_request_content_lua value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tcp_request_content_use_service": schema.StringAttribute{
				MarkdownDescription: "Action tcp_request_content_use_service value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tcp_request_inspect_delay": schema.StringAttribute{
				MarkdownDescription: "Action tcp_request_inspect_delay value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tcp_response_content_lua": schema.StringAttribute{
				MarkdownDescription: "Action tcp_response_content_lua value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tcp_response_inspect_delay": schema.StringAttribute{
				MarkdownDescription: "Action tcp_response_inspect_delay value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"map_data_use_backend_file": schema.StringAttribute{
				MarkdownDescription: "Action map_data_use_backend_file (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"map_data_use_backend_default": schema.StringAttribute{
				MarkdownDescription: "Action map_data_use_backend_default (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"map_data_use_backend_input": schema.StringAttribute{
				MarkdownDescription: "Action map_data_use_backend_input value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"map_use_backend_file": schema.StringAttribute{
				MarkdownDescription: "Action map_use_backend_file (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"map_use_backend_default": schema.StringAttribute{
				MarkdownDescription: "Action map_use_backend_default (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"compression_algo_res": schema.SetAttribute{
				MarkdownDescription: "Selected compression_algo_res values for this action. One or more of: , gzip, deflate, raw-deflate.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"compression_algo_req": schema.SetAttribute{
				MarkdownDescription: "Selected compression_algo_req values for this action. One or more of: , gzip, deflate, raw-deflate.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"compression_mime_res": schema.SetAttribute{
				MarkdownDescription: "Selected compression_mime_res values for this action. One or more of: .",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"compression_mime_req": schema.SetAttribute{
				MarkdownDescription: "Selected compression_mime_req values for this action. One or more of: .",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"compression_offloading": schema.BoolAttribute{
				MarkdownDescription: "Enable the compression_offloading option for this action.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"compression_minsize_res": schema.StringAttribute{
				MarkdownDescription: "Action compression_minsize_res value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"compression_minsize_req": schema.StringAttribute{
				MarkdownDescription: "Action compression_minsize_req value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"compression_direction": schema.StringAttribute{
				MarkdownDescription: "Action compression_direction option. One of: , response, request, both.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "response", "request", "both")},
			},
			"gpc_number": schema.StringAttribute{
				MarkdownDescription: "Action gpc_number value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"gpt_number": schema.StringAttribute{
				MarkdownDescription: "Action gpt_number value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_number": schema.StringAttribute{
				MarkdownDescription: "Action sc_number value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"mapfile": schema.StringAttribute{
				MarkdownDescription: "Action mapfile (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"map_default": schema.StringAttribute{
				MarkdownDescription: "Action map_default value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sample_fetch": schema.StringAttribute{
				MarkdownDescription: "Action sample_fetch value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
		},
	}
}

func actionDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense HAProxy rule/action.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the haproxy action.",
				Required: true,
				Validators: []validator.String{stringvalidator.LengthAtLeast(1)},
			},
			"enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the enabled option for this action.", Computed: true},
			"name": dschema.StringAttribute{MarkdownDescription: "Action name value.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Action description value.", Computed: true},
			"test_type": dschema.StringAttribute{MarkdownDescription: "Action test_type option. One of: if, unless.", Computed: true},
			"linked_acls": dschema.SetAttribute{MarkdownDescription: "List of linked_acls values for this action.", ElementType: types.StringType, Computed: true},
			"operator": dschema.StringAttribute{MarkdownDescription: "Action operator option. One of: , and, or.", Computed: true},
			"type": dschema.StringAttribute{MarkdownDescription: "Action type option. One of: compression, fcgi_pass_header, fcgi_set_param, http-after-response, http-request, http-response, map_data_use_backend, map_use_backend, monitor_fail, tcp-request, tcp-response, use_backend, use_server, custom.", Computed: true},
			"use_backend": dschema.StringAttribute{MarkdownDescription: "Action use_backend (UUID reference).", Computed: true},
			"use_server": dschema.StringAttribute{MarkdownDescription: "Action use_server (UUID reference).", Computed: true},
			"fcgi_pass_header": dschema.StringAttribute{MarkdownDescription: "Action fcgi_pass_header value.", Computed: true},
			"fcgi_set_param": dschema.StringAttribute{MarkdownDescription: "Action fcgi_set_param value.", Computed: true},
			"monitor_fail_uri": dschema.StringAttribute{MarkdownDescription: "Action monitor_fail_uri value.", Computed: true},
			"custom": dschema.StringAttribute{MarkdownDescription: "Action custom value.", Computed: true},
			"http_after_response_action": dschema.StringAttribute{MarkdownDescription: "Action http_after_response_action option. One of: , add-header, allow, capture, del-header, del-map, do-log, replace-header, replace-value, sc-add-gpc, sc-inc-gpc, sc-inc-gpc0, sc-inc-gpc1, sc-set-gpt, sc-set-gpt0, set-header, set-log-level, set-map, set-status, set-var, set-var-fmt, strict-mode, unset-var.", Computed: true},
			"http_after_response_option": dschema.StringAttribute{MarkdownDescription: "Action http_after_response_option value.", Computed: true},
			"http_request_action": dschema.StringAttribute{MarkdownDescription: "Action http_request_action option. One of: , add-acl, add-header, allow, auth, cache-use, capture, del-acl, del-header, del-map, deny, disable-l7-retry, do-log, do-resolve, early-hint, lua, normalize-uri, redirect, reject, replace-header, replace-path, replace-pathq, replace-uri, replace-value, return, sc-add-gpc, sc-inc-gpc, sc-inc-gpc0, sc-inc-gpc1, sc-set-gpt, sc-set-gpt0, send-spoe-group, set-dst, set-dst-port, set-fc-mark, set-fc-tos, set-header, set-log-level, set-map, set-method, set-nice, set-path, set-pathq, set-priority-class, set-priority-offset, set-query, set-src, set-src-port, set-timeout, set-uri, set-var, set-var-fmt, silent-drop, strict-mode, tarpit, track-sc0, track-sc1, track-sc2, unset-var, use-service, wait-for-body, wait-for-handshake.", Computed: true},
			"http_request_option": dschema.StringAttribute{MarkdownDescription: "Action http_request_option value.", Computed: true},
			"http_response_action": dschema.StringAttribute{MarkdownDescription: "Action http_response_action option. One of: , add-acl, add-header, allow, cache-store, capture, del-acl, del-header, del-map, deny, do-log, lua, redirect, replace-header, replace-value, return, sc-add-gpc, sc-inc-gpc, sc-inc-gpc0, sc-inc-gpc1, sc-set-gpt, sc-set-gpt0, send-spoe-group, set-fc-mark, set-fc-tos, set-header, set-log-level, set-map, set-nice, set-status, set-timeout, set-var, set-var-fmt, silent-drop, strict-mode, track-sc0, track-sc1, track-sc2, unset-var, wait-for-body.", Computed: true},
			"http_response_option": dschema.StringAttribute{MarkdownDescription: "Action http_response_option value.", Computed: true},
			"tcp_request_action": dschema.StringAttribute{MarkdownDescription: "Action tcp_request_action option. One of: , connection_accept, connection_expect-netscaler-cip, connection_expect-proxy, connection_fc-silent-drop, connection_reject, connection_sc-add-gpc, connection_sc-inc-gpc, connection_sc-inc-gpc0, connection_sc-inc-gpc1, connection_sc-set-gpt, connection_sc-set-gpt0, connection_send-spoe-group, connection_set-dst, connection_set-dst-port, connection_set-fc-mark, connection_set-fc-tos, connection_set-log-level, connection_set-src, connection_set-src-port, connection_set-var, connection_set-var-fmt, connection_silent-drop, connection_track-sc0, connection_track-sc1, connection_track-sc2, connection_unset-var, content_accept, content_capture, content_do-resolve, content_lua, content_reject, content_sc-add-gpc, content_sc-inc-gpc, content_sc-inc-gpc0, content_sc-inc-gpc1, content_sc-set-gpt, content_sc-set-gpt0, content_send-spoe-group, content_set-dst, content_set-dst-port, content_set-fc-mark, content_set-fc-tos, content_set-log-level, content_set-nice, content_set-priority-class, content_set-priority-offset, content_set-src, content_set-src-port, content_set-var, content_set-var-fmt, content_silent-drop, content_switch-mode, content_track-sc0, content_track-sc1, content_track-sc2, content_unset-var, content_use-service, inspect-delay, session_accept, session_attach-srv, session_reject, session_sc-add-gpc, session_sc-inc-gpc, session_sc-inc-gpc0, session_sc-inc-gpc1, session_sc-set-gpt, session_sc-set-gpt0, session_send-spoe-group, session_set-dst, session_set-dst-port, session_set-fc-mark, session_set-fc-tos, session_set-log-level, session_set-src, session_set-src-port, session_set-var, session_set-var-fmt, session_silent-drop, session_track-sc0, session_track-sc1, session_track-sc2, session_unset-var.", Computed: true},
			"tcp_request_option": dschema.StringAttribute{MarkdownDescription: "Action tcp_request_option value.", Computed: true},
			"tcp_response_action": dschema.StringAttribute{MarkdownDescription: "Action tcp_response_action option. One of: , content_accept, content_close, content_lua, content_reject, content_sc-add-gpc, content_sc-inc-gpc, content_sc-inc-gpc0, content_sc-inc-gpc1, content_sc-set-gpt, content_sc-set-gpt0, content_send-spoe-group, content_set-fc-mark, content_set-fc-tos, content_set-log-level, content_set-nice, content_set-var, content_set-var-fmt, content_silent-drop, content_unset-var, inspect-delay.", Computed: true},
			"tcp_response_option": dschema.StringAttribute{MarkdownDescription: "Action tcp_response_option value.", Computed: true},
			"http_request_auth": dschema.StringAttribute{MarkdownDescription: "Action http_request_auth value.", Computed: true},
			"http_request_deny_status": dschema.StringAttribute{MarkdownDescription: "Action http_request_deny_status value.", Computed: true},
			"http_request_redirect": dschema.StringAttribute{MarkdownDescription: "Action http_request_redirect value.", Computed: true},
			"http_request_lua": dschema.StringAttribute{MarkdownDescription: "Action http_request_lua value.", Computed: true},
			"http_request_use_service": dschema.StringAttribute{MarkdownDescription: "Action http_request_use_service value.", Computed: true},
			"http_request_add_header_name": dschema.StringAttribute{MarkdownDescription: "Action http_request_add_header_name value.", Computed: true},
			"http_request_add_header_content": dschema.StringAttribute{MarkdownDescription: "Action http_request_add_header_content value.", Computed: true},
			"http_request_set_header_name": dschema.StringAttribute{MarkdownDescription: "Action http_request_set_header_name value.", Computed: true},
			"http_request_set_header_content": dschema.StringAttribute{MarkdownDescription: "Action http_request_set_header_content value.", Computed: true},
			"http_request_del_header_name": dschema.StringAttribute{MarkdownDescription: "Action http_request_del_header_name value.", Computed: true},
			"http_request_replace_header_name": dschema.StringAttribute{MarkdownDescription: "Action http_request_replace_header_name value.", Computed: true},
			"http_request_replace_header_regex": dschema.StringAttribute{MarkdownDescription: "Action http_request_replace_header_regex value.", Computed: true},
			"http_request_replace_value_name": dschema.StringAttribute{MarkdownDescription: "Action http_request_replace_value_name value.", Computed: true},
			"http_request_replace_value_regex": dschema.StringAttribute{MarkdownDescription: "Action http_request_replace_value_regex value.", Computed: true},
			"http_request_set_path": dschema.StringAttribute{MarkdownDescription: "Action http_request_set_path value.", Computed: true},
			"http_request_set_var_scope": dschema.StringAttribute{MarkdownDescription: "Action http_request_set_var_scope option. One of: , proc, sess, txn, req, res.", Computed: true},
			"http_request_set_var_name": dschema.StringAttribute{MarkdownDescription: "Action http_request_set_var_name value.", Computed: true},
			"http_request_set_var_expr": dschema.StringAttribute{MarkdownDescription: "Action http_request_set_var_expr value.", Computed: true},
			"http_response_lua": dschema.StringAttribute{MarkdownDescription: "Action http_response_lua value.", Computed: true},
			"http_response_add_header_name": dschema.StringAttribute{MarkdownDescription: "Action http_response_add_header_name value.", Computed: true},
			"http_response_add_header_content": dschema.StringAttribute{MarkdownDescription: "Action http_response_add_header_content value.", Computed: true},
			"http_response_set_header_name": dschema.StringAttribute{MarkdownDescription: "Action http_response_set_header_name value.", Computed: true},
			"http_response_set_header_content": dschema.StringAttribute{MarkdownDescription: "Action http_response_set_header_content value.", Computed: true},
			"http_response_del_header_name": dschema.StringAttribute{MarkdownDescription: "Action http_response_del_header_name value.", Computed: true},
			"http_response_replace_header_name": dschema.StringAttribute{MarkdownDescription: "Action http_response_replace_header_name value.", Computed: true},
			"http_response_replace_header_regex": dschema.StringAttribute{MarkdownDescription: "Action http_response_replace_header_regex value.", Computed: true},
			"http_response_replace_value_name": dschema.StringAttribute{MarkdownDescription: "Action http_response_replace_value_name value.", Computed: true},
			"http_response_replace_value_regex": dschema.StringAttribute{MarkdownDescription: "Action http_response_replace_value_regex value.", Computed: true},
			"http_response_set_status_code": dschema.StringAttribute{MarkdownDescription: "Action http_response_set_status_code value.", Computed: true},
			"http_response_set_status_reason": dschema.StringAttribute{MarkdownDescription: "Action http_response_set_status_reason value.", Computed: true},
			"http_response_set_var_scope": dschema.StringAttribute{MarkdownDescription: "Action http_response_set_var_scope option. One of: , proc, sess, txn, req, res.", Computed: true},
			"http_response_set_var_name": dschema.StringAttribute{MarkdownDescription: "Action http_response_set_var_name value.", Computed: true},
			"http_response_set_var_expr": dschema.StringAttribute{MarkdownDescription: "Action http_response_set_var_expr value.", Computed: true},
			"tcp_request_content_lua": dschema.StringAttribute{MarkdownDescription: "Action tcp_request_content_lua value.", Computed: true},
			"tcp_request_content_use_service": dschema.StringAttribute{MarkdownDescription: "Action tcp_request_content_use_service value.", Computed: true},
			"tcp_request_inspect_delay": dschema.StringAttribute{MarkdownDescription: "Action tcp_request_inspect_delay value.", Computed: true},
			"tcp_response_content_lua": dschema.StringAttribute{MarkdownDescription: "Action tcp_response_content_lua value.", Computed: true},
			"tcp_response_inspect_delay": dschema.StringAttribute{MarkdownDescription: "Action tcp_response_inspect_delay value.", Computed: true},
			"map_data_use_backend_file": dschema.StringAttribute{MarkdownDescription: "Action map_data_use_backend_file (UUID reference).", Computed: true},
			"map_data_use_backend_default": dschema.StringAttribute{MarkdownDescription: "Action map_data_use_backend_default (UUID reference).", Computed: true},
			"map_data_use_backend_input": dschema.StringAttribute{MarkdownDescription: "Action map_data_use_backend_input value.", Computed: true},
			"map_use_backend_file": dschema.StringAttribute{MarkdownDescription: "Action map_use_backend_file (UUID reference).", Computed: true},
			"map_use_backend_default": dschema.StringAttribute{MarkdownDescription: "Action map_use_backend_default (UUID reference).", Computed: true},
			"compression_algo_res": dschema.SetAttribute{MarkdownDescription: "Selected compression_algo_res values for this action. One or more of: , gzip, deflate, raw-deflate.", ElementType: types.StringType, Computed: true},
			"compression_algo_req": dschema.SetAttribute{MarkdownDescription: "Selected compression_algo_req values for this action. One or more of: , gzip, deflate, raw-deflate.", ElementType: types.StringType, Computed: true},
			"compression_mime_res": dschema.SetAttribute{MarkdownDescription: "Selected compression_mime_res values for this action. One or more of: .", ElementType: types.StringType, Computed: true},
			"compression_mime_req": dschema.SetAttribute{MarkdownDescription: "Selected compression_mime_req values for this action. One or more of: .", ElementType: types.StringType, Computed: true},
			"compression_offloading": dschema.BoolAttribute{MarkdownDescription: "Enable the compression_offloading option for this action.", Computed: true},
			"compression_minsize_res": dschema.StringAttribute{MarkdownDescription: "Action compression_minsize_res value.", Computed: true},
			"compression_minsize_req": dschema.StringAttribute{MarkdownDescription: "Action compression_minsize_req value.", Computed: true},
			"compression_direction": dschema.StringAttribute{MarkdownDescription: "Action compression_direction option. One of: , response, request, both.", Computed: true},
			"gpc_number": dschema.StringAttribute{MarkdownDescription: "Action gpc_number value.", Computed: true},
			"gpt_number": dschema.StringAttribute{MarkdownDescription: "Action gpt_number value.", Computed: true},
			"sc_number": dschema.StringAttribute{MarkdownDescription: "Action sc_number value.", Computed: true},
			"mapfile": dschema.StringAttribute{MarkdownDescription: "Action mapfile (UUID reference).", Computed: true},
			"map_default": dschema.StringAttribute{MarkdownDescription: "Action map_default value.", Computed: true},
			"sample_fetch": dschema.StringAttribute{MarkdownDescription: "Action sample_fetch value.", Computed: true},
		},
	}
}
