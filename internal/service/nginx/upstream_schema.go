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

type nginxUpstreamResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Description            types.String `tfsdk:"description"`
	ServerEntries          types.Set    `tfsdk:"server_entries"`
	LoadBalancingAlgorithm types.String `tfsdk:"load_balancing_algorithm"`
	ProxyProtocol          types.Bool   `tfsdk:"proxy_protocol"`
	Keepalive              types.String `tfsdk:"keepalive"`
	KeepaliveRequests      types.String `tfsdk:"keepalive_requests"`
	KeepaliveTimeout       types.String `tfsdk:"keepalive_timeout"`
	HostPort               types.String `tfsdk:"host_port"`
	XForwardedHostVerbatim types.Bool   `tfsdk:"x_forwarded_host_verbatim"`
	TLSEnable              types.Bool   `tfsdk:"tls_enable"`
	TLSClientCertificate   types.String `tfsdk:"tls_client_certificate"`
	TLSNameOverride        types.String `tfsdk:"tls_name_override"`
	TLSProtocolVersions    types.Set    `tfsdk:"tls_protocol_versions"`
	TLSSessionReuse        types.Bool   `tfsdk:"tls_session_reuse"`
	TLSTrustedCertificate  types.String `tfsdk:"tls_trusted_certificate"`
	TLSVerify              types.Bool   `tfsdk:"tls_verify"`
	TLSVerifyDepth         types.String `tfsdk:"tls_verify_depth"`
	Store                  types.Bool   `tfsdk:"store"`
}

func nginxUpstreamResourceSchema() schema.Schema {
	defaultTLSProtocols, _ := types.SetValue(types.StringType, []attr.Value{
		types.StringValue("TLSv1.2"),
		types.StringValue("TLSv1.3"),
	})

	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense Nginx upstream.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Human readable description for the upstream.",
				Required:            true,
			},
			"server_entries": schema.SetAttribute{
				MarkdownDescription: "UUIDs of upstream servers that belong to this upstream.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"load_balancing_algorithm": schema.StringAttribute{
				MarkdownDescription: "Load balancing algorithm (`` or `ip_hash`).",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators: []validator.String{
					stringvalidator.OneOf("", "ip_hash"),
				},
			},
			"proxy_protocol": schema.BoolAttribute{
				MarkdownDescription: "Enable proxy protocol when connecting to upstream servers.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"keepalive": schema.StringAttribute{
				MarkdownDescription: "Number of idle keepalive connections to upstream servers.",
				Optional:            true,
			},
			"keepalive_requests": schema.StringAttribute{
				MarkdownDescription: "Maximum number of requests served through one keepalive connection.",
				Optional:            true,
			},
			"keepalive_timeout": schema.StringAttribute{
				MarkdownDescription: "Keepalive timeout in seconds.",
				Optional:            true,
			},
			"host_port": schema.StringAttribute{
				MarkdownDescription: "Optional host:port override for upstream requests.",
				Optional:            true,
			},
			"x_forwarded_host_verbatim": schema.BoolAttribute{
				MarkdownDescription: "Forward original host header verbatim.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tls_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable TLS when connecting to upstream servers.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tls_client_certificate": schema.StringAttribute{
				MarkdownDescription: "Client certificate UUID used for upstream TLS connections.",
				Optional:            true,
			},
			"tls_name_override": schema.StringAttribute{
				MarkdownDescription: "Override the SNI/hostname used for upstream TLS handshakes.",
				Optional:            true,
			},
			"tls_protocol_versions": schema.SetAttribute{
				MarkdownDescription: "TLS protocol versions allowed for upstream connections.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(defaultTLSProtocols),
			},
			"tls_session_reuse": schema.BoolAttribute{
				MarkdownDescription: "Enable TLS session reuse for upstream connections.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"tls_trusted_certificate": schema.StringAttribute{
				MarkdownDescription: "Trusted certificate store UUID used to verify upstream TLS connections.",
				Optional:            true,
			},
			"tls_verify": schema.BoolAttribute{
				MarkdownDescription: "Verify upstream TLS certificates.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tls_verify_depth": schema.StringAttribute{
				MarkdownDescription: "Maximum verification depth for upstream TLS certificates.",
				Optional:            true,
			},
			"store": schema.BoolAttribute{
				MarkdownDescription: "Persist upstream state in the cache store.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

func nginxUpstreamDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense Nginx upstream.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the upstream.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description of the upstream.",
				Computed:            true,
			},
			"server_entries": dschema.SetAttribute{
				MarkdownDescription: "UUIDs of upstream servers that belong to this upstream.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"load_balancing_algorithm": dschema.StringAttribute{
				MarkdownDescription: "Load balancing algorithm used by the upstream.",
				Computed:            true,
			},
			"proxy_protocol": dschema.BoolAttribute{
				MarkdownDescription: "Whether proxy protocol is enabled for upstream connections.",
				Computed:            true,
			},
			"keepalive": dschema.StringAttribute{
				MarkdownDescription: "Number of idle keepalive connections to upstream servers.",
				Computed:            true,
			},
			"keepalive_requests": dschema.StringAttribute{
				MarkdownDescription: "Maximum number of requests served through one keepalive connection.",
				Computed:            true,
			},
			"keepalive_timeout": dschema.StringAttribute{
				MarkdownDescription: "Keepalive timeout in seconds.",
				Computed:            true,
			},
			"host_port": dschema.StringAttribute{
				MarkdownDescription: "Host:port override for upstream requests.",
				Computed:            true,
			},
			"x_forwarded_host_verbatim": dschema.BoolAttribute{
				MarkdownDescription: "Whether the original host header is forwarded verbatim.",
				Computed:            true,
			},
			"tls_enable": dschema.BoolAttribute{
				MarkdownDescription: "Whether TLS is enabled for upstream connections.",
				Computed:            true,
			},
			"tls_client_certificate": dschema.StringAttribute{
				MarkdownDescription: "Client certificate UUID used for upstream TLS connections.",
				Computed:            true,
			},
			"tls_name_override": dschema.StringAttribute{
				MarkdownDescription: "Override for the SNI/hostname used for upstream TLS.",
				Computed:            true,
			},
			"tls_protocol_versions": dschema.SetAttribute{
				MarkdownDescription: "TLS protocol versions allowed for upstream connections.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"tls_session_reuse": dschema.BoolAttribute{
				MarkdownDescription: "Whether TLS session reuse is enabled.",
				Computed:            true,
			},
			"tls_trusted_certificate": dschema.StringAttribute{
				MarkdownDescription: "Trusted certificate store UUID used to verify upstream TLS connections.",
				Computed:            true,
			},
			"tls_verify": dschema.BoolAttribute{
				MarkdownDescription: "Whether upstream TLS certificates are verified.",
				Computed:            true,
			},
			"tls_verify_depth": dschema.StringAttribute{
				MarkdownDescription: "Maximum verification depth for upstream TLS certificates.",
				Computed:            true,
			},
			"store": dschema.BoolAttribute{
				MarkdownDescription: "Whether upstream state is persisted in the cache store.",
				Computed:            true,
			},
		},
	}
}
