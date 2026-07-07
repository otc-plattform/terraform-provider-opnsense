package haproxy

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &haproxyObjectResource{}
var _ resource.ResourceWithConfigure = &haproxyObjectResource{}
var _ resource.ResourceWithImportState = &haproxyObjectResource{}

func newHAProxyBackendResource() resource.Resource {
	return &haproxyObjectResource{kind: haproxyBackendKind}
}

func newHAProxyFrontendResource() resource.Resource {
	return &haproxyObjectResource{kind: haproxyFrontendKind}
}

func newHAProxyACLResource() resource.Resource {
	return &haproxyObjectResource{kind: haproxyACLKind}
}

func newHAProxyActionResource() resource.Resource {
	return &haproxyObjectResource{kind: haproxyActionKind}
}

type haproxyObjectResource struct {
	client opnsense.Client
	kind   haproxyObjectKind
}

func (r *haproxyObjectResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.kind.typeSuffix
}

func (r *haproxyObjectResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = haproxyObjectResourceSchema(r.kind.resourceDesc)
}

func (r *haproxyObjectResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *haproxyObjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *haproxyObjectResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	object, err := objectFromModel(ctx, data)
	if err != nil {
		resp.Diagnostics.AddError("Config Error", err.Error())
		return
	}

	result, err := r.kind.add(ctx, r.client.HAProxy(), object)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s, got error: %s", r.kind.summaryName, err))
		return
	}
	if result == nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s: empty response received from API.", r.kind.summaryName))
		return
	}
	if result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("create "+r.kind.summaryName, result))
		return
	}

	id := result.UUID
	if id == "" {
		name := apiValueToString(object["name"])
		var found bool
		id, found, err = findObjectIDByName(ctx, r.kind, r.client.HAProxy(), name)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to find %s after create, got error: %s", r.kind.summaryName, err))
			return
		}
		if !found {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create %s: API response did not include a UUID and the object could not be found by name.", r.kind.summaryName))
			return
		}
	}

	model, err := r.fetchModel(ctx, id)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read %s after create, got error: %s", r.kind.summaryName, err))
		return
	}

	tflog.Trace(ctx, "created HAProxy object", map[string]any{"kind": r.kind.typeSuffix, "id": id})
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *haproxyObjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *haproxyObjectResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.State.RemoveResource(ctx)
		return
	}

	model, err := r.fetchModel(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read %s, got error: %s", r.kind.summaryName, err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *haproxyObjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *haproxyObjectResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("%s update requires a valid id.", r.kind.summaryName))
		return
	}

	object, err := objectFromModel(ctx, data)
	if err != nil {
		resp.Diagnostics.AddError("Config Error", err.Error())
		return
	}

	result, err := r.kind.edit(ctx, r.client.HAProxy(), data.Id.ValueString(), object)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update %s, got error: %s", r.kind.summaryName, err))
		return
	}
	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("update "+r.kind.summaryName, result))
		return
	}

	model, err := r.fetchModel(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read %s after update, got error: %s", r.kind.summaryName, err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *haproxyObjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *haproxyObjectResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		return
	}

	result, err := r.kind.delete(ctx, r.client.HAProxy(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete %s, got error: %s", r.kind.summaryName, err))
		return
	}
	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("delete "+r.kind.summaryName, result))
		return
	}
}

func (r *haproxyObjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *haproxyObjectResource) fetchModel(ctx context.Context, id string) (haproxyObjectResourceModel, error) {
	object, err := r.kind.get(ctx, r.client.HAProxy(), id)
	if err != nil {
		return haproxyObjectResourceModel{}, err
	}
	return objectModelFromAPI(ctx, id, object), nil
}
