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

var _ resource.Resource = &nginxUpstreamServerResource{}
var _ resource.ResourceWithConfigure = &nginxUpstreamServerResource{}
var _ resource.ResourceWithImportState = &nginxUpstreamServerResource{}

func newNginxUpstreamServerResource() resource.Resource {
	return &nginxUpstreamServerResource{}
}

type nginxUpstreamServerResource struct {
	client opnsense.Client
}

func (r *nginxUpstreamServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nginx_upstream_server"
}

func (r *nginxUpstreamServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = nginxUpstreamServerResourceSchema()
}

func (r *nginxUpstreamServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *nginxUpstreamServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *nginxUpstreamServerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	upstreamServer := data.toUpstreamServer()

	result, err := r.client.Nginx().NginxAddUpstreamServer(ctx, upstreamServer)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nginx upstream server, got error: %s", err))
		return
	}

	if result == nil {
		resp.Diagnostics.AddError("Client Error", "Unable to create nginx upstream server: empty response received from API.")
		return
	}

	if result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("create Nginx upstream server", result))
		return
	}

	if result.UUID == "" {
		resp.Diagnostics.AddError("Client Error", "Unable to create nginx upstream server: API response did not include a UUID.")
		return
	}

	model, err := fetchUpstreamServerModel(ctx, r.client.Nginx(), result.UUID)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx upstream server after create, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "created nginx upstream server", map[string]any{"id": result.UUID})

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *nginxUpstreamServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *nginxUpstreamServerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.State.RemoveResource(ctx)
		return
	}

	model, err := fetchUpstreamServerModel(ctx, r.client.Nginx(), data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("nginx upstream server %s not present in remote, removing from state", data.Id.ValueString()))
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx upstream server, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *nginxUpstreamServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *nginxUpstreamServerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Client Error", "Nginx upstream server update requires a valid id.")
		return
	}

	upstreamServer := data.toUpstreamServer()

	result, err := r.client.Nginx().NginxEditUpstreamServer(ctx, data.Id.ValueString(), upstreamServer)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nginx upstream server, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("update Nginx upstream server", result))
		return
	}

	model, err := fetchUpstreamServerModel(ctx, r.client.Nginx(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx upstream server after update, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *nginxUpstreamServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *nginxUpstreamServerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		return
	}

	result, err := r.client.Nginx().NginxDeleteUpstreamServer(ctx, data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			return
		}

		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nginx upstream server, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("delete Nginx upstream server", result))
		return
	}
}

func (r *nginxUpstreamServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
