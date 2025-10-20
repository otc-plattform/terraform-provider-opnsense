package nginx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type nginxUpstreamServerResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Description types.String `tfsdk:"description"`
	Server      types.String `tfsdk:"server"`
	Port        types.String `tfsdk:"port"`
	Priority    types.String `tfsdk:"priority"`
	MaxConns    types.String `tfsdk:"max_conns"`
	MaxFails    types.String `tfsdk:"max_fails"`
	FailTimeout types.String `tfsdk:"fail_timeout"`
	NoUse       types.String `tfsdk:"no_use"`
}

func nginxUpstreamServerResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense Nginx upstream server entry.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for the upstream server entry.",
				Optional:            true,
			},
			"server": schema.StringAttribute{
				MarkdownDescription: "Target server hostname or IP address.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"port": schema.StringAttribute{
				MarkdownDescription: "Target server port.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"priority": schema.StringAttribute{
				MarkdownDescription: "Priority value influencing load balancing weight.",
				Optional:            true,
			},
			"max_conns": schema.StringAttribute{
				MarkdownDescription: "Maximum concurrent connections to this upstream server.",
				Optional:            true,
			},
			"max_fails": schema.StringAttribute{
				MarkdownDescription: "Maximum number of failed attempts before marking server as unavailable.",
				Optional:            true,
			},
			"fail_timeout": schema.StringAttribute{
				MarkdownDescription: "Time in seconds after which a failed server is retried.",
				Optional:            true,
			},
			"no_use": schema.StringAttribute{
				MarkdownDescription: "Advanced upstream server mode (`down`, `backup`, or unset).",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "down", "backup"),
				},
			},
		},
	}
}

func nginxUpstreamServerDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense Nginx upstream server entry.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the upstream server entry.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description for the upstream server entry.",
				Computed:            true,
			},
			"server": dschema.StringAttribute{
				MarkdownDescription: "Target server hostname or IP address.",
				Computed:            true,
			},
			"port": dschema.StringAttribute{
				MarkdownDescription: "Target server port.",
				Computed:            true,
			},
			"priority": dschema.StringAttribute{
				MarkdownDescription: "Priority value influencing load balancing weight.",
				Computed:            true,
			},
			"max_conns": dschema.StringAttribute{
				MarkdownDescription: "Maximum concurrent connections to this upstream server.",
				Computed:            true,
			},
			"max_fails": dschema.StringAttribute{
				MarkdownDescription: "Maximum number of failed attempts before marking server as unavailable.",
				Computed:            true,
			},
			"fail_timeout": dschema.StringAttribute{
				MarkdownDescription: "Time in seconds after which a failed server is retried.",
				Computed:            true,
			},
			"no_use": dschema.StringAttribute{
				MarkdownDescription: "Advanced upstream server mode (`down`, `backup`, or unset).",
				Computed:            true,
			},
		},
	}
}
