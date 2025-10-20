package nginx

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func upstreamServerResponseToModel(id string, resp *nginx.UpstreamServerGetResponse) nginxUpstreamServerResourceModel {
	server := resp.UpstreamServer

	return nginxUpstreamServerResourceModel{
		Id:          types.StringValue(id),
		Description: tools.StringOrNull(server.Description),
		Server:      tools.StringOrNull(server.Server),
		Port:        tools.StringOrNull(server.Port),
		Priority:    tools.StringOrNull(server.Priority),
		MaxConns:    tools.StringOrNull(server.MaxConns),
		MaxFails:    tools.StringOrNull(server.MaxFails),
		FailTimeout: tools.StringOrNull(server.FailTimeout),
		NoUse:       fieldOptionsFirstOption(server.NoUse),
	}
}

func (m *nginxUpstreamServerResourceModel) toUpstreamServer() nginx.UpstreamServer {
	if m == nil {
		return nginx.UpstreamServer{}
	}

	return nginx.UpstreamServer{
		Description: stringValue(m.Description),
		Server:      stringValue(m.Server),
		Port:        stringValue(m.Port),
		Priority:    stringValue(m.Priority),
		MaxConns:    stringValue(m.MaxConns),
		MaxFails:    stringValue(m.MaxFails),
		FailTimeout: stringValue(m.FailTimeout),
		NoUse:       stringValue(m.NoUse),
	}
}

func fetchUpstreamServerModel(ctx context.Context, controller *nginx.Controller, id string) (nginxUpstreamServerResourceModel, error) {
	resp, err := controller.NginxGetUpstreamServer(ctx, id)
	if err != nil {
		return nginxUpstreamServerResourceModel{}, err
	}

	return upstreamServerResponseToModel(id, resp), nil
}

func fieldOptionsFirstOption(options api.FieldOptions) types.String {
	for key, option := range options {
		if option.Selected == 1 {
			if key == "" {
				return types.StringNull()
			}
			return types.StringValue(key)
		}
	}

	return types.StringNull()
}
