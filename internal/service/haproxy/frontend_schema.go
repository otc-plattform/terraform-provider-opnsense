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

type frontendResourceModel struct {
	Id types.String `tfsdk:"id"`
	InternalId types.String `tfsdk:"internal_id"`
	Enabled types.Bool `tfsdk:"enabled"`
	Name types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Bind types.Set `tfsdk:"bind"`
	BindOptions types.String `tfsdk:"bind_options"`
	Mode types.String `tfsdk:"mode"`
	DefaultBackend types.String `tfsdk:"default_backend"`
	SslEnabled types.Bool `tfsdk:"ssl_enabled"`
	SslCertificates types.Set `tfsdk:"ssl_certificates"`
	SslDefaultCertificate types.String `tfsdk:"ssl_default_certificate"`
	SslCustomOptions types.String `tfsdk:"ssl_custom_options"`
	SslAdvancedEnabled types.Bool `tfsdk:"ssl_advanced_enabled"`
	SslBindOptions types.Set `tfsdk:"ssl_bind_options"`
	SslMinVersion types.String `tfsdk:"ssl_min_version"`
	SslMaxVersion types.String `tfsdk:"ssl_max_version"`
	SslCipherList types.String `tfsdk:"ssl_cipher_list"`
	SslCipherSuites types.String `tfsdk:"ssl_cipher_suites"`
	SslHstsEnabled types.Bool `tfsdk:"ssl_hsts_enabled"`
	SslHstsIncludeSubDomains types.Bool `tfsdk:"ssl_hsts_include_sub_domains"`
	SslHstsPreload types.Bool `tfsdk:"ssl_hsts_preload"`
	SslHstsMaxAge types.String `tfsdk:"ssl_hsts_max_age"`
	SslClientAuthEnabled types.Bool `tfsdk:"ssl_client_auth_enabled"`
	SslClientAuthVerify types.String `tfsdk:"ssl_client_auth_verify"`
	SslClientAuthCAs types.Set `tfsdk:"ssl_client_auth_c_as"`
	SslClientAuthCRLs types.Set `tfsdk:"ssl_client_auth_c_r_ls"`
	BasicAuthEnabled types.Bool `tfsdk:"basic_auth_enabled"`
	BasicAuthUsers types.Set `tfsdk:"basic_auth_users"`
	BasicAuthGroups types.Set `tfsdk:"basic_auth_groups"`
	TuningMaxConnections types.String `tfsdk:"tuning_max_connections"`
	TuningTimeoutClient types.String `tfsdk:"tuning_timeout_client"`
	TuningTimeoutHttpReq types.String `tfsdk:"tuning_timeout_http_req"`
	TuningTimeoutHttpKeepAlive types.String `tfsdk:"tuning_timeout_http_keep_alive"`
	LinkedCpuAffinityRules types.Set `tfsdk:"linked_cpu_affinity_rules"`
	TuningShards types.String `tfsdk:"tuning_shards"`
	LoggingDontLogNull types.Bool `tfsdk:"logging_dont_log_null"`
	LoggingDontLogNormal types.Bool `tfsdk:"logging_dont_log_normal"`
	LoggingLogSeparateErrors types.Bool `tfsdk:"logging_log_separate_errors"`
	LoggingDetailedLog types.Bool `tfsdk:"logging_detailed_log"`
	LoggingSocketStats types.Bool `tfsdk:"logging_socket_stats"`
	StickinessPattern types.String `tfsdk:"stickiness_pattern"`
	StickinessDataTypes types.Set `tfsdk:"stickiness_data_types"`
	StickinessExpire types.String `tfsdk:"stickiness_expire"`
	StickinessSize types.String `tfsdk:"stickiness_size"`
	StickinessCounter types.String `tfsdk:"stickiness_counter"`
	StickinessCounterKey types.String `tfsdk:"stickiness_counter_key"`
	StickinessLength types.String `tfsdk:"stickiness_length"`
	StickinessConnRatePeriod types.String `tfsdk:"stickiness_conn_rate_period"`
	StickinessSessRatePeriod types.String `tfsdk:"stickiness_sess_rate_period"`
	StickinessHttpReqRatePeriod types.String `tfsdk:"stickiness_http_req_rate_period"`
	StickinessHttpErrRatePeriod types.String `tfsdk:"stickiness_http_err_rate_period"`
	StickinessBytesInRatePeriod types.String `tfsdk:"stickiness_bytes_in_rate_period"`
	StickinessBytesOutRatePeriod types.String `tfsdk:"stickiness_bytes_out_rate_period"`
	StickinessGpcElements types.String `tfsdk:"stickiness_gpc_elements"`
	StickinessGpcRatePeriod types.String `tfsdk:"stickiness_gpc_rate_period"`
	StickinessGptElements types.String `tfsdk:"stickiness_gpt_elements"`
	StickinessHttpFailRatePeriod types.String `tfsdk:"stickiness_http_fail_rate_period"`
	StickinessGlitchRatePeriod types.String `tfsdk:"stickiness_glitch_rate_period"`
	Http2Enabled types.Bool `tfsdk:"http2_enabled"`
	Http2EnabledNontls types.Bool `tfsdk:"http2_enabled_nontls"`
	AdvertisedProtocols types.Set `tfsdk:"advertised_protocols"`
	ForwardFor types.String `tfsdk:"forward_for"`
	PrometheusEnabled types.Bool `tfsdk:"prometheus_enabled"`
	PrometheusPath types.String `tfsdk:"prometheus_path"`
	ConnectionBehaviour types.String `tfsdk:"connection_behaviour"`
	CustomOptions types.String `tfsdk:"custom_options"`
	LinkedActions types.Set `tfsdk:"linked_actions"`
	LinkedErrorfiles types.Set `tfsdk:"linked_errorfiles"`
}

func frontendResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense HAProxy public service.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID of the haproxy frontend.",
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"internal_id": schema.StringAttribute{
				MarkdownDescription: "Internal OPNsense id assigned to this haproxy frontend.",
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the enabled option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Frontend name value.",
				Required: true,
				Validators: []validator.String{stringvalidator.LengthAtLeast(1)},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Frontend description value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"bind": schema.SetAttribute{
				MarkdownDescription: "List of bind values for this frontend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"bind_options": schema.StringAttribute{
				MarkdownDescription: "Frontend bind_options value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Frontend mode option. One of: http, ssl, tcp.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("http", "ssl", "tcp")},
			},
			"default_backend": schema.StringAttribute{
				MarkdownDescription: "Frontend default_backend (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the ssl_enabled option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"ssl_certificates": schema.SetAttribute{
				MarkdownDescription: "List of ssl_certificates values for this frontend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"ssl_default_certificate": schema.StringAttribute{
				MarkdownDescription: "Frontend ssl_default_certificate (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_custom_options": schema.StringAttribute{
				MarkdownDescription: "Frontend ssl_custom_options value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_advanced_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the ssl_advanced_enabled option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"ssl_bind_options": schema.SetAttribute{
				MarkdownDescription: "Selected ssl_bind_options values for this frontend. One or more of: no-sslv3, no-tlsv10, no-tlsv11, no-tlsv12, no-tlsv13, no-tls-tickets, force-sslv3, force-tlsv10, force-tlsv11, force-tlsv12, force-tlsv13, prefer-client-ciphers, strict-sni.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"ssl_min_version": schema.StringAttribute{
				MarkdownDescription: "Frontend ssl_min_version option. One of: , SSLv3, TLSv1.0, TLSv1.1, TLSv1.2, TLSv1.3.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "SSLv3", "TLSv1.0", "TLSv1.1", "TLSv1.2", "TLSv1.3")},
			},
			"ssl_max_version": schema.StringAttribute{
				MarkdownDescription: "Frontend ssl_max_version option. One of: , SSLv3, TLSv1.0, TLSv1.1, TLSv1.2, TLSv1.3.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "SSLv3", "TLSv1.0", "TLSv1.1", "TLSv1.2", "TLSv1.3")},
			},
			"ssl_cipher_list": schema.StringAttribute{
				MarkdownDescription: "Frontend ssl_cipher_list value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_cipher_suites": schema.StringAttribute{
				MarkdownDescription: "Frontend ssl_cipher_suites value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_hsts_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the ssl_hsts_enabled option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"ssl_hsts_include_sub_domains": schema.BoolAttribute{
				MarkdownDescription: "Enable the ssl_hsts_include_sub_domains option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"ssl_hsts_preload": schema.BoolAttribute{
				MarkdownDescription: "Enable the ssl_hsts_preload option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"ssl_hsts_max_age": schema.StringAttribute{
				MarkdownDescription: "Maximum age in seconds for the Strict-Transport-Security header. Required by OPNsense when HSTS is enabled.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString("15768000"),
			},
			"ssl_client_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the ssl_client_auth_enabled option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"ssl_client_auth_verify": schema.StringAttribute{
				MarkdownDescription: "Frontend ssl_client_auth_verify option. One of: , none, optional, required.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "none", "optional", "required")},
			},
			"ssl_client_auth_c_as": schema.SetAttribute{
				MarkdownDescription: "List of ssl_client_auth_c_as values for this frontend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"ssl_client_auth_c_r_ls": schema.SetAttribute{
				MarkdownDescription: "List of ssl_client_auth_c_r_ls values for this frontend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"basic_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the basic_auth_enabled option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"basic_auth_users": schema.SetAttribute{
				MarkdownDescription: "List of basic_auth_users values for this frontend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"basic_auth_groups": schema.SetAttribute{
				MarkdownDescription: "List of basic_auth_groups values for this frontend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"tuning_max_connections": schema.StringAttribute{
				MarkdownDescription: "Frontend tuning_max_connections value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tuning_timeout_client": schema.StringAttribute{
				MarkdownDescription: "Frontend tuning_timeout_client value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tuning_timeout_http_req": schema.StringAttribute{
				MarkdownDescription: "Frontend tuning_timeout_http_req value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tuning_timeout_http_keep_alive": schema.StringAttribute{
				MarkdownDescription: "Frontend tuning_timeout_http_keep_alive value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"linked_cpu_affinity_rules": schema.SetAttribute{
				MarkdownDescription: "List of linked_cpu_affinity_rules values for this frontend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"tuning_shards": schema.StringAttribute{
				MarkdownDescription: "Frontend tuning_shards value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"logging_dont_log_null": schema.BoolAttribute{
				MarkdownDescription: "Enable the logging_dont_log_null option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"logging_dont_log_normal": schema.BoolAttribute{
				MarkdownDescription: "Enable the logging_dont_log_normal option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"logging_log_separate_errors": schema.BoolAttribute{
				MarkdownDescription: "Enable the logging_log_separate_errors option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"logging_detailed_log": schema.BoolAttribute{
				MarkdownDescription: "Enable the logging_detailed_log option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"logging_socket_stats": schema.BoolAttribute{
				MarkdownDescription: "Enable the logging_socket_stats option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"stickiness_pattern": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_pattern option. One of: , binary, integer, ipv4, ipv6, string.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "binary", "integer", "ipv4", "ipv6", "string")},
			},
			"stickiness_data_types": schema.SetAttribute{
				MarkdownDescription: "Selected stickiness_data_types values for this frontend. One or more of: bytes_in_cnt, bytes_in_rate, bytes_out_cnt, bytes_out_rate, conn_cnt, conn_cur, conn_rate, glitch_cnt, glitch_rate, gpc, gpc_rate, gpc0, gpc0_rate, gpc1, gpc1_rate, gpt, gpt0, http_err_cnt, http_err_rate, http_fail_cnt, http_fail_rate, http_req_cnt, http_req_rate, server_id, sess_cnt, sess_rate.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"stickiness_expire": schema.StringAttribute{
				MarkdownDescription: "Expiration time for the stickiness table (e.g. `30m`). Required by OPNsense when stickiness is configured.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString("30m"),
			},
			"stickiness_size": schema.StringAttribute{
				MarkdownDescription: "Size of the stickiness table (e.g. `50k`). Required by OPNsense when stickiness is configured.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString("50k"),
			},
			"stickiness_counter": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_counter value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_counter_key": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_counter_key value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_length": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_length value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_conn_rate_period": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_conn_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_sess_rate_period": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_sess_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_http_req_rate_period": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_http_req_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_http_err_rate_period": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_http_err_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_bytes_in_rate_period": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_bytes_in_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_bytes_out_rate_period": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_bytes_out_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_gpc_elements": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_gpc_elements value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_gpc_rate_period": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_gpc_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_gpt_elements": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_gpt_elements value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_http_fail_rate_period": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_http_fail_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_glitch_rate_period": schema.StringAttribute{
				MarkdownDescription: "Frontend stickiness_glitch_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"http2_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the http2_enabled option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"http2_enabled_nontls": schema.BoolAttribute{
				MarkdownDescription: "Enable the http2_enabled_nontls option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"advertised_protocols": schema.SetAttribute{
				MarkdownDescription: "Selected advertised_protocols values for this frontend. One or more of: h3, h2, http11, http10.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"forward_for": schema.StringAttribute{
				MarkdownDescription: "Frontend forward_for value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"prometheus_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the prometheus_enabled option for this frontend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"prometheus_path": schema.StringAttribute{
				MarkdownDescription: "Frontend prometheus_path value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"connection_behaviour": schema.StringAttribute{
				MarkdownDescription: "Frontend connection_behaviour option. One of: http-keep-alive, httpclose, http-server-close.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("http-keep-alive", "httpclose", "http-server-close")},
			},
			"custom_options": schema.StringAttribute{
				MarkdownDescription: "Frontend custom_options value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"linked_actions": schema.SetAttribute{
				MarkdownDescription: "List of linked_actions values for this frontend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"linked_errorfiles": schema.SetAttribute{
				MarkdownDescription: "List of linked_errorfiles values for this frontend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
		},
	}
}

func frontendDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense HAProxy public service.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the haproxy frontend.",
				Required: true,
				Validators: []validator.String{stringvalidator.LengthAtLeast(1)},
			},
			"internal_id": dschema.StringAttribute{MarkdownDescription: "Frontend internal_id value.", Computed: true},
			"enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the enabled option for this frontend.", Computed: true},
			"name": dschema.StringAttribute{MarkdownDescription: "Frontend name value.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Frontend description value.", Computed: true},
			"bind": dschema.SetAttribute{MarkdownDescription: "List of bind values for this frontend.", ElementType: types.StringType, Computed: true},
			"bind_options": dschema.StringAttribute{MarkdownDescription: "Frontend bind_options value.", Computed: true},
			"mode": dschema.StringAttribute{MarkdownDescription: "Frontend mode option. One of: http, ssl, tcp.", Computed: true},
			"default_backend": dschema.StringAttribute{MarkdownDescription: "Frontend default_backend (UUID reference).", Computed: true},
			"ssl_enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the ssl_enabled option for this frontend.", Computed: true},
			"ssl_certificates": dschema.SetAttribute{MarkdownDescription: "List of ssl_certificates values for this frontend.", ElementType: types.StringType, Computed: true},
			"ssl_default_certificate": dschema.StringAttribute{MarkdownDescription: "Frontend ssl_default_certificate (UUID reference).", Computed: true},
			"ssl_custom_options": dschema.StringAttribute{MarkdownDescription: "Frontend ssl_custom_options value.", Computed: true},
			"ssl_advanced_enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the ssl_advanced_enabled option for this frontend.", Computed: true},
			"ssl_bind_options": dschema.SetAttribute{MarkdownDescription: "Selected ssl_bind_options values for this frontend. One or more of: no-sslv3, no-tlsv10, no-tlsv11, no-tlsv12, no-tlsv13, no-tls-tickets, force-sslv3, force-tlsv10, force-tlsv11, force-tlsv12, force-tlsv13, prefer-client-ciphers, strict-sni.", ElementType: types.StringType, Computed: true},
			"ssl_min_version": dschema.StringAttribute{MarkdownDescription: "Frontend ssl_min_version option. One of: , SSLv3, TLSv1.0, TLSv1.1, TLSv1.2, TLSv1.3.", Computed: true},
			"ssl_max_version": dschema.StringAttribute{MarkdownDescription: "Frontend ssl_max_version option. One of: , SSLv3, TLSv1.0, TLSv1.1, TLSv1.2, TLSv1.3.", Computed: true},
			"ssl_cipher_list": dschema.StringAttribute{MarkdownDescription: "Frontend ssl_cipher_list value.", Computed: true},
			"ssl_cipher_suites": dschema.StringAttribute{MarkdownDescription: "Frontend ssl_cipher_suites value.", Computed: true},
			"ssl_hsts_enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the ssl_hsts_enabled option for this frontend.", Computed: true},
			"ssl_hsts_include_sub_domains": dschema.BoolAttribute{MarkdownDescription: "Enable the ssl_hsts_include_sub_domains option for this frontend.", Computed: true},
			"ssl_hsts_preload": dschema.BoolAttribute{MarkdownDescription: "Enable the ssl_hsts_preload option for this frontend.", Computed: true},
			"ssl_hsts_max_age": dschema.StringAttribute{MarkdownDescription: "Frontend ssl_hsts_max_age value.", Computed: true},
			"ssl_client_auth_enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the ssl_client_auth_enabled option for this frontend.", Computed: true},
			"ssl_client_auth_verify": dschema.StringAttribute{MarkdownDescription: "Frontend ssl_client_auth_verify option. One of: , none, optional, required.", Computed: true},
			"ssl_client_auth_c_as": dschema.SetAttribute{MarkdownDescription: "List of ssl_client_auth_c_as values for this frontend.", ElementType: types.StringType, Computed: true},
			"ssl_client_auth_c_r_ls": dschema.SetAttribute{MarkdownDescription: "List of ssl_client_auth_c_r_ls values for this frontend.", ElementType: types.StringType, Computed: true},
			"basic_auth_enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the basic_auth_enabled option for this frontend.", Computed: true},
			"basic_auth_users": dschema.SetAttribute{MarkdownDescription: "List of basic_auth_users values for this frontend.", ElementType: types.StringType, Computed: true},
			"basic_auth_groups": dschema.SetAttribute{MarkdownDescription: "List of basic_auth_groups values for this frontend.", ElementType: types.StringType, Computed: true},
			"tuning_max_connections": dschema.StringAttribute{MarkdownDescription: "Frontend tuning_max_connections value.", Computed: true},
			"tuning_timeout_client": dschema.StringAttribute{MarkdownDescription: "Frontend tuning_timeout_client value.", Computed: true},
			"tuning_timeout_http_req": dschema.StringAttribute{MarkdownDescription: "Frontend tuning_timeout_http_req value.", Computed: true},
			"tuning_timeout_http_keep_alive": dschema.StringAttribute{MarkdownDescription: "Frontend tuning_timeout_http_keep_alive value.", Computed: true},
			"linked_cpu_affinity_rules": dschema.SetAttribute{MarkdownDescription: "List of linked_cpu_affinity_rules values for this frontend.", ElementType: types.StringType, Computed: true},
			"tuning_shards": dschema.StringAttribute{MarkdownDescription: "Frontend tuning_shards value.", Computed: true},
			"logging_dont_log_null": dschema.BoolAttribute{MarkdownDescription: "Enable the logging_dont_log_null option for this frontend.", Computed: true},
			"logging_dont_log_normal": dschema.BoolAttribute{MarkdownDescription: "Enable the logging_dont_log_normal option for this frontend.", Computed: true},
			"logging_log_separate_errors": dschema.BoolAttribute{MarkdownDescription: "Enable the logging_log_separate_errors option for this frontend.", Computed: true},
			"logging_detailed_log": dschema.BoolAttribute{MarkdownDescription: "Enable the logging_detailed_log option for this frontend.", Computed: true},
			"logging_socket_stats": dschema.BoolAttribute{MarkdownDescription: "Enable the logging_socket_stats option for this frontend.", Computed: true},
			"stickiness_pattern": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_pattern option. One of: , binary, integer, ipv4, ipv6, string.", Computed: true},
			"stickiness_data_types": dschema.SetAttribute{MarkdownDescription: "Selected stickiness_data_types values for this frontend. One or more of: bytes_in_cnt, bytes_in_rate, bytes_out_cnt, bytes_out_rate, conn_cnt, conn_cur, conn_rate, glitch_cnt, glitch_rate, gpc, gpc_rate, gpc0, gpc0_rate, gpc1, gpc1_rate, gpt, gpt0, http_err_cnt, http_err_rate, http_fail_cnt, http_fail_rate, http_req_cnt, http_req_rate, server_id, sess_cnt, sess_rate.", ElementType: types.StringType, Computed: true},
			"stickiness_expire": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_expire value.", Computed: true},
			"stickiness_size": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_size value.", Computed: true},
			"stickiness_counter": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_counter value.", Computed: true},
			"stickiness_counter_key": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_counter_key value.", Computed: true},
			"stickiness_length": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_length value.", Computed: true},
			"stickiness_conn_rate_period": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_conn_rate_period value.", Computed: true},
			"stickiness_sess_rate_period": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_sess_rate_period value.", Computed: true},
			"stickiness_http_req_rate_period": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_http_req_rate_period value.", Computed: true},
			"stickiness_http_err_rate_period": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_http_err_rate_period value.", Computed: true},
			"stickiness_bytes_in_rate_period": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_bytes_in_rate_period value.", Computed: true},
			"stickiness_bytes_out_rate_period": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_bytes_out_rate_period value.", Computed: true},
			"stickiness_gpc_elements": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_gpc_elements value.", Computed: true},
			"stickiness_gpc_rate_period": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_gpc_rate_period value.", Computed: true},
			"stickiness_gpt_elements": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_gpt_elements value.", Computed: true},
			"stickiness_http_fail_rate_period": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_http_fail_rate_period value.", Computed: true},
			"stickiness_glitch_rate_period": dschema.StringAttribute{MarkdownDescription: "Frontend stickiness_glitch_rate_period value.", Computed: true},
			"http2_enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the http2_enabled option for this frontend.", Computed: true},
			"http2_enabled_nontls": dschema.BoolAttribute{MarkdownDescription: "Enable the http2_enabled_nontls option for this frontend.", Computed: true},
			"advertised_protocols": dschema.SetAttribute{MarkdownDescription: "Selected advertised_protocols values for this frontend. One or more of: h3, h2, http11, http10.", ElementType: types.StringType, Computed: true},
			"forward_for": dschema.StringAttribute{MarkdownDescription: "Frontend forward_for value.", Computed: true},
			"prometheus_enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the prometheus_enabled option for this frontend.", Computed: true},
			"prometheus_path": dschema.StringAttribute{MarkdownDescription: "Frontend prometheus_path value.", Computed: true},
			"connection_behaviour": dschema.StringAttribute{MarkdownDescription: "Frontend connection_behaviour option. One of: http-keep-alive, httpclose, http-server-close.", Computed: true},
			"custom_options": dschema.StringAttribute{MarkdownDescription: "Frontend custom_options value.", Computed: true},
			"linked_actions": dschema.SetAttribute{MarkdownDescription: "List of linked_actions values for this frontend.", ElementType: types.StringType, Computed: true},
			"linked_errorfiles": dschema.SetAttribute{MarkdownDescription: "List of linked_errorfiles values for this frontend.", ElementType: types.StringType, Computed: true},
		},
	}
}
