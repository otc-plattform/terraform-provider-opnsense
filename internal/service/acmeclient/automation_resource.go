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

var _ resource.Resource = &acmeclientAutomationResource{}
var _ resource.ResourceWithConfigure = &acmeclientAutomationResource{}
var _ resource.ResourceWithImportState = &acmeclientAutomationResource{}

func newACMEClientAutomationResource() resource.Resource {
	return &acmeclientAutomationResource{}
}

type acmeclientAutomationResource struct {
	client opnsense.Client
}

func (r *acmeclientAutomationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acmeclient_automation"
}

func (r *acmeclientAutomationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = acmeclientAutomationResourceSchema()
}

func (r *acmeclientAutomationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *acmeclientAutomationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *acmeclientAutomationResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	action := data.toAction()
	result, err := r.client.Acmeclient().ACMEClientAddAutomation(ctx, action)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create acmeclient automation, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("create ACME client automation", result))
		return
	}

	if result != nil && result.UUID != "" {
		data.Id = types.StringValue(result.UUID)
	}

	if data.Id.IsNull() || data.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", "API did not return an identifier for the new automation.")
		return
	}

	model, err := fetchAutomationModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read automation after create, got error: %s", err))
		return
	}

	model.Id = data.Id
	mergeAutomationSensitiveFields(&model, data)

	tflog.Trace(ctx, "created acmeclient automation", map[string]any{
		"id": model.Id.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *acmeclientAutomationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *acmeclientAutomationResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	model, err := fetchAutomationModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		var notFound *errs.NotFoundError
		if errors.As(err, &notFound) {
			tflog.Warn(ctx, "acmeclient automation not present in remote, removing from state", map[string]any{
				"id": data.Id.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient automation, got error: %s", err))
		return
	}

	model.Id = data.Id
	mergeAutomationSensitiveFields(&model, data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *acmeclientAutomationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *acmeclientAutomationResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	action := data.toAction()
	result, err := r.client.Acmeclient().ACMEClientEditAutomation(ctx, data.Id.ValueString(), action)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update acmeclient automation, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("update ACME client automation", result))
		return
	}

	model, err := fetchAutomationModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient automation after update, got error: %s", err))
		return
	}

	model.Id = data.Id
	mergeAutomationSensitiveFields(&model, data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *acmeclientAutomationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *acmeclientAutomationResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Acmeclient().ACMEClientDeleteAutomation(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete acmeclient automation, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("delete ACME client automation", result))
		return
	}
}

func (r *acmeclientAutomationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
