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

var _ resource.Resource = &nginxHTTPServerResource{}
var _ resource.ResourceWithConfigure = &nginxHTTPServerResource{}
var _ resource.ResourceWithImportState = &nginxHTTPServerResource{}

func newNginxHTTPServerResource() resource.Resource {
	return &nginxHTTPServerResource{}
}

type nginxHTTPServerResource struct {
	client opnsense.Client
}

func (r *nginxHTTPServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nginx_http_server"
}

func (r *nginxHTTPServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = nginxHTTPServerResourceSchema()
}

func (r *nginxHTTPServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *nginxHTTPServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *nginxHTTPServerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpServer := data.toHTTPServer()

	result, err := r.client.Nginx().NginxAddHTTPServer(ctx, httpServer)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nginx HTTP server, got error: %s", err))
		return
	}

	if result == nil {
		resp.Diagnostics.AddError("Client Error", "Unable to create nginx HTTP server: empty response received from API.")
		return
	}

	if result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("create Nginx HTTP server", result))
		return
	}

	if result.UUID == "" {
		resp.Diagnostics.AddError("Client Error", "Unable to create nginx HTTP server: API response did not include a UUID.")
		return
	}

	model, err := fetchHTTPServerModel(ctx, r.client.Nginx(), result.UUID)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx HTTP server after create, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "created nginx HTTP server", map[string]any{"id": result.UUID})

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *nginxHTTPServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *nginxHTTPServerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.State.RemoveResource(ctx)
		return
	}

	model, err := fetchHTTPServerModel(ctx, r.client.Nginx(), data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("nginx HTTP server %s not present in remote, removing from state", data.Id.ValueString()))
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx HTTP server, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *nginxHTTPServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *nginxHTTPServerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		resp.Diagnostics.AddError("Client Error", "Nginx HTTP server update requires a valid id.")
		return
	}

	httpServer := data.toHTTPServer()

	result, err := r.client.Nginx().NginxEditHTTPServer(ctx, data.Id.ValueString(), httpServer)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nginx HTTP server, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("update Nginx HTTP server", result))
		return
	}

	model, err := fetchHTTPServerModel(ctx, r.client.Nginx(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx HTTP server after update, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *nginxHTTPServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *nginxHTTPServerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data == nil || data.Id.IsNull() || data.Id.IsUnknown() {
		return
	}

	result, err := r.client.Nginx().NginxDeleteHTTPServer(ctx, data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			return
		}

		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nginx HTTP server, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("delete Nginx HTTP server", result))
		return
	}
}

func (r *nginxHTTPServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
