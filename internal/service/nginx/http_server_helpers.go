package nginx

import (
	"context"
	"sort"
	"strings"

	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func httpServerResponseToModel(id string, resp *nginx.HTTPServerGetResponse) nginxHTTPServerResourceModel {
	server := resp.HTTPServer

	return nginxHTTPServerResourceModel{
		Id:                             types.StringValue(id),
		ServerName:                     stringValueFromOptionMap(server.ServerName),
		ListenHTTPAddress:              stringValueFromOptionMap(server.ListenHTTPAddress),
		ListenHTTPSAddress:             stringValueFromOptionMap(server.ListenHTTPSAddress),
		DefaultServer:                  types.BoolValue(tools.StringToBool(server.DefaultServer)),
		TLSRejectHandshake:             types.BoolValue(tools.StringToBool(server.TLSRejectHandshake)),
		SyslogTargets:                  stringSliceToSet(server.SyslogTargets),
		ProxyProtocol:                  types.BoolValue(tools.StringToBool(server.ProxyProtocol)),
		TrustedProxies:                 stringSliceToSet(server.TrustedProxies),
		TrustedProxiesAlias:            stringValueFromOptionMap(server.TrustedProxiesAlias),
		RealIPSource:                   stringValueFromOptionMap(server.RealIPSource),
		Locations:                      optionMapToSet(server.Locations),
		Rewrites:                       stringSliceToSet(server.Rewrites),
		Root:                           tools.StringOrNull(server.Root),
		MaxBodySize:                    tools.StringOrNull(server.MaxBodySize),
		BodyBufferSize:                 tools.StringOrNull(server.BodyBufferSize),
		Certificate:                    stringValueFromOptionMap(server.Certificate),
		CA:                             stringValueFromOptionMap(server.CA),
		VerifyClient:                   stringValueFromOptionMap(server.VerifyClient),
		ZeroRTT:                        types.BoolValue(tools.StringToBool(server.ZeroRTT)),
		AccessLogFormat:                stringValueFromOptionMap(server.AccessLogFormat),
		ErrorLogLevel:                  stringValueFromOptionMap(server.ErrorLogLevel),
		LogHandshakes:                  types.BoolValue(tools.StringToBool(server.LogHandshakes)),
		EnableACMESupport:              types.BoolValue(tools.StringToBool(server.EnableACMESupport)),
		Charset:                        stringValueFromOptionMap(server.Charset),
		HTTPSOnly:                      types.BoolValue(tools.StringToBool(server.HTTPSOnly)),
		TLSProtocols:                   optionMapToSet(server.TLSProtocols),
		TLSCiphers:                     tools.StringOrNull(server.TLSCiphers),
		TLSECDHCurve:                   tools.StringOrNull(server.TLSECDHCurve),
		TLSPreferServerCiphers:         types.BoolValue(tools.StringToBool(server.TLSPreferServerCiphers)),
		Resolver:                       stringValueFromOptionMap(server.Resolver),
		OCSPStapling:                   types.BoolValue(tools.StringToBool(server.OCSPStapling)),
		OCSPVerify:                     types.BoolValue(tools.StringToBool(server.OCSPVerify)),
		BlockNonpublicData:             types.BoolValue(tools.StringToBool(server.BlockNonpublicData)),
		DisableGzip:                    types.BoolValue(tools.StringToBool(server.DisableGzip)),
		DisableBotProtection:           types.BoolValue(tools.StringToBool(server.DisableBotProtection)),
		IPACL:                          stringValueFromOptionMap(server.IPACL),
		AdvancedACLServer:              stringValueFromOptionMap(server.AdvancedACLServer),
		Satisfy:                        stringValueFromOptionMap(server.Satisfy),
		NaxsiWhitelistSrcIP:            stringSliceToSet(server.NaxsiWhitelistSrcIP),
		NaxsiExtensiveLog:              types.BoolValue(tools.StringToBool(server.NaxsiExtensiveLog)),
		Sendfile:                       types.BoolValue(tools.StringToBool(server.Sendfile)),
		ClientHeaderBufferSize:         tools.StringOrNull(server.ClientHeaderBufferSize),
		LargeClientHeaderBuffersNumber: tools.StringOrNull(server.LargeClientHeaderBuffersNumber),
		LargeClientHeaderBuffersSize:   tools.StringOrNull(server.LargeClientHeaderBuffersSize),
		SecurityHeader:                 stringValueFromOptionMap(server.SecurityHeader),
		LimitRequestConnections:        stringSliceToSet(server.LimitRequestConnections),
		ErrorPages:                     stringSliceToSet(server.ErrorPages),
	}
}

func (m *nginxHTTPServerResourceModel) toHTTPServer() nginx.HTTPServer {
	if m == nil {
		return nginx.HTTPServer{}
	}

	return nginx.HTTPServer{
		ListenHTTPAddress:              stringValue(m.ListenHTTPAddress),
		ListenHTTPSAddress:             stringValue(m.ListenHTTPSAddress),
		DefaultServer:                  boolToAPIString(m.DefaultServer),
		TLSRejectHandshake:             boolToAPIString(m.TLSRejectHandshake),
		SyslogTargets:                  joinStringSet(m.SyslogTargets),
		ProxyProtocol:                  boolToAPIString(m.ProxyProtocol),
		TrustedProxies:                 joinStringSet(m.TrustedProxies),
		TrustedProxiesAlias:            stringValue(m.TrustedProxiesAlias),
		RealIPSource:                   stringValue(m.RealIPSource),
		ServerName:                     stringValue(m.ServerName),
		Locations:                      joinStringSet(m.Locations),
		Rewrites:                       joinStringSet(m.Rewrites),
		Root:                           stringValue(m.Root),
		MaxBodySize:                    stringValue(m.MaxBodySize),
		BodyBufferSize:                 stringValue(m.BodyBufferSize),
		Certificate:                    stringValue(m.Certificate),
		CA:                             stringValue(m.CA),
		VerifyClient:                   stringValue(m.VerifyClient),
		ZeroRTT:                        boolToAPIString(m.ZeroRTT),
		AccessLogFormat:                stringValue(m.AccessLogFormat),
		ErrorLogLevel:                  stringValue(m.ErrorLogLevel),
		LogHandshakes:                  boolToAPIString(m.LogHandshakes),
		EnableACMESupport:              boolToAPIString(m.EnableACMESupport),
		Charset:                        stringValue(m.Charset),
		HTTPSOnly:                      boolToAPIString(m.HTTPSOnly),
		TLSProtocols:                   joinStringSetComma(m.TLSProtocols),
		TLSCiphers:                     stringValue(m.TLSCiphers),
		TLSECDHCurve:                   stringValue(m.TLSECDHCurve),
		TLSPreferServerCiphers:         boolToAPIString(m.TLSPreferServerCiphers),
		Resolver:                       stringValue(m.Resolver),
		OCSPStapling:                   boolToAPIString(m.OCSPStapling),
		OCSPVerify:                     boolToAPIString(m.OCSPVerify),
		BlockNonpublicData:             boolToAPIString(m.BlockNonpublicData),
		DisableGzip:                    boolToAPIString(m.DisableGzip),
		DisableBotProtection:           boolToAPIString(m.DisableBotProtection),
		IPACL:                          stringValue(m.IPACL),
		AdvancedACLServer:              stringValue(m.AdvancedACLServer),
		Satisfy:                        stringValue(m.Satisfy),
		NaxsiWhitelistSrcIP:            joinStringSet(m.NaxsiWhitelistSrcIP),
		NaxsiExtensiveLog:              boolToAPIString(m.NaxsiExtensiveLog),
		Sendfile:                       boolToAPIString(m.Sendfile),
		ClientHeaderBufferSize:         stringValue(m.ClientHeaderBufferSize),
		LargeClientHeaderBuffersNumber: stringValue(m.LargeClientHeaderBuffersNumber),
		LargeClientHeaderBuffersSize:   stringValue(m.LargeClientHeaderBuffersSize),
		SecurityHeader:                 stringValue(m.SecurityHeader),
		LimitRequestConnections:        joinStringSet(m.LimitRequestConnections),
		ErrorPages:                     joinStringSet(m.ErrorPages),
	}
}

func fetchHTTPServerModel(ctx context.Context, controller *nginx.Controller, id string) (nginxHTTPServerResourceModel, error) {
	resp, err := controller.NginxGetHTTPServer(ctx, id)
	if err != nil {
		return nginxHTTPServerResourceModel{}, err
	}

	return httpServerResponseToModel(id, resp), nil
}

func joinStringSetComma(set types.Set) string {
	if set.IsNull() || set.IsUnknown() {
		return ""
	}

	values := tools.SetToStringSlice(set)
	if len(values) == 0 {
		return ""
	}

	sort.Strings(values)
	return strings.Join(values, ",")
}
