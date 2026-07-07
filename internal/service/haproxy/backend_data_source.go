package haproxy

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = &backendDataSource{}
var _ datasource.DataSourceWithConfigure = &backendDataSource{}

func newHAProxyBackendDataSource() datasource.DataSource {
	return &backendDataSource{}
}

type backendDataSource struct {
	client opnsense.Client
}

func (d *backendDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_haproxy_backend"
}

func (d *backendDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = backendDataSourceSchema()
}

func (d *backendDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil { return }
	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	d.client = opnsense.NewClient(apiClient)
}

func (d *backendDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *backendResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() || data.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", "HAProxy backend data source requires a valid id."); return
	}
	model, err := fetchBackendModel(ctx, d.client.HAProxy(), data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy backend, got error: %s", err)); return }
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
