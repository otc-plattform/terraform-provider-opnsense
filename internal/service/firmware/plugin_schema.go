package firmware

import (
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// firmwarePluginResourceModel describes the resource data model.
type firmwarePluginResourceModel struct {
	Installed types.Bool   `tfsdk:"installed"`
	Package   types.String `tfsdk:"package"`
}

func firmwarePluginResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense firmware plugin.",
		Attributes: map[string]schema.Attribute{
			"installed": schema.BoolAttribute{
				MarkdownDescription: "Whether this plugin is installed.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"package": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the package.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func firmwarePluginDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense firmware plugin.",
		Attributes: map[string]dschema.Attribute{
			"package": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Name of the package.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"installed": dschema.BoolAttribute{
				MarkdownDescription: "Whether this firmware job is enabled.",
				Computed:            true,
			},
		},
	}
}
