package nginx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// nginxSettingsResourceModel describes the Terraform model for the general settings.
type nginxSettingsResourceModel struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	BanTTL  types.Int64 `tfsdk:"ban_ttl"`
}

func nginxSettingsResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage OPNsense Nginx plugin general settings.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the Nginx plugin service.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ban_ttl": schema.Int64Attribute{
				MarkdownDescription: "Ban duration in seconds before a blocked client is released.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
		},
	}
}

func nginxSettingsDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read OPNsense Nginx plugin general settings.",
		Attributes: map[string]dschema.Attribute{
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether the Nginx plugin service is enabled.",
				Computed:            true,
			},
			"ban_ttl": dschema.Int64Attribute{
				MarkdownDescription: "Ban duration in seconds before a blocked client is released.",
				Computed:            true,
			},
		},
	}
}
