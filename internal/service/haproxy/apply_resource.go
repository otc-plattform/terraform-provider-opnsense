package haproxy

import (
	"context"
	"fmt"
	"strings"

	"github.com/browningluke/opnsense-go/pkg/api"
	client "github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &haproxyApplyResource{}
var _ resource.ResourceWithConfigure = &haproxyApplyResource{}

type haproxyApplyResource struct {
	client opnsense.Client
}

type haproxyApplyModel struct {
	ID       types.String `tfsdk:"id"`
	Triggers types.Map    `tfsdk:"triggers"`
}

func newHAProxyApplyResource() resource.Resource { return &haproxyApplyResource{} }

func (r *haproxyApplyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_haproxy_apply"
}

func (r *haproxyApplyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Validate and activate the staged HAProxy configuration. Changes to triggers rerun the action; destroying the resource does not change HAProxy.",
		Attributes: map[string]schema.Attribute{
			"id":       schema.StringAttribute{Computed: true, MarkdownDescription: "Identifier for the apply action.", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"triggers": schema.MapAttribute{Required: true, ElementType: types.StringType, MarkdownDescription: "Values whose changes trigger another syntax check and reconfiguration. Use a fingerprint of all relevant HAProxy configuration, and explicitly depend on the resources it configures."},
		},
	}
}

func (r *haproxyApplyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *api.Client, got %T.", req.ProviderData))
		return
	}
	r.client = opnsense.NewClient(apiClient)
}

// applyHAProxy follows the HAProxy UI's configtest-before-reconfigure sequence.
func applyHAProxy(ctx context.Context, controller *client.Controller) (string, error) {
	output, err := controller.HAProxyConfigTest(ctx)
	if err != nil {
		return "", err
	}
	output = strings.TrimSpace(output)
	upper := strings.ToUpper(output)
	if strings.Contains(upper, "ALERT") || strings.Contains(upper, "FATAL") {
		return "", fmt.Errorf("HAProxy syntax check failed; configuration was not activated:\n%s", output)
	}
	// OPNsense can return an empty result for a valid configuration with no
	// active listeners. A missing result is rejected by HAProxyConfigTest.
	if output != "" && !strings.Contains(strings.ToLower(output), "configuration file is valid") && !strings.Contains(upper, "WARNING") {
		return "", fmt.Errorf("HAProxy syntax check returned unrecognized output; configuration was not activated. Output: %q", output)
	}
	if err := controller.HAProxyReconfigure(ctx); err != nil {
		return "", err
	}
	if strings.Contains(upper, "WARNING") {
		return output, nil
	}
	return "", nil
}

func (r *haproxyApplyResource) apply(ctx context.Context, data *haproxyApplyModel, resp *resource.CreateResponse) {
	warning, err := applyHAProxy(ctx, r.client.HAProxy())
	if err != nil {
		resp.Diagnostics.AddError("HAProxy Apply Error", err.Error())
		return
	}
	if warning != "" {
		resp.Diagnostics.AddWarning("HAProxy Configuration Warning", warning)
	}
	data.ID = types.StringValue("haproxy")
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *haproxyApplyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data haproxyApplyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if !resp.Diagnostics.HasError() {
		r.apply(ctx, &data, resp)
	}
}

func (r *haproxyApplyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// This is an action, not a remote object. Preserve the last successful triggers.
}

func (r *haproxyApplyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data haproxyApplyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	warning, err := applyHAProxy(ctx, r.client.HAProxy())
	if err != nil {
		resp.Diagnostics.AddError("HAProxy Apply Error", err.Error())
		return
	}
	if warning != "" {
		resp.Diagnostics.AddWarning("HAProxy Configuration Warning", warning)
	}
	data.ID = types.StringValue("haproxy")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *haproxyApplyResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Removing the action from state must not stop or modify HAProxy.
}
