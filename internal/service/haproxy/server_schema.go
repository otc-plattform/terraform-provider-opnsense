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

type haproxyServerResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	InternalId           types.String `tfsdk:"internal_id"`
	Enabled              types.Bool   `tfsdk:"enabled"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	Type                 types.String `tfsdk:"type"`
	Address              types.String `tfsdk:"address"`
	ServiceName          types.String `tfsdk:"service_name"`
	Number               types.String `tfsdk:"number"`
	LinkedResolver       types.String `tfsdk:"linked_resolver"`
	ResolverOpts         types.Set    `tfsdk:"resolver_opts"`
	UnixSocket           types.String `tfsdk:"unix_socket"`
	Port                 types.String `tfsdk:"port"`
	Mode                 types.String `tfsdk:"mode"`
	MultiplexerProtocol  types.String `tfsdk:"multiplexer_protocol"`
	ResolvePrefer        types.String `tfsdk:"resolve_prefer"`
	SSL                  types.Bool   `tfsdk:"ssl"`
	SSLSNI               types.String `tfsdk:"ssl_sni"`
	SSLSNIExpr           types.String `tfsdk:"ssl_sni_expr"`
	SSLVerify            types.Bool   `tfsdk:"ssl_verify"`
	SSLCA                types.String `tfsdk:"ssl_ca"`
	SSLCRL               types.String `tfsdk:"ssl_crl"`
	SSLClientCertificate types.String `tfsdk:"ssl_client_certificate"`
	MaxConnections       types.String `tfsdk:"max_connections"`
	Weight               types.String `tfsdk:"weight"`
	CheckInterval        types.String `tfsdk:"check_interval"`
	CheckDownInterval    types.String `tfsdk:"check_down_interval"`
	Checkport            types.String `tfsdk:"checkport"`
	Source               types.String `tfsdk:"source"`
	Advanced             types.String `tfsdk:"advanced"`
}

func haproxyServerResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense HAProxy real server.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID of the HAProxy real server.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"internal_id": schema.StringAttribute{
				MarkdownDescription: "Internal HAProxy server id assigned by OPNsense.",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this HAProxy real server.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Unique HAProxy real server name.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"description":     stringResourceAttr("Optional description."),
			"type":            stringEnumResourceAttr("Server type.", "static", "static", "template", "unix"),
			"address":         stringResourceAttr("Server address for static or template servers."),
			"service_name":    stringResourceAttr("Service name for template servers."),
			"number":          stringResourceAttr("Number of servers for template servers."),
			"linked_resolver": stringResourceAttr("Linked HAProxy resolver UUID."),
			"resolver_opts": schema.SetAttribute{
				MarkdownDescription: "Resolver options.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"unix_socket":          stringResourceAttr("Linked Unix socket UUID."),
			"port":                 stringResourceAttr("Server port."),
			"mode":                 stringEnumResourceAttr("Server mode.", "active", "", "active", "backup", "disabled"),
			"multiplexer_protocol": stringEnumResourceAttr("HAProxy multiplexer protocol.", "unspecified", "", "unspecified", "fcgi", "h2", "h1"),
			"resolve_prefer":       stringEnumResourceAttr("Preferred resolved address family.", "", "", "ipv4", "ipv6"),
			"ssl": schema.BoolAttribute{
				MarkdownDescription: "Enable SSL when connecting to this server.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_sni":      stringResourceAttr("Static SNI value for SSL connections."),
			"ssl_sni_expr": stringResourceAttr("HAProxy expression used to derive SNI."),
			"ssl_verify": schema.BoolAttribute{
				MarkdownDescription: "Verify the server certificate.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ssl_ca":                 stringResourceAttr("CA certificate UUID for server verification."),
			"ssl_crl":                stringResourceAttr("CRL UUID for server verification."),
			"ssl_client_certificate": stringResourceAttr("Client certificate UUID for SSL connections."),
			"max_connections":        stringResourceAttr("Maximum concurrent connections to this server."),
			"weight":                 stringResourceAttr("Server weight."),
			"check_interval":         stringResourceAttr("Health check interval."),
			"check_down_interval":    stringResourceAttr("Health check interval while down."),
			"checkport":              stringResourceAttr("Health check port override."),
			"source":                 stringResourceAttr("Source address for connections to this server."),
			"advanced":               stringResourceAttr("Advanced HAProxy server options."),
		},
	}
}

func haproxyServerDataSourceSchema() dschema.Schema {
	attrs := map[string]dschema.Attribute{
		"id": dschema.StringAttribute{
			MarkdownDescription: "UUID of the HAProxy real server.",
			Required:            true,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
		},
		"internal_id":            dschema.StringAttribute{MarkdownDescription: "Internal HAProxy server id assigned by OPNsense.", Computed: true},
		"enabled":                dschema.BoolAttribute{MarkdownDescription: "Whether this HAProxy real server is enabled.", Computed: true},
		"name":                   dschema.StringAttribute{MarkdownDescription: "HAProxy real server name.", Computed: true},
		"description":            dschema.StringAttribute{MarkdownDescription: "Optional description.", Computed: true},
		"type":                   dschema.StringAttribute{MarkdownDescription: "Server type.", Computed: true},
		"address":                dschema.StringAttribute{MarkdownDescription: "Server address.", Computed: true},
		"service_name":           dschema.StringAttribute{MarkdownDescription: "Service name for template servers.", Computed: true},
		"number":                 dschema.StringAttribute{MarkdownDescription: "Number of servers for template servers.", Computed: true},
		"linked_resolver":        dschema.StringAttribute{MarkdownDescription: "Linked HAProxy resolver UUID.", Computed: true},
		"resolver_opts":          dschema.SetAttribute{MarkdownDescription: "Resolver options.", ElementType: types.StringType, Computed: true},
		"unix_socket":            dschema.StringAttribute{MarkdownDescription: "Linked Unix socket UUID.", Computed: true},
		"port":                   dschema.StringAttribute{MarkdownDescription: "Server port.", Computed: true},
		"mode":                   dschema.StringAttribute{MarkdownDescription: "Server mode.", Computed: true},
		"multiplexer_protocol":   dschema.StringAttribute{MarkdownDescription: "HAProxy multiplexer protocol.", Computed: true},
		"resolve_prefer":         dschema.StringAttribute{MarkdownDescription: "Preferred resolved address family.", Computed: true},
		"ssl":                    dschema.BoolAttribute{MarkdownDescription: "Whether SSL is enabled.", Computed: true},
		"ssl_sni":                dschema.StringAttribute{MarkdownDescription: "Static SNI value.", Computed: true},
		"ssl_sni_expr":           dschema.StringAttribute{MarkdownDescription: "SNI expression.", Computed: true},
		"ssl_verify":             dschema.BoolAttribute{MarkdownDescription: "Whether server certificate verification is enabled.", Computed: true},
		"ssl_ca":                 dschema.StringAttribute{MarkdownDescription: "CA certificate UUID.", Computed: true},
		"ssl_crl":                dschema.StringAttribute{MarkdownDescription: "CRL UUID.", Computed: true},
		"ssl_client_certificate": dschema.StringAttribute{MarkdownDescription: "Client certificate UUID.", Computed: true},
		"max_connections":        dschema.StringAttribute{MarkdownDescription: "Maximum concurrent connections.", Computed: true},
		"weight":                 dschema.StringAttribute{MarkdownDescription: "Server weight.", Computed: true},
		"check_interval":         dschema.StringAttribute{MarkdownDescription: "Health check interval.", Computed: true},
		"check_down_interval":    dschema.StringAttribute{MarkdownDescription: "Health check interval while down.", Computed: true},
		"checkport":              dschema.StringAttribute{MarkdownDescription: "Health check port override.", Computed: true},
		"source":                 dschema.StringAttribute{MarkdownDescription: "Source address.", Computed: true},
		"advanced":               dschema.StringAttribute{MarkdownDescription: "Advanced HAProxy server options.", Computed: true},
	}

	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense HAProxy real server.",
		Attributes:          attrs,
	}
}

func stringResourceAttr(description string) schema.StringAttribute {
	return schema.StringAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
	}
}

func stringEnumResourceAttr(description, defaultValue string, values ...string) schema.StringAttribute {
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
