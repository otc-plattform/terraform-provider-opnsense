package haproxy

import (
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resolverResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Enabled             types.Bool   `tfsdk:"enabled"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	Nameservers         types.Set    `tfsdk:"nameservers"`
	ParseResolvConf     types.Bool   `tfsdk:"parse_resolv_conf"`
	ResolveRetries      types.String `tfsdk:"resolve_retries"`
	TimeoutResolve      types.String `tfsdk:"timeout_resolve"`
	TimeoutRetry        types.String `tfsdk:"timeout_retry"`
	AcceptedPayloadSize types.String `tfsdk:"accepted_payload_size"`
	HoldValid           types.String `tfsdk:"hold_valid"`
	HoldObsolete        types.String `tfsdk:"hold_obsolete"`
	HoldRefused         types.String `tfsdk:"hold_refused"`
	HoldNX              types.String `tfsdk:"hold_nx"`
	HoldTimeout         types.String `tfsdk:"hold_timeout"`
	HoldOther           types.String `tfsdk:"hold_other"`
}

func resolverResourceSchema() schema.Schema {
	str := func(description, defaultValue string) schema.StringAttribute {
		return schema.StringAttribute{MarkdownDescription: description, Optional: true, Computed: true, Default: stringdefault.StaticString(defaultValue)}
	}
	return schema.Schema{MarkdownDescription: "Manage an OPNsense HAProxy DNS resolver.", Attributes: map[string]schema.Attribute{
		"id":                    schema.StringAttribute{Computed: true, MarkdownDescription: "Resolver UUID.", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
		"enabled":               schema.BoolAttribute{Optional: true, Computed: true, Default: booldefault.StaticBool(true)},
		"name":                  schema.StringAttribute{Required: true, MarkdownDescription: "Resolver name."},
		"description":           str("Description.", ""),
		"nameservers":           schema.SetAttribute{Optional: true, ElementType: types.StringType, MarkdownDescription: "Nameservers, such as 127.0.0.1:53."},
		"parse_resolv_conf":     schema.BoolAttribute{Optional: true, Computed: true, Default: booldefault.StaticBool(false)},
		"resolve_retries":       str("Number of DNS retries.", "3"),
		"timeout_resolve":       str("DNS resolve timeout.", "1s"),
		"timeout_retry":         str("DNS retry timeout.", "1s"),
		"accepted_payload_size": str("Maximum accepted DNS payload size.", ""),
		"hold_valid":            str("Valid response hold time.", ""),
		"hold_obsolete":         str("Obsolete response hold time.", ""),
		"hold_refused":          str("Refused response hold time.", ""),
		"hold_nx":               str("NXDOMAIN response hold time.", ""),
		"hold_timeout":          str("Timeout response hold time.", ""),
		"hold_other":            str("Other response hold time.", ""),
	}}
}

func resolverDataSourceSchema() dschema.Schema {
	attrs := map[string]dschema.Attribute{
		"id":                dschema.StringAttribute{Required: true},
		"enabled":           dschema.BoolAttribute{Computed: true},
		"parse_resolv_conf": dschema.BoolAttribute{Computed: true},
		"nameservers":       dschema.SetAttribute{Computed: true, ElementType: types.StringType},
	}
	for _, name := range []string{"name", "description", "resolve_retries", "timeout_resolve", "timeout_retry", "accepted_payload_size", "hold_valid", "hold_obsolete", "hold_refused", "hold_nx", "hold_timeout", "hold_other"} {
		attrs[name] = dschema.StringAttribute{Computed: true}
	}
	return dschema.Schema{MarkdownDescription: "Read an OPNsense HAProxy DNS resolver.", Attributes: attrs}
}
