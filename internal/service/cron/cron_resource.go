package cron

import (
	"context"
	"errors"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/cron"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &cronResource{}
var _ resource.ResourceWithConfigure = &cronResource{}
var _ resource.ResourceWithImportState = &cronResource{}

func newCRONResource() resource.Resource {
	return &cronResource{}
}

// cronResource defines the resource implementation.
type cronResource struct {
	client opnsense.Client
}

func (r *cronResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_cron"
}

func (r *cronResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = cronResourceSchema()
}

func (r *cronResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *cronResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *cronResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Convert TF schema OPNsense struct
	resourceStruct, err := convertCRONSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse firwall cron, got error: %s", err))
		return
	}

	// Add settings cron to unbound
	res, err := r.client.Cron().SettingsAddJob(ctx, cron.JobRequest{
		Origin:      "cron",
		Enabled:     resourceStruct.Enabled,
		Minutes:     resourceStruct.Minutes,
		Hours:       resourceStruct.Hours,
		Days:        resourceStruct.Days,
		Months:      resourceStruct.Months,
		Weekdays:    resourceStruct.Weekdays,
		Command:     resourceStruct.Command,
		Parameters:  resourceStruct.Parameters,
		Description: resourceStruct.Description,
	})
	if err != nil {
		if res != nil && res.UUID != "" {
			// Tag new resource with ID from OPNsense
			data.Id = types.StringValue(res.UUID)

			// Save data into Terraform state
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create settings cron, got error: %s", err))
		return
	}

	if res.UUID == "" || len(res.Validations) > 0 {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create settings cron, got resp: %v", res))
		return
	}

	// Tag new resource with ID from OPNsense
	data.Id = types.StringValue(res.UUID)

	// Write logs using the tflog package
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *cronResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *cronResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get settings cron from OPNsense unbound API
	resourceStruct, err := r.client.Cron().SettingsGetJob(ctx, data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, "settings cron not present in remote, removing from state")
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read settings cron, got error: %s", err))
		return
	}

	job := cronItemFromGet(resourceStruct)
	job.UUID = data.Id.ValueString()

	// Convert OPNsense struct to TF schema
	resourceModel, err := convertCRONStructToSchema(job)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read settings cron, got error: %s", err))
		return
	}

	// ID cannot be added by convert... func, have to add here
	resourceModel.Id = data.Id

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &resourceModel)...)
}

func (r *cronResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *cronResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Convert TF schema OPNsense struct
	resourceStruct, err := convertCRONSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse settings cron, got error: %s", err))
		return
	}

	// Update settings cron in unbound
	res, err := r.client.Cron().SettingsSetJob(ctx, cron.JobRequest{
		Origin:      "user",
		Enabled:     resourceStruct.Enabled,
		Minutes:     resourceStruct.Minutes,
		Hours:       resourceStruct.Hours,
		Days:        resourceStruct.Days,
		Months:      resourceStruct.Months,
		Weekdays:    resourceStruct.Weekdays,
		Parameters:  resourceStruct.Parameters,
		Description: resourceStruct.Description,
		Command:     resourceStruct.Command,
	}, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update settings cron, got error: %s", err))
		return
	}
	if res.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update settings cron, got resp: %v", res))
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *cronResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *cronResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Cron().SettingsDeleteJob(ctx, data.Id.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete settings cron, got error: %s", err))
		return
	}
}

func (r *cronResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
