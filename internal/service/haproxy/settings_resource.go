package haproxy

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

var _ resource.Resource = &haproxySettingsResource{}
var _ resource.ResourceWithConfigure = &haproxySettingsResource{}
var _ resource.ResourceWithImportState = &haproxySettingsResource{}

func newHAProxySettingsResource() resource.Resource {
	return &haproxySettingsResource{}
}

type haproxySettingsResource struct {
	client opnsense.Client
}

func (r *haproxySettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_haproxy_settings"
}

func (r *haproxySettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = haproxySettingsResourceSchema()
}

func (r *haproxySettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *haproxySettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *haproxySettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	settingsReq := data.toSettingsSetRequest()
	res, err := r.client.HAProxy().HAProxySetSettings(ctx, settingsReq)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create HAProxy settings, got error: %s", err))
		return
	}
	if res != nil && res.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("create HAProxy settings", res))
		return
	}

	tflog.Trace(ctx, "created HAProxy settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *haproxySettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *haproxySettingsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	settings, err := r.client.HAProxy().HAProxyGetSettings(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read HAProxy settings, got error: %s", err))
		return
	}

	resourceModel := settingsResponseToModel(settings)
	resp.Diagnostics.Append(resp.State.Set(ctx, &resourceModel)...)
}

func (r *haproxySettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *haproxySettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	settingsReq := data.toSettingsSetRequest()
	res, err := r.client.HAProxy().HAProxySetSettings(ctx, settingsReq)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update HAProxy settings, got error: %s", err))
		return
	}
	if res != nil && res.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("update HAProxy settings", res))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *haproxySettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *haproxySettingsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	settings, err := r.client.HAProxy().HAProxyGetSettings(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read current HAProxy settings during delete, got error: %s", err))
		return
	}

	resourceModel := settingsResponseToModel(settings)
	resourceModel.Enabled = types.BoolValue(false)

	settingsReq := resourceModel.toSettingsSetRequest()
	settingsReq.General.Enabled = boolToAPIString(types.BoolValue(false))

	res, err := r.client.HAProxy().HAProxySetSettings(ctx, settingsReq)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to disable HAProxy settings, got error: %s", err))
		return
	}
	if res != nil && res.Result == "failed" {
		resp.Diagnostics.AddError("Client Error", formatActionResultFailure("disable HAProxy settings", res))
		return
	}
}

func (r *haproxySettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
