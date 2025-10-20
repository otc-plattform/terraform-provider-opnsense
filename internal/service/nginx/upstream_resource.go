package nginx

import (
	"context"
	"errors"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &nginxUpstreamResource{}
var _ resource.ResourceWithConfigure = &nginxUpstreamResource{}
var _ resource.ResourceWithImportState = &nginxUpstreamResource{}

func newNginxUpstreamResource() resource.Resource {
	return &nginxUpstreamResource{}
}

type nginxUpstreamResource struct {
	client opnsense.Client
}

func (r *nginxUpstreamResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nginx_upstream"
}

func (r *nginxUpstreamResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = nginxUpstreamResourceSchema()
}

func (r *nginxUpstreamResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *nginxUpstreamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *nginxUpstreamResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	upstream := data.toUpstream()

	result, err := r.client.Nginx().NginxAddUpstream(ctx, upstream)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nginx upstream, got error: %s", err))
		return
	}

	if result == nil {
		resp.Diagnostics.AddError("Client Error", "Unable to create nginx upstream: empty response received from API.")
		return
	}

	if result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("create Nginx upstream", result))
		return
	}

	if result.UUID == "" {
		resp.Diagnostics.AddError("Client Error", "Unable to create nginx upstream: API response did not include a UUID.")
		return
	}

	model, err := fetchUpstreamModel(ctx, r.client.Nginx(), result.UUID)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx upstream after create, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "created nginx upstream", map[string]any{"id": result.UUID})

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *nginxUpstreamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *nginxUpstreamResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.State.RemoveResource(ctx)
		return
	}

	model, err := fetchUpstreamModel(ctx, r.client.Nginx(), data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("nginx upstream %s not present in remote, removing from state", data.Id.ValueString()))
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx upstream, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *nginxUpstreamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *nginxUpstreamResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Client Error", "Nginx upstream update requires a valid id.")
		return
	}

	upstream := data.toUpstream()

	result, err := r.client.Nginx().NginxEditUpstream(ctx, data.Id.ValueString(), upstream)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nginx upstream, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("update Nginx upstream", result))
		return
	}

	model, err := fetchUpstreamModel(ctx, r.client.Nginx(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx upstream after update, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *nginxUpstreamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *nginxUpstreamResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		return
	}

	result, err := r.client.Nginx().NginxDeleteUpstream(ctx, data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			return
		}

		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nginx upstream, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("delete Nginx upstream", result))
		return
	}
}

func (r *nginxUpstreamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
