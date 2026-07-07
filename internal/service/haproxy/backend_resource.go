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

var _ resource.Resource = &backendResource{}
var _ resource.ResourceWithConfigure = &backendResource{}
var _ resource.ResourceWithImportState = &backendResource{}

func newHAProxyBackendResource() resource.Resource {
	return &backendResource{}
}

type backendResource struct {
	client opnsense.Client
}

func (r *backendResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_haproxy_backend"
}

func (r *backendResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = backendResourceSchema()
}

func (r *backendResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil { return }
	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = opnsense.NewClient(apiClient)
}

func (r *backendResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *backendResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	result, err := r.client.HAProxy().HAProxyAddBackend(ctx, data.toBackendObject())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create HAProxy backend, got error: %s", err)); return }
	if result == nil { resp.Diagnostics.AddError("Client Error", "Unable to create HAProxy backend: empty response received from API."); return }
	if result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("create HAProxy backend", result)); return }
	id := result.UUID
	if id == "" {
		var found bool
		id, found, err = findBackendIDByName(ctx, r.client.HAProxy(), data.Name.ValueString())
		if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to find HAProxy backend after create, got error: %s", err)); return }
		if !found { resp.Diagnostics.AddError("Client Error", "Unable to create HAProxy backend: API response did not include a UUID and the object could not be found by name."); return }
	}
	model, err := fetchBackendModel(ctx, r.client.HAProxy(), id)
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy backend after create, got error: %s", err)); return }
	tflog.Trace(ctx, "created HAProxy backend", map[string]any{"id": id})
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *backendResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *backendResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() { resp.State.RemoveResource(ctx); return }
	model, err := fetchBackendModel(ctx, r.client.HAProxy(), data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy backend, got error: %s", err)); return }
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *backendResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *backendResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Client Error", "HAProxy backend update requires a valid id."); return
	}
	result, err := r.client.HAProxy().HAProxyEditBackend(ctx, data.Id.ValueString(), data.toBackendObject())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update HAProxy backend, got error: %s", err)); return }
	if result != nil && result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("update HAProxy backend", result)); return }
	model, err := fetchBackendModel(ctx, r.client.HAProxy(), data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy backend after update, got error: %s", err)); return }
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *backendResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *backendResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() { return }
	result, err := r.client.HAProxy().HAProxyDeleteBackend(ctx, data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete HAProxy backend, got error: %s", err)); return }
	if result != nil && result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("delete HAProxy backend", result)); return }
}

func (r *backendResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
