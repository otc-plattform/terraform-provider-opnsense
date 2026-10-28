package haproxy

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = &healthMonitorDataSource{}
var _ datasource.DataSourceWithConfigure = &healthMonitorDataSource{}

func newHAProxyHealthMonitorDataSource() datasource.DataSource {
	return &healthMonitorDataSource{}
}

type healthMonitorDataSource struct {
	client opnsense.Client
}

func (d *healthMonitorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_haproxy_health_monitor"
}

func (d *healthMonitorDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = healthMonitorDataSourceSchema()
}

func (d *healthMonitorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	d.client = opnsense.NewClient(apiClient)
}

func (d *healthMonitorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *healthMonitorResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() || data.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", "HAProxy health monitor data source requires a valid id.")
		return
	}
	model, err := fetchHealthMonitorModel(ctx, d.client.HAProxy(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy health monitor, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
