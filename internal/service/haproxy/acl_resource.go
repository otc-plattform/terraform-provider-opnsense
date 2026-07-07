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

var _ resource.Resource = &aclResource{}
var _ resource.ResourceWithConfigure = &aclResource{}
var _ resource.ResourceWithImportState = &aclResource{}

func newHAProxyACLResource() resource.Resource {
	return &aclResource{}
}

type aclResource struct {
	client opnsense.Client
}

func (r *aclResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_haproxy_acl"
}

func (r *aclResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = aclResourceSchema()
}

func (r *aclResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil { return }
	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = opnsense.NewClient(apiClient)
}

func (r *aclResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *aclResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	result, err := r.client.HAProxy().HAProxyAddACL(ctx, data.toACLObject())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create HAProxy ACL, got error: %s", err)); return }
	if result == nil { resp.Diagnostics.AddError("Client Error", "Unable to create HAProxy ACL: empty response received from API."); return }
	if result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("create HAProxy ACL", result)); return }
	id := result.UUID
	if id == "" {
		var found bool
		id, found, err = findACLIDByName(ctx, r.client.HAProxy(), data.Name.ValueString())
		if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to find HAProxy ACL after create, got error: %s", err)); return }
		if !found { resp.Diagnostics.AddError("Client Error", "Unable to create HAProxy ACL: API response did not include a UUID and the object could not be found by name."); return }
	}
	model, err := fetchACLModel(ctx, r.client.HAProxy(), id)
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy ACL after create, got error: %s", err)); return }
	tflog.Trace(ctx, "created HAProxy ACL", map[string]any{"id": id})
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *aclResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *aclResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() { resp.State.RemoveResource(ctx); return }
	model, err := fetchACLModel(ctx, r.client.HAProxy(), data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy ACL, got error: %s", err)); return }
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *aclResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *aclResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Client Error", "HAProxy ACL update requires a valid id."); return
	}
	result, err := r.client.HAProxy().HAProxyEditACL(ctx, data.Id.ValueString(), data.toACLObject())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update HAProxy ACL, got error: %s", err)); return }
	if result != nil && result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("update HAProxy ACL", result)); return }
	model, err := fetchACLModel(ctx, r.client.HAProxy(), data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy ACL after update, got error: %s", err)); return }
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *aclResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *aclResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() { return }
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() { return }
	result, err := r.client.HAProxy().HAProxyDeleteACL(ctx, data.Id.ValueString())
	if err != nil { resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete HAProxy ACL, got error: %s", err)); return }
	if result != nil && result.Result == "failed" { resp.Diagnostics.AddError("Client Error", formatActionResultFailure("delete HAProxy ACL", result)); return }
}

func (r *aclResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
