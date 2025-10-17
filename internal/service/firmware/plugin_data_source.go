package firmware

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &firmwarePluginDataSource{}
var _ datasource.DataSourceWithConfigure = &firmwarePluginDataSource{}

func newFirmwarePluginDataSource() datasource.DataSource {
	return &firmwarePluginDataSource{}
}

// firmwarePluginDataSource defines the data source implementation.
type firmwarePluginDataSource struct {
	client opnsense.Client
}

func (d *firmwarePluginDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_firmware_plugin"
}

func (d *firmwarePluginDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = firmwarePluginDataSourceSchema()
}

func (d *firmwarePluginDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *opnsense.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = opnsense.NewClient(apiClient)
}

func (d *firmwarePluginDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *firmwarePluginResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get firmwarePlugin from OPNsense unbound API
	installed, err := isInstalled(ctx, d.client, data.Package.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read firewall firmwarePlugin, got error: %s", err))
		return
	}

	// Convert OPNsense struct to TF schema
	resourceModel := firmwarePluginResourceModel{
		Package:   types.StringValue(data.Package.String()),
		Installed: types.BoolValue(installed),
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &resourceModel)...)
}

func isInstalled(ctx context.Context, client opnsense.Client, pkg string) (bool, error) {
	resourceStruct, err := client.Firmware().FirmwareInfo(ctx)
	if err != nil {
		return false, fmt.Errorf("Unable to read firewall firmwarePlugin, got error: %s", err)
	}

	installed := false
	for i := range resourceStruct.Package {
		if resourceStruct.Package[i].Name == pkg {
			installed = resourceStruct.Package[i].Installed == "1"
			break
		}
	}

	return installed, nil
}
