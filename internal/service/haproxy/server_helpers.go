package haproxy

import (
	"context"
	"strings"

	ophaproxy "github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func serverResponseToModel(id string, resp *ophaproxy.ServerGetResponse) haproxyServerResourceModel {
	server := resp.Server

	return haproxyServerResourceModel{
		Id:                   types.StringValue(id),
		InternalId:           types.StringValue(server.ID),
		Enabled:              types.BoolValue(tools.StringToBool(server.Enabled)),
		Name:                 types.StringValue(server.Name),
		Description:          types.StringValue(server.Description),
		Type:                 types.StringValue(selectedOptionKey(server.Type)),
		Address:              types.StringValue(server.Address),
		ServiceName:          types.StringValue(server.ServiceName),
		Number:               types.StringValue(server.Number),
		LinkedResolver:       types.StringValue(selectedOptionKey(server.LinkedResolver)),
		ResolverOpts:         tools.StringSliceToSet(selectedOptionKeys(server.ResolverOpts)),
		UnixSocket:           types.StringValue(selectedOptionKey(server.UnixSocket)),
		Port:                 types.StringValue(server.Port),
		Mode:                 types.StringValue(selectedOptionKey(server.Mode)),
		MultiplexerProtocol:  types.StringValue(selectedOptionKey(server.MultiplexerProtocol)),
		ResolvePrefer:        types.StringValue(selectedOptionKey(server.ResolvePrefer)),
		SSL:                  types.BoolValue(tools.StringToBool(server.SSL)),
		SSLSNI:               types.StringValue(server.SSLSNI),
		SSLSNIExpr:           types.StringValue(server.SSLSNIExpr),
		SSLVerify:            types.BoolValue(tools.StringToBool(server.SSLVerify)),
		SSLCA:                types.StringValue(selectedOptionKey(server.SSLCA)),
		SSLCRL:               types.StringValue(selectedOptionKey(server.SSLCRL)),
		SSLClientCertificate: types.StringValue(selectedOptionKey(server.SSLClientCertificate)),
		MaxConnections:       types.StringValue(server.MaxConnections),
		Weight:               types.StringValue(server.Weight),
		CheckInterval:        types.StringValue(server.CheckInterval),
		CheckDownInterval:    types.StringValue(server.CheckDownInterval),
		Checkport:            types.StringValue(server.Checkport),
		Source:               types.StringValue(server.Source),
		Advanced:             types.StringValue(server.Advanced),
	}
}

func (m *haproxyServerResourceModel) toServer() ophaproxy.Server {
	if m == nil {
		return ophaproxy.Server{}
	}

	return ophaproxy.Server{
		Enabled:              boolToAPIString(m.Enabled),
		Name:                 stringToAPIValue(m.Name),
		Description:          stringToAPIValue(m.Description),
		Type:                 stringToAPIValue(m.Type),
		Address:              stringToAPIValue(m.Address),
		ServiceName:          stringToAPIValue(m.ServiceName),
		Number:               stringToAPIValue(m.Number),
		LinkedResolver:       stringToAPIValue(m.LinkedResolver),
		ResolverOpts:         strings.Join(tools.SetToStringSlice(m.ResolverOpts), ","),
		UnixSocket:           stringToAPIValue(m.UnixSocket),
		Port:                 stringToAPIValue(m.Port),
		Mode:                 stringToAPIValue(m.Mode),
		MultiplexerProtocol:  stringToAPIValue(m.MultiplexerProtocol),
		ResolvePrefer:        stringToAPIValue(m.ResolvePrefer),
		SSL:                  boolToAPIString(m.SSL),
		SSLSNI:               stringToAPIValue(m.SSLSNI),
		SSLSNIExpr:           stringToAPIValue(m.SSLSNIExpr),
		SSLVerify:            boolToAPIString(m.SSLVerify),
		SSLCA:                stringToAPIValue(m.SSLCA),
		SSLCRL:               stringToAPIValue(m.SSLCRL),
		SSLClientCertificate: stringToAPIValue(m.SSLClientCertificate),
		MaxConnections:       stringToAPIValue(m.MaxConnections),
		Weight:               stringToAPIValue(m.Weight),
		CheckInterval:        stringToAPIValue(m.CheckInterval),
		CheckDownInterval:    stringToAPIValue(m.CheckDownInterval),
		Checkport:            stringToAPIValue(m.Checkport),
		Source:               stringToAPIValue(m.Source),
		Advanced:             stringToAPIValue(m.Advanced),
	}
}

func fetchServerModel(ctx context.Context, controller *ophaproxy.Controller, id string) (haproxyServerResourceModel, error) {
	resp, err := controller.HAProxyGetServer(ctx, id)
	if err != nil {
		return haproxyServerResourceModel{}, err
	}

	return serverResponseToModel(id, resp), nil
}

func findServerIDByName(ctx context.Context, controller *ophaproxy.Controller, name string) (string, bool, error) {
	result, err := controller.HAProxySearchServers(ctx)
	if err != nil {
		return "", false, err
	}

	for _, row := range result.Rows {
		if row.Name == name {
			return row.UUID, true, nil
		}
	}

	return "", false, nil
}
