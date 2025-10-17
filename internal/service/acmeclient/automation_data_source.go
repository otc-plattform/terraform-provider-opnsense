package acmeclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &acmeclientAutomationDataSource{}
var _ datasource.DataSourceWithConfigure = &acmeclientAutomationDataSource{}

func newACMEClientAutomationDataSource() datasource.DataSource {
	return &acmeclientAutomationDataSource{}
}

type acmeclientAutomationDataSource struct {
	client opnsense.Client
}

func (d *acmeclientAutomationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acmeclient_automation"
}

func (d *acmeclientAutomationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = acmeclientAutomationDataSourceSchema()
}

func (d *acmeclientAutomationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *acmeclientAutomationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *acmeclientAutomationResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil {
		resp.Diagnostics.AddError("Client Error", "Failed to decode ACME automation data source configuration.")
		return
	}

	if data.Id.IsUnknown() {
		return
	}

	if data.Id.IsNull() || data.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", "ACME automation data source requires a valid id.")
		return
	}

	model, err := fetchAutomationModel(ctx, d.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		var notFound *errs.NotFoundError
		if errors.As(err, &notFound) {
			resp.Diagnostics.AddError("Client Error",
				fmt.Sprintf("ACME client automation with ID %s not found.", data.Id.ValueString()))
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient automation, got error: %s", err))
		return
	}

	model.Id = types.StringValue(data.Id.ValueString())

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
