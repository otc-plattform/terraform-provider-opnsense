package nginx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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

type nginxHTTPServerResourceModel struct {
	Id                             types.String `tfsdk:"id"`
	ServerName                     types.String `tfsdk:"server_name"`
	ListenHTTPAddress              types.String `tfsdk:"listen_http_address"`
	ListenHTTPSAddress             types.String `tfsdk:"listen_https_address"`
	DefaultServer                  types.Bool   `tfsdk:"default_server"`
	TLSRejectHandshake             types.Bool   `tfsdk:"tls_reject_handshake"`
	SyslogTargets                  types.Set    `tfsdk:"syslog_targets"`
	ProxyProtocol                  types.Bool   `tfsdk:"proxy_protocol"`
	TrustedProxies                 types.Set    `tfsdk:"trusted_proxies"`
	TrustedProxiesAlias            types.String `tfsdk:"trusted_proxies_alias"`
	RealIPSource                   types.String `tfsdk:"real_ip_source"`
	Locations                      types.Set    `tfsdk:"locations"`
	Rewrites                       types.Set    `tfsdk:"rewrites"`
	Root                           types.String `tfsdk:"root"`
	MaxBodySize                    types.String `tfsdk:"max_body_size"`
	BodyBufferSize                 types.String `tfsdk:"body_buffer_size"`
	Certificate                    types.String `tfsdk:"certificate"`
	CA                             types.String `tfsdk:"ca"`
	VerifyClient                   types.String `tfsdk:"verify_client"`
	ZeroRTT                        types.Bool   `tfsdk:"zero_rtt"`
	AccessLogFormat                types.String `tfsdk:"access_log_format"`
	ErrorLogLevel                  types.String `tfsdk:"error_log_level"`
	LogHandshakes                  types.Bool   `tfsdk:"log_handshakes"`
	EnableACMESupport              types.Bool   `tfsdk:"enable_acme_support"`
	Charset                        types.String `tfsdk:"charset"`
	HTTPSOnly                      types.Bool   `tfsdk:"https_only"`
	TLSProtocols                   types.Set    `tfsdk:"tls_protocols"`
	TLSCiphers                     types.String `tfsdk:"tls_ciphers"`
	TLSECDHCurve                   types.String `tfsdk:"tls_ecdh_curve"`
	TLSPreferServerCiphers         types.Bool   `tfsdk:"tls_prefer_server_ciphers"`
	Resolver                       types.String `tfsdk:"resolver"`
	OCSPStapling                   types.Bool   `tfsdk:"ocsp_stapling"`
	OCSPVerify                     types.Bool   `tfsdk:"ocsp_verify"`
	BlockNonpublicData             types.Bool   `tfsdk:"block_nonpublic_data"`
	DisableGzip                    types.Bool   `tfsdk:"disable_gzip"`
	DisableBotProtection           types.Bool   `tfsdk:"disable_bot_protection"`
	IPACL                          types.String `tfsdk:"ip_acl"`
	AdvancedACLServer              types.String `tfsdk:"advanced_acl_server"`
	Satisfy                        types.String `tfsdk:"satisfy"`
	NaxsiWhitelistSrcIP            types.Set    `tfsdk:"naxsi_whitelist_src_ip"`
	NaxsiExtensiveLog              types.Bool   `tfsdk:"naxsi_extensive_log"`
	Sendfile                       types.Bool   `tfsdk:"sendfile"`
	ClientHeaderBufferSize         types.String `tfsdk:"client_header_buffer_size"`
	LargeClientHeaderBuffersNumber types.String `tfsdk:"large_client_header_buffers_number"`
	LargeClientHeaderBuffersSize   types.String `tfsdk:"large_client_header_buffers_size"`
	SecurityHeader                 types.String `tfsdk:"security_header"`
	LimitRequestConnections        types.Set    `tfsdk:"limit_request_connections"`
	ErrorPages                     types.Set    `tfsdk:"error_pages"`
}

func nginxHTTPServerResourceSchema() schema.Schema {
	defaultTLSProtocols, _ := types.SetValue(types.StringType, []attr.Value{
		types.StringValue("TLSv1.2"),
		types.StringValue("TLSv1.3"),
	})

	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense Nginx HTTP server.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"server_name": schema.StringAttribute{
				MarkdownDescription: "Server name identifier.",
				Required:            true,
			},
			"listen_http_address": schema.StringAttribute{
				MarkdownDescription: "Address (including port) for HTTP listeners.",
				Optional:            true,
			},
			"listen_https_address": schema.StringAttribute{
				MarkdownDescription: "Address (including port) for HTTPS listeners.",
				Optional:            true,
			},
			"default_server": schema.BoolAttribute{
				MarkdownDescription: "Mark this server as the default virtual host.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tls_reject_handshake": schema.BoolAttribute{
				MarkdownDescription: "Reject TLS handshakes for this server.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"syslog_targets": schema.SetAttribute{
				MarkdownDescription: "Syslog targets associated with the server.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"proxy_protocol": schema.BoolAttribute{
				MarkdownDescription: "Enable proxy protocol handling.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"trusted_proxies": schema.SetAttribute{
				MarkdownDescription: "Trusted proxy addresses.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"trusted_proxies_alias": schema.StringAttribute{
				MarkdownDescription: "Alias containing trusted proxy addresses.",
				Optional:            true,
			},
			"real_ip_source": schema.StringAttribute{
				MarkdownDescription: "Source to retrieve the real client IP.",
				Optional:            true,
			},
			"locations": schema.SetAttribute{
				MarkdownDescription: "Location UUIDs served by this HTTP server.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"rewrites": schema.SetAttribute{
				MarkdownDescription: "Rewrite rules applied before locations are evaluated.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"root": schema.StringAttribute{
				MarkdownDescription: "Root directory served by this HTTP server.",
				Optional:            true,
			},
			"max_body_size": schema.StringAttribute{
				MarkdownDescription: "Maximum request body size (e.g. 10m).",
				Optional:            true,
			},
			"body_buffer_size": schema.StringAttribute{
				MarkdownDescription: "Buffer size for request bodies.",
				Optional:            true,
			},
			"certificate": schema.StringAttribute{
				MarkdownDescription: "Certificate UUID used for TLS.",
				Optional:            true,
			},
			"ca": schema.StringAttribute{
				MarkdownDescription: "Certificate Authority UUID used for client verification.",
				Optional:            true,
			},
			"verify_client": schema.StringAttribute{
				MarkdownDescription: "Client verification mode.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("off"),
				Validators: []validator.String{
					stringvalidator.OneOf("off", "on", "optional", "optional_no_ca"),
				},
			},
			"zero_rtt": schema.BoolAttribute{
				MarkdownDescription: "Enable TLS 0-RTT.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"access_log_format": schema.StringAttribute{
				MarkdownDescription: "Access log format identifier.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("main"),
				Validators: []validator.String{
					stringvalidator.OneOf("main", "main_ext", "anonymized", "disabled"),
				},
			},
			"error_log_level": schema.StringAttribute{
				MarkdownDescription: "Error log verbosity level.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("error"),
				Validators: []validator.String{
					stringvalidator.OneOf("emerg", "alert", "crit", "error", "warn", "notice", "info"),
				},
			},
			"log_handshakes": schema.BoolAttribute{
				MarkdownDescription: "Log TLS handshakes.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_acme_support": schema.BoolAttribute{
				MarkdownDescription: "Enable ACME challenge support on this server.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"charset": schema.StringAttribute{
				MarkdownDescription: "Character set used for responses.",
				Optional:            true,
			},
			"https_only": schema.BoolAttribute{
				MarkdownDescription: "Redirect HTTP requests to HTTPS.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tls_protocols": schema.SetAttribute{
				MarkdownDescription: "TLS protocols allowed for this server.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(defaultTLSProtocols),
			},
			"tls_ciphers": schema.StringAttribute{
				MarkdownDescription: "TLS cipher string.",
				Optional:            true,
			},
			"tls_ecdh_curve": schema.StringAttribute{
				MarkdownDescription: "Elliptic curves used for TLS.",
				Optional:            true,
			},
			"tls_prefer_server_ciphers": schema.BoolAttribute{
				MarkdownDescription: "Prefer server ciphers over client selection.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"resolver": schema.StringAttribute{
				MarkdownDescription: "Resolver configuration identifier.",
				Optional:            true,
			},
			"ocsp_stapling": schema.BoolAttribute{
				MarkdownDescription: "Enable OCSP stapling.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ocsp_verify": schema.BoolAttribute{
				MarkdownDescription: "Verify OCSP responses.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"block_nonpublic_data": schema.BoolAttribute{
				MarkdownDescription: "Block non-public data.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disable_gzip": schema.BoolAttribute{
				MarkdownDescription: "Disable gzip compression.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disable_bot_protection": schema.BoolAttribute{
				MarkdownDescription: "Disable bot protection.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ip_acl": schema.StringAttribute{
				MarkdownDescription: "IP ACL applied to the server.",
				Optional:            true,
			},
			"advanced_acl_server": schema.StringAttribute{
				MarkdownDescription: "Advanced ACL server option.",
				Optional:            true,
			},
			"satisfy": schema.StringAttribute{
				MarkdownDescription: "Access satisfy directive.",
				Optional:            true,
			},
			"naxsi_whitelist_src_ip": schema.SetAttribute{
				MarkdownDescription: "NAXSI whitelist source IPs.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"naxsi_extensive_log": schema.BoolAttribute{
				MarkdownDescription: "Enable extensive NAXSI logging.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"sendfile": schema.BoolAttribute{
				MarkdownDescription: "Enable sendfile support.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"client_header_buffer_size": schema.StringAttribute{
				MarkdownDescription: "Size of the client header buffer.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1"),
			},
			"large_client_header_buffers_number": schema.StringAttribute{
				MarkdownDescription: "Number of large client header buffers.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("4"),
			},
			"large_client_header_buffers_size": schema.StringAttribute{
				MarkdownDescription: "Size of large client header buffers.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("8"),
			},
			"security_header": schema.StringAttribute{
				MarkdownDescription: "Security header policy.",
				Optional:            true,
			},
			"limit_request_connections": schema.SetAttribute{
				MarkdownDescription: "Limit request connection identifiers applied to the server.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"error_pages": schema.SetAttribute{
				MarkdownDescription: "Custom error pages.",
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func nginxHTTPServerDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense Nginx HTTP server.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the HTTP server.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"server_name": dschema.StringAttribute{
				MarkdownDescription: "Server name identifier.",
				Computed:            true,
			},
			"listen_http_address": dschema.StringAttribute{
				MarkdownDescription: "Address used for HTTP listeners.",
				Computed:            true,
			},
			"listen_https_address": dschema.StringAttribute{
				MarkdownDescription: "Address used for HTTPS listeners.",
				Computed:            true,
			},
			"default_server": dschema.BoolAttribute{
				MarkdownDescription: "Whether this server is the default virtual host.",
				Computed:            true,
			},
			"tls_reject_handshake": dschema.BoolAttribute{
				MarkdownDescription: "Whether TLS handshakes are rejected.",
				Computed:            true,
			},
			"syslog_targets": dschema.SetAttribute{
				MarkdownDescription: "Syslog targets associated with the server.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"proxy_protocol": dschema.BoolAttribute{
				MarkdownDescription: "Whether proxy protocol handling is enabled.",
				Computed:            true,
			},
			"trusted_proxies": dschema.SetAttribute{
				MarkdownDescription: "Trusted proxy addresses.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"trusted_proxies_alias": dschema.StringAttribute{
				MarkdownDescription: "Alias containing trusted proxies.",
				Computed:            true,
			},
			"real_ip_source": dschema.StringAttribute{
				MarkdownDescription: "Source to retrieve the real client IP.",
				Computed:            true,
			},
			"locations": dschema.SetAttribute{
				MarkdownDescription: "Location UUIDs served by this HTTP server.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"rewrites": dschema.SetAttribute{
				MarkdownDescription: "Rewrite rules applied before locations are evaluated.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"root": dschema.StringAttribute{
				MarkdownDescription: "Root directory served by this HTTP server.",
				Computed:            true,
			},
			"max_body_size": dschema.StringAttribute{
				MarkdownDescription: "Maximum request body size.",
				Computed:            true,
			},
			"body_buffer_size": dschema.StringAttribute{
				MarkdownDescription: "Buffer size for request bodies.",
				Computed:            true,
			},
			"certificate": dschema.StringAttribute{
				MarkdownDescription: "Certificate UUID used for TLS.",
				Computed:            true,
			},
			"ca": dschema.StringAttribute{
				MarkdownDescription: "Certificate Authority UUID used for client verification.",
				Computed:            true,
			},
			"verify_client": dschema.StringAttribute{
				MarkdownDescription: "Client verification mode.",
				Computed:            true,
			},
			"zero_rtt": dschema.BoolAttribute{
				MarkdownDescription: "Whether TLS 0-RTT is enabled.",
				Computed:            true,
			},
			"access_log_format": dschema.StringAttribute{
				MarkdownDescription: "Access log format identifier.",
				Computed:            true,
			},
			"error_log_level": dschema.StringAttribute{
				MarkdownDescription: "Error log verbosity level.",
				Computed:            true,
			},
			"log_handshakes": dschema.BoolAttribute{
				MarkdownDescription: "Whether TLS handshakes are logged.",
				Computed:            true,
			},
			"enable_acme_support": dschema.BoolAttribute{
				MarkdownDescription: "Whether ACME support is enabled.",
				Computed:            true,
			},
			"charset": dschema.StringAttribute{
				MarkdownDescription: "Character set used for responses.",
				Computed:            true,
			},
			"https_only": dschema.BoolAttribute{
				MarkdownDescription: "Whether HTTP is redirected to HTTPS.",
				Computed:            true,
			},
			"tls_protocols": dschema.SetAttribute{
				MarkdownDescription: "TLS protocols allowed for this server.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"tls_ciphers": dschema.StringAttribute{
				MarkdownDescription: "TLS cipher string.",
				Computed:            true,
			},
			"tls_ecdh_curve": dschema.StringAttribute{
				MarkdownDescription: "Elliptic curves for TLS.",
				Computed:            true,
			},
			"tls_prefer_server_ciphers": dschema.BoolAttribute{
				MarkdownDescription: "Whether server ciphers are preferred.",
				Computed:            true,
			},
			"resolver": dschema.StringAttribute{
				MarkdownDescription: "Resolver configuration identifier.",
				Computed:            true,
			},
			"ocsp_stapling": dschema.BoolAttribute{
				MarkdownDescription: "Whether OCSP stapling is enabled.",
				Computed:            true,
			},
			"ocsp_verify": dschema.BoolAttribute{
				MarkdownDescription: "Whether OCSP responses are verified.",
				Computed:            true,
			},
			"block_nonpublic_data": dschema.BoolAttribute{
				MarkdownDescription: "Whether non-public data is blocked.",
				Computed:            true,
			},
			"disable_gzip": dschema.BoolAttribute{
				MarkdownDescription: "Whether gzip is disabled.",
				Computed:            true,
			},
			"disable_bot_protection": dschema.BoolAttribute{
				MarkdownDescription: "Whether bot protection is disabled.",
				Computed:            true,
			},
			"ip_acl": dschema.StringAttribute{
				MarkdownDescription: "IP ACL applied to the server.",
				Computed:            true,
			},
			"advanced_acl_server": dschema.StringAttribute{
				MarkdownDescription: "Advanced ACL server option.",
				Computed:            true,
			},
			"satisfy": dschema.StringAttribute{
				MarkdownDescription: "Access satisfy directive.",
				Computed:            true,
			},
			"naxsi_whitelist_src_ip": dschema.SetAttribute{
				MarkdownDescription: "NAXSI whitelist source IPs.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"naxsi_extensive_log": dschema.BoolAttribute{
				MarkdownDescription: "Whether extensive NAXSI logging is enabled.",
				Computed:            true,
			},
			"sendfile": dschema.BoolAttribute{
				MarkdownDescription: "Whether sendfile is enabled.",
				Computed:            true,
			},
			"client_header_buffer_size": dschema.StringAttribute{
				MarkdownDescription: "Size of the client header buffer.",
				Computed:            true,
			},
			"large_client_header_buffers_number": dschema.StringAttribute{
				MarkdownDescription: "Number of large client header buffers.",
				Computed:            true,
			},
			"large_client_header_buffers_size": dschema.StringAttribute{
				MarkdownDescription: "Size of large client header buffers.",
				Computed:            true,
			},
			"security_header": dschema.StringAttribute{
				MarkdownDescription: "Security header policy.",
				Computed:            true,
			},
			"limit_request_connections": dschema.SetAttribute{
				MarkdownDescription: "Limit request connection identifiers applied to the server.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"error_pages": dschema.SetAttribute{
				MarkdownDescription: "Custom error pages.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}
