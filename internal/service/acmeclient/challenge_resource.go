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
var _ resource.Resource = &acmeclientChallengeResource{}
var _ resource.ResourceWithConfigure = &acmeclientChallengeResource{}
var _ resource.ResourceWithImportState = &acmeclientChallengeResource{}

func newACMEClientChallengeResource() resource.Resource {
	return &acmeclientChallengeResource{}
}

// acmeclientChallengeResource defines the resource implementation.
type acmeclientChallengeResource struct {
	client opnsense.Client
}

func (r *acmeclientChallengeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acmeclient_challenge"
}

func (r *acmeclientChallengeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = acmeclientChallengeResourceSchema()
}

func (r *acmeclientChallengeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *acmeclientChallengeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *acmeclientChallengeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	validation, err := data.toValidation(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to prepare acmeclient challenge payload, got error: %s", err))
		return
	}

	result, err := r.client.Acmeclient().ACMEClientAddChallengeType(ctx, validation)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create acmeclient challenge, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("create ACME client challenge", result))
		return
	}

	if result != nil && result.UUID != "" {
		data.Id = types.StringValue(result.UUID)
	}

	if data.Id.IsNull() || data.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", "API did not return an identifier for the new acmeclient challenge.")
		return
	}

	challengeModel, err := fetchChallengeModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient challenge after create, got error: %s", err))
		return
	}

	challengeModel.Id = data.Id
	mergeChallengeParameters(ctx, &challengeModel, data)

	tflog.Trace(ctx, "created acmeclient challenge", map[string]any{
		"id": challengeModel.Id.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &challengeModel)...)
}

func (r *acmeclientChallengeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *acmeclientChallengeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	challengeModel, err := fetchChallengeModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, "acmeclient challenge not present in remote, removing from state", map[string]any{
				"id": data.Id.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient challenge, got error: %s", err))
		return
	}

	challengeModel.Id = data.Id
	mergeChallengeParameters(ctx, &challengeModel, data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &challengeModel)...)
}

func (r *acmeclientChallengeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *acmeclientChallengeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	validation, err := data.toValidation(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to prepare acmeclient challenge payload, got error: %s", err))
		return
	}

	result, err := r.client.Acmeclient().ACMEClientEditChallengeType(ctx, data.Id.ValueString(), validation)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update acmeclient challenge, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("update ACME client challenge", result))
		return
	}

	challengeModel, err := fetchChallengeModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient challenge after update, got error: %s", err))
		return
	}

	challengeModel.Id = data.Id
	mergeChallengeParameters(ctx, &challengeModel, data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &challengeModel)...)
}

func (r *acmeclientChallengeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *acmeclientChallengeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Acmeclient().ACMEClientDeleteChallengeType(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete acmeclient challenge, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("delete ACME client challenge", result))
		return
	}
}

func (r *acmeclientChallengeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func mergeChallengeParameters(ctx context.Context, state *acmeclientChallengeResourceModel, plan *acmeclientChallengeResourceModel) {
	if state == nil || plan == nil {
		return
	}

	var planMap map[string]string
	if plan.Parameters.IsNull() || plan.Parameters.IsUnknown() {
		return
	}

	if err := plan.Parameters.ElementsAs(ctx, &planMap, false); err != nil {
		return
	}

	var stateMap map[string]string
	if !state.Parameters.IsNull() && !state.Parameters.IsUnknown() {
		if err := state.Parameters.ElementsAs(ctx, &stateMap, false); err != nil {
			stateMap = map[string]string{}
		}
	}

	if stateMap == nil {
		stateMap = map[string]string{}
	}

	changed := false
	for key, value := range planMap {
		if existing, ok := stateMap[key]; !ok || existing == "" {
			stateMap[key] = value
			changed = true
		}
	}

	if changed {
		state.Parameters = stringMapToTypesMap(stateMap)
	}
}
