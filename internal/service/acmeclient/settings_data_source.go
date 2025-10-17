package acmeclient

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &acmeclientSettingsDataSource{}
var _ datasource.DataSourceWithConfigure = &acmeclientSettingsDataSource{}

func newACMEClientSettingsDataSource() datasource.DataSource {
	return &acmeclientSettingsDataSource{}
}

// acmeclientSettingsDataSource defines the data source implementation.
type acmeclientSettingsDataSource struct {
	client opnsense.Client
}

func (d *acmeclientSettingsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acmeclient_settings"
}

func (d *acmeclientSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = acmeclientSettingsDataSourceSchema()
}

func (d *acmeclientSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *acmeclientSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *acmeclientSettingsResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get acmeclientSettings from OPNsense unbound API
	settings, err := d.client.Acmeclient().ACMEClientGetSettings(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient settings, got error: %s", err))
		return
	}

	// Convert OPNsense struct to TF schema
	resourceModel := acmeclientSettingsResourceModel{
		Enabled: types.BoolValue(tools.StringToBool(settings.ACMEClient.Settings.Enabled)),

		// TODO: convert rest of fields
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &resourceModel)...)
}
