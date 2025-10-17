package firmware

import (
	"context"
	"fmt"
	"time"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &firmwarePluginResource{}
var _ resource.ResourceWithConfigure = &firmwarePluginResource{}
var _ resource.ResourceWithImportState = &firmwarePluginResource{}

func newFirmwarePluginResource() resource.Resource {
	return &firmwarePluginResource{}
}

// firmwarePluginResource defines the resource implementation.
type firmwarePluginResource struct {
	client opnsense.Client
}

func (r *firmwarePluginResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_firmware_plugin"
}

func (r *firmwarePluginResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = firmwarePluginResourceSchema()
}

func (r *firmwarePluginResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *firmwarePluginResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *firmwarePluginResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add settings firmwarePlugin to unbound
	_, err := r.client.Firmware().FirmwareInstall(ctx, data.Package.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create settings firmwarePlugin, got error: %s (%s)", err, data.Package.ValueString()))
		return
	}

	installed := false
	for range 10 {
		installed, err = isInstalled(ctx, r.client, data.Package.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Client Error",
				fmt.Sprintf("Unable to read firewall firmwarePlugin, got error: %s", err))
			return
		}

		if installed {
			break
		}

		time.Sleep(time.Second * 5)
	}

	if !installed {
		resp.Diagnostics.AddError("Client Error", "Unable to install package: timeout")
		return
	}

	// Write logs using the tflog package
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *firmwarePluginResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *firmwarePluginResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get settings firmwarePlugin from OPNsense unbound API
	installed, err := isInstalled(ctx, r.client, data.Package.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read settings firmwarePlugin, got error: %s", err))
		return
	}

	// Convert OPNsense struct to TF schema
	resourceModel := firmwarePluginResourceModel{
		Package:   data.Package,
		Installed: types.BoolValue(installed),
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &resourceModel)...)
}

func (r *firmwarePluginResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *firmwarePluginResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Convert TF schema OPNsense struct
	if data.Installed.ValueBool() {

	} else {

	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *firmwarePluginResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *firmwarePluginResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Firmware().FirmwareRemove(ctx, data.Package.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete settings firmwarePlugin, got error: %s", err))
		return
	}
}

func (r *firmwarePluginResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
