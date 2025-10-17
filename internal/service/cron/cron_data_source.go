package cron

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/cron"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &cronDataSource{}
var _ datasource.DataSourceWithConfigure = &cronDataSource{}

func newCRONDataSource() datasource.DataSource {
	return &cronDataSource{}
}

// cronDataSource defines the data source implementation.
type cronDataSource struct {
	client opnsense.Client
}

func (d *cronDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_cron"
}

func (d *cronDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = cronDataSourceSchema()
}

func (d *cronDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = opnsense.NewClient(apiClient)
}

func (d *cronDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *cronResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get cron from OPNsense unbound API
	resourceStruct, err := d.client.Cron().SettingsGetJob(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read firewall cron, got error: %s", err))
		return
	}

	job := cronItemFromGet(resourceStruct)
	job.UUID = data.Id.ValueString()

	// Convert OPNsense struct to TF schema
	resourceModel, err := convertCRONStructToSchema(job)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read firewall cron, got error: %s", err))
		return
	}

	// ID cannot be added by convert... func, have to add here
	resourceModel.Id = data.Id

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &resourceModel)...)
}

func cronItemFromGet(get *cron.GetJobResponse) *cron.Item {
	job := &cron.Item{
		Origin:      get.Job.Origin,
		Enabled:     get.Job.Enabled,
		Minutes:     get.Job.Minutes,
		Hours:       get.Job.Hours,
		Days:        get.Job.Days,
		Months:      get.Job.Months,
		Weekdays:    get.Job.Weekdays,
		Who:         get.Job.Who,
		Parameters:  get.Job.Parameters,
		Description: get.Job.Description,
	}

	for k, v := range get.Job.Command {
		if v.Selected == 1 {
			job.Command = k
			break
		}
	}

	return job
}
