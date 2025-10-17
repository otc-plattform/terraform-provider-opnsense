package acmeclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &acmeclientAccountResource{}
var _ resource.ResourceWithConfigure = &acmeclientAccountResource{}
var _ resource.ResourceWithImportState = &acmeclientAccountResource{}

func newACMEClientAccountResource() resource.Resource {
	return &acmeclientAccountResource{}
}

// acmeclientAccountResource defines the resource implementation.
type acmeclientAccountResource struct {
	client opnsense.Client
}

func (r *acmeclientAccountResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acmeclient_account"
}

func (r *acmeclientAccountResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = acmeclientAccountResourceSchema()
}

func (r *acmeclientAccountResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = opnsense.NewClient(apiClient)
}

func (r *acmeclientAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *acmeclientAccountResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := data.toAccountCreateRequest()
	result, err := r.client.Acmeclient().ACMEClientAddAccount(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create acmeclient account, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("create ACME client account", result))
		return
	}

	if result != nil && result.UUID != "" {
		data.Id = types.StringValue(result.UUID)
	}

	if data.Id.IsNull() || data.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", "API did not return an identifier for the new acmeclient account.")
		return
	}

	accountModel, err := fetchAccountModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient account after create, got error: %s", err))
		return
	}

	if accountModel.EABHMAC.IsNull() || accountModel.EABHMAC.ValueString() == "" {
		accountModel.EABHMAC = data.EABHMAC
	}

	tflog.Trace(ctx, "created acmeclient account", map[string]any{
		"id": accountModel.Id.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &accountModel)...)
}

func (r *acmeclientAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *acmeclientAccountResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountModel, err := fetchAccountModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, "acmeclient account not present in remote, removing from state", map[string]any{
				"id": data.Id.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient account, got error: %s", err))
		return
	}

	accountModel.Id = data.Id

	if accountModel.EABHMAC.IsNull() || accountModel.EABHMAC.ValueString() == "" {
		accountModel.EABHMAC = data.EABHMAC
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &accountModel)...)
}

func (r *acmeclientAccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *acmeclientAccountResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := data.toAccountEditRequest()
	result, err := r.client.Acmeclient().ACMEClientEditAccount(ctx, data.Id.ValueString(), request)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update acmeclient account, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("update ACME client account", result))
		return
	}

	accountModel, err := fetchAccountModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient account after update, got error: %s", err))
		return
	}

	accountModel.Id = data.Id

	if accountModel.EABHMAC.IsNull() || accountModel.EABHMAC.ValueString() == "" {
		accountModel.EABHMAC = data.EABHMAC
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &accountModel)...)
}

func (r *acmeclientAccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *acmeclientAccountResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Acmeclient().ACMEClientDeleteAccount(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete acmeclient account, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("delete ACME client account", result))
		return
	}
}

func (r *acmeclientAccountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
