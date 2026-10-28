package haproxy

import (
	"context"

	ophaproxy "github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m *healthMonitorResourceModel) toHealthMonitorObject() ophaproxy.HAProxyObject {
	return ophaproxy.HAProxyObject{
		"name": stringToAPIValue(m.Name), "description": stringToAPIValue(m.Description), "type": stringToAPIValue(m.Type), "interval": stringToAPIValue(m.Interval),
		"ssl": stringToAPIValue(m.SSL), "sslSNI": stringToAPIValue(m.SSLSNI), "checkport": stringToAPIValue(m.CheckPort),
		"http_method": stringToAPIValue(m.HTTPMethod), "http_uri": stringToAPIValue(m.HTTPURI), "http_version": stringToAPIValue(m.HTTPVersion), "http_host": stringToAPIValue(m.HTTPHost),
		"http_expressionEnabled": boolToAPIString(m.HTTPExpressionEnabled), "http_expression": stringToAPIValue(m.HTTPExpression), "http_negate": boolToAPIString(m.HTTPNegate), "http_value": stringToAPIValue(m.HTTPValue),
		"tcp_enabled": boolToAPIString(m.TCPEnabled), "tcp_sendValue": stringToAPIValue(m.TCPSendValue), "tcp_matchType": stringToAPIValue(m.TCPMatchType), "tcp_negate": boolToAPIString(m.TCPNegate), "tcp_matchValue": stringToAPIValue(m.TCPMatchValue),
		"agent_port": stringToAPIValue(m.AgentPort), "mysql_user": stringToAPIValue(m.MySQLUser), "mysql_post41": boolToAPIString(m.MySQLPost41), "pgsql_user": stringToAPIValue(m.PostgresUser),
		"smtp_domain": stringToAPIValue(m.SMTPDomain), "esmtp_domain": stringToAPIValue(m.ESMTPDomain),
	}
}

func healthMonitorResponseToModel(id string, obj ophaproxy.HAProxyObject) healthMonitorResourceModel {
	value := func(key string) types.String { return types.StringValue(apiValueToString(obj[key])) }
	flag := func(key string) types.Bool { return types.BoolValue(tools.StringToBool(apiValueToString(obj[key]))) }
	return healthMonitorResourceModel{
		Id: types.StringValue(id), Name: value("name"), Description: value("description"), Type: value("type"), Interval: value("interval"),
		SSL: value("ssl"), SSLSNI: value("sslSNI"), CheckPort: value("checkport"), HTTPMethod: value("http_method"), HTTPURI: value("http_uri"), HTTPVersion: value("http_version"), HTTPHost: value("http_host"),
		HTTPExpressionEnabled: flag("http_expressionEnabled"), HTTPExpression: value("http_expression"), HTTPNegate: flag("http_negate"), HTTPValue: value("http_value"),
		TCPEnabled: flag("tcp_enabled"), TCPSendValue: value("tcp_sendValue"), TCPMatchType: value("tcp_matchType"), TCPNegate: flag("tcp_negate"), TCPMatchValue: value("tcp_matchValue"),
		AgentPort: value("agent_port"), MySQLUser: value("mysql_user"), MySQLPost41: flag("mysql_post41"), PostgresUser: value("pgsql_user"), SMTPDomain: value("smtp_domain"), ESMTPDomain: value("esmtp_domain"),
	}
}

func fetchHealthMonitorModel(ctx context.Context, c *ophaproxy.Controller, id string) (healthMonitorResourceModel, error) {
	obj, err := c.HAProxyGetHealthMonitor(ctx, id)
	if err != nil {
		return healthMonitorResourceModel{}, err
	}
	return healthMonitorResponseToModel(id, obj), nil
}

func findHealthMonitorIDByName(ctx context.Context, c *ophaproxy.Controller, name string) (string, bool, error) {
	result, err := c.HAProxySearchHealthMonitors(ctx)
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
