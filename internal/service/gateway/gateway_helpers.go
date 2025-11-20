package gateway

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/gateway"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func convertGatewaySchemaToRequest(d *gatewayResourceModel) gateway.GatewayRequest {
	if d == nil {
		return gateway.GatewayRequest{}
	}

	return gateway.GatewayRequest{
		Disabled:                  boolToAPIString(d.Disabled),
		Name:                      stringValue(d.Name),
		Interface:                 stringValue(d.Interface),
		IPProtocol:                stringValue(d.IpProtocol),
		Priority:                  int64ToAPIString(d.Priority),
		Gateway:                   stringValue(d.Gateway),
		DefaultGW:                 boolToAPIString(d.DefaultGateway),
		FarGW:                     boolToAPIString(d.FarGateway),
		MonitorDisable:            boolToAPIString(d.MonitorDisable),
		MonitorNoRoute:            boolToAPIString(d.MonitorNoRoute),
		MonitorKillStates:         boolToAPIString(d.MonitorKillStates),
		MonitorKillStatesPriority: int64ToAPIString(d.MonitorKillStatesPriority),
		Monitor:                   stringValue(d.Monitor),
		ForceDown:                 boolToAPIString(d.ForceDown),
		Weight:                    int64ToAPIString(d.Weight),
		LatencyLow:                int64ToAPIString(d.LatencyLow),
		LatencyHigh:               int64ToAPIString(d.LatencyHigh),
		LossLow:                   int64ToAPIString(d.LossLow),
		LossHigh:                  int64ToAPIString(d.LossHigh),
		Interval:                  int64ToAPIString(d.Interval),
		TimePeriod:                int64ToAPIString(d.TimePeriod),
		LossInterval:              int64ToAPIString(d.LossInterval),
		DataLength:                int64ToAPIString(d.DataLength),
		Description:               stringValue(d.Description),
	}
}

func gatewayResponseToModel(id string, resp *gateway.GetGatewayResponse) *gatewayResourceModel {
	if resp == nil {
		return nil
	}

	data := resp.Gateway

	model := &gatewayResourceModel{
		Id:                        types.StringValue(id),
		Name:                      types.StringValue(data.Name),
		Description:               tools.StringOrNull(data.Description),
		Interface:                 types.StringValue(selectedOptionKey(data.Interface)),
		IpProtocol:                types.StringValue(selectedOptionKey(data.IPProtocol)),
		Gateway:                   optionalStringValue(data.Gateway),
		Disabled:                  types.BoolValue(tools.StringToBool(data.Disabled)),
		DefaultGateway:            types.BoolValue(tools.StringToBool(data.DefaultGW)),
		FarGateway:                types.BoolValue(tools.StringToBool(data.FarGW)),
		MonitorDisable:            types.BoolValue(tools.StringToBool(data.MonitorDisable)),
		MonitorNoRoute:            types.BoolValue(tools.StringToBool(data.MonitorNoRoute)),
		MonitorKillStates:         types.BoolValue(tools.StringToBool(data.MonitorKillStates)),
		MonitorKillStatesPriority: tools.StringToInt64Null(data.MonitorKillStatesPriority),
		Monitor:                   optionalStringValue(data.Monitor),
		ForceDown:                 types.BoolValue(tools.StringToBool(data.ForceDown)),
		Priority:                  tools.StringToInt64Null(data.Priority),
		Weight:                    tools.StringToInt64Null(data.Weight),
		LatencyLow:                tools.StringToInt64Null(data.LatencyLow),
		LatencyHigh:               tools.StringToInt64Null(data.LatencyHigh),
		LossLow:                   tools.StringToInt64Null(data.LossLow),
		LossHigh:                  tools.StringToInt64Null(data.LossHigh),
		Interval:                  tools.StringToInt64Null(data.Interval),
		TimePeriod:                tools.StringToInt64Null(data.TimePeriod),
		LossInterval:              tools.StringToInt64Null(data.LossInterval),
		DataLength:                tools.StringToInt64Null(data.DataLength),
	}

	return model
}

func fetchGatewayModel(ctx context.Context, ctrl *gateway.Controller, id string) (*gatewayResourceModel, error) {
	resp, err := ctrl.SettingsGetGateway(ctx, id)
	if err != nil {
		return nil, err
	}

	return gatewayResponseToModel(id, resp), nil
}

func boolToAPIString(value types.Bool) string {
	if value.IsNull() || value.IsUnknown() {
		return tools.BoolToString(false)
	}
	return tools.BoolToString(value.ValueBool())
}

func stringValue(value types.String) string {
	if value.IsNull() || value.IsUnknown() {
		return ""
	}
	return value.ValueString()
}

func int64ToAPIString(value types.Int64) string {
	if value.IsNull() || value.IsUnknown() {
		return ""
	}
	return tools.Int64ToString(value.ValueInt64())
}

func selectedOptionKey(options api.FieldOptions) string {
	for key, option := range options {
		if option.Selected == 1 {
			return key
		}
	}
	return ""
}

func optionalStringValue(value string) types.String {
	if value == "" {
		return types.StringNull()
	}
	return types.StringValue(value)
}
