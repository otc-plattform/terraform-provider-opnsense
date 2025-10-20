package acmeclient

import (
	"github.com/browningluke/opnsense-go/pkg/acmeclient"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func settingsResponseToModel(resp *acmeclient.SettingsGetResponse) acmeclientSettingsResourceModel {
	settings := resp.ACMEClient.Settings

	return acmeclientSettingsResourceModel{
		Enabled:            types.BoolValue(tools.StringToBool(settings.Enabled)),
		AutoRenewal:        types.BoolValue(tools.StringToBool(settings.AutoRenewal)),
		HAProxyIntegration: types.BoolValue(tools.StringToBool(settings.HAProxyIntegration)),
		LogLevel:           stringValueFromOptionMap(settings.LogLevel),
		ShowIntro:          types.BoolValue(tools.StringToBool(settings.ShowIntro)),
		ChallengePort:      tools.StringToInt64Null(settings.ChallengePort),
		TLSChallengePort:   tools.StringToInt64Null(settings.TLSChallengePort),
		RestartTimeout:     tools.StringToInt64Null(settings.RestartTimeout),
	}
}

func (m *acmeclientSettingsResourceModel) toSettingsSetRequest() acmeclient.SettingsSetRequest {
	if m == nil {
		return acmeclient.SettingsSetRequest{}
	}

	return acmeclient.SettingsSetRequest{
		Settings: acmeclient.SettingsSetSettings{
			Enabled:            boolToAPIString(m.Enabled),
			AutoRenewal:        boolToAPIString(m.AutoRenewal),
			HAProxyIntegration: boolToAPIString(m.HAProxyIntegration),
			LogLevel:           stringToAPIValue(m.LogLevel),
			ShowIntro:          boolToAPIString(m.ShowIntro),
			ChallengePort:      int64ToAPIString(m.ChallengePort),
			TLSChallengePort:   int64ToAPIString(m.TLSChallengePort),
			RestartTimeout:     int64ToAPIString(m.RestartTimeout),
		},
	}
}

func stringValueFromOptionMap(options api.FieldOptions) types.String {
	value := selectedOptionKey(options)
	if value == "" {
		return types.StringNull()
	}
	return types.StringValue(value)
}

func selectedOptionKey(options api.FieldOptions) string {
	for key, option := range options {
		if option.Selected == 1 {
			return key
		}
	}
	return ""
}

func boolToAPIString(value types.Bool) string {
	if value.IsNull() || value.IsUnknown() {
		return tools.BoolToString(false)
	}
	return tools.BoolToString(value.ValueBool())
}

func int64ToAPIString(value types.Int64) string {
	if value.IsNull() || value.IsUnknown() {
		return ""
	}
	return tools.Int64ToString(value.ValueInt64())
}

func stringToAPIValue(value types.String) string {
	if value.IsNull() || value.IsUnknown() {
		return ""
	}
	return value.ValueString()
}
