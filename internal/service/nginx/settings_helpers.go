package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func settingsResponseToModel(resp *nginx.SettingsGetResponse) nginxSettingsResourceModel {
	general := resp.NGinx.General

	return nginxSettingsResourceModel{
		Enabled: types.BoolValue(tools.StringToBool(general.Enabled)),
		BanTTL:  tools.StringToInt64Null(general.BanTTL),
	}
}

func (m *nginxSettingsResourceModel) toSettingsSetRequest() nginx.SettingsSetRequest {
	if m == nil {
		return nginx.SettingsSetRequest{}
	}

	return nginx.SettingsSetRequest{
		General: nginx.General{
			Enabled: boolToAPIString(m.Enabled),
			BanTTL:  int64ToAPIString(m.BanTTL),
		},
	}
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
