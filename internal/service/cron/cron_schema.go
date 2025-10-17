package cron

import (
	"github.com/browningluke/opnsense-go/pkg/cron"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// cronResourceModel describes the resource data model.
type cronResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Minutes     types.String `tfsdk:"minutes"`
	Hours       types.String `tfsdk:"hours"`
	Days        types.String `tfsdk:"days"`
	Months      types.String `tfsdk:"months"`
	Weekdays    types.String `tfsdk:"weekdays"`
	Who         types.String `tfsdk:"who"`
	Command     types.String `tfsdk:"command"`
	Parameters  types.String `tfsdk:"parameters"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func cronResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense cron job.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether this cron job is enabled. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},

			// Standard 5-field cron schedule strings.
			"minutes": schema.StringAttribute{
				MarkdownDescription: "Cron minutes field (e.g., `*`, `0`, `*/5`, `0,30`). Defaults to `*`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("*"),
			},
			"hours": schema.StringAttribute{
				MarkdownDescription: "Cron hours field (e.g., `*`, `0`, `0-23`, `*/2`). Defaults to `*`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("*"),
			},
			"days": schema.StringAttribute{
				MarkdownDescription: "Cron day-of-month field (e.g., `*`, `1`, `1-31`). Defaults to `*`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("*"),
			},
			"months": schema.StringAttribute{
				MarkdownDescription: "Cron month field (e.g., `*`, `1`, `1-12`). Defaults to `*`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("*"),
			},
			"weekdays": schema.StringAttribute{
				MarkdownDescription: "Cron day-of-week field (e.g., `*`, `0`, `1-5`, `mon-fri`). Defaults to `*`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("*"),
			},

			"who": schema.StringAttribute{
				MarkdownDescription: "User to run the command as. Defaults to `root`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("root"),
			},

			"command": schema.StringAttribute{
				MarkdownDescription: "Command or action to execute (e.g., script path or named action).",
				Required:            true,
			},

			"parameters": schema.StringAttribute{
				MarkdownDescription: "Optional parameters passed to the command.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},

			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description for your reference.",
				Optional:            true,
			},

			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the cron job.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func cronDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense cron job by UUID.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the cron job.",
				Required:            true,
			},
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether this cron job is enabled.",
				Computed:            true,
			},
			"minutes": dschema.StringAttribute{
				MarkdownDescription: "Cron minutes field.",
				Computed:            true,
			},
			"hours": dschema.StringAttribute{
				MarkdownDescription: "Cron hours field.",
				Computed:            true,
			},
			"days": dschema.StringAttribute{
				MarkdownDescription: "Cron day-of-month field.",
				Computed:            true,
			},
			"months": dschema.StringAttribute{
				MarkdownDescription: "Cron month field.",
				Computed:            true,
			},
			"weekdays": dschema.StringAttribute{
				MarkdownDescription: "Cron day-of-week field.",
				Computed:            true,
			},
			"who": dschema.StringAttribute{
				MarkdownDescription: "User to run the command as.",
				Computed:            true,
			},
			"command": dschema.StringAttribute{
				MarkdownDescription: "Command or action to execute.",
				Computed:            true,
			},
			"parameters": dschema.StringAttribute{
				MarkdownDescription: "Optional parameters passed to the command.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description for your reference.",
				Computed:            true,
			},
		},
	}
}

func convertCRONSchemaToStruct(d *cronResourceModel) (*cron.Item, error) {
	return &cron.Item{
		UUID:        d.Id.ValueString(),
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Minutes:     d.Minutes.ValueString(),
		Hours:       d.Hours.ValueString(),
		Days:        d.Days.ValueString(),
		Months:      d.Months.ValueString(),
		Weekdays:    d.Weekdays.ValueString(),
		Who:         d.Who.ValueString(),
		Command:     d.Command.ValueString(),
		Parameters:  d.Parameters.ValueString(),
		Description: d.Description.ValueString(),
	}, nil
}

func convertCRONStructToSchema(d *cron.Item) (*cronResourceModel, error) {
	model := &cronResourceModel{
		Id:          types.StringValue(d.UUID),
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Minutes:     types.StringValue(d.Minutes),
		Hours:       types.StringValue(d.Hours),
		Days:        types.StringValue(d.Days),
		Months:      types.StringValue(d.Months),
		Weekdays:    types.StringValue(d.Weekdays),
		Who:         types.StringValue(d.Who),
		Command:     types.StringValue(d.Command),
		Parameters:  types.StringValue(d.Parameters),
		Description: types.StringValue(d.Description),
	}
	return model, nil
}
