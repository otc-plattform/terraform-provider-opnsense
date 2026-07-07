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

var _ resource.Resource = &actionResource{}
var _ resource.ResourceWithConfigure = &actionResource{}
var _ resource.ResourceWithImportState = &actionResource{}

func newHAProxyActionResource() resource.Resource {
	return &actionResource{}
}

type actionResource struct {
	client opnsense.Client
}

func (r *actionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_haproxy_action"
}

func (r *actionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = actionResourceSchema()
}

func (r *actionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil { return }
	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = opnsense.NewClient(apiClient)
}

func (r *actionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *actionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	result, err := r.client.HAProxy().HAProxyAddAction(ctx, data.toActionObject())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create HAProxy action, got error: %s", err)); return }
	if result == nil { resp.Diagnostics.AddError("Client Error", "Unable to create HAProxy action: empty response received from API."); return }
	if result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("create HAProxy action", result)); return }
	id := result.UUID
	if id == "" {
		var found bool
		id, found, err = findActionIDByName(ctx, r.client.HAProxy(), data.Name.ValueString())
		if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to find HAProxy action after create, got error: %s", err)); return }
		if !found { resp.Diagnostics.AddError("Client Error", "Unable to create HAProxy action: API response did not include a UUID and the object could not be found by name."); return }
	}
	model, err := fetchActionModel(ctx, r.client.HAProxy(), id)
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy action after create, got error: %s", err)); return }
	tflog.Trace(ctx, "created HAProxy action", map[string]any{"id": id})
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *actionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *actionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() { resp.State.RemoveResource(ctx); return }
	model, err := fetchActionModel(ctx, r.client.HAProxy(), data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy action, got error: %s", err)); return }
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *actionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *actionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Client Error", "HAProxy action update requires a valid id."); return
	}
	result, err := r.client.HAProxy().HAProxyEditAction(ctx, data.Id.ValueString(), data.toActionObject())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update HAProxy action, got error: %s", err)); return }
	if result != nil && result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("update HAProxy action", result)); return }
	model, err := fetchActionModel(ctx, r.client.HAProxy(), data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy action after update, got error: %s", err)); return }
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *actionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *actionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() { return }
	result, err := r.client.HAProxy().HAProxyDeleteAction(ctx, data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete HAProxy action, got error: %s", err)); return }
	if result != nil && result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("delete HAProxy action", result)); return }
}

func (r *actionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
