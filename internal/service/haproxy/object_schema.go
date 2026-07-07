package haproxy

import (
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type haproxyObjectResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Config types.Map    `tfsdk:"config"`
}

func haproxyObjectResourceSchema(description string) schema.Schema {
	return schema.Schema{
		MarkdownDescription: description,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID assigned by OPNsense.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"config": schema.MapAttribute{
				MarkdownDescription: "Raw HAProxy settings payload as string values. Keys match OPNsense API field names.",
				ElementType:         types.StringType,
				Required:            true,
			},
		},
	}
}

func haproxyObjectDataSourceSchema(description string) dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: description,
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID assigned by OPNsense.",
				Required:            true,
			},
			"config": dschema.MapAttribute{
				MarkdownDescription: "Raw HAProxy settings payload as string values. Keys match OPNsense API field names.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}
