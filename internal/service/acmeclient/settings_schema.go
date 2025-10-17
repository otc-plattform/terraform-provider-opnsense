package acmeclient

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// acmeclientSettingsResourceModel describes the resource data model.
type acmeclientSettingsResourceModel struct {
	Enabled            types.Bool   `tfsdk:"enabled"`
	AutoRenewal        types.Bool   `tfsdk:"auto_renewal"`
	HAProxyIntegration types.Bool   `tfsdk:"haproxy_integration"`
	LogLevel           types.String `tfsdk:"log_level"`
	ShowIntro          types.Bool   `tfsdk:"show_intro"`
	ChallengePort      types.Int64  `tfsdk:"challenge_port"`
	TLSChallengePort   types.Int64  `tfsdk:"tls_challenge_port"`
	RestartTimeout     types.Int64  `tfsdk:"restart_timeout"`
}

func acmeclientSettingsResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense acmeclient plugin settings.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the ACME client.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"auto_renewal": schema.BoolAttribute{
				MarkdownDescription: "Automatically renew certificates.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"haproxy_integration": schema.BoolAttribute{
				MarkdownDescription: "Enable HAProxy integration.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"log_level": schema.StringAttribute{
				MarkdownDescription: "Log level for ACME client (`debug`, `info`, `warning`, `error`).",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("info"),
				Validators: []validator.String{
					stringvalidator.OneOf("debug", "info", "warning", "error"),
				},
			},
			"show_intro": schema.BoolAttribute{
				MarkdownDescription: "Show the introductory/help text in the UI.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"challenge_port": schema.Int64Attribute{
				MarkdownDescription: "Port used for HTTP-01 challenge. Defaults to 80.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(80),
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"tls_challenge_port": schema.Int64Attribute{
				MarkdownDescription: "Port used for TLS-ALPN-01 challenge. Defaults to 443.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(443),
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"restart_timeout": schema.Int64Attribute{
				MarkdownDescription: "Seconds to wait for service restart operations. Defaults to 60.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(60),
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
		},
	}
}

func acmeclientSettingsDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense acmeclient plugin settings.",
		Attributes: map[string]dschema.Attribute{
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether the ACME client is enabled.",
				Computed:            true,
			},
			"auto_renewal": dschema.BoolAttribute{
				MarkdownDescription: "Whether automatic renewal is enabled.",
				Computed:            true,
			},
			"haproxy_integration": dschema.BoolAttribute{
				MarkdownDescription: "Whether HAProxy integration is enabled.",
				Computed:            true,
			},
			"log_level": dschema.StringAttribute{
				MarkdownDescription: "Log level for the ACME client.",
				Computed:            true,
			},
			"show_intro": dschema.BoolAttribute{
				MarkdownDescription: "Whether the introductory/help text is shown.",
				Computed:            true,
			},
			"challenge_port": dschema.Int64Attribute{
				MarkdownDescription: "Port used for HTTP-01 challenge.",
				Computed:            true,
			},
			"tls_challenge_port": dschema.Int64Attribute{
				MarkdownDescription: "Port used for TLS-ALPN-01 challenge.",
				Computed:            true,
			},
			"restart_timeout": dschema.Int64Attribute{
				MarkdownDescription: "Seconds to wait for service restart operations.",
				Computed:            true,
			},
		},
	}
}
