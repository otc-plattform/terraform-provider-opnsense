package haproxy

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = &haproxyObjectDataSource{}
var _ datasource.DataSourceWithConfigure = &haproxyObjectDataSource{}

func newHAProxyBackendDataSource() datasource.DataSource {
	return &haproxyObjectDataSource{kind: haproxyBackendKind}
}

func newHAProxyFrontendDataSource() datasource.DataSource {
	return &haproxyObjectDataSource{kind: haproxyFrontendKind}
}

func newHAProxyACLDataSource() datasource.DataSource {
	return &haproxyObjectDataSource{kind: haproxyACLKind}
}

func newHAProxyActionDataSource() datasource.DataSource {
	return &haproxyObjectDataSource{kind: haproxyActionKind}
}

type haproxyObjectDataSource struct {
	client opnsense.Client
	kind   haproxyObjectKind
}

func (d *haproxyObjectDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + d.kind.typeSuffix
}

func (d *haproxyObjectDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = haproxyObjectDataSourceSchema(d.kind.dataSourceDesc)
}

func (d *haproxyObjectDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *haproxyObjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *haproxyObjectResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() || data.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("%s data source requires a valid id.", d.kind.summaryName))
		return
	}

	object, err := d.kind.get(ctx, d.client.HAProxy(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read %s, got error: %s", d.kind.summaryName, err))
		return
	}

	model := objectModelFromAPI(ctx, data.Id.ValueString(), object)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
