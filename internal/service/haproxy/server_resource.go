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

var _ resource.Resource = &haproxyServerResource{}
var _ resource.ResourceWithConfigure = &haproxyServerResource{}
var _ resource.ResourceWithImportState = &haproxyServerResource{}

func newHAProxyServerResource() resource.Resource {
	return &haproxyServerResource{}
}

type haproxyServerResource struct {
	client opnsense.Client
}

func (r *haproxyServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_haproxy_server"
}

func (r *haproxyServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = haproxyServerResourceSchema()
}

func (r *haproxyServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *haproxyServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *haproxyServerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.HAProxy().HAProxyAddServer(ctx, data.toServer())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create HAProxy server, got error: %s", err))
		return
	}
	if result == nil {
		resp.Diagnostics.AddError("Client Error", "Unable to create HAProxy server: empty response received from API.")
		return
	}
	if result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("create HAProxy server", result))
		return
	}
	id := result.UUID
	if id == "" {
		var found bool
		id, found, err = findServerIDByName(ctx, r.client.HAProxy(), data.Name.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to find HAProxy server after create, got error: %s", err))
			return
		}
		if !found {
			resp.Diagnostics.AddError("Client Error", "Unable to create HAProxy server: API response did not include a UUID and the server could not be found by name.")
			return
		}
	}

	model, err := fetchServerModel(ctx, r.client.HAProxy(), id)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy server after create, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "created HAProxy server", map[string]any{"id": id})

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *haproxyServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *haproxyServerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.State.RemoveResource(ctx)
		return
	}

	model, err := fetchServerModel(ctx, r.client.HAProxy(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy server, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *haproxyServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *haproxyServerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Client Error", "HAProxy server update requires a valid id.")
		return
	}

	result, err := r.client.HAProxy().HAProxyEditServer(ctx, data.Id.ValueString(), data.toServer())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update HAProxy server, got error: %s", err))
		return
	}
	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("update HAProxy server", result))
		return
	}

	model, err := fetchServerModel(ctx, r.client.HAProxy(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy server after update, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *haproxyServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *haproxyServerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		return
	}

	result, err := r.client.HAProxy().HAProxyDeleteServer(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete HAProxy server, got error: %s", err))
		return
	}
	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("delete HAProxy server", result))
		return
	}
}

func (r *haproxyServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
