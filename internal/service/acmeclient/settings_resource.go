package acmeclient

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/acmeclient"
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

	// Add settings acmeclientSettings to unbound
	r.client.Acmeclient().ACMEClientSetSettings(ctx, acmeclient.SettingsSetRequest{
		Settings: acmeclient.SettingsSetSettings{
			Enabled: tools.BoolToString(data.Enabled.ValueBool()),

			// TODO: fill other fields
		},
	})

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
	resourceModel := acmeclientSettingsResourceModel{
		Enabled: types.BoolValue(tools.StringToBool(settings.ACMEClient.Settings.Enabled)),

		// TODO: convert rest of fields
	}

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

	// Convert TF schema OPNsense struct

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

	_, err := r.client.Acmeclient().ACMEClientSetSettings(ctx, acmeclient.SettingsSetRequest{
		Settings: acmeclient.SettingsSetSettings{
			Enabled: tools.BoolToString(false),

			// TODO: fill other fields
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete settings acmeclientSettings, got error: %s", err))
		return
	}
}

func (r *acmeclientSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
