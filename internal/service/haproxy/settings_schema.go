package haproxy

import (
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	defaultSSLCipherList   = "ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256"
	defaultSSLCipherSuites = "TLS_AES_128_GCM_SHA256:TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256"
)

type haproxySettingsResourceModel struct {
	Enabled                         types.Bool   `tfsdk:"enabled"`
	GracefulStop                    types.Bool   `tfsdk:"graceful_stop"`
	HardStopAfter                   types.String `tfsdk:"hard_stop_after"`
	CloseSpreadTime                 types.String `tfsdk:"close_spread_time"`
	SeamlessReload                  types.Bool   `tfsdk:"seamless_reload"`
	StoreOCSP                       types.Bool   `tfsdk:"store_ocsp"`
	ShowIntro                       types.Bool   `tfsdk:"show_intro"`
	TuningRoot                      types.Bool   `tfsdk:"tuning_root"`
	TuningMaxConnections            types.Int64  `tfsdk:"tuning_max_connections"`
	TuningNbthread                  types.Int64  `tfsdk:"tuning_nbthread"`
	TuningResolversPrefer           types.String `tfsdk:"tuning_resolvers_prefer"`
	TuningSSLServerVerify           types.String `tfsdk:"tuning_ssl_server_verify"`
	TuningMaxDHSize                 types.Int64  `tfsdk:"tuning_max_dh_size"`
	TuningBufferSize                types.Int64  `tfsdk:"tuning_buffer_size"`
	TuningSpreadChecks              types.Int64  `tfsdk:"tuning_spread_checks"`
	TuningBogusProxyEnabled         types.Bool   `tfsdk:"tuning_bogus_proxy_enabled"`
	TuningLuaMaxMem                 types.Int64  `tfsdk:"tuning_lua_max_mem"`
	TuningCustomOptions             types.String `tfsdk:"tuning_custom_options"`
	TuningOCSPUpdateEnabled         types.Bool   `tfsdk:"tuning_ocsp_update_enabled"`
	TuningOCSPUpdateMinDelay        types.Int64  `tfsdk:"tuning_ocsp_update_min_delay"`
	TuningOCSPUpdateMaxDelay        types.Int64  `tfsdk:"tuning_ocsp_update_max_delay"`
	TuningSSLDefaultsEnabled        types.Bool   `tfsdk:"tuning_ssl_defaults_enabled"`
	TuningSSLBindOptions            types.Set    `tfsdk:"tuning_ssl_bind_options"`
	TuningSSLMinVersion             types.String `tfsdk:"tuning_ssl_min_version"`
	TuningSSLMaxVersion             types.String `tfsdk:"tuning_ssl_max_version"`
	TuningSSLCipherList             types.String `tfsdk:"tuning_ssl_cipher_list"`
	TuningSSLCipherSuites           types.String `tfsdk:"tuning_ssl_cipher_suites"`
	TuningH2InitialWindowSize       types.Int64  `tfsdk:"tuning_h2_initial_window_size"`
	TuningH2InitialWindowSizeOut    types.Int64  `tfsdk:"tuning_h2_initial_window_size_outgoing"`
	TuningH2InitialWindowSizeIn     types.Int64  `tfsdk:"tuning_h2_initial_window_size_incoming"`
	TuningH2MaxConcurrentStreams    types.Int64  `tfsdk:"tuning_h2_max_concurrent_streams"`
	TuningH2MaxConcurrentStreamsOut types.Int64  `tfsdk:"tuning_h2_max_concurrent_streams_outgoing"`
	TuningH2MaxConcurrentStreamsIn  types.Int64  `tfsdk:"tuning_h2_max_concurrent_streams_incoming"`
	DefaultsMaxConnections          types.Int64  `tfsdk:"defaults_max_connections"`
	DefaultsMaxConnectionsServers   types.Int64  `tfsdk:"defaults_max_connections_servers"`
	DefaultsTimeoutClient           types.String `tfsdk:"defaults_timeout_client"`
	DefaultsTimeoutConnect          types.String `tfsdk:"defaults_timeout_connect"`
	DefaultsTimeoutCheck            types.String `tfsdk:"defaults_timeout_check"`
	DefaultsTimeoutServer           types.String `tfsdk:"defaults_timeout_server"`
	DefaultsRetries                 types.Int64  `tfsdk:"defaults_retries"`
	DefaultsRedispatch              types.String `tfsdk:"defaults_redispatch"`
	DefaultsInitAddr                types.Set    `tfsdk:"defaults_init_addr"`
	DefaultsCustomOptions           types.String `tfsdk:"defaults_custom_options"`
	LoggingHost                     types.String `tfsdk:"logging_host"`
	LoggingFacility                 types.String `tfsdk:"logging_facility"`
	LoggingLevel                    types.String `tfsdk:"logging_level"`
	LoggingLength                   types.Int64  `tfsdk:"logging_length"`
	StatsEnabled                    types.Bool   `tfsdk:"stats_enabled"`
	StatsPort                       types.Int64  `tfsdk:"stats_port"`
	StatsRemoteEnabled              types.Bool   `tfsdk:"stats_remote_enabled"`
	StatsRemoteBind                 types.String `tfsdk:"stats_remote_bind"`
	StatsAuthEnabled                types.Bool   `tfsdk:"stats_auth_enabled"`
	StatsAllowedUsers               types.Set    `tfsdk:"stats_allowed_users"`
	StatsAllowedGroups              types.Set    `tfsdk:"stats_allowed_groups"`
	StatsCustomOptions              types.String `tfsdk:"stats_custom_options"`
	StatsPrometheusEnabled          types.Bool   `tfsdk:"stats_prometheus_enabled"`
	StatsPrometheusBind             types.String `tfsdk:"stats_prometheus_bind"`
	StatsPrometheusPath             types.String `tfsdk:"stats_prometheus_path"`
	CacheEnabled                    types.Bool   `tfsdk:"cache_enabled"`
	CacheTotalMaxSize               types.Int64  `tfsdk:"cache_total_max_size"`
	CacheMaxAge                     types.Int64  `tfsdk:"cache_max_age"`
	CacheMaxObjectSize              types.Int64  `tfsdk:"cache_max_object_size"`
	CacheProcessVary                types.Bool   `tfsdk:"cache_process_vary"`
	CacheMaxSecondaryEntries        types.Int64  `tfsdk:"cache_max_secondary_entries"`
	PeersEnabled                    types.Bool   `tfsdk:"peers_enabled"`
	PeersName1                      types.String `tfsdk:"peers_name1"`
	PeersListen1                    types.String `tfsdk:"peers_listen1"`
	PeersPort1                      types.Int64  `tfsdk:"peers_port1"`
	PeersName2                      types.String `tfsdk:"peers_name2"`
	PeersListen2                    types.String `tfsdk:"peers_listen2"`
	PeersPort2                      types.Int64  `tfsdk:"peers_port2"`
}

func haproxySettingsResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage OPNsense HAProxy plugin service and global parameter settings.",
		Attributes: map[string]schema.Attribute{
			"enabled":                      boolAttr("Enable the HAProxy service.", false),
			"graceful_stop":                boolAttr("Use graceful stop for HAProxy.", false),
			"hard_stop_after":              stringAttr("Maximum time before HAProxy is hard-stopped.", "60s"),
			"close_spread_time":            stringAttr("Time window used to spread connection closes.", ""),
			"seamless_reload":              boolAttr("Use seamless reloads.", true),
			"store_ocsp":                   boolAttr("Store OCSP responses.", false),
			"show_intro":                   boolAttr("Show the introductory/help text in the UI.", true),
			"tuning_root":                  boolAttr("Run HAProxy with root privileges.", false),
			"tuning_max_connections":       optionalInt64Attr("Maximum concurrent connections."),
			"tuning_nbthread":              int64Attr("Number of HAProxy threads.", 1, int64validator.AtLeast(1)),
			"tuning_resolvers_prefer":      stringEnumAttr("Preferred resolver IP protocol.", "ipv4", "", "ipv4", "ipv6"),
			"tuning_ssl_server_verify":     stringEnumAttr("Default server certificate verification behavior.", "ignore", "ignore", "required", "none"),
			"tuning_max_dh_size":           int64Attr("Maximum Diffie-Hellman parameter size.", 2048, int64validator.AtLeast(0)),
			"tuning_buffer_size":           int64Attr("HAProxy buffer size.", 16384, int64validator.AtLeast(0)),
			"tuning_spread_checks":         int64Attr("Health check spreading percentage.", 2, int64validator.AtLeast(0)),
			"tuning_bogus_proxy_enabled":   boolAttr("Enable bogus PROXY protocol support.", false),
			"tuning_lua_max_mem":           int64Attr("Lua maximum memory in megabytes.", 0, int64validator.AtLeast(0)),
			"tuning_custom_options":        stringAttr("Custom HAProxy global options.", ""),
			"tuning_ocsp_update_enabled":   boolAttr("Enable automatic OCSP updates.", false),
			"tuning_ocsp_update_min_delay": int64Attr("Minimum OCSP update delay in seconds.", 300, int64validator.AtLeast(0)),
			"tuning_ocsp_update_max_delay": int64Attr("Maximum OCSP update delay in seconds.", 3600, int64validator.AtLeast(0)),
			"tuning_ssl_defaults_enabled":  boolAttr("Enable global SSL default settings.", false),
			"tuning_ssl_bind_options": schema.SetAttribute{
				MarkdownDescription: "Global SSL bind options.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.StringSliceToSet([]string{"prefer-client-ciphers"})),
			},
			"tuning_ssl_min_version":                    stringEnumAttr("Minimum SSL/TLS version.", "TLSv1.2", "", "SSLv3", "TLSv1.0", "TLSv1.1", "TLSv1.2", "TLSv1.3"),
			"tuning_ssl_max_version":                    stringEnumAttr("Maximum SSL/TLS version.", "", "", "SSLv3", "TLSv1.0", "TLSv1.1", "TLSv1.2", "TLSv1.3"),
			"tuning_ssl_cipher_list":                    stringAttr("TLS 1.2 and earlier cipher list.", defaultSSLCipherList),
			"tuning_ssl_cipher_suites":                  stringAttr("TLS 1.3 cipher suites.", defaultSSLCipherSuites),
			"tuning_h2_initial_window_size":             optionalInt64Attr("HTTP/2 initial window size."),
			"tuning_h2_initial_window_size_outgoing":    optionalInt64Attr("HTTP/2 outgoing initial window size."),
			"tuning_h2_initial_window_size_incoming":    optionalInt64Attr("HTTP/2 incoming initial window size."),
			"tuning_h2_max_concurrent_streams":          optionalInt64Attr("HTTP/2 maximum concurrent streams."),
			"tuning_h2_max_concurrent_streams_outgoing": optionalInt64Attr("HTTP/2 outgoing maximum concurrent streams."),
			"tuning_h2_max_concurrent_streams_incoming": optionalInt64Attr("HTTP/2 incoming maximum concurrent streams."),
			"defaults_max_connections":                  optionalInt64Attr("Default maximum concurrent connections."),
			"defaults_max_connections_servers":          optionalInt64Attr("Default maximum concurrent server connections."),
			"defaults_timeout_client":                   stringAttr("Default client timeout.", "30s"),
			"defaults_timeout_connect":                  stringAttr("Default connect timeout.", "30s"),
			"defaults_timeout_check":                    stringAttr("Default health check timeout.", ""),
			"defaults_timeout_server":                   stringAttr("Default server timeout.", "30s"),
			"defaults_retries":                          int64Attr("Default retry count.", 3, int64validator.AtLeast(0)),
			"defaults_redispatch":                       stringEnumAttr("Default redispatch behavior.", "x-1", "", "x3", "x2", "x1", "x0", "x-1", "x-2", "x-3"),
			"defaults_init_addr": schema.SetAttribute{
				MarkdownDescription: "Default server init-addr methods.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.StringSliceToSet([]string{"last", "libc"})),
			},
			"defaults_custom_options": stringAttr("Custom HAProxy default options.", ""),
			"logging_host":            stringAttr("Syslog host for HAProxy logs.", "127.0.0.1"),
			"logging_facility":        stringEnumAttr("Syslog facility for HAProxy logs.", "local0", "alert", "audit", "auth2", "auth", "cron2", "cron", "daemon", "ftp", "kern", "local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7", "lpr", "mail", "news", "ntp", "syslog", "user", "uucp"),
			"logging_level":           stringEnumAttr("Syslog level for HAProxy logs.", "info", "", "alert", "crit", "debug", "emerg", "err", "info", "notice", "warning"),
			"logging_length":          optionalInt64Attr("Maximum log line length."),
			"stats_enabled":           boolAttr("Enable the HAProxy statistics page.", false),
			"stats_port":              int64Attr("HAProxy statistics page port.", 8822, int64validator.Between(1, 65535)),
			"stats_remote_enabled":    boolAttr("Enable remote access to HAProxy statistics.", false),
			"stats_remote_bind":       stringAttr("Remote bind address for HAProxy statistics.", ""),
			"stats_auth_enabled":      boolAttr("Enable HAProxy statistics authentication.", false),
			"stats_allowed_users": schema.SetAttribute{
				MarkdownDescription: "Allowed HAProxy statistics users.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"stats_allowed_groups": schema.SetAttribute{
				MarkdownDescription: "Allowed HAProxy statistics groups.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"stats_custom_options":        stringAttr("Custom HAProxy statistics options.", ""),
			"stats_prometheus_enabled":    boolAttr("Enable the HAProxy Prometheus exporter.", false),
			"stats_prometheus_bind":       stringAttr("HAProxy Prometheus exporter bind address.", "*:8404"),
			"stats_prometheus_path":       stringAttr("HAProxy Prometheus exporter path.", "/metrics"),
			"cache_enabled":               boolAttr("Enable HAProxy cache.", false),
			"cache_total_max_size":        int64Attr("Total HAProxy cache size in megabytes.", 4, int64validator.AtLeast(0)),
			"cache_max_age":               int64Attr("HAProxy cache maximum age in seconds.", 60, int64validator.AtLeast(0)),
			"cache_max_object_size":       optionalInt64Attr("HAProxy cache maximum object size."),
			"cache_process_vary":          boolAttr("Process Vary headers in HAProxy cache.", false),
			"cache_max_secondary_entries": int64Attr("Maximum HAProxy cache secondary entries.", 10, int64validator.AtLeast(0)),
			"peers_enabled":               boolAttr("Enable HAProxy peers.", false),
			"peers_name1":                 stringAttr("First HAProxy peer name.", ""),
			"peers_listen1":               stringAttr("First HAProxy peer listen address.", ""),
			"peers_port1":                 int64Attr("First HAProxy peer port.", 1024, int64validator.Between(1, 65535)),
			"peers_name2":                 stringAttr("Second HAProxy peer name.", ""),
			"peers_listen2":               stringAttr("Second HAProxy peer listen address.", ""),
			"peers_port2":                 int64Attr("Second HAProxy peer port.", 1024, int64validator.Between(1, 65535)),
		},
	}
}

func haproxySettingsDataSourceSchema() dschema.Schema {
	attrs := map[string]dschema.Attribute{}
	for name, description := range haproxySettingsAttributeDescriptions() {
		switch name {
		case "enabled", "graceful_stop", "seamless_reload", "store_ocsp", "show_intro", "tuning_root", "tuning_bogus_proxy_enabled", "tuning_ocsp_update_enabled", "tuning_ssl_defaults_enabled", "stats_enabled", "stats_remote_enabled", "stats_auth_enabled", "stats_prometheus_enabled", "cache_enabled", "cache_process_vary", "peers_enabled":
			attrs[name] = dschema.BoolAttribute{MarkdownDescription: description, Computed: true}
		case "tuning_max_connections", "tuning_nbthread", "tuning_max_dh_size", "tuning_buffer_size", "tuning_spread_checks", "tuning_lua_max_mem", "tuning_ocsp_update_min_delay", "tuning_ocsp_update_max_delay", "tuning_h2_initial_window_size", "tuning_h2_initial_window_size_outgoing", "tuning_h2_initial_window_size_incoming", "tuning_h2_max_concurrent_streams", "tuning_h2_max_concurrent_streams_outgoing", "tuning_h2_max_concurrent_streams_incoming", "defaults_max_connections", "defaults_max_connections_servers", "defaults_retries", "logging_length", "stats_port", "cache_total_max_size", "cache_max_age", "cache_max_object_size", "cache_max_secondary_entries", "peers_port1", "peers_port2":
			attrs[name] = dschema.Int64Attribute{MarkdownDescription: description, Computed: true}
		case "tuning_ssl_bind_options", "defaults_init_addr", "stats_allowed_users", "stats_allowed_groups":
			attrs[name] = dschema.SetAttribute{MarkdownDescription: description, ElementType: types.StringType, Computed: true}
		default:
			attrs[name] = dschema.StringAttribute{MarkdownDescription: description, Computed: true}
		}
	}

	return dschema.Schema{
		MarkdownDescription: "Read OPNsense HAProxy plugin service and global parameter settings.",
		Attributes:          attrs,
	}
}

func haproxySettingsAttributeDescriptions() map[string]string {
	return map[string]string{
		"enabled":                                   "Whether the HAProxy service is enabled.",
		"graceful_stop":                             "Whether graceful stop is enabled.",
		"hard_stop_after":                           "Maximum time before HAProxy is hard-stopped.",
		"close_spread_time":                         "Time window used to spread connection closes.",
		"seamless_reload":                           "Whether seamless reloads are enabled.",
		"store_ocsp":                                "Whether OCSP responses are stored.",
		"show_intro":                                "Whether the introductory/help text is shown.",
		"tuning_root":                               "Whether HAProxy runs with root privileges.",
		"tuning_max_connections":                    "Maximum concurrent connections.",
		"tuning_nbthread":                           "Number of HAProxy threads.",
		"tuning_resolvers_prefer":                   "Preferred resolver IP protocol.",
		"tuning_ssl_server_verify":                  "Default server certificate verification behavior.",
		"tuning_max_dh_size":                        "Maximum Diffie-Hellman parameter size.",
		"tuning_buffer_size":                        "HAProxy buffer size.",
		"tuning_spread_checks":                      "Health check spreading percentage.",
		"tuning_bogus_proxy_enabled":                "Whether bogus PROXY protocol support is enabled.",
		"tuning_lua_max_mem":                        "Lua maximum memory in megabytes.",
		"tuning_custom_options":                     "Custom HAProxy global options.",
		"tuning_ocsp_update_enabled":                "Whether automatic OCSP updates are enabled.",
		"tuning_ocsp_update_min_delay":              "Minimum OCSP update delay in seconds.",
		"tuning_ocsp_update_max_delay":              "Maximum OCSP update delay in seconds.",
		"tuning_ssl_defaults_enabled":               "Whether global SSL defaults are enabled.",
		"tuning_ssl_bind_options":                   "Global SSL bind options.",
		"tuning_ssl_min_version":                    "Minimum SSL/TLS version.",
		"tuning_ssl_max_version":                    "Maximum SSL/TLS version.",
		"tuning_ssl_cipher_list":                    "TLS 1.2 and earlier cipher list.",
		"tuning_ssl_cipher_suites":                  "TLS 1.3 cipher suites.",
		"tuning_h2_initial_window_size":             "HTTP/2 initial window size.",
		"tuning_h2_initial_window_size_outgoing":    "HTTP/2 outgoing initial window size.",
		"tuning_h2_initial_window_size_incoming":    "HTTP/2 incoming initial window size.",
		"tuning_h2_max_concurrent_streams":          "HTTP/2 maximum concurrent streams.",
		"tuning_h2_max_concurrent_streams_outgoing": "HTTP/2 outgoing maximum concurrent streams.",
		"tuning_h2_max_concurrent_streams_incoming": "HTTP/2 incoming maximum concurrent streams.",
		"defaults_max_connections":                  "Default maximum concurrent connections.",
		"defaults_max_connections_servers":          "Default maximum concurrent server connections.",
		"defaults_timeout_client":                   "Default client timeout.",
		"defaults_timeout_connect":                  "Default connect timeout.",
		"defaults_timeout_check":                    "Default health check timeout.",
		"defaults_timeout_server":                   "Default server timeout.",
		"defaults_retries":                          "Default retry count.",
		"defaults_redispatch":                       "Default redispatch behavior.",
		"defaults_init_addr":                        "Default server init-addr methods.",
		"defaults_custom_options":                   "Custom HAProxy default options.",
		"logging_host":                              "Syslog host for HAProxy logs.",
		"logging_facility":                          "Syslog facility for HAProxy logs.",
		"logging_level":                             "Syslog level for HAProxy logs.",
		"logging_length":                            "Maximum log line length.",
		"stats_enabled":                             "Whether the HAProxy statistics page is enabled.",
		"stats_port":                                "HAProxy statistics page port.",
		"stats_remote_enabled":                      "Whether remote access to HAProxy statistics is enabled.",
		"stats_remote_bind":                         "Remote bind address for HAProxy statistics.",
		"stats_auth_enabled":                        "Whether HAProxy statistics authentication is enabled.",
		"stats_allowed_users":                       "Allowed HAProxy statistics users.",
		"stats_allowed_groups":                      "Allowed HAProxy statistics groups.",
		"stats_custom_options":                      "Custom HAProxy statistics options.",
		"stats_prometheus_enabled":                  "Whether the HAProxy Prometheus exporter is enabled.",
		"stats_prometheus_bind":                     "HAProxy Prometheus exporter bind address.",
		"stats_prometheus_path":                     "HAProxy Prometheus exporter path.",
		"cache_enabled":                             "Whether HAProxy cache is enabled.",
		"cache_total_max_size":                      "Total HAProxy cache size in megabytes.",
		"cache_max_age":                             "HAProxy cache maximum age in seconds.",
		"cache_max_object_size":                     "HAProxy cache maximum object size.",
		"cache_process_vary":                        "Whether Vary headers are processed in HAProxy cache.",
		"cache_max_secondary_entries":               "Maximum HAProxy cache secondary entries.",
		"peers_enabled":                             "Whether HAProxy peers are enabled.",
		"peers_name1":                               "First HAProxy peer name.",
		"peers_listen1":                             "First HAProxy peer listen address.",
		"peers_port1":                               "First HAProxy peer port.",
		"peers_name2":                               "Second HAProxy peer name.",
		"peers_listen2":                             "Second HAProxy peer listen address.",
		"peers_port2":                               "Second HAProxy peer port.",
	}
}

func boolAttr(description string, defaultValue bool) schema.BoolAttribute {
	return schema.BoolAttribute{MarkdownDescription: description, Optional: true, Computed: true, Default: booldefault.StaticBool(defaultValue)}
}

func stringAttr(description, defaultValue string) schema.StringAttribute {
	return schema.StringAttribute{MarkdownDescription: description, Optional: true, Computed: true, Default: stringdefault.StaticString(defaultValue)}
}

func stringEnumAttr(description, defaultValue string, values ...string) schema.StringAttribute {
	return schema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(defaultValue),
		Validators: []validator.String{
			stringvalidator.OneOf(values...),
		},
	}
}

func int64Attr(description string, defaultValue int64, validators ...validator.Int64) schema.Int64Attribute {
	return schema.Int64Attribute{MarkdownDescription: description, Optional: true, Computed: true, Default: int64default.StaticInt64(defaultValue), Validators: validators}
}

func optionalInt64Attr(description string) schema.Int64Attribute {
	return schema.Int64Attribute{MarkdownDescription: description, Optional: true}
}
