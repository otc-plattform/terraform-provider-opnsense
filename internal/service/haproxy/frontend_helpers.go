package haproxy

import (
	"context"
	"strings"

	ophaproxy "github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m *frontendResourceModel) toFrontendObject() ophaproxy.HAProxyObject {
	if m == nil {
		return ophaproxy.HAProxyObject{}
	}
	obj := ophaproxy.HAProxyObject{}
	obj["enabled"] = boolToAPIString(m.Enabled)
	obj["name"] = stringToAPIValue(m.Name)
	obj["description"] = stringToAPIValue(m.Description)
	obj["bind"] = strings.Join(tools.SetToStringSlice(m.Bind), ",")
	obj["bindOptions"] = stringToAPIValue(m.BindOptions)
	obj["mode"] = stringToAPIValue(m.Mode)
	obj["defaultBackend"] = stringToAPIValue(m.DefaultBackend)
	obj["ssl_enabled"] = boolToAPIString(m.SslEnabled)
	obj["ssl_certificates"] = strings.Join(tools.SetToStringSlice(m.SslCertificates), ",")
	obj["ssl_default_certificate"] = stringToAPIValue(m.SslDefaultCertificate)
	obj["ssl_customOptions"] = stringToAPIValue(m.SslCustomOptions)
	obj["ssl_advancedEnabled"] = boolToAPIString(m.SslAdvancedEnabled)
	obj["ssl_bindOptions"] = strings.Join(tools.SetToStringSlice(m.SslBindOptions), ",")
	obj["ssl_minVersion"] = stringToAPIValue(m.SslMinVersion)
	obj["ssl_maxVersion"] = stringToAPIValue(m.SslMaxVersion)
	obj["ssl_cipherList"] = stringToAPIValue(m.SslCipherList)
	obj["ssl_cipherSuites"] = stringToAPIValue(m.SslCipherSuites)
	obj["ssl_hstsEnabled"] = boolToAPIString(m.SslHstsEnabled)
	obj["ssl_hstsIncludeSubDomains"] = boolToAPIString(m.SslHstsIncludeSubDomains)
	obj["ssl_hstsPreload"] = boolToAPIString(m.SslHstsPreload)
	obj["ssl_hstsMaxAge"] = stringToAPIValue(m.SslHstsMaxAge)
	obj["ssl_clientAuthEnabled"] = boolToAPIString(m.SslClientAuthEnabled)
	obj["ssl_clientAuthVerify"] = stringToAPIValue(m.SslClientAuthVerify)
	obj["ssl_clientAuthCAs"] = strings.Join(tools.SetToStringSlice(m.SslClientAuthCAs), ",")
	obj["ssl_clientAuthCRLs"] = strings.Join(tools.SetToStringSlice(m.SslClientAuthCRLs), ",")
	obj["basicAuthEnabled"] = boolToAPIString(m.BasicAuthEnabled)
	obj["basicAuthUsers"] = strings.Join(tools.SetToStringSlice(m.BasicAuthUsers), ",")
	obj["basicAuthGroups"] = strings.Join(tools.SetToStringSlice(m.BasicAuthGroups), ",")
	obj["tuning_maxConnections"] = stringToAPIValue(m.TuningMaxConnections)
	obj["tuning_timeoutClient"] = stringToAPIValue(m.TuningTimeoutClient)
	obj["tuning_timeoutHttpReq"] = stringToAPIValue(m.TuningTimeoutHttpReq)
	obj["tuning_timeoutHttpKeepAlive"] = stringToAPIValue(m.TuningTimeoutHttpKeepAlive)
	obj["linkedCpuAffinityRules"] = strings.Join(tools.SetToStringSlice(m.LinkedCpuAffinityRules), ",")
	obj["tuning_shards"] = stringToAPIValue(m.TuningShards)
	obj["logging_dontLogNull"] = boolToAPIString(m.LoggingDontLogNull)
	obj["logging_dontLogNormal"] = boolToAPIString(m.LoggingDontLogNormal)
	obj["logging_logSeparateErrors"] = boolToAPIString(m.LoggingLogSeparateErrors)
	obj["logging_detailedLog"] = boolToAPIString(m.LoggingDetailedLog)
	obj["logging_socketStats"] = boolToAPIString(m.LoggingSocketStats)
	obj["stickiness_pattern"] = stringToAPIValue(m.StickinessPattern)
	obj["stickiness_dataTypes"] = strings.Join(tools.SetToStringSlice(m.StickinessDataTypes), ",")
	obj["stickiness_expire"] = stringToAPIValue(m.StickinessExpire)
	obj["stickiness_size"] = stringToAPIValue(m.StickinessSize)
	obj["stickiness_counter"] = stringToAPIValue(m.StickinessCounter)
	obj["stickiness_counter_key"] = stringToAPIValue(m.StickinessCounterKey)
	obj["stickiness_length"] = stringToAPIValue(m.StickinessLength)
	obj["stickiness_connRatePeriod"] = stringToAPIValue(m.StickinessConnRatePeriod)
	obj["stickiness_sessRatePeriod"] = stringToAPIValue(m.StickinessSessRatePeriod)
	obj["stickiness_httpReqRatePeriod"] = stringToAPIValue(m.StickinessHttpReqRatePeriod)
	obj["stickiness_httpErrRatePeriod"] = stringToAPIValue(m.StickinessHttpErrRatePeriod)
	obj["stickiness_bytesInRatePeriod"] = stringToAPIValue(m.StickinessBytesInRatePeriod)
	obj["stickiness_bytesOutRatePeriod"] = stringToAPIValue(m.StickinessBytesOutRatePeriod)
	obj["stickiness_gpcElements"] = stringToAPIValue(m.StickinessGpcElements)
	obj["stickiness_gpcRatePeriod"] = stringToAPIValue(m.StickinessGpcRatePeriod)
	obj["stickiness_gptElements"] = stringToAPIValue(m.StickinessGptElements)
	obj["stickiness_httpFailRatePeriod"] = stringToAPIValue(m.StickinessHttpFailRatePeriod)
	obj["stickiness_glitchRatePeriod"] = stringToAPIValue(m.StickinessGlitchRatePeriod)
	obj["http2Enabled"] = boolToAPIString(m.Http2Enabled)
	obj["http2Enabled_nontls"] = boolToAPIString(m.Http2EnabledNontls)
	obj["advertised_protocols"] = strings.Join(tools.SetToStringSlice(m.AdvertisedProtocols), ",")
	obj["forwardFor"] = stringToAPIValue(m.ForwardFor)
	obj["prometheus_enabled"] = boolToAPIString(m.PrometheusEnabled)
	obj["prometheus_path"] = stringToAPIValue(m.PrometheusPath)
	obj["connectionBehaviour"] = stringToAPIValue(m.ConnectionBehaviour)
	obj["customOptions"] = stringToAPIValue(m.CustomOptions)
	obj["linkedActions"] = strings.Join(tools.SetToStringSlice(m.LinkedActions), ",")
	obj["linkedErrorfiles"] = strings.Join(tools.SetToStringSlice(m.LinkedErrorfiles), ",")
	return obj
}

func frontendResponseToModel(ctx context.Context, id string, obj ophaproxy.HAProxyObject) frontendResourceModel {
	m := frontendResourceModel{}
	m.Id = types.StringValue(id)
	m.InternalId = types.StringValue(apiValueToString(obj["id"]))
	m.Enabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["enabled"])))
	m.Name = types.StringValue(apiValueToString(obj["name"]))
	m.Description = types.StringValue(apiValueToString(obj["description"]))
	m.Bind = tools.StringSliceToSet(objectSetKeys(obj["bind"]))
	m.BindOptions = types.StringValue(apiValueToString(obj["bindOptions"]))
	m.Mode = types.StringValue(apiValueToString(obj["mode"]))
	m.DefaultBackend = types.StringValue(apiValueToString(obj["defaultBackend"]))
	m.SslEnabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["ssl_enabled"])))
	m.SslCertificates = tools.StringSliceToSet(objectSetKeys(obj["ssl_certificates"]))
	m.SslDefaultCertificate = types.StringValue(apiValueToString(obj["ssl_default_certificate"]))
	m.SslCustomOptions = types.StringValue(apiValueToString(obj["ssl_customOptions"]))
	m.SslAdvancedEnabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["ssl_advancedEnabled"])))
	m.SslBindOptions = tools.StringSliceToSet(objectSetKeys(obj["ssl_bindOptions"]))
	m.SslMinVersion = types.StringValue(apiValueToString(obj["ssl_minVersion"]))
	m.SslMaxVersion = types.StringValue(apiValueToString(obj["ssl_maxVersion"]))
	m.SslCipherList = types.StringValue(apiValueToString(obj["ssl_cipherList"]))
	m.SslCipherSuites = types.StringValue(apiValueToString(obj["ssl_cipherSuites"]))
	m.SslHstsEnabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["ssl_hstsEnabled"])))
	m.SslHstsIncludeSubDomains = types.BoolValue(tools.StringToBool(apiValueToString(obj["ssl_hstsIncludeSubDomains"])))
	m.SslHstsPreload = types.BoolValue(tools.StringToBool(apiValueToString(obj["ssl_hstsPreload"])))
	m.SslHstsMaxAge = types.StringValue(apiValueToString(obj["ssl_hstsMaxAge"]))
	m.SslClientAuthEnabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["ssl_clientAuthEnabled"])))
	m.SslClientAuthVerify = types.StringValue(apiValueToString(obj["ssl_clientAuthVerify"]))
	m.SslClientAuthCAs = tools.StringSliceToSet(objectSetKeys(obj["ssl_clientAuthCAs"]))
	m.SslClientAuthCRLs = tools.StringSliceToSet(objectSetKeys(obj["ssl_clientAuthCRLs"]))
	m.BasicAuthEnabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["basicAuthEnabled"])))
	m.BasicAuthUsers = tools.StringSliceToSet(objectSetKeys(obj["basicAuthUsers"]))
	m.BasicAuthGroups = tools.StringSliceToSet(objectSetKeys(obj["basicAuthGroups"]))
	m.TuningMaxConnections = types.StringValue(apiValueToString(obj["tuning_maxConnections"]))
	m.TuningTimeoutClient = types.StringValue(apiValueToString(obj["tuning_timeoutClient"]))
	m.TuningTimeoutHttpReq = types.StringValue(apiValueToString(obj["tuning_timeoutHttpReq"]))
	m.TuningTimeoutHttpKeepAlive = types.StringValue(apiValueToString(obj["tuning_timeoutHttpKeepAlive"]))
	m.LinkedCpuAffinityRules = tools.StringSliceToSet(objectSetKeys(obj["linkedCpuAffinityRules"]))
	m.TuningShards = types.StringValue(apiValueToString(obj["tuning_shards"]))
	m.LoggingDontLogNull = types.BoolValue(tools.StringToBool(apiValueToString(obj["logging_dontLogNull"])))
	m.LoggingDontLogNormal = types.BoolValue(tools.StringToBool(apiValueToString(obj["logging_dontLogNormal"])))
	m.LoggingLogSeparateErrors = types.BoolValue(tools.StringToBool(apiValueToString(obj["logging_logSeparateErrors"])))
	m.LoggingDetailedLog = types.BoolValue(tools.StringToBool(apiValueToString(obj["logging_detailedLog"])))
	m.LoggingSocketStats = types.BoolValue(tools.StringToBool(apiValueToString(obj["logging_socketStats"])))
	m.StickinessPattern = types.StringValue(apiValueToString(obj["stickiness_pattern"]))
	m.StickinessDataTypes = tools.StringSliceToSet(objectSetKeys(obj["stickiness_dataTypes"]))
	m.StickinessExpire = types.StringValue(apiValueToString(obj["stickiness_expire"]))
	m.StickinessSize = types.StringValue(apiValueToString(obj["stickiness_size"]))
	m.StickinessCounter = types.StringValue(apiValueToString(obj["stickiness_counter"]))
	m.StickinessCounterKey = types.StringValue(apiValueToString(obj["stickiness_counter_key"]))
	m.StickinessLength = types.StringValue(apiValueToString(obj["stickiness_length"]))
	m.StickinessConnRatePeriod = types.StringValue(apiValueToString(obj["stickiness_connRatePeriod"]))
	m.StickinessSessRatePeriod = types.StringValue(apiValueToString(obj["stickiness_sessRatePeriod"]))
	m.StickinessHttpReqRatePeriod = types.StringValue(apiValueToString(obj["stickiness_httpReqRatePeriod"]))
	m.StickinessHttpErrRatePeriod = types.StringValue(apiValueToString(obj["stickiness_httpErrRatePeriod"]))
	m.StickinessBytesInRatePeriod = types.StringValue(apiValueToString(obj["stickiness_bytesInRatePeriod"]))
	m.StickinessBytesOutRatePeriod = types.StringValue(apiValueToString(obj["stickiness_bytesOutRatePeriod"]))
	m.StickinessGpcElements = types.StringValue(apiValueToString(obj["stickiness_gpcElements"]))
	m.StickinessGpcRatePeriod = types.StringValue(apiValueToString(obj["stickiness_gpcRatePeriod"]))
	m.StickinessGptElements = types.StringValue(apiValueToString(obj["stickiness_gptElements"]))
	m.StickinessHttpFailRatePeriod = types.StringValue(apiValueToString(obj["stickiness_httpFailRatePeriod"]))
	m.StickinessGlitchRatePeriod = types.StringValue(apiValueToString(obj["stickiness_glitchRatePeriod"]))
	m.Http2Enabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["http2Enabled"])))
	m.Http2EnabledNontls = types.BoolValue(tools.StringToBool(apiValueToString(obj["http2Enabled_nontls"])))
	m.AdvertisedProtocols = tools.StringSliceToSet(objectSetKeys(obj["advertised_protocols"]))
	m.ForwardFor = types.StringValue(apiValueToString(obj["forwardFor"]))
	m.PrometheusEnabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["prometheus_enabled"])))
	m.PrometheusPath = types.StringValue(apiValueToString(obj["prometheus_path"]))
	m.ConnectionBehaviour = types.StringValue(apiValueToString(obj["connectionBehaviour"]))
	m.CustomOptions = types.StringValue(apiValueToString(obj["customOptions"]))
	m.LinkedActions = tools.StringSliceToSet(objectSetKeys(obj["linkedActions"]))
	m.LinkedErrorfiles = tools.StringSliceToSet(objectSetKeys(obj["linkedErrorfiles"]))
	return m
}

func fetchFrontendModel(ctx context.Context, c *ophaproxy.Controller, id string) (frontendResourceModel, error) {
	obj, err := c.HAProxyGetFrontend(ctx, id)
	if err != nil {
		return frontendResourceModel{}, err
	}
	return frontendResponseToModel(ctx, id, obj), nil
}

func findFrontendIDByName(ctx context.Context, c *ophaproxy.Controller, name string) (string, bool, error) {
	result, err := c.HAProxySearchFrontends(ctx)
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
