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

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &acmeclientChallengeDataSource{}
var _ datasource.DataSourceWithConfigure = &acmeclientChallengeDataSource{}

func newACMEClientChallengeDataSource() datasource.DataSource {
	return &acmeclientChallengeDataSource{}
}

type acmeclientChallengeDataSource struct {
	client opnsense.Client
}

func (d *acmeclientChallengeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acmeclient_challenge"
}

func (d *acmeclientChallengeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = acmeclientChallengeDataSourceSchema()
}

func (d *acmeclientChallengeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = opnsense.NewClient(apiClient)
}

func (d *acmeclientChallengeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *acmeclientChallengeResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil {
		resp.Diagnostics.AddError("Client Error", "Failed to decode ACME challenge data source configuration.")
		return
	}

	if data.Id.IsUnknown() {
		return
	}

	if data.Id.IsNull() || data.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", "ACME challenge data source requires a valid id.")
		return
	}

	model, err := fetchChallengeModel(ctx, d.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			resp.Diagnostics.AddError("Client Error",
				fmt.Sprintf("ACME client challenge with ID %s not found.", data.Id.ValueString()))
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient challenge, got error: %s", err))
		return
	}

	model.Id = types.StringValue(data.Id.ValueString())

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
