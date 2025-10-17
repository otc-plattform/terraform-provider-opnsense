package nginx

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = &nginxSettingsDataSource{}
var _ datasource.DataSourceWithConfigure = &nginxSettingsDataSource{}

func newNginxSettingsDataSource() datasource.DataSource {
	return &nginxSettingsDataSource{}
}

type nginxSettingsDataSource struct {
	client opnsense.Client
}

func (d *nginxSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nginx_settings"
}

func (d *nginxSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = nginxSettingsDataSourceSchema()
}

func (d *nginxSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = opnsense.NewClient(apiClient)
}

func (d *nginxSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *nginxSettingsResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	settings, err := d.client.Nginx().NginxGetSettings(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx settings, got error: %s", err))
		return
	}

	resourceModel := settingsResponseToModel(settings)

	resp.Diagnostics.Append(resp.State.Set(ctx, &resourceModel)...)
}
