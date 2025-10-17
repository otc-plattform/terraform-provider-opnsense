package acmeclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &acmeclientCertificateResource{}
var _ resource.ResourceWithConfigure = &acmeclientCertificateResource{}
var _ resource.ResourceWithImportState = &acmeclientCertificateResource{}

func newACMEClientCertificateResource() resource.Resource {
	return &acmeclientCertificateResource{}
}

type acmeclientCertificateResource struct {
	client opnsense.Client
}

func (r *acmeclientCertificateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acmeclient_certificate"
}

func (r *acmeclientCertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = acmeclientCertificateResourceSchema()
}

func (r *acmeclientCertificateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *acmeclientCertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *acmeclientCertificateResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	certificate, err := data.toCertificate(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to prepare acmeclient certificate payload, got error: %s", err))
		return
	}

	result, err := r.client.Acmeclient().ACMEClientAddCert(ctx, certificate)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create acmeclient certificate, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("create ACME client certificate", result))
		return
	}

	if result != nil && result.UUID != "" {
		data.Id = types.StringValue(result.UUID)
	}

	if data.Id.IsNull() || data.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", "API did not return an identifier for the new acmeclient certificate.")
		return
	}

	certModel, err := fetchCertificateModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient certificate after create, got error: %s", err))
		return
	}

	certModel.Id = data.Id

	tflog.Trace(ctx, "created acmeclient certificate", map[string]any{
		"id": certModel.Id.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &certModel)...)
}

func (r *acmeclientCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *acmeclientCertificateResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	certModel, err := fetchCertificateModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, "acmeclient certificate not present in remote, removing from state", map[string]any{
				"id": data.Id.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient certificate, got error: %s", err))
		return
	}

	certModel.Id = data.Id

	resp.Diagnostics.Append(resp.State.Set(ctx, &certModel)...)
}

func (r *acmeclientCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *acmeclientCertificateResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	certificate, err := data.toCertificate(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to prepare acmeclient certificate payload, got error: %s", err))
		return
	}

	result, err := r.client.Acmeclient().ACMEClientEditCert(ctx, data.Id.ValueString(), certificate)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update acmeclient certificate, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("update ACME client certificate", result))
		return
	}

	certModel, err := fetchCertificateModel(ctx, r.client.Acmeclient(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read acmeclient certificate after update, got error: %s", err))
		return
	}

	certModel.Id = data.Id

	resp.Diagnostics.Append(resp.State.Set(ctx, &certModel)...)
}

func (r *acmeclientCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *acmeclientCertificateResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Acmeclient().ACMEClientDeleteCert(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete acmeclient certificate, got error: %s", err))
		return
	}

	if result != nil && result.Result == "failed" {
		resp.Diagnostics.AddError("Client Error",
			formatActionResultFailure("delete ACME client certificate", result))
		return
	}
}

func (r *acmeclientCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
