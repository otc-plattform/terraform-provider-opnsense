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

var _ resource.Resource = &frontendResource{}
var _ resource.ResourceWithConfigure = &frontendResource{}
var _ resource.ResourceWithImportState = &frontendResource{}

func newHAProxyFrontendResource() resource.Resource {
	return &frontendResource{}
}

type frontendResource struct {
	client opnsense.Client
}

func (r *frontendResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_haproxy_frontend"
}

func (r *frontendResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = frontendResourceSchema()
}

func (r *frontendResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil { return }
	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = opnsense.NewClient(apiClient)
}

func (r *frontendResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *frontendResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	result, err := r.client.HAProxy().HAProxyAddFrontend(ctx, data.toFrontendObject())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create HAProxy frontend, got error: %s", err)); return }
	if result == nil { resp.Diagnostics.AddError("Client Error", "Unable to create HAProxy frontend: empty response received from API."); return }
	if result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("create HAProxy frontend", result)); return }
	id := result.UUID
	if id == "" {
		var found bool
		id, found, err = findFrontendIDByName(ctx, r.client.HAProxy(), data.Name.ValueString())
		if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to find HAProxy frontend after create, got error: %s", err)); return }
		if !found { resp.Diagnostics.AddError("Client Error", "Unable to create HAProxy frontend: API response did not include a UUID and the object could not be found by name."); return }
	}
	model, err := fetchFrontendModel(ctx, r.client.HAProxy(), id)
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy frontend after create, got error: %s", err)); return }
	tflog.Trace(ctx, "created HAProxy frontend", map[string]any{"id": id})
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *frontendResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *frontendResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() { resp.State.RemoveResource(ctx); return }
	model, err := fetchFrontendModel(ctx, r.client.HAProxy(), data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy frontend, got error: %s", err)); return }
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *frontendResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *frontendResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Client Error", "HAProxy frontend update requires a valid id."); return
	}
	result, err := r.client.HAProxy().HAProxyEditFrontend(ctx, data.Id.ValueString(), data.toFrontendObject())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update HAProxy frontend, got error: %s", err)); return }
	if result != nil && result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("update HAProxy frontend", result)); return }
	model, err := fetchFrontendModel(ctx, r.client.HAProxy(), data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy frontend after update, got error: %s", err)); return }
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *frontendResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *frontendResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() { return }
	result, err := r.client.HAProxy().HAProxyDeleteFrontend(ctx, data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete HAProxy frontend, got error: %s", err)); return }
	if result != nil && result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("delete HAProxy frontend", result)); return }
}

func (r *frontendResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
