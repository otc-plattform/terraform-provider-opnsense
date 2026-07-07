package haproxy

import (
	"context"
	"strings"

	ophaproxy "github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m *backendResourceModel) toBackendObject() ophaproxy.HAProxyObject {
	if m == nil {
		return ophaproxy.HAProxyObject{}
	}
	obj := ophaproxy.HAProxyObject{}
	obj["enabled"] = boolToAPIString(m.Enabled)
	obj["name"] = stringToAPIValue(m.Name)
	obj["description"] = stringToAPIValue(m.Description)
	obj["mode"] = stringToAPIValue(m.Mode)
	obj["algorithm"] = stringToAPIValue(m.Algorithm)
	obj["random_draws"] = stringToAPIValue(m.RandomDraws)
	obj["proxyProtocol"] = stringToAPIValue(m.ProxyProtocol)
	obj["linkedServers"] = strings.Join(tools.SetToStringSlice(m.LinkedServers), ",")
	obj["linkedFcgi"] = stringToAPIValue(m.LinkedFcgi)
	obj["linkedResolver"] = stringToAPIValue(m.LinkedResolver)
	obj["resolverOpts"] = strings.Join(tools.SetToStringSlice(m.ResolverOpts), ",")
	obj["resolvePrefer"] = stringToAPIValue(m.ResolvePrefer)
	obj["source"] = stringToAPIValue(m.Source)
	obj["healthCheckEnabled"] = boolToAPIString(m.HealthCheckEnabled)
	obj["healthCheck"] = stringToAPIValue(m.HealthCheck)
	obj["healthCheckLogStatus"] = stringToAPIValue(m.HealthCheckLogStatus)
	obj["checkInterval"] = stringToAPIValue(m.CheckInterval)
	obj["checkDownInterval"] = stringToAPIValue(m.CheckDownInterval)
	obj["healthCheckFall"] = stringToAPIValue(m.HealthCheckFall)
	obj["healthCheckRise"] = stringToAPIValue(m.HealthCheckRise)
	obj["linkedMailer"] = stringToAPIValue(m.LinkedMailer)
	obj["healthCheckProxyProto"] = stringToAPIValue(m.HealthCheckProxyProto)
	obj["http2Enabled"] = boolToAPIString(m.Http2Enabled)
	obj["http2Enabled_nontls"] = boolToAPIString(m.Http2EnabledNontls)
	obj["ba_advertised_protocols"] = strings.Join(tools.SetToStringSlice(m.BaAdvertisedProtocols), ",")
	obj["forwardFor"] = stringToAPIValue(m.ForwardFor)
	obj["forwardedHeader"] = stringToAPIValue(m.ForwardedHeader)
	obj["forwardedHeaderParameters"] = strings.Join(tools.SetToStringSlice(m.ForwardedHeaderParameters), ",")
	obj["persistence"] = stringToAPIValue(m.Persistence)
	obj["persistence_cookiemode"] = stringToAPIValue(m.PersistenceCookiemode)
	obj["persistence_cookiename"] = stringToAPIValue(m.PersistenceCookiename)
	obj["persistence_stripquotes"] = boolToAPIString(m.PersistenceStripquotes)
	obj["stickiness_pattern"] = stringToAPIValue(m.StickinessPattern)
	obj["stickiness_dataTypes"] = strings.Join(tools.SetToStringSlice(m.StickinessDataTypes), ",")
	obj["stickiness_expire"] = stringToAPIValue(m.StickinessExpire)
	obj["stickiness_size"] = stringToAPIValue(m.StickinessSize)
	obj["stickiness_cookiename"] = stringToAPIValue(m.StickinessCookiename)
	obj["stickiness_cookielength"] = stringToAPIValue(m.StickinessCookielength)
	obj["stickiness_length"] = stringToAPIValue(m.StickinessLength)
	obj["stickiness_connRatePeriod"] = stringToAPIValue(m.StickinessConnRatePeriod)
	obj["stickiness_sessRatePeriod"] = stringToAPIValue(m.StickinessSessRatePeriod)
	obj["stickiness_httpReqRatePeriod"] = stringToAPIValue(m.StickinessHttpReqRatePeriod)
	obj["stickiness_httpErrRatePeriod"] = stringToAPIValue(m.StickinessHttpErrRatePeriod)
	obj["stickiness_bytesInRatePeriod"] = stringToAPIValue(m.StickinessBytesInRatePeriod)
	obj["stickiness_bytesOutRatePeriod"] = stringToAPIValue(m.StickinessBytesOutRatePeriod)
	obj["stickiness_gpcElements"] = stringToAPIValue(m.StickinessGpcElements)
	obj["stickiness_gptElements"] = stringToAPIValue(m.StickinessGptElements)
	obj["stickiness_gpcRatePeriod"] = stringToAPIValue(m.StickinessGpcRatePeriod)
	obj["stickiness_httpFailRatePeriod"] = stringToAPIValue(m.StickinessHttpFailRatePeriod)
	obj["stickiness_glitchRatePeriod"] = stringToAPIValue(m.StickinessGlitchRatePeriod)
	obj["basicAuthEnabled"] = boolToAPIString(m.BasicAuthEnabled)
	obj["basicAuthUsers"] = strings.Join(tools.SetToStringSlice(m.BasicAuthUsers), ",")
	obj["basicAuthGroups"] = strings.Join(tools.SetToStringSlice(m.BasicAuthGroups), ",")
	obj["tuning_timeoutConnect"] = stringToAPIValue(m.TuningTimeoutConnect)
	obj["tuning_timeoutCheck"] = stringToAPIValue(m.TuningTimeoutCheck)
	obj["tuning_timeoutServer"] = stringToAPIValue(m.TuningTimeoutServer)
	obj["tuning_retries"] = stringToAPIValue(m.TuningRetries)
	obj["customOptions"] = stringToAPIValue(m.CustomOptions)
	obj["tuning_defaultserver"] = stringToAPIValue(m.TuningDefaultserver)
	obj["tuning_noport"] = boolToAPIString(m.TuningNoport)
	obj["tuning_httpreuse"] = stringToAPIValue(m.TuningHttpreuse)
	obj["tuning_caching"] = boolToAPIString(m.TuningCaching)
	obj["linkedActions"] = strings.Join(tools.SetToStringSlice(m.LinkedActions), ",")
	obj["linkedErrorfiles"] = strings.Join(tools.SetToStringSlice(m.LinkedErrorfiles), ",")
	return obj
}

func backendResponseToModel(ctx context.Context, id string, obj ophaproxy.HAProxyObject) backendResourceModel {
	m := backendResourceModel{}
	m.Id = types.StringValue(id)
	m.InternalId = types.StringValue(apiValueToString(obj["id"]))
	m.Enabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["enabled"])))
	m.Name = types.StringValue(apiValueToString(obj["name"]))
	m.Description = types.StringValue(apiValueToString(obj["description"]))
	m.Mode = types.StringValue(apiValueToString(obj["mode"]))
	m.Algorithm = types.StringValue(apiValueToString(obj["algorithm"]))
	m.RandomDraws = types.StringValue(apiValueToString(obj["random_draws"]))
	m.ProxyProtocol = types.StringValue(apiValueToString(obj["proxyProtocol"]))
	m.LinkedServers = tools.StringSliceToSet(objectSetKeys(obj["linkedServers"]))
	m.LinkedFcgi = types.StringValue(apiValueToString(obj["linkedFcgi"]))
	m.LinkedResolver = types.StringValue(apiValueToString(obj["linkedResolver"]))
	m.ResolverOpts = tools.StringSliceToSet(objectSetKeys(obj["resolverOpts"]))
	m.ResolvePrefer = types.StringValue(apiValueToString(obj["resolvePrefer"]))
	m.Source = types.StringValue(apiValueToString(obj["source"]))
	m.HealthCheckEnabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["healthCheckEnabled"])))
	m.HealthCheck = types.StringValue(apiValueToString(obj["healthCheck"]))
	m.HealthCheckLogStatus = types.StringValue(apiValueToString(obj["healthCheckLogStatus"]))
	m.CheckInterval = types.StringValue(apiValueToString(obj["checkInterval"]))
	m.CheckDownInterval = types.StringValue(apiValueToString(obj["checkDownInterval"]))
	m.HealthCheckFall = types.StringValue(apiValueToString(obj["healthCheckFall"]))
	m.HealthCheckRise = types.StringValue(apiValueToString(obj["healthCheckRise"]))
	m.LinkedMailer = types.StringValue(apiValueToString(obj["linkedMailer"]))
	m.HealthCheckProxyProto = types.StringValue(apiValueToString(obj["healthCheckProxyProto"]))
	m.Http2Enabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["http2Enabled"])))
	m.Http2EnabledNontls = types.BoolValue(tools.StringToBool(apiValueToString(obj["http2Enabled_nontls"])))
	m.BaAdvertisedProtocols = tools.StringSliceToSet(objectSetKeys(obj["ba_advertised_protocols"]))
	m.ForwardFor = types.StringValue(apiValueToString(obj["forwardFor"]))
	m.ForwardedHeader = types.StringValue(apiValueToString(obj["forwardedHeader"]))
	m.ForwardedHeaderParameters = tools.StringSliceToSet(objectSetKeys(obj["forwardedHeaderParameters"]))
	m.Persistence = types.StringValue(apiValueToString(obj["persistence"]))
	m.PersistenceCookiemode = types.StringValue(apiValueToString(obj["persistence_cookiemode"]))
	m.PersistenceCookiename = types.StringValue(apiValueToString(obj["persistence_cookiename"]))
	m.PersistenceStripquotes = types.BoolValue(tools.StringToBool(apiValueToString(obj["persistence_stripquotes"])))
	m.StickinessPattern = types.StringValue(apiValueToString(obj["stickiness_pattern"]))
	m.StickinessDataTypes = tools.StringSliceToSet(objectSetKeys(obj["stickiness_dataTypes"]))
	m.StickinessExpire = types.StringValue(apiValueToString(obj["stickiness_expire"]))
	m.StickinessSize = types.StringValue(apiValueToString(obj["stickiness_size"]))
	m.StickinessCookiename = types.StringValue(apiValueToString(obj["stickiness_cookiename"]))
	m.StickinessCookielength = types.StringValue(apiValueToString(obj["stickiness_cookielength"]))
	m.StickinessLength = types.StringValue(apiValueToString(obj["stickiness_length"]))
	m.StickinessConnRatePeriod = types.StringValue(apiValueToString(obj["stickiness_connRatePeriod"]))
	m.StickinessSessRatePeriod = types.StringValue(apiValueToString(obj["stickiness_sessRatePeriod"]))
	m.StickinessHttpReqRatePeriod = types.StringValue(apiValueToString(obj["stickiness_httpReqRatePeriod"]))
	m.StickinessHttpErrRatePeriod = types.StringValue(apiValueToString(obj["stickiness_httpErrRatePeriod"]))
	m.StickinessBytesInRatePeriod = types.StringValue(apiValueToString(obj["stickiness_bytesInRatePeriod"]))
	m.StickinessBytesOutRatePeriod = types.StringValue(apiValueToString(obj["stickiness_bytesOutRatePeriod"]))
	m.StickinessGpcElements = types.StringValue(apiValueToString(obj["stickiness_gpcElements"]))
	m.StickinessGptElements = types.StringValue(apiValueToString(obj["stickiness_gptElements"]))
	m.StickinessGpcRatePeriod = types.StringValue(apiValueToString(obj["stickiness_gpcRatePeriod"]))
	m.StickinessHttpFailRatePeriod = types.StringValue(apiValueToString(obj["stickiness_httpFailRatePeriod"]))
	m.StickinessGlitchRatePeriod = types.StringValue(apiValueToString(obj["stickiness_glitchRatePeriod"]))
	m.BasicAuthEnabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["basicAuthEnabled"])))
	m.BasicAuthUsers = tools.StringSliceToSet(objectSetKeys(obj["basicAuthUsers"]))
	m.BasicAuthGroups = tools.StringSliceToSet(objectSetKeys(obj["basicAuthGroups"]))
	m.TuningTimeoutConnect = types.StringValue(apiValueToString(obj["tuning_timeoutConnect"]))
	m.TuningTimeoutCheck = types.StringValue(apiValueToString(obj["tuning_timeoutCheck"]))
	m.TuningTimeoutServer = types.StringValue(apiValueToString(obj["tuning_timeoutServer"]))
	m.TuningRetries = types.StringValue(apiValueToString(obj["tuning_retries"]))
	m.CustomOptions = types.StringValue(apiValueToString(obj["customOptions"]))
	m.TuningDefaultserver = types.StringValue(apiValueToString(obj["tuning_defaultserver"]))
	m.TuningNoport = types.BoolValue(tools.StringToBool(apiValueToString(obj["tuning_noport"])))
	m.TuningHttpreuse = types.StringValue(apiValueToString(obj["tuning_httpreuse"]))
	m.TuningCaching = types.BoolValue(tools.StringToBool(apiValueToString(obj["tuning_caching"])))
	m.LinkedActions = tools.StringSliceToSet(objectSetKeys(obj["linkedActions"]))
	m.LinkedErrorfiles = tools.StringSliceToSet(objectSetKeys(obj["linkedErrorfiles"]))
	return m
}

func fetchBackendModel(ctx context.Context, c *ophaproxy.Controller, id string) (backendResourceModel, error) {
	obj, err := c.HAProxyGetBackend(ctx, id)
	if err != nil {
		return backendResourceModel{}, err
	}
	return backendResponseToModel(ctx, id, obj), nil
}

func findBackendIDByName(ctx context.Context, c *ophaproxy.Controller, name string) (string, bool, error) {
	result, err := c.HAProxySearchBackends(ctx)
	if err != nil {
		return "", false, err
	}
	for _, row := range result.Rows {
		if apiValueToString(row["name"]) == name {
			return apiValueToString(row["uuid"]), true, nil
		}
	}
	return "", false, nil
}
