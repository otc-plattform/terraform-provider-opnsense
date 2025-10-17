package acmeclient

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &acmeclientSettingsResource{}
var _ resource.ResourceWithConfigure = &acmeclientSettingsResource{}
var _ resource.ResourceWithImportState = &acmeclientSettingsResource{}

func newACMEClientSettingsResource() resource.Resource {
	return &acmeclientSettingsResource{}
}

// acmeclientSettingsResource defines the resource implementation.
type acmeclientSettingsResource struct {
	client opnsense.Client
}

func (r *acmeclientSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acmeclient_settings"
}

func (r *acmeclientSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = acmeclientSettingsResourceSchema()
}

func (r *acmeclientSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *opnsense.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = opnsense.NewClient(apiClient)
}

func (r *acmeclientSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *acmeclientSettingsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add settings acmeclientSettings to OPNsense
	settingsReq := data.toSettingsSetRequest()
	res, err := r.client.Acmeclient().ACMEClientSetSettings(ctx, settingsReq)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create acmeclient settings, got error: %s", err))
		return
	}
	if res != nil && res.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create acmeclient settings, got response: %v", res))
		return
	}

	// Write logs using the tflog package
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *acmeclientSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *acmeclientSettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get settings acmeclientSettings from OPNsense unbound API
	settings, err := r.client.Acmeclient().ACMEClientGetSettings(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient settings, got error: %s", err))
		return
	}

	// Convert OPNsense struct to TF schema
	resourceModel := settingsResponseToModel(settings)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &resourceModel)...)
}

func (r *acmeclientSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *acmeclientSettingsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	settingsReq := data.toSettingsSetRequest()
	res, err := r.client.Acmeclient().ACMEClientSetSettings(ctx, settingsReq)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update acmeclient settings, got error: %s", err))
		return
	}
	if res != nil && res.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update acmeclient settings, got response: %v", res))
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *acmeclientSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *acmeclientSettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	settings, err := r.client.Acmeclient().ACMEClientGetSettings(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read current acmeclient settings during delete, got error: %s", err))
		return
	}

	resourceModel := settingsResponseToModel(settings)
	resourceModel.Enabled = types.BoolValue(false)

	settingsReq := resourceModel.toSettingsSetRequest()
	settingsReq.Settings.Enabled = tools.BoolToString(false)

	res, err := r.client.Acmeclient().ACMEClientSetSettings(ctx, settingsReq)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete settings acmeclientSettings, got error: %s", err))
		return
	}
	if res != nil && res.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete settings acmeclientSettings, got response: %v", res))
		return
	}
}

func (r *acmeclientSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
