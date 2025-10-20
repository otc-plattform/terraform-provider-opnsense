package nginx

import (
	"context"
	"errors"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = &nginxUpstreamDataSource{}
var _ datasource.DataSourceWithConfigure = &nginxUpstreamDataSource{}

func newNginxUpstreamDataSource() datasource.DataSource {
	return &nginxUpstreamDataSource{}
}

type nginxUpstreamDataSource struct {
	client opnsense.Client
}

func (d *nginxUpstreamDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nginx_upstream"
}

func (d *nginxUpstreamDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = nginxUpstreamDataSourceSchema()
}

func (d *nginxUpstreamDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *nginxUpstreamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *nginxUpstreamResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil {
		resp.Diagnostics.AddError("Client Error", "Failed to decode nginx upstream data source configuration.")
		return
	}

	if data.Id.IsUnknown() {
		return
	}

	if data.Id.IsNull() || data.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", "Nginx upstream data source requires a valid id.")
		return
	}

	model, err := fetchUpstreamModel(ctx, d.client.Nginx(), data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Nginx upstream with ID %s not found.", data.Id.ValueString()))
			return
		}

		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx upstream, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
