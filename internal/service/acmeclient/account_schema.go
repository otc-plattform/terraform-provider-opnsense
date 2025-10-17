package acmeclient

import (
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// acmeclientAccountResourceModel describes the resource data model.
type acmeclientAccountResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	Email            types.String `tfsdk:"email"`
	CA               types.String `tfsdk:"ca"`
	CustomCA         types.String `tfsdk:"custom_ca"`
	EABKID           types.String `tfsdk:"eab_kid"`
	EABHMAC          types.String `tfsdk:"eab_hmac"`
	Key              types.String `tfsdk:"key"`
	StatusCode       types.String `tfsdk:"status_code"`
	StatusLastUpdate types.String `tfsdk:"status_last_update"`
}

func acmeclientAccountResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage ACME client accounts on OPNsense.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID of the ACME client account.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the ACME account is enabled.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Display name of the ACME account.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description for the ACME account.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"email": schema.StringAttribute{
				MarkdownDescription: "Contact email address used for ACME registration.",
				Required:            true,
			},
			"ca": schema.StringAttribute{
				MarkdownDescription: "Certification Authority identifier (for example `letsencrypt`).",
				Required:            true,
			},
			"custom_ca": schema.StringAttribute{
				MarkdownDescription: "Custom CA configuration (used when `ca` is `custom`).",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"eab_kid": schema.StringAttribute{
				MarkdownDescription: "External Account Binding key identifier, if required by the CA.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"eab_hmac": schema.StringAttribute{
				MarkdownDescription: "External Account Binding HMAC key, if required by the CA.",
				Optional:            true,
				Sensitive:           true,
			},
			"key": schema.StringAttribute{
				MarkdownDescription: "Raw ACME account key material.",
				Computed:            true,
				Sensitive:           true,
			},
			"status_code": schema.StringAttribute{
				MarkdownDescription: "Latest ACME account status code reported by OPNsense.",
				Computed:            true,
			},
			"status_last_update": schema.StringAttribute{
				MarkdownDescription: "Timestamp of the most recent ACME account status update.",
				Computed:            true,
			},
		},
	}
}

func acmeclientAccountDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read ACME client account details from OPNsense.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the ACME client account.",
				Required:            true,
			},
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether the ACME account is enabled.",
				Computed:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "Display name of the ACME account.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description of the ACME account.",
				Computed:            true,
			},
			"email": dschema.StringAttribute{
				MarkdownDescription: "Contact email address used for ACME registration.",
				Computed:            true,
			},
			"ca": dschema.StringAttribute{
				MarkdownDescription: "Certification Authority identifier.",
				Computed:            true,
			},
			"custom_ca": dschema.StringAttribute{
				MarkdownDescription: "Custom CA configuration value.",
				Computed:            true,
			},
			"eab_kid": dschema.StringAttribute{
				MarkdownDescription: "External Account Binding key identifier.",
				Computed:            true,
			},
			"eab_hmac": dschema.StringAttribute{
				MarkdownDescription: "External Account Binding HMAC key.",
				Computed:            true,
				Sensitive:           true,
			},
			"key": dschema.StringAttribute{
				MarkdownDescription: "Raw ACME account key material.",
				Computed:            true,
				Sensitive:           true,
			},
			"status_code": dschema.StringAttribute{
				MarkdownDescription: "Latest ACME account status code reported by OPNsense.",
				Computed:            true,
			},
			"status_last_update": dschema.StringAttribute{
				MarkdownDescription: "Timestamp of the most recent ACME account status update.",
				Computed:            true,
			},
		},
	}
}
