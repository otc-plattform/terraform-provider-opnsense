package nginx

import (
	"context"
	"sort"
	"strings"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func locationResponseToModel(id string, resp *nginx.LocationGetResponse) nginxLocationResourceModel {
	location := resp.Location

	return nginxLocationResourceModel{
		Id:                      types.StringValue(id),
		Description:             tools.StringOrNull(location.Description),
		URLPattern:              tools.StringOrNull(location.URLPattern),
		MatchType:               stringValueFromOptionMap(location.MatchType),
		PathPrefix:              tools.StringOrNull(location.PathPrefix),
		Upstream:                stringValueFromOptionMap(location.Upstream),
		EnableSecRules:          types.BoolValue(tools.StringToBool(location.EnableSecRules)),
		EnableLearningMode:      types.BoolValue(tools.StringToBool(location.EnableLearningMode)),
		SecRulesErrorPage:       stringValueFromOptionMap(location.SecRulesErrorPage),
		XSSBlockScore:           tools.StringToInt64Null(location.XSSBlockScore),
		SQLiBlockScore:          tools.StringToInt64Null(location.SQLiBlockScore),
		CustomPolicy:            stringSliceToSet(location.CustomPolicy),
		CachePath:               stringValueFromOptionMap(location.CachePath),
		CacheUseStale:           stringValueFromOptionMap(location.CacheUseStale),
		CacheMethods:            stringValueFromOptionMap(location.CacheMethods),
		CacheMinUses:            tools.StringToInt64Null(location.CacheMinUses),
		CacheValid:              tools.StringOrNull(location.CacheValid),
		CacheBackgroundUpdate:   types.BoolValue(tools.StringToBool(location.CacheBackgroundUpdate)),
		CacheLock:               types.BoolValue(tools.StringToBool(location.CacheLock)),
		CacheRevalidate:         types.BoolValue(tools.StringToBool(location.CacheRevalidate)),
		Root:                    tools.StringOrNull(location.Root),
		Rewrites:                stringSliceToSet(location.Rewrites),
		Index:                   stringSliceToSet(location.Index),
		AutoIndex:               types.BoolValue(tools.StringToBool(location.AutoIndex)),
		AuthBasic:               types.BoolValue(tools.StringToBool(location.AuthBasic)),
		AuthBasicUserFile:       stringValueFromOptionMap(location.AuthBasicUserFile),
		AdvancedACL:             types.BoolValue(tools.StringToBool(location.AdvancedACL)),
		ForceHTTPS:              types.BoolValue(tools.StringToBool(location.ForceHTTPS)),
		PHPEnable:               types.BoolValue(tools.StringToBool(location.PHPEnable)),
		PHPOverrideScriptName:   tools.StringOrNull(location.PHPOverrideScriptName),
		LimitRequestConnections: stringSliceToSet(location.LimitRequestConnections),
		MaxBodySize:             tools.StringOrNull(location.MaxBodySize),
		BodyBufferSize:          tools.StringOrNull(location.BodyBufferSize),
		Honeypot:                types.BoolValue(tools.StringToBool(location.Honeypot)),
		WebSocket:               types.BoolValue(tools.StringToBool(location.Websocket)),
		UpstreamKeepalive:       types.BoolValue(tools.StringToBool(location.UpstreamKeepalive)),
		ProxyBufferSize:         tools.StringOrNull(location.ProxyBufferSize),
		ProxyBuffersCount:       tools.StringOrNull(location.ProxyBuffersCount),
		ProxyBuffersSize:        tools.StringOrNull(location.ProxyBuffersSize),
		ProxyBusyBuffersSize:    tools.StringOrNull(location.ProxyBusyBuffersSize),
		ProxyIgnoreClientAbort:  types.BoolValue(tools.StringToBool(location.ProxyIgnoreClientAbort)),
		ProxyRequestBuffering:   types.BoolValue(tools.StringToBool(location.ProxyRequestBuffering)),
		ProxyBuffering:          types.BoolValue(tools.StringToBool(location.ProxyBuffering)),
		ProxyReadTimeout:        tools.StringOrNull(location.ProxyReadTimeout),
		ProxySendTimeout:        tools.StringOrNull(location.ProxySendTimeout),
		IPACL:                   stringValueFromOptionMap(location.IPACL),
		Satisfy:                 stringValueFromOptionMap(location.Satisfy),
		ProxyMaxTempFileSize:    tools.StringOrNull(location.ProxyMaxTempFileSize),
		ProxySSLServerName:      types.BoolValue(tools.StringToBool(location.ProxySSLServerName)),
		ErrorPages:              stringSliceToSet(location.ErrorPages),
	}
}

func (m *nginxLocationResourceModel) toLocation() nginx.Location {
	if m == nil {
		return nginx.Location{}
	}

	return nginx.Location{
		AdvancedACL:             boolToAPIString(m.AdvancedACL),
		AuthBasic:               boolToAPIString(m.AuthBasic),
		AuthBasicUserFile:       stringValue(m.AuthBasicUserFile),
		AutoIndex:               boolToAPIString(m.AutoIndex),
		BodyBufferSize:          stringValue(m.BodyBufferSize),
		CacheBackgroundUpdate:   boolToAPIString(m.CacheBackgroundUpdate),
		CacheLock:               boolToAPIString(m.CacheLock),
		CacheMethods:            stringValue(m.CacheMethods),
		CacheMinUses:            int64ToAPIString(m.CacheMinUses),
		CachePath:               stringValue(m.CachePath),
		CacheRevalidate:         boolToAPIString(m.CacheRevalidate),
		CacheUseStale:           stringValue(m.CacheUseStale),
		CacheValid:              stringValue(m.CacheValid),
		CustomPolicy:            joinStringSet(m.CustomPolicy),
		Description:             stringValue(m.Description),
		EnableLearningMode:      boolToAPIString(m.EnableLearningMode),
		EnableSecRules:          boolToAPIString(m.EnableSecRules),
		ErrorPages:              joinStringSet(m.ErrorPages),
		ForceHTTPS:              boolToAPIString(m.ForceHTTPS),
		Honeypot:                boolToAPIString(m.Honeypot),
		Index:                   joinStringSet(m.Index),
		IPACL:                   stringValue(m.IPACL),
		LimitRequestConnections: joinStringSet(m.LimitRequestConnections),
		MatchType:               stringValue(m.MatchType),
		MaxBodySize:             stringValue(m.MaxBodySize),
		PathPrefix:              stringValue(m.PathPrefix),
		PHPEnable:               boolToAPIString(m.PHPEnable),
		PHPOverrideScriptName:   stringValue(m.PHPOverrideScriptName),
		ProxyBuffering:          boolToAPIString(m.ProxyBuffering),
		ProxyBufferSize:         stringValue(m.ProxyBufferSize),
		ProxyBuffersCount:       stringValue(m.ProxyBuffersCount),
		ProxyBuffersSize:        stringValue(m.ProxyBuffersSize),
		ProxyBusyBuffersSize:    stringValue(m.ProxyBusyBuffersSize),
		ProxyIgnoreClientAbort:  boolToAPIString(m.ProxyIgnoreClientAbort),
		ProxyMaxTempFileSize:    stringValue(m.ProxyMaxTempFileSize),
		ProxyReadTimeout:        stringValue(m.ProxyReadTimeout),
		ProxyRequestBuffering:   boolToAPIString(m.ProxyRequestBuffering),
		ProxySendTimeout:        stringValue(m.ProxySendTimeout),
		ProxySSLServerName:      boolToAPIString(m.ProxySSLServerName),
		Rewrites:                joinStringSet(m.Rewrites),
		Root:                    stringValue(m.Root),
		Satisfy:                 stringValue(m.Satisfy),
		SecRulesErrorPage:       stringValue(m.SecRulesErrorPage),
		SQLiBlockScore:          int64ToAPIString(m.SQLiBlockScore),
		Upstream:                stringValue(m.Upstream),
		UpstreamKeepalive:       boolToAPIString(m.UpstreamKeepalive),
		URLPattern:              stringValue(m.URLPattern),
		WebSocket:               boolToAPIString(m.WebSocket),
		XSSBlockScore:           int64ToAPIString(m.XSSBlockScore),
	}
}

func fetchLocationModel(ctx context.Context, controller *nginx.Controller, id string) (nginxLocationResourceModel, error) {
	resp, err := controller.NginxGetLocation(ctx, id)
	if err != nil {
		return nginxLocationResourceModel{}, err
	}

	return locationResponseToModel(id, resp), nil
}

func joinStringSet(set types.Set) string {
	if set.IsNull() || set.IsUnknown() {
		return ""
	}

	values := tools.SetToStringSlice(set)
	return strings.Join(values, "\n")
}

func stringSliceToSet(values []string) types.Set {
	if len(values) == 0 {
		return types.SetNull(types.StringType)
	}

	return tools.StringSliceToSet(values)
}

func stringValue(value types.String) string {
	if value.IsNull() || value.IsUnknown() {
		return ""
	}

	return value.ValueString()
}

func stringValueFromOptionMap(options api.FieldOptions) types.String {
	key := selectedOptionKey(options)
	if key == "" {
		return types.StringNull()
	}

	return types.StringValue(key)
}

func optionMapToSet(options api.FieldOptions) types.Set {
	keys := selectedOptionKeys(options)
	if len(keys) == 0 {
		return types.SetNull(types.StringType)
	}

	return tools.StringSliceToSet(keys)
}

func selectedOptionKeys(options api.FieldOptions) []string {
	var keys []string
	for key, option := range options {
		if option.Selected == 1 {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	return keys
}

func selectedOptionKey(options api.FieldOptions) string {
	keys := selectedOptionKeys(options)
	if len(keys) == 0 {
		return ""
	}
	return keys[0]
}
