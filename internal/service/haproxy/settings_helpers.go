package haproxy

import (
	"strings"

	"github.com/browningluke/opnsense-go/pkg/api"
	ophaproxy "github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func settingsResponseToModel(resp *ophaproxy.SettingsGetResponse) haproxySettingsResourceModel {
	general := resp.HAProxy.General
	tuning := general.Tuning
	defaults := general.Defaults
	logging := general.Logging
	stats := general.Stats
	cache := general.Cache
	peers := general.Peers

	return haproxySettingsResourceModel{
		Enabled:                         types.BoolValue(tools.StringToBool(general.Enabled)),
		GracefulStop:                    types.BoolValue(tools.StringToBool(general.GracefulStop)),
		HardStopAfter:                   types.StringValue(general.HardStopAfter),
		CloseSpreadTime:                 types.StringValue(general.CloseSpreadTime),
		SeamlessReload:                  types.BoolValue(tools.StringToBool(general.SeamlessReload)),
		StoreOCSP:                       types.BoolValue(tools.StringToBool(general.StoreOCSP)),
		ShowIntro:                       types.BoolValue(tools.StringToBool(general.ShowIntro)),
		TuningRoot:                      types.BoolValue(tools.StringToBool(tuning.Root)),
		TuningMaxConnections:            tools.StringToInt64Null(tuning.MaxConnections),
		TuningNbthread:                  tools.StringToInt64Null(tuning.Nbthread),
		TuningResolversPrefer:           types.StringValue(selectedOptionKey(tuning.ResolversPrefer)),
		TuningSSLServerVerify:           types.StringValue(selectedOptionKey(tuning.SSLServerVerify)),
		TuningMaxDHSize:                 tools.StringToInt64Null(tuning.MaxDHSize),
		TuningBufferSize:                tools.StringToInt64Null(tuning.BufferSize),
		TuningSpreadChecks:              tools.StringToInt64Null(tuning.SpreadChecks),
		TuningBogusProxyEnabled:         types.BoolValue(tools.StringToBool(tuning.BogusProxyEnabled)),
		TuningLuaMaxMem:                 tools.StringToInt64Null(tuning.LuaMaxMem),
		TuningCustomOptions:             types.StringValue(tuning.CustomOptions),
		TuningOCSPUpdateEnabled:         types.BoolValue(tools.StringToBool(tuning.OCSPUpdateEnabled)),
		TuningOCSPUpdateMinDelay:        tools.StringToInt64Null(tuning.OCSPUpdateMinDelay),
		TuningOCSPUpdateMaxDelay:        tools.StringToInt64Null(tuning.OCSPUpdateMaxDelay),
		TuningSSLDefaultsEnabled:        types.BoolValue(tools.StringToBool(tuning.SSLDefaultsEnabled)),
		TuningSSLBindOptions:            tools.StringSliceToSet(selectedOptionKeys(tuning.SSLBindOptions)),
		TuningSSLMinVersion:             types.StringValue(selectedOptionKey(tuning.SSLMinVersion)),
		TuningSSLMaxVersion:             types.StringValue(selectedOptionKey(tuning.SSLMaxVersion)),
		TuningSSLCipherList:             types.StringValue(tuning.SSLCipherList),
		TuningSSLCipherSuites:           types.StringValue(tuning.SSLCipherSuites),
		TuningH2InitialWindowSize:       tools.StringToInt64Null(tuning.H2InitialWindowSize),
		TuningH2InitialWindowSizeOut:    tools.StringToInt64Null(tuning.H2InitialWindowSizeOutgoing),
		TuningH2InitialWindowSizeIn:     tools.StringToInt64Null(tuning.H2InitialWindowSizeIncoming),
		TuningH2MaxConcurrentStreams:    tools.StringToInt64Null(tuning.H2MaxConcurrentStreams),
		TuningH2MaxConcurrentStreamsOut: tools.StringToInt64Null(tuning.H2MaxConcurrentStreamsOutgoing),
		TuningH2MaxConcurrentStreamsIn:  tools.StringToInt64Null(tuning.H2MaxConcurrentStreamsIncoming),
		DefaultsMaxConnections:          tools.StringToInt64Null(defaults.MaxConnections),
		DefaultsMaxConnectionsServers:   tools.StringToInt64Null(defaults.MaxConnectionsServers),
		DefaultsTimeoutClient:           types.StringValue(defaults.TimeoutClient),
		DefaultsTimeoutConnect:          types.StringValue(defaults.TimeoutConnect),
		DefaultsTimeoutCheck:            types.StringValue(defaults.TimeoutCheck),
		DefaultsTimeoutServer:           types.StringValue(defaults.TimeoutServer),
		DefaultsRetries:                 tools.StringToInt64Null(defaults.Retries),
		DefaultsRedispatch:              types.StringValue(selectedOptionKey(defaults.Redispatch)),
		DefaultsInitAddr:                tools.StringSliceToSet(selectedOptionKeys(defaults.InitAddr)),
		DefaultsCustomOptions:           types.StringValue(defaults.CustomOptions),
		LoggingHost:                     types.StringValue(logging.Host),
		LoggingFacility:                 types.StringValue(selectedOptionKey(logging.Facility)),
		LoggingLevel:                    types.StringValue(selectedOptionKey(logging.Level)),
		LoggingLength:                   tools.StringToInt64Null(logging.Length),
		StatsEnabled:                    types.BoolValue(tools.StringToBool(stats.Enabled)),
		StatsPort:                       tools.StringToInt64Null(stats.Port),
		StatsRemoteEnabled:              types.BoolValue(tools.StringToBool(stats.RemoteEnabled)),
		StatsRemoteBind:                 types.StringValue(selectedOptionKey(stats.RemoteBind)),
		StatsAuthEnabled:                types.BoolValue(tools.StringToBool(stats.AuthEnabled)),
		StatsAllowedUsers:               tools.StringSliceToSet(stats.AllowedUsers),
		StatsAllowedGroups:              tools.StringSliceToSet(stats.AllowedGroups),
		StatsCustomOptions:              types.StringValue(stats.CustomOptions),
		StatsPrometheusEnabled:          types.BoolValue(tools.StringToBool(stats.PrometheusEnabled)),
		StatsPrometheusBind:             types.StringValue(selectedOptionKey(stats.PrometheusBind)),
		StatsPrometheusPath:             types.StringValue(stats.PrometheusPath),
		CacheEnabled:                    types.BoolValue(tools.StringToBool(cache.Enabled)),
		CacheTotalMaxSize:               tools.StringToInt64Null(cache.TotalMaxSize),
		CacheMaxAge:                     tools.StringToInt64Null(cache.MaxAge),
		CacheMaxObjectSize:              tools.StringToInt64Null(cache.MaxObjectSize),
		CacheProcessVary:                types.BoolValue(tools.StringToBool(cache.ProcessVary)),
		CacheMaxSecondaryEntries:        tools.StringToInt64Null(cache.MaxSecondaryEntries),
		PeersEnabled:                    types.BoolValue(tools.StringToBool(peers.Enabled)),
		PeersName1:                      types.StringValue(peers.Name1),
		PeersListen1:                    types.StringValue(peers.Listen1),
		PeersPort1:                      tools.StringToInt64Null(peers.Port1),
		PeersName2:                      types.StringValue(peers.Name2),
		PeersListen2:                    types.StringValue(peers.Listen2),
		PeersPort2:                      tools.StringToInt64Null(peers.Port2),
	}
}

func (m *haproxySettingsResourceModel) toSettingsSetRequest() ophaproxy.SettingsSetRequest {
	if m == nil {
		return ophaproxy.SettingsSetRequest{}
	}

	return ophaproxy.SettingsSetRequest{
		General: ophaproxy.GeneralSet{
			Enabled:         boolToAPIString(m.Enabled),
			GracefulStop:    boolToAPIString(m.GracefulStop),
			HardStopAfter:   stringToAPIValue(m.HardStopAfter),
			CloseSpreadTime: stringToAPIValue(m.CloseSpreadTime),
			SeamlessReload:  boolToAPIString(m.SeamlessReload),
			StoreOCSP:       boolToAPIString(m.StoreOCSP),
			ShowIntro:       boolToAPIString(m.ShowIntro),
			Tuning: &ophaproxy.TuningSet{
				Root:                           boolToAPIString(m.TuningRoot),
				MaxConnections:                 int64ToAPIString(m.TuningMaxConnections),
				Nbthread:                       int64ToAPIString(m.TuningNbthread),
				ResolversPrefer:                stringToAPIValue(m.TuningResolversPrefer),
				SSLServerVerify:                stringToAPIValue(m.TuningSSLServerVerify),
				MaxDHSize:                      int64ToAPIString(m.TuningMaxDHSize),
				BufferSize:                     int64ToAPIString(m.TuningBufferSize),
				SpreadChecks:                   int64ToAPIString(m.TuningSpreadChecks),
				BogusProxyEnabled:              boolToAPIString(m.TuningBogusProxyEnabled),
				LuaMaxMem:                      int64ToAPIString(m.TuningLuaMaxMem),
				CustomOptions:                  stringToAPIValue(m.TuningCustomOptions),
				OCSPUpdateEnabled:              boolToAPIString(m.TuningOCSPUpdateEnabled),
				OCSPUpdateMinDelay:             int64ToAPIString(m.TuningOCSPUpdateMinDelay),
				OCSPUpdateMaxDelay:             int64ToAPIString(m.TuningOCSPUpdateMaxDelay),
				SSLDefaultsEnabled:             boolToAPIString(m.TuningSSLDefaultsEnabled),
				SSLBindOptions:                 strings.Join(tools.SetToStringSlice(m.TuningSSLBindOptions), ","),
				SSLMinVersion:                  stringToAPIValue(m.TuningSSLMinVersion),
				SSLMaxVersion:                  stringToAPIValue(m.TuningSSLMaxVersion),
				SSLCipherList:                  stringToAPIValue(m.TuningSSLCipherList),
				SSLCipherSuites:                stringToAPIValue(m.TuningSSLCipherSuites),
				H2InitialWindowSize:            int64ToAPIString(m.TuningH2InitialWindowSize),
				H2InitialWindowSizeOutgoing:    int64ToAPIString(m.TuningH2InitialWindowSizeOut),
				H2InitialWindowSizeIncoming:    int64ToAPIString(m.TuningH2InitialWindowSizeIn),
				H2MaxConcurrentStreams:         int64ToAPIString(m.TuningH2MaxConcurrentStreams),
				H2MaxConcurrentStreamsOutgoing: int64ToAPIString(m.TuningH2MaxConcurrentStreamsOut),
				H2MaxConcurrentStreamsIncoming: int64ToAPIString(m.TuningH2MaxConcurrentStreamsIn),
			},
			Defaults: &ophaproxy.DefaultsSet{
				MaxConnections:        int64ToAPIString(m.DefaultsMaxConnections),
				MaxConnectionsServers: int64ToAPIString(m.DefaultsMaxConnectionsServers),
				TimeoutClient:         stringToAPIValue(m.DefaultsTimeoutClient),
				TimeoutConnect:        stringToAPIValue(m.DefaultsTimeoutConnect),
				TimeoutCheck:          stringToAPIValue(m.DefaultsTimeoutCheck),
				TimeoutServer:         stringToAPIValue(m.DefaultsTimeoutServer),
				Retries:               int64ToAPIString(m.DefaultsRetries),
				Redispatch:            stringToAPIValue(m.DefaultsRedispatch),
				InitAddr:              strings.Join(tools.SetToStringSlice(m.DefaultsInitAddr), ","),
				CustomOptions:         stringToAPIValue(m.DefaultsCustomOptions),
			},
			Logging: &ophaproxy.LoggingSet{
				Host:     stringToAPIValue(m.LoggingHost),
				Facility: stringToAPIValue(m.LoggingFacility),
				Level:    stringToAPIValue(m.LoggingLevel),
				Length:   int64ToAPIString(m.LoggingLength),
			},
			Stats: &ophaproxy.StatsSet{
				Enabled:           boolToAPIString(m.StatsEnabled),
				Port:              int64ToAPIString(m.StatsPort),
				RemoteEnabled:     boolToAPIString(m.StatsRemoteEnabled),
				RemoteBind:        stringToAPIValue(m.StatsRemoteBind),
				AuthEnabled:       boolToAPIString(m.StatsAuthEnabled),
				AllowedUsers:      strings.Join(tools.SetToStringSlice(m.StatsAllowedUsers), ","),
				AllowedGroups:     strings.Join(tools.SetToStringSlice(m.StatsAllowedGroups), ","),
				CustomOptions:     stringToAPIValue(m.StatsCustomOptions),
				PrometheusEnabled: boolToAPIString(m.StatsPrometheusEnabled),
				PrometheusBind:    stringToAPIValue(m.StatsPrometheusBind),
				PrometheusPath:    stringToAPIValue(m.StatsPrometheusPath),
			},
			Cache: &ophaproxy.Cache{
				Enabled:             boolToAPIString(m.CacheEnabled),
				TotalMaxSize:        int64ToAPIString(m.CacheTotalMaxSize),
				MaxAge:              int64ToAPIString(m.CacheMaxAge),
				MaxObjectSize:       int64ToAPIString(m.CacheMaxObjectSize),
				ProcessVary:         boolToAPIString(m.CacheProcessVary),
				MaxSecondaryEntries: int64ToAPIString(m.CacheMaxSecondaryEntries),
			},
			Peers: &ophaproxy.Peers{
				Enabled: boolToAPIString(m.PeersEnabled),
				Name1:   stringToAPIValue(m.PeersName1),
				Listen1: stringToAPIValue(m.PeersListen1),
				Port1:   int64ToAPIString(m.PeersPort1),
				Name2:   stringToAPIValue(m.PeersName2),
				Listen2: stringToAPIValue(m.PeersListen2),
				Port2:   int64ToAPIString(m.PeersPort2),
			},
		},
	}
}

func selectedOptionKeys(options api.FieldOptions) []string {
	keys := make([]string, 0, len(options))
	for key, option := range options {
		if option.Selected == 1 {
			keys = append(keys, key)
		}
	}
	return keys
}

func selectedOptionKey(options api.FieldOptions) string {
	keys := selectedOptionKeys(options)
	if len(keys) == 0 {
		return ""
	}
	return keys[0]
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
