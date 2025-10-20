package nginx

import (
	"context"
	"sort"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func upstreamResponseToModel(id string, resp *nginx.UpstreamGetResponse) nginxUpstreamResourceModel {
	upstream := resp.Upstream

	return nginxUpstreamResourceModel{
		Id:                     types.StringValue(id),
		Description:            tools.StringOrNull(upstream.Description),
		ServerEntries:          fieldOptionsToSet(upstream.ServerEntries),
		LoadBalancingAlgorithm: fieldOptionsFirstKey(upstream.LoadBalancingAlgorithm),
		ProxyProtocol:          types.BoolValue(tools.StringToBool(upstream.ProxyProtocol)),
		Keepalive:              tools.StringOrNull(upstream.Keepalive),
		KeepaliveRequests:      tools.StringOrNull(upstream.KeepaliveRequests),
		KeepaliveTimeout:       tools.StringOrNull(upstream.KeepaliveTimeout),
		HostPort:               tools.StringOrNull(upstream.HostPort),
		XForwardedHostVerbatim: types.BoolValue(tools.StringToBool(upstream.XForwardedHostVerbatim)),
		TLSEnable:              types.BoolValue(tools.StringToBool(upstream.TLSEnable)),
		TLSClientCertificate:   fieldOptionsFirstKey(upstream.TLSClientCertificate),
		TLSNameOverride:        tools.StringOrNull(upstream.TLSNameOverride),
		TLSProtocolVersions:    fieldOptionsToSet(upstream.TLSProtocolVersions),
		TLSSessionReuse:        types.BoolValue(tools.StringToBool(upstream.TLSSessionReuse)),
		TLSTrustedCertificate:  fieldOptionsFirstKey(upstream.TLSTrustedCertificate),
		TLSVerify:              types.BoolValue(tools.StringToBool(upstream.TLSVerify)),
		TLSVerifyDepth:         tools.StringOrNull(upstream.TLSVerifyDepth),
		Store:                  types.BoolValue(tools.StringToBool(upstream.Store)),
	}
}

func (m *nginxUpstreamResourceModel) toUpstream() nginx.Upstream {
	if m == nil {
		return nginx.Upstream{}
	}

	return nginx.Upstream{
		Description:            stringValue(m.Description),
		ServerEntries:          joinStringSet(m.ServerEntries),
		LoadBalancingAlgorithm: stringValue(m.LoadBalancingAlgorithm),
		ProxyProtocol:          boolToAPIString(m.ProxyProtocol),
		Keepalive:              stringValue(m.Keepalive),
		KeepaliveRequests:      stringValue(m.KeepaliveRequests),
		KeepaliveTimeout:       stringValue(m.KeepaliveTimeout),
		HostPort:               stringValue(m.HostPort),
		XForwardedHostVerbatim: boolToAPIString(m.XForwardedHostVerbatim),
		TLSEnable:              boolToAPIString(m.TLSEnable),
		TLSClientCertificate:   stringValue(m.TLSClientCertificate),
		TLSNameOverride:        stringValue(m.TLSNameOverride),
		TLSProtocolVersions:    joinStringSetComma(m.TLSProtocolVersions),
		TLSSessionReuse:        boolToAPIString(m.TLSSessionReuse),
		TLSTrustedCertificate:  stringValue(m.TLSTrustedCertificate),
		TLSVerify:              boolToAPIString(m.TLSVerify),
		TLSVerifyDepth:         stringValue(m.TLSVerifyDepth),
		Store:                  boolToAPIString(m.Store),
	}
}

func fetchUpstreamModel(ctx context.Context, controller *nginx.Controller, id string) (nginxUpstreamResourceModel, error) {
	resp, err := controller.NginxGetUpstream(ctx, id)
	if err != nil {
		return nginxUpstreamResourceModel{}, err
	}

	return upstreamResponseToModel(id, resp), nil
}

func fieldOptionsToSet(options api.FieldOptions) types.Set {
	keys := selectedFieldOptionKeys(options)
	if len(keys) == 0 {
		return types.SetNull(types.StringType)
	}

	return tools.StringSliceToSet(keys)
}

func fieldOptionsFirstKey(options api.FieldOptions) types.String {
	keys := selectedFieldOptionKeys(options)
	if len(keys) == 0 {
		return types.StringNull()
	}
	return types.StringValue(keys[0])
}

func selectedFieldOptionKeys(options api.FieldOptions) []string {
	if len(options) == 0 {
		return []string{}
	}

	keys := make([]string, 0, len(options))
	for key, option := range options {
		if option.Selected == 1 {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)
	return keys
}
