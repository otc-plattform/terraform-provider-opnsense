package nginx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type nginxLocationResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Description             types.String `tfsdk:"description"`
	URLPattern              types.String `tfsdk:"url_pattern"`
	MatchType               types.String `tfsdk:"match_type"`
	PathPrefix              types.String `tfsdk:"path_prefix"`
	Upstream                types.String `tfsdk:"upstream"`
	EnableSecRules          types.Bool   `tfsdk:"enable_sec_rules"`
	EnableLearningMode      types.Bool   `tfsdk:"enable_learning_mode"`
	SecRulesErrorPage       types.String `tfsdk:"sec_rules_error_page"`
	XSSBlockScore           types.Int64  `tfsdk:"xss_block_score"`
	SQLiBlockScore          types.Int64  `tfsdk:"sqli_block_score"`
	CustomPolicy            types.Set    `tfsdk:"custom_policy"`
	CachePath               types.String `tfsdk:"cache_path"`
	CacheUseStale           types.String `tfsdk:"cache_use_stale"`
	CacheMethods            types.String `tfsdk:"cache_methods"`
	CacheMinUses            types.Int64  `tfsdk:"cache_min_uses"`
	CacheValid              types.String `tfsdk:"cache_valid"`
	CacheBackgroundUpdate   types.Bool   `tfsdk:"cache_background_update"`
	CacheLock               types.Bool   `tfsdk:"cache_lock"`
	CacheRevalidate         types.Bool   `tfsdk:"cache_revalidate"`
	Root                    types.String `tfsdk:"root"`
	Rewrites                types.Set    `tfsdk:"rewrites"`
	Index                   types.Set    `tfsdk:"index"`
	AutoIndex               types.Bool   `tfsdk:"auto_index"`
	AuthBasic               types.Bool   `tfsdk:"auth_basic"`
	AuthBasicUserFile       types.String `tfsdk:"auth_basic_user_file"`
	AdvancedACL             types.Bool   `tfsdk:"advanced_acl"`
	ForceHTTPS              types.Bool   `tfsdk:"force_https"`
	PHPEnable               types.Bool   `tfsdk:"php_enable"`
	PHPOverrideScriptName   types.String `tfsdk:"php_override_script_name"`
	LimitRequestConnections types.Set    `tfsdk:"limit_request_connections"`
	MaxBodySize             types.String `tfsdk:"max_body_size"`
	BodyBufferSize          types.String `tfsdk:"body_buffer_size"`
	Honeypot                types.Bool   `tfsdk:"honeypot"`
	WebSocket               types.Bool   `tfsdk:"websocket"`
	UpstreamKeepalive       types.Bool   `tfsdk:"upstream_keepalive"`
	ProxyBufferSize         types.String `tfsdk:"proxy_buffer_size"`
	ProxyBuffersCount       types.String `tfsdk:"proxy_buffers_count"`
	ProxyBuffersSize        types.String `tfsdk:"proxy_buffers_size"`
	ProxyBusyBuffersSize    types.String `tfsdk:"proxy_busy_buffers_size"`
	ProxyIgnoreClientAbort  types.Bool   `tfsdk:"proxy_ignore_client_abort"`
	ProxyRequestBuffering   types.Bool   `tfsdk:"proxy_request_buffering"`
	ProxyBuffering          types.Bool   `tfsdk:"proxy_buffering"`
	ProxyReadTimeout        types.String `tfsdk:"proxy_read_timeout"`
	ProxySendTimeout        types.String `tfsdk:"proxy_send_timeout"`
	IPACL                   types.String `tfsdk:"ip_acl"`
	Satisfy                 types.String `tfsdk:"satisfy"`
	ProxyMaxTempFileSize    types.String `tfsdk:"proxy_max_temp_file_size"`
	ProxySSLServerName      types.Bool   `tfsdk:"proxy_ssl_server_name"`
	ErrorPages              types.Set    `tfsdk:"error_pages"`
}

func nginxLocationResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense Nginx location.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Human readable description for the location.",
				Required:            true,
			},
			"url_pattern": schema.StringAttribute{
				MarkdownDescription: "URL pattern that the location will match.",
				Required:            true,
			},
			"match_type": schema.StringAttribute{
				MarkdownDescription: "Match strategy for the URL pattern.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators: []validator.String{
					stringvalidator.OneOf("", "=", "~", "~*", "^~"),
				},
			},
			"path_prefix": schema.StringAttribute{
				MarkdownDescription: "Optional path prefix when matching upstream requests.",
				Optional:            true,
			},
			"upstream": schema.StringAttribute{
				MarkdownDescription: "UUID of the upstream to service this location.",
				Optional:            true,
			},
			"enable_sec_rules": schema.BoolAttribute{
				MarkdownDescription: "Enable the web application firewall for this location.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_learning_mode": schema.BoolAttribute{
				MarkdownDescription: "Enable learning mode for waf security rules.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"sec_rules_error_page": schema.StringAttribute{
				MarkdownDescription: "Error page to use when a WAF rule blocks a request.",
				Optional:            true,
			},
			"xss_block_score": schema.Int64Attribute{
				MarkdownDescription: "Threshold score that will block XSS requests. Defaults to 100.",
				Optional:            true,
			},
			"sqli_block_score": schema.Int64Attribute{
				MarkdownDescription: "Threshold score that will block SQL injection requests. Defaults to 100.",
				Optional:            true,
			},
			"custom_policy": schema.SetAttribute{
				MarkdownDescription: "Custom WAF policy identifiers to apply.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"cache_path": schema.StringAttribute{
				MarkdownDescription: "Cache path identifier.",
				Optional:            true,
			},
			"cache_use_stale": schema.StringAttribute{
				MarkdownDescription: "Conditions under which a stale cache may be served.",
				Optional:            true,
			},
			"cache_methods": schema.StringAttribute{
				MarkdownDescription: "Request methods that should be cached.",
				Optional:            true,
			},
			"cache_min_uses": schema.Int64Attribute{
				MarkdownDescription: "Minimum number of uses before caching is enabled.",
				Optional:            true,
			},
			"cache_valid": schema.StringAttribute{
				MarkdownDescription: "Cache validity configuration.",
				Optional:            true,
			},
			"cache_background_update": schema.BoolAttribute{
				MarkdownDescription: "Allow cache background updates.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cache_lock": schema.BoolAttribute{
				MarkdownDescription: "Enable cache lock to avoid cache stampede.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cache_revalidate": schema.BoolAttribute{
				MarkdownDescription: "Enable cache revalidation.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"root": schema.StringAttribute{
				MarkdownDescription: "Root directory for the location.",
				Optional:            true,
			},
			"rewrites": schema.SetAttribute{
				MarkdownDescription: "Rewrite rules to apply within the location.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"index": schema.SetAttribute{
				MarkdownDescription: "Index files to serve for directory requests.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"auto_index": schema.BoolAttribute{
				MarkdownDescription: "Enable directory listing.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"auth_basic": schema.BoolAttribute{
				MarkdownDescription: "Require HTTP basic authentication.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"auth_basic_user_file": schema.StringAttribute{
				MarkdownDescription: "Identifier of the user file used for HTTP basic auth.",
				Optional:            true,
			},
			"advanced_acl": schema.BoolAttribute{
				MarkdownDescription: "Enable the advanced ACL definition for this location.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"force_https": schema.BoolAttribute{
				MarkdownDescription: "Force HTTPS redirects.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"php_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable PHP handling for this location.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"php_override_script_name": schema.StringAttribute{
				MarkdownDescription: "Override script name passed to PHP.",
				Optional:            true,
			},
			"limit_request_connections": schema.SetAttribute{
				MarkdownDescription: "Limit request connection identifiers to apply.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"max_body_size": schema.StringAttribute{
				MarkdownDescription: "Maximum accepted body size for requests (e.g. `10m`).",
				Optional:            true,
			},
			"body_buffer_size": schema.StringAttribute{
				MarkdownDescription: "Buffer size used for request bodies.",
				Optional:            true,
			},
			"honeypot": schema.BoolAttribute{
				MarkdownDescription: "Enable honeypot mode.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"websocket": schema.BoolAttribute{
				MarkdownDescription: "Enable websocket support.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"upstream_keepalive": schema.BoolAttribute{
				MarkdownDescription: "Enable keepalive connections for the upstream.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"proxy_buffer_size": schema.StringAttribute{
				MarkdownDescription: "Proxy buffer size.",
				Optional:            true,
			},
			"proxy_buffers_count": schema.StringAttribute{
				MarkdownDescription: "Number of proxy buffers.",
				Optional:            true,
			},
			"proxy_buffers_size": schema.StringAttribute{
				MarkdownDescription: "Size of each proxy buffer.",
				Optional:            true,
			},
			"proxy_busy_buffers_size": schema.StringAttribute{
				MarkdownDescription: "Size of busy proxy buffers.",
				Optional:            true,
			},
			"proxy_ignore_client_abort": schema.BoolAttribute{
				MarkdownDescription: "Ignore client aborts.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"proxy_request_buffering": schema.BoolAttribute{
				MarkdownDescription: "Enable proxy request buffering.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"proxy_buffering": schema.BoolAttribute{
				MarkdownDescription: "Enable proxy buffering.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"proxy_read_timeout": schema.StringAttribute{
				MarkdownDescription: "Proxy read timeout.",
				Optional:            true,
			},
			"proxy_send_timeout": schema.StringAttribute{
				MarkdownDescription: "Proxy send timeout.",
				Optional:            true,
			},
			"ip_acl": schema.StringAttribute{
				MarkdownDescription: "Identifier of the IP ACL to apply.",
				Optional:            true,
			},
			"satisfy": schema.StringAttribute{
				MarkdownDescription: "Satisfy directive for access control.",
				Optional:            true,
			},
			"proxy_max_temp_file_size": schema.StringAttribute{
				MarkdownDescription: "Maximum size of proxy temporary files.",
				Optional:            true,
			},
			"proxy_ssl_server_name": schema.BoolAttribute{
				MarkdownDescription: "Enable proxy SSL server name.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"error_pages": schema.SetAttribute{
				MarkdownDescription: "Custom error pages for the location.",
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func nginxLocationDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense Nginx location.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the location.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Human readable description for the location.",
				Computed:            true,
			},
			"url_pattern": dschema.StringAttribute{
				MarkdownDescription: "URL pattern that the location matches.",
				Computed:            true,
			},
			"match_type": dschema.StringAttribute{
				MarkdownDescription: "Match strategy for the URL pattern.",
				Computed:            true,
			},
			"path_prefix": dschema.StringAttribute{
				MarkdownDescription: "Optional path prefix when matching upstream requests.",
				Computed:            true,
			},
			"upstream": dschema.StringAttribute{
				MarkdownDescription: "UUID of the upstream servicing this location.",
				Computed:            true,
			},
			"enable_sec_rules": dschema.BoolAttribute{
				MarkdownDescription: "Whether the web application firewall is enabled.",
				Computed:            true,
			},
			"enable_learning_mode": dschema.BoolAttribute{
				MarkdownDescription: "Whether learning mode is enabled for waf rules.",
				Computed:            true,
			},
			"sec_rules_error_page": dschema.StringAttribute{
				MarkdownDescription: "Error page used when a WAF rule blocks a request.",
				Computed:            true,
			},
			"xss_block_score": dschema.Int64Attribute{
				MarkdownDescription: "Threshold score that triggers XSS blocking.",
				Computed:            true,
			},
			"sqli_block_score": dschema.Int64Attribute{
				MarkdownDescription: "Threshold score that triggers SQLi blocking.",
				Computed:            true,
			},
			"custom_policy": dschema.SetAttribute{
				MarkdownDescription: "Custom WAF policy identifiers applied to the location.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"cache_path": dschema.StringAttribute{
				MarkdownDescription: "Cache path identifier.",
				Computed:            true,
			},
			"cache_use_stale": dschema.StringAttribute{
				MarkdownDescription: "Conditions under which a stale cache may be served.",
				Computed:            true,
			},
			"cache_methods": dschema.StringAttribute{
				MarkdownDescription: "Request methods that should be cached.",
				Computed:            true,
			},
			"cache_min_uses": dschema.Int64Attribute{
				MarkdownDescription: "Minimum number of uses before caching is enabled.",
				Computed:            true,
			},
			"cache_valid": dschema.StringAttribute{
				MarkdownDescription: "Cache validity configuration.",
				Computed:            true,
			},
			"cache_background_update": dschema.BoolAttribute{
				MarkdownDescription: "Whether cache background updates are enabled.",
				Computed:            true,
			},
			"cache_lock": dschema.BoolAttribute{
				MarkdownDescription: "Whether cache lock is enabled.",
				Computed:            true,
			},
			"cache_revalidate": dschema.BoolAttribute{
				MarkdownDescription: "Whether cache revalidation is enabled.",
				Computed:            true,
			},
			"root": dschema.StringAttribute{
				MarkdownDescription: "Root directory for the location.",
				Computed:            true,
			},
			"rewrites": dschema.SetAttribute{
				MarkdownDescription: "Rewrite rules applied within the location.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"index": dschema.SetAttribute{
				MarkdownDescription: "Index files served for directory requests.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"auto_index": dschema.BoolAttribute{
				MarkdownDescription: "Whether directory listing is enabled.",
				Computed:            true,
			},
			"auth_basic": dschema.BoolAttribute{
				MarkdownDescription: "Whether HTTP basic authentication is required.",
				Computed:            true,
			},
			"auth_basic_user_file": dschema.StringAttribute{
				MarkdownDescription: "Identifier of the user file used for HTTP basic auth.",
				Computed:            true,
			},
			"advanced_acl": dschema.BoolAttribute{
				MarkdownDescription: "Whether the advanced ACL definition is enabled.",
				Computed:            true,
			},
			"force_https": dschema.BoolAttribute{
				MarkdownDescription: "Whether HTTPS redirects are forced.",
				Computed:            true,
			},
			"php_enable": dschema.BoolAttribute{
				MarkdownDescription: "Whether PHP handling is enabled.",
				Computed:            true,
			},
			"php_override_script_name": dschema.StringAttribute{
				MarkdownDescription: "Override script name passed to PHP.",
				Computed:            true,
			},
			"limit_request_connections": dschema.SetAttribute{
				MarkdownDescription: "Limit request connection identifiers applied.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"max_body_size": dschema.StringAttribute{
				MarkdownDescription: "Maximum accepted body size for requests.",
				Computed:            true,
			},
			"body_buffer_size": dschema.StringAttribute{
				MarkdownDescription: "Buffer size used for request bodies.",
				Computed:            true,
			},
			"honeypot": dschema.BoolAttribute{
				MarkdownDescription: "Whether honeypot mode is enabled.",
				Computed:            true,
			},
			"websocket": dschema.BoolAttribute{
				MarkdownDescription: "Whether websocket support is enabled.",
				Computed:            true,
			},
			"upstream_keepalive": dschema.BoolAttribute{
				MarkdownDescription: "Whether upstream keepalive connections are enabled.",
				Computed:            true,
			},
			"proxy_buffer_size": dschema.StringAttribute{
				MarkdownDescription: "Proxy buffer size.",
				Computed:            true,
			},
			"proxy_buffers_count": dschema.StringAttribute{
				MarkdownDescription: "Number of proxy buffers.",
				Computed:            true,
			},
			"proxy_buffers_size": dschema.StringAttribute{
				MarkdownDescription: "Size of each proxy buffer.",
				Computed:            true,
			},
			"proxy_busy_buffers_size": dschema.StringAttribute{
				MarkdownDescription: "Size of busy proxy buffers.",
				Computed:            true,
			},
			"proxy_ignore_client_abort": dschema.BoolAttribute{
				MarkdownDescription: "Whether client aborts are ignored.",
				Computed:            true,
			},
			"proxy_request_buffering": dschema.BoolAttribute{
				MarkdownDescription: "Whether proxy request buffering is enabled.",
				Computed:            true,
			},
			"proxy_buffering": dschema.BoolAttribute{
				MarkdownDescription: "Whether proxy buffering is enabled.",
				Computed:            true,
			},
			"proxy_read_timeout": dschema.StringAttribute{
				MarkdownDescription: "Proxy read timeout.",
				Computed:            true,
			},
			"proxy_send_timeout": dschema.StringAttribute{
				MarkdownDescription: "Proxy send timeout.",
				Computed:            true,
			},
			"ip_acl": dschema.StringAttribute{
				MarkdownDescription: "Identifier of the IP ACL applied.",
				Computed:            true,
			},
			"satisfy": dschema.StringAttribute{
				MarkdownDescription: "Satisfy directive for access control.",
				Computed:            true,
			},
			"proxy_max_temp_file_size": dschema.StringAttribute{
				MarkdownDescription: "Maximum size of proxy temporary files.",
				Computed:            true,
			},
			"proxy_ssl_server_name": dschema.BoolAttribute{
				MarkdownDescription: "Whether proxy SSL server name is enabled.",
				Computed:            true,
			},
			"error_pages": dschema.SetAttribute{
				MarkdownDescription: "Custom error pages for the location.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}
