package acmeclient

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// acmeclientChallengeResourceModel describes the challenge resource data model.
type acmeclientChallengeResourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Enabled                  types.Bool   `tfsdk:"enabled"`
	Name                     types.String `tfsdk:"name"`
	Description              types.String `tfsdk:"description"`
	Method                   types.String `tfsdk:"method"`
	HTTPService              types.String `tfsdk:"http_service"`
	HTTPOpnAutodiscovery     types.Bool   `tfsdk:"http_opn_autodiscovery"`
	HTTPOpnInterface         types.String `tfsdk:"http_opn_interface"`
	HTTPOpnIPAddresses       types.Set    `tfsdk:"http_opn_ipaddresses"`
	HTTPHAProxyInject        types.Bool   `tfsdk:"http_haproxy_inject"`
	HTTPHAProxyFrontends     types.Set    `tfsdk:"http_haproxy_frontends"`
	TLSALPNService           types.String `tfsdk:"tlsalpn_service"`
	TLSALPNAcmeAutodiscovery types.Bool   `tfsdk:"tlsalpn_acme_autodiscovery"`
	TLSALPNAcmeInterface     types.String `tfsdk:"tlsalpn_acme_interface"`
	TLSALPNAcmeIPAddresses   types.Set    `tfsdk:"tlsalpn_acme_ipaddresses"`
	DNSService               types.String `tfsdk:"dns_service"`
	DNSSleep                 types.Int64  `tfsdk:"dns_sleep"`
	Parameters               types.Map    `tfsdk:"parameters"`
}

func acmeclientChallengeResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage ACME client validation challenges on OPNsense.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID of the ACME challenge.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether this challenge is enabled.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Display name of the challenge.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description of the challenge.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"method": schema.StringAttribute{
				MarkdownDescription: "Validation method (`http01`, `dns01`, `tlsalpn01`).",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("http01", "dns01", "tlsalpn01"),
				},
			},
			"http_service": schema.StringAttribute{
				MarkdownDescription: "HTTP service integration when using http-01 (`opnsense`, `haproxy`, etc.).",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("opnsense"),
			},
			"http_opn_autodiscovery": schema.BoolAttribute{
				MarkdownDescription: "Automatically discover OPNsense interfaces for http-01.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_opn_interface": schema.StringAttribute{
				MarkdownDescription: "Specific OPNsense interface to use for http-01.",
				Optional:            true,
			},
			"http_opn_ipaddresses": schema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "Specific IPs to bind for http-01.",
				Optional:            true,
			},
			"http_haproxy_inject": schema.BoolAttribute{
				MarkdownDescription: "Whether to inject HTTP-01 validation responses into HAProxy.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_haproxy_frontends": schema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "HAProxy frontends to use when http-01 is integrated.",
				Optional:            true,
			},
			"tlsalpn_service": schema.StringAttribute{
				MarkdownDescription: "TLS-ALPN service integration (`acme`, `nginx`, etc.).",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("acme"),
			},
			"tlsalpn_acme_autodiscovery": schema.BoolAttribute{
				MarkdownDescription: "Automatically discover interfaces for TLS-ALPN.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"tlsalpn_acme_interface": schema.StringAttribute{
				MarkdownDescription: "Specific interface for TLS-ALPN validation.",
				Optional:            true,
			},
			"tlsalpn_acme_ipaddresses": schema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "Specific IPs for TLS-ALPN validation.",
				Optional:            true,
			},
			"dns_service": schema.StringAttribute{
				MarkdownDescription: "DNS provider integration identifier when using dns-01.",
				Optional:            true,
			},
			"dns_sleep": schema.Int64Attribute{
				MarkdownDescription: "Number of seconds to wait after updating DNS (dns-01).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
			},
			"parameters": schema.MapAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "Additional provider-specific parameters (exact keys as expected by OPNsense, e.g. `dns_cf_token`).",
				Optional:            true,
			},
		},
	}
}

func acmeclientChallengeDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read ACME client validation challenge details from OPNsense.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the ACME challenge.",
				Required:            true,
			},
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether this challenge is enabled.",
				Computed:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "Display name of the challenge.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description of the challenge.",
				Computed:            true,
			},
			"method": dschema.StringAttribute{
				MarkdownDescription: "Validation method.",
				Computed:            true,
			},
			"http_service": dschema.StringAttribute{
				MarkdownDescription: "HTTP service integration.",
				Computed:            true,
			},
			"http_opn_autodiscovery": dschema.BoolAttribute{
				MarkdownDescription: "Whether OPNsense interface autodiscovery is enabled for http-01.",
				Computed:            true,
			},
			"http_opn_interface": dschema.StringAttribute{
				MarkdownDescription: "Specific interface for http-01.",
				Computed:            true,
			},
			"http_opn_ipaddresses": dschema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "IP addresses for http-01.",
				Computed:            true,
			},
			"http_haproxy_inject": dschema.BoolAttribute{
				MarkdownDescription: "Whether HAProxy injection is enabled for http-01.",
				Computed:            true,
			},
			"http_haproxy_frontends": dschema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "HAProxy frontends used for http-01.",
				Computed:            true,
			},
			"tlsalpn_service": dschema.StringAttribute{
				MarkdownDescription: "TLS-ALPN service integration.",
				Computed:            true,
			},
			"tlsalpn_acme_autodiscovery": dschema.BoolAttribute{
				MarkdownDescription: "Whether TLS-ALPN autodiscovery is enabled.",
				Computed:            true,
			},
			"tlsalpn_acme_interface": dschema.StringAttribute{
				MarkdownDescription: "Interface used for TLS-ALPN.",
				Computed:            true,
			},
			"tlsalpn_acme_ipaddresses": dschema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "IP addresses used for TLS-ALPN.",
				Computed:            true,
			},
			"dns_service": dschema.StringAttribute{
				MarkdownDescription: "DNS provider identifier.",
				Computed:            true,
			},
			"dns_sleep": dschema.Int64Attribute{
				MarkdownDescription: "Wait time after DNS updates.",
				Computed:            true,
			},
			"parameters": dschema.MapAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "Provider-specific parameters.",
				Computed:            true,
			},
		},
	}
}
