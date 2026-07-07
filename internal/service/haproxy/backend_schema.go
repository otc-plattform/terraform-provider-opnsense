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

type backendResourceModel struct {
	Id types.String `tfsdk:"id"`
	InternalId types.String `tfsdk:"internal_id"`
	Enabled types.Bool `tfsdk:"enabled"`
	Name types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Mode types.String `tfsdk:"mode"`
	Algorithm types.String `tfsdk:"algorithm"`
	RandomDraws types.String `tfsdk:"random_draws"`
	ProxyProtocol types.String `tfsdk:"proxy_protocol"`
	LinkedServers types.Set `tfsdk:"linked_servers"`
	LinkedFcgi types.String `tfsdk:"linked_fcgi"`
	LinkedResolver types.String `tfsdk:"linked_resolver"`
	ResolverOpts types.Set `tfsdk:"resolver_opts"`
	ResolvePrefer types.String `tfsdk:"resolve_prefer"`
	Source types.String `tfsdk:"source"`
	HealthCheckEnabled types.Bool `tfsdk:"health_check_enabled"`
	HealthCheck types.String `tfsdk:"health_check"`
	HealthCheckLogStatus types.String `tfsdk:"health_check_log_status"`
	CheckInterval types.String `tfsdk:"check_interval"`
	CheckDownInterval types.String `tfsdk:"check_down_interval"`
	HealthCheckFall types.String `tfsdk:"health_check_fall"`
	HealthCheckRise types.String `tfsdk:"health_check_rise"`
	LinkedMailer types.String `tfsdk:"linked_mailer"`
	HealthCheckProxyProto types.String `tfsdk:"health_check_proxy_proto"`
	Http2Enabled types.Bool `tfsdk:"http2_enabled"`
	Http2EnabledNontls types.Bool `tfsdk:"http2_enabled_nontls"`
	BaAdvertisedProtocols types.Set `tfsdk:"ba_advertised_protocols"`
	ForwardFor types.String `tfsdk:"forward_for"`
	ForwardedHeader types.String `tfsdk:"forwarded_header"`
	ForwardedHeaderParameters types.Set `tfsdk:"forwarded_header_parameters"`
	Persistence types.String `tfsdk:"persistence"`
	PersistenceCookiemode types.String `tfsdk:"persistence_cookiemode"`
	PersistenceCookiename types.String `tfsdk:"persistence_cookiename"`
	PersistenceStripquotes types.Bool `tfsdk:"persistence_stripquotes"`
	StickinessPattern types.String `tfsdk:"stickiness_pattern"`
	StickinessDataTypes types.Set `tfsdk:"stickiness_data_types"`
	StickinessExpire types.String `tfsdk:"stickiness_expire"`
	StickinessSize types.String `tfsdk:"stickiness_size"`
	StickinessCookiename types.String `tfsdk:"stickiness_cookiename"`
	StickinessCookielength types.String `tfsdk:"stickiness_cookielength"`
	StickinessLength types.String `tfsdk:"stickiness_length"`
	StickinessConnRatePeriod types.String `tfsdk:"stickiness_conn_rate_period"`
	StickinessSessRatePeriod types.String `tfsdk:"stickiness_sess_rate_period"`
	StickinessHttpReqRatePeriod types.String `tfsdk:"stickiness_http_req_rate_period"`
	StickinessHttpErrRatePeriod types.String `tfsdk:"stickiness_http_err_rate_period"`
	StickinessBytesInRatePeriod types.String `tfsdk:"stickiness_bytes_in_rate_period"`
	StickinessBytesOutRatePeriod types.String `tfsdk:"stickiness_bytes_out_rate_period"`
	StickinessGpcElements types.String `tfsdk:"stickiness_gpc_elements"`
	StickinessGptElements types.String `tfsdk:"stickiness_gpt_elements"`
	StickinessGpcRatePeriod types.String `tfsdk:"stickiness_gpc_rate_period"`
	StickinessHttpFailRatePeriod types.String `tfsdk:"stickiness_http_fail_rate_period"`
	StickinessGlitchRatePeriod types.String `tfsdk:"stickiness_glitch_rate_period"`
	BasicAuthEnabled types.Bool `tfsdk:"basic_auth_enabled"`
	BasicAuthUsers types.Set `tfsdk:"basic_auth_users"`
	BasicAuthGroups types.Set `tfsdk:"basic_auth_groups"`
	TuningTimeoutConnect types.String `tfsdk:"tuning_timeout_connect"`
	TuningTimeoutCheck types.String `tfsdk:"tuning_timeout_check"`
	TuningTimeoutServer types.String `tfsdk:"tuning_timeout_server"`
	TuningRetries types.String `tfsdk:"tuning_retries"`
	CustomOptions types.String `tfsdk:"custom_options"`
	TuningDefaultserver types.String `tfsdk:"tuning_defaultserver"`
	TuningNoport types.Bool `tfsdk:"tuning_noport"`
	TuningHttpreuse types.String `tfsdk:"tuning_httpreuse"`
	TuningCaching types.Bool `tfsdk:"tuning_caching"`
	LinkedActions types.Set `tfsdk:"linked_actions"`
	LinkedErrorfiles types.Set `tfsdk:"linked_errorfiles"`
}

func backendResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense HAProxy backend pool.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID of the haproxy backend.",
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"internal_id": schema.StringAttribute{
				MarkdownDescription: "Internal OPNsense id assigned to this haproxy backend.",
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the enabled option for this backend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Backend name value.",
				Required: true,
				Validators: []validator.String{stringvalidator.LengthAtLeast(1)},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Backend description value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Backend mode option. One of: http, tcp.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("http", "tcp")},
			},
			"algorithm": schema.StringAttribute{
				MarkdownDescription: "Backend algorithm option. One of: source, roundrobin, static-rr, leastconn, uri, random.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("source", "roundrobin", "static-rr", "leastconn", "uri", "random")},
			},
			"random_draws": schema.StringAttribute{
				MarkdownDescription: "Backend random_draws value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"proxy_protocol": schema.StringAttribute{
				MarkdownDescription: "Backend proxy_protocol option. One of: , v1, v2.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "v1", "v2")},
			},
			"linked_servers": schema.SetAttribute{
				MarkdownDescription: "List of linked_servers values for this backend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"linked_fcgi": schema.StringAttribute{
				MarkdownDescription: "Backend linked_fcgi (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"linked_resolver": schema.StringAttribute{
				MarkdownDescription: "Backend linked_resolver (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"resolver_opts": schema.SetAttribute{
				MarkdownDescription: "Selected resolver_opts values for this backend. One or more of: allow-dup-ip, ignore-weight, prevent-dup-ip.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"resolve_prefer": schema.StringAttribute{
				MarkdownDescription: "Backend resolve_prefer option. One of: , ipv4, ipv6.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "ipv4", "ipv6")},
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "Backend source value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"health_check_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the health_check_enabled option for this backend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"health_check": schema.StringAttribute{
				MarkdownDescription: "Backend health_check (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"health_check_log_status": schema.StringAttribute{
				MarkdownDescription: "Backend health_check_log_status value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"check_interval": schema.StringAttribute{
				MarkdownDescription: "Backend check_interval value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"check_down_interval": schema.StringAttribute{
				MarkdownDescription: "Backend check_down_interval value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"health_check_fall": schema.StringAttribute{
				MarkdownDescription: "Backend health_check_fall value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"health_check_rise": schema.StringAttribute{
				MarkdownDescription: "Backend health_check_rise value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"linked_mailer": schema.StringAttribute{
				MarkdownDescription: "Backend linked_mailer (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"health_check_proxy_proto": schema.StringAttribute{
				MarkdownDescription: "Backend health_check_proxy_proto option. One of: , backend, enable, disable.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "backend", "enable", "disable")},
			},
			"http2_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the http2_enabled option for this backend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"http2_enabled_nontls": schema.BoolAttribute{
				MarkdownDescription: "Enable the http2_enabled_nontls option for this backend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"ba_advertised_protocols": schema.SetAttribute{
				MarkdownDescription: "Selected ba_advertised_protocols values for this backend. One or more of: h2, http11, http10.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"forward_for": schema.StringAttribute{
				MarkdownDescription: "Backend forward_for value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"forwarded_header": schema.StringAttribute{
				MarkdownDescription: "Backend forwarded_header value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"forwarded_header_parameters": schema.SetAttribute{
				MarkdownDescription: "Selected forwarded_header_parameters values for this backend. One or more of: proto, host, by, by_port, for, for_port.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"persistence": schema.StringAttribute{
				MarkdownDescription: "Backend persistence option. One of: , sticktable, cookie.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "sticktable", "cookie")},
			},
			"persistence_cookiemode": schema.StringAttribute{
				MarkdownDescription: "Backend persistence_cookiemode option. One of: piggyback, new.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("piggyback", "new")},
			},
			"persistence_cookiename": schema.StringAttribute{
				MarkdownDescription: "Backend persistence_cookiename value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"persistence_stripquotes": schema.BoolAttribute{
				MarkdownDescription: "Enable the persistence_stripquotes option for this backend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"stickiness_pattern": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_pattern option. One of: , binary, cookievalue, integer, rdpcookie, sourceipv4, sourceipv6, string.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "binary", "cookievalue", "integer", "rdpcookie", "sourceipv4", "sourceipv6", "string")},
			},
			"stickiness_data_types": schema.SetAttribute{
				MarkdownDescription: "Selected stickiness_data_types values for this backend. One or more of: bytes_in_cnt, bytes_in_rate, bytes_out_cnt, bytes_out_rate, conn_cnt, conn_cur, conn_rate, glitch_cnt, glitch_rate, gpc, gpc_rate, gpc0, gpc0_rate, gpc1, gpc1_rate, gpt, gpt0, http_err_cnt, http_err_rate, http_fail_cnt, http_fail_rate, http_req_cnt, http_req_rate, server_id, sess_cnt, sess_rate.",
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
			"stickiness_cookiename": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_cookiename value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_cookielength": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_cookielength value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_length": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_length value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_conn_rate_period": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_conn_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_sess_rate_period": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_sess_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_http_req_rate_period": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_http_req_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_http_err_rate_period": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_http_err_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_bytes_in_rate_period": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_bytes_in_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_bytes_out_rate_period": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_bytes_out_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_gpc_elements": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_gpc_elements value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_gpt_elements": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_gpt_elements value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_gpc_rate_period": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_gpc_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_http_fail_rate_period": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_http_fail_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"stickiness_glitch_rate_period": schema.StringAttribute{
				MarkdownDescription: "Backend stickiness_glitch_rate_period value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"basic_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the basic_auth_enabled option for this backend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"basic_auth_users": schema.SetAttribute{
				MarkdownDescription: "List of basic_auth_users values for this backend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"basic_auth_groups": schema.SetAttribute{
				MarkdownDescription: "List of basic_auth_groups values for this backend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"tuning_timeout_connect": schema.StringAttribute{
				MarkdownDescription: "Backend tuning_timeout_connect value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tuning_timeout_check": schema.StringAttribute{
				MarkdownDescription: "Backend tuning_timeout_check value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tuning_timeout_server": schema.StringAttribute{
				MarkdownDescription: "Backend tuning_timeout_server value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tuning_retries": schema.StringAttribute{
				MarkdownDescription: "Backend tuning_retries value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"custom_options": schema.StringAttribute{
				MarkdownDescription: "Backend custom_options value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tuning_defaultserver": schema.StringAttribute{
				MarkdownDescription: "Backend tuning_defaultserver value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"tuning_noport": schema.BoolAttribute{
				MarkdownDescription: "Enable the tuning_noport option for this backend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"tuning_httpreuse": schema.StringAttribute{
				MarkdownDescription: "Backend tuning_httpreuse option. One of: , never, safe, aggressive, always.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "never", "safe", "aggressive", "always")},
			},
			"tuning_caching": schema.BoolAttribute{
				MarkdownDescription: "Enable the tuning_caching option for this backend.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"linked_actions": schema.SetAttribute{
				MarkdownDescription: "List of linked_actions values for this backend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"linked_errorfiles": schema.SetAttribute{
				MarkdownDescription: "List of linked_errorfiles values for this backend.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
		},
	}
}

func backendDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense HAProxy backend pool.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the haproxy backend.",
				Required: true,
				Validators: []validator.String{stringvalidator.LengthAtLeast(1)},
			},
			"internal_id": dschema.StringAttribute{MarkdownDescription: "Backend internal_id value.", Computed: true},
			"enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the enabled option for this backend.", Computed: true},
			"name": dschema.StringAttribute{MarkdownDescription: "Backend name value.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Backend description value.", Computed: true},
			"mode": dschema.StringAttribute{MarkdownDescription: "Backend mode option. One of: http, tcp.", Computed: true},
			"algorithm": dschema.StringAttribute{MarkdownDescription: "Backend algorithm option. One of: source, roundrobin, static-rr, leastconn, uri, random.", Computed: true},
			"random_draws": dschema.StringAttribute{MarkdownDescription: "Backend random_draws value.", Computed: true},
			"proxy_protocol": dschema.StringAttribute{MarkdownDescription: "Backend proxy_protocol option. One of: , v1, v2.", Computed: true},
			"linked_servers": dschema.SetAttribute{MarkdownDescription: "List of linked_servers values for this backend.", ElementType: types.StringType, Computed: true},
			"linked_fcgi": dschema.StringAttribute{MarkdownDescription: "Backend linked_fcgi (UUID reference).", Computed: true},
			"linked_resolver": dschema.StringAttribute{MarkdownDescription: "Backend linked_resolver (UUID reference).", Computed: true},
			"resolver_opts": dschema.SetAttribute{MarkdownDescription: "Selected resolver_opts values for this backend. One or more of: allow-dup-ip, ignore-weight, prevent-dup-ip.", ElementType: types.StringType, Computed: true},
			"resolve_prefer": dschema.StringAttribute{MarkdownDescription: "Backend resolve_prefer option. One of: , ipv4, ipv6.", Computed: true},
			"source": dschema.StringAttribute{MarkdownDescription: "Backend source value.", Computed: true},
			"health_check_enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the health_check_enabled option for this backend.", Computed: true},
			"health_check": dschema.StringAttribute{MarkdownDescription: "Backend health_check (UUID reference).", Computed: true},
			"health_check_log_status": dschema.StringAttribute{MarkdownDescription: "Backend health_check_log_status value.", Computed: true},
			"check_interval": dschema.StringAttribute{MarkdownDescription: "Backend check_interval value.", Computed: true},
			"check_down_interval": dschema.StringAttribute{MarkdownDescription: "Backend check_down_interval value.", Computed: true},
			"health_check_fall": dschema.StringAttribute{MarkdownDescription: "Backend health_check_fall value.", Computed: true},
			"health_check_rise": dschema.StringAttribute{MarkdownDescription: "Backend health_check_rise value.", Computed: true},
			"linked_mailer": dschema.StringAttribute{MarkdownDescription: "Backend linked_mailer (UUID reference).", Computed: true},
			"health_check_proxy_proto": dschema.StringAttribute{MarkdownDescription: "Backend health_check_proxy_proto option. One of: , backend, enable, disable.", Computed: true},
			"http2_enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the http2_enabled option for this backend.", Computed: true},
			"http2_enabled_nontls": dschema.BoolAttribute{MarkdownDescription: "Enable the http2_enabled_nontls option for this backend.", Computed: true},
			"ba_advertised_protocols": dschema.SetAttribute{MarkdownDescription: "Selected ba_advertised_protocols values for this backend. One or more of: h2, http11, http10.", ElementType: types.StringType, Computed: true},
			"forward_for": dschema.StringAttribute{MarkdownDescription: "Backend forward_for value.", Computed: true},
			"forwarded_header": dschema.StringAttribute{MarkdownDescription: "Backend forwarded_header value.", Computed: true},
			"forwarded_header_parameters": dschema.SetAttribute{MarkdownDescription: "Selected forwarded_header_parameters values for this backend. One or more of: proto, host, by, by_port, for, for_port.", ElementType: types.StringType, Computed: true},
			"persistence": dschema.StringAttribute{MarkdownDescription: "Backend persistence option. One of: , sticktable, cookie.", Computed: true},
			"persistence_cookiemode": dschema.StringAttribute{MarkdownDescription: "Backend persistence_cookiemode option. One of: piggyback, new.", Computed: true},
			"persistence_cookiename": dschema.StringAttribute{MarkdownDescription: "Backend persistence_cookiename value.", Computed: true},
			"persistence_stripquotes": dschema.BoolAttribute{MarkdownDescription: "Enable the persistence_stripquotes option for this backend.", Computed: true},
			"stickiness_pattern": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_pattern option. One of: , binary, cookievalue, integer, rdpcookie, sourceipv4, sourceipv6, string.", Computed: true},
			"stickiness_data_types": dschema.SetAttribute{MarkdownDescription: "Selected stickiness_data_types values for this backend. One or more of: bytes_in_cnt, bytes_in_rate, bytes_out_cnt, bytes_out_rate, conn_cnt, conn_cur, conn_rate, glitch_cnt, glitch_rate, gpc, gpc_rate, gpc0, gpc0_rate, gpc1, gpc1_rate, gpt, gpt0, http_err_cnt, http_err_rate, http_fail_cnt, http_fail_rate, http_req_cnt, http_req_rate, server_id, sess_cnt, sess_rate.", ElementType: types.StringType, Computed: true},
			"stickiness_expire": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_expire value.", Computed: true},
			"stickiness_size": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_size value.", Computed: true},
			"stickiness_cookiename": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_cookiename value.", Computed: true},
			"stickiness_cookielength": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_cookielength value.", Computed: true},
			"stickiness_length": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_length value.", Computed: true},
			"stickiness_conn_rate_period": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_conn_rate_period value.", Computed: true},
			"stickiness_sess_rate_period": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_sess_rate_period value.", Computed: true},
			"stickiness_http_req_rate_period": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_http_req_rate_period value.", Computed: true},
			"stickiness_http_err_rate_period": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_http_err_rate_period value.", Computed: true},
			"stickiness_bytes_in_rate_period": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_bytes_in_rate_period value.", Computed: true},
			"stickiness_bytes_out_rate_period": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_bytes_out_rate_period value.", Computed: true},
			"stickiness_gpc_elements": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_gpc_elements value.", Computed: true},
			"stickiness_gpt_elements": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_gpt_elements value.", Computed: true},
			"stickiness_gpc_rate_period": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_gpc_rate_period value.", Computed: true},
			"stickiness_http_fail_rate_period": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_http_fail_rate_period value.", Computed: true},
			"stickiness_glitch_rate_period": dschema.StringAttribute{MarkdownDescription: "Backend stickiness_glitch_rate_period value.", Computed: true},
			"basic_auth_enabled": dschema.BoolAttribute{MarkdownDescription: "Enable the basic_auth_enabled option for this backend.", Computed: true},
			"basic_auth_users": dschema.SetAttribute{MarkdownDescription: "List of basic_auth_users values for this backend.", ElementType: types.StringType, Computed: true},
			"basic_auth_groups": dschema.SetAttribute{MarkdownDescription: "List of basic_auth_groups values for this backend.", ElementType: types.StringType, Computed: true},
			"tuning_timeout_connect": dschema.StringAttribute{MarkdownDescription: "Backend tuning_timeout_connect value.", Computed: true},
			"tuning_timeout_check": dschema.StringAttribute{MarkdownDescription: "Backend tuning_timeout_check value.", Computed: true},
			"tuning_timeout_server": dschema.StringAttribute{MarkdownDescription: "Backend tuning_timeout_server value.", Computed: true},
			"tuning_retries": dschema.StringAttribute{MarkdownDescription: "Backend tuning_retries value.", Computed: true},
			"custom_options": dschema.StringAttribute{MarkdownDescription: "Backend custom_options value.", Computed: true},
			"tuning_defaultserver": dschema.StringAttribute{MarkdownDescription: "Backend tuning_defaultserver value.", Computed: true},
			"tuning_noport": dschema.BoolAttribute{MarkdownDescription: "Enable the tuning_noport option for this backend.", Computed: true},
			"tuning_httpreuse": dschema.StringAttribute{MarkdownDescription: "Backend tuning_httpreuse option. One of: , never, safe, aggressive, always.", Computed: true},
			"tuning_caching": dschema.BoolAttribute{MarkdownDescription: "Enable the tuning_caching option for this backend.", Computed: true},
			"linked_actions": dschema.SetAttribute{MarkdownDescription: "List of linked_actions values for this backend.", ElementType: types.StringType, Computed: true},
			"linked_errorfiles": dschema.SetAttribute{MarkdownDescription: "List of linked_errorfiles values for this backend.", ElementType: types.StringType, Computed: true},
		},
	}
}
