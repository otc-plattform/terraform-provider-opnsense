package nginx

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &nginxSettingsResource{}
var _ resource.ResourceWithConfigure = &nginxSettingsResource{}
var _ resource.ResourceWithImportState = &nginxSettingsResource{}

func newNginxSettingsResource() resource.Resource {
	return &nginxSettingsResource{}
}

type nginxSettingsResource struct {
	client opnsense.Client
}

func (r *nginxSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nginx_settings"
}

func (r *nginxSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = nginxSettingsResourceSchema()
}

func (r *nginxSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *nginxSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *nginxSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	settingsReq := data.toSettingsSetRequest()
	res, err := r.client.Nginx().NginxSetSettings(ctx, settingsReq)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nginx settings, got error: %s", err))
		return
	}
	if res != nil && res.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("create Nginx settings", res))
		return
	}

	tflog.Trace(ctx, "created nginx settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *nginxSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *nginxSettingsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	settings, err := r.client.Nginx().NginxGetSettings(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nginx settings, got error: %s", err))
		return
	}

	resourceModel := settingsResponseToModel(settings)

	resp.Diagnostics.Append(resp.State.Set(ctx, &resourceModel)...)
}

func (r *nginxSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *nginxSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	settingsReq := data.toSettingsSetRequest()
	res, err := r.client.Nginx().NginxSetSettings(ctx, settingsReq)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nginx settings, got error: %s", err))
		return
	}
	if res != nil && res.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("update Nginx settings", res))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *nginxSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *nginxSettingsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	settings, err := r.client.Nginx().NginxGetSettings(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read current nginx settings during delete, got error: %s", err))
		return
	}

	resourceModel := settingsResponseToModel(settings)
	resourceModel.Enabled = types.BoolValue(false)

	settingsReq := resourceModel.toSettingsSetRequest()
	settingsReq.General.Enabled = boolToAPIString(types.BoolValue(false))

	res, err := r.client.Nginx().NginxSetSettings(ctx, settingsReq)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to disable nginx settings, got error: %s", err))
		return
	}
	if res != nil && res.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("disable Nginx settings", res))
		return
	}
}

func (r *nginxSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
