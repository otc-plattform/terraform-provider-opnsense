package gateway

import (
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type gatewayResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Name                      types.String `tfsdk:"name"`
	Description               types.String `tfsdk:"description"`
	Interface                 types.String `tfsdk:"interface"`
	IpProtocol                types.String `tfsdk:"ip_protocol"`
	Gateway                   types.String `tfsdk:"gateway"`
	Disabled                  types.Bool   `tfsdk:"disabled"`
	DefaultGateway            types.Bool   `tfsdk:"default_gateway"`
	FarGateway                types.Bool   `tfsdk:"far_gateway"`
	MonitorDisable            types.Bool   `tfsdk:"monitor_disable"`
	MonitorNoRoute            types.Bool   `tfsdk:"monitor_no_route"`
	MonitorKillStates         types.Bool   `tfsdk:"monitor_kill_states"`
	MonitorKillStatesPriority types.Int64  `tfsdk:"monitor_kill_states_priority"`
	Monitor                   types.String `tfsdk:"monitor"`
	ForceDown                 types.Bool   `tfsdk:"force_down"`
	Priority                  types.Int64  `tfsdk:"priority"`
	Weight                    types.Int64  `tfsdk:"weight"`
	LatencyLow                types.Int64  `tfsdk:"latency_low"`
	LatencyHigh               types.Int64  `tfsdk:"latency_high"`
	LossLow                   types.Int64  `tfsdk:"loss_low"`
	LossHigh                  types.Int64  `tfsdk:"loss_high"`
	Interval                  types.Int64  `tfsdk:"interval"`
	TimePeriod                types.Int64  `tfsdk:"time_period"`
	LossInterval              types.Int64  `tfsdk:"loss_interval"`
	DataLength                types.Int64  `tfsdk:"data_length"`
}

func gatewayResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage individual gateways under **System → Gateways → Single**.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID of the gateway.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Unique name for the gateway, e.g. `WAN_DHCP`.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional free-form description for your reference.",
				Optional:            true,
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "Interface identifier (e.g., `wan`, `lan`) the gateway belongs to.",
				Required:            true,
			},
			"ip_protocol": schema.StringAttribute{
				MarkdownDescription: "IP protocol of the gateway, usually `inet` for IPv4 or `inet6` for IPv6.",
				Required:            true,
			},
			"gateway": schema.StringAttribute{
				MarkdownDescription: "Gateway IP address. Leave empty for dynamic gateways that derive their address from the interface.",
				Optional:            true,
			},
			"disabled": schema.BoolAttribute{
				MarkdownDescription: "Disable this gateway. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"default_gateway": schema.BoolAttribute{
				MarkdownDescription: "Set this gateway as default for the selected protocol. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"far_gateway": schema.BoolAttribute{
				MarkdownDescription: "Allow the gateway to exist outside the interface subnet. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"monitor_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable gateway monitoring. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"monitor_no_route": schema.BoolAttribute{
				MarkdownDescription: "Do not create monitoring routes automatically. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"monitor_kill_states": schema.BoolAttribute{
				MarkdownDescription: "Kill states when the gateway transitions to down. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"monitor_kill_states_priority": schema.Int64Attribute{
				MarkdownDescription: "Optional priority to control state killing order across gateways.",
				Optional:            true,
			},
			"monitor": schema.StringAttribute{
				MarkdownDescription: "Optional monitor IP address or host.",
				Optional:            true,
			},
			"force_down": schema.BoolAttribute{
				MarkdownDescription: "Manually force the gateway status to down. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Gateway priority (lower is preferred).",
				Optional:            true,
			},
			"weight": schema.Int64Attribute{
				MarkdownDescription: "Load balancing weight for the gateway.",
				Optional:            true,
			},
			"latency_low": schema.Int64Attribute{
				MarkdownDescription: "Low latency threshold in milliseconds.",
				Optional:            true,
			},
			"latency_high": schema.Int64Attribute{
				MarkdownDescription: "High latency threshold in milliseconds.",
				Optional:            true,
			},
			"loss_low": schema.Int64Attribute{
				MarkdownDescription: "Low packet loss threshold as percentage.",
				Optional:            true,
			},
			"loss_high": schema.Int64Attribute{
				MarkdownDescription: "High packet loss threshold as percentage.",
				Optional:            true,
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: "Ping interval in seconds for monitoring.",
				Optional:            true,
			},
			"time_period": schema.Int64Attribute{
				MarkdownDescription: "Monitoring time period in seconds.",
				Optional:            true,
			},
			"loss_interval": schema.Int64Attribute{
				MarkdownDescription: "Monitoring loss interval in milliseconds.",
				Optional:            true,
			},
			"data_length": schema.Int64Attribute{
				MarkdownDescription: "ICMP payload size in bytes for monitoring pings.",
				Optional:            true,
			},
		},
	}
}

func gatewayDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Look up an existing gateway under **System → Gateways → Single** by UUID.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the gateway.",
				Required:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "Gateway name.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Gateway description.",
				Computed:            true,
			},
			"interface": dschema.StringAttribute{
				MarkdownDescription: "Interface identifier.",
				Computed:            true,
			},
			"ip_protocol": dschema.StringAttribute{
				MarkdownDescription: "IP protocol handled by the gateway.",
				Computed:            true,
			},
			"gateway": dschema.StringAttribute{
				MarkdownDescription: "Gateway IP address.",
				Computed:            true,
			},
			"disabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether the gateway is disabled.",
				Computed:            true,
			},
			"default_gateway": dschema.BoolAttribute{
				MarkdownDescription: "Whether the gateway is the default for this protocol.",
				Computed:            true,
			},
			"far_gateway": dschema.BoolAttribute{
				MarkdownDescription: "Whether the gateway may exist outside the interface subnet.",
				Computed:            true,
			},
			"monitor_disable": dschema.BoolAttribute{
				MarkdownDescription: "Whether gateway monitoring is disabled.",
				Computed:            true,
			},
			"monitor_no_route": dschema.BoolAttribute{
				MarkdownDescription: "Whether monitoring routes are skipped.",
				Computed:            true,
			},
			"monitor_kill_states": dschema.BoolAttribute{
				MarkdownDescription: "Whether states are cleared when the gateway is down.",
				Computed:            true,
			},
			"monitor_kill_states_priority": dschema.Int64Attribute{
				MarkdownDescription: "Priority for the kill-states action.",
				Computed:            true,
			},
			"monitor": dschema.StringAttribute{
				MarkdownDescription: "Monitor IP or host.",
				Computed:            true,
			},
			"force_down": dschema.BoolAttribute{
				MarkdownDescription: "Whether the gateway is forced down.",
				Computed:            true,
			},
			"priority": dschema.Int64Attribute{
				MarkdownDescription: "Gateway priority.",
				Computed:            true,
			},
			"weight": dschema.Int64Attribute{
				MarkdownDescription: "Gateway weight.",
				Computed:            true,
			},
			"latency_low": dschema.Int64Attribute{
				MarkdownDescription: "Low latency threshold.",
				Computed:            true,
			},
			"latency_high": dschema.Int64Attribute{
				MarkdownDescription: "High latency threshold.",
				Computed:            true,
			},
			"loss_low": dschema.Int64Attribute{
				MarkdownDescription: "Low loss threshold.",
				Computed:            true,
			},
			"loss_high": dschema.Int64Attribute{
				MarkdownDescription: "High loss threshold.",
				Computed:            true,
			},
			"interval": dschema.Int64Attribute{
				MarkdownDescription: "Monitoring interval.",
				Computed:            true,
			},
			"time_period": dschema.Int64Attribute{
				MarkdownDescription: "Monitoring time period.",
				Computed:            true,
			},
			"loss_interval": dschema.Int64Attribute{
				MarkdownDescription: "Monitoring loss interval.",
				Computed:            true,
			},
			"data_length": dschema.Int64Attribute{
				MarkdownDescription: "Monitoring data length.",
				Computed:            true,
			},
		},
	}
}
