package haproxy

import (
	"context"
	"strings"

	ophaproxy "github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m *resolverResourceModel) toResolverObject() ophaproxy.HAProxyObject {
	return ophaproxy.HAProxyObject{
		"enabled": boolToAPIString(m.Enabled), "name": stringToAPIValue(m.Name), "description": stringToAPIValue(m.Description),
		"nameservers": strings.Join(tools.SetToStringSlice(m.Nameservers), ","), "parse_resolv_conf": boolToAPIString(m.ParseResolvConf),
		"resolve_retries": stringToAPIValue(m.ResolveRetries), "timeout_resolve": stringToAPIValue(m.TimeoutResolve), "timeout_retry": stringToAPIValue(m.TimeoutRetry),
		"accepted_payload_size": stringToAPIValue(m.AcceptedPayloadSize), "hold_valid": stringToAPIValue(m.HoldValid), "hold_obsolete": stringToAPIValue(m.HoldObsolete),
		"hold_refused": stringToAPIValue(m.HoldRefused), "hold_nx": stringToAPIValue(m.HoldNX), "hold_timeout": stringToAPIValue(m.HoldTimeout), "hold_other": stringToAPIValue(m.HoldOther),
	}
}

func resolverResponseToModel(id string, obj ophaproxy.HAProxyObject) resolverResourceModel {
	value := func(key string) types.String { return types.StringValue(apiValueToString(obj[key])) }
	var nameservers []string
	for _, name := range objectSetKeys(obj["nameservers"]) {
		if name != "" {
			nameservers = append(nameservers, name)
		}
	}
	set := types.SetNull(types.StringType)
	if len(nameservers) > 0 {
		set = tools.StringSliceToSet(nameservers)
	}
	return resolverResourceModel{
		Id: types.StringValue(id), Enabled: types.BoolValue(apiValueToString(obj["enabled"]) == "1"), Name: value("name"), Description: value("description"),
		Nameservers: set, ParseResolvConf: types.BoolValue(apiValueToString(obj["parse_resolv_conf"]) == "1"),
		ResolveRetries: value("resolve_retries"), TimeoutResolve: value("timeout_resolve"), TimeoutRetry: value("timeout_retry"), AcceptedPayloadSize: value("accepted_payload_size"),
		HoldValid: value("hold_valid"), HoldObsolete: value("hold_obsolete"), HoldRefused: value("hold_refused"), HoldNX: value("hold_nx"), HoldTimeout: value("hold_timeout"), HoldOther: value("hold_other"),
	}
}

func fetchResolverModel(ctx context.Context, c *ophaproxy.Controller, id string) (resolverResourceModel, error) {
	obj, err := c.HAProxyGetResolver(ctx, id)
	if err != nil {
		return resolverResourceModel{}, err
	}
	return resolverResponseToModel(id, obj), nil
}

func findResolverIDByName(ctx context.Context, c *ophaproxy.Controller, name string) (string, bool, error) {
	result, err := c.HAProxySearchResolvers(ctx)
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
