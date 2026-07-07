package haproxy

import (
	"context"
	"strings"

	ophaproxy "github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m *actionResourceModel) toActionObject() ophaproxy.HAProxyObject {
	if m == nil {
		return ophaproxy.HAProxyObject{}
	}
	obj := ophaproxy.HAProxyObject{}
	obj["enabled"] = boolToAPIString(m.Enabled)
	obj["name"] = stringToAPIValue(m.Name)
	obj["description"] = stringToAPIValue(m.Description)
	obj["testType"] = stringToAPIValue(m.TestType)
	obj["linkedAcls"] = strings.Join(tools.SetToStringSlice(m.LinkedAcls), ",")
	obj["operator"] = stringToAPIValue(m.Operator)
	obj["type"] = stringToAPIValue(m.Type_)
	obj["use_backend"] = stringToAPIValue(m.UseBackend)
	obj["use_server"] = stringToAPIValue(m.UseServer)
	obj["fcgi_pass_header"] = stringToAPIValue(m.FcgiPassHeader)
	obj["fcgi_set_param"] = stringToAPIValue(m.FcgiSetParam)
	obj["monitor_fail_uri"] = stringToAPIValue(m.MonitorFailUri)
	obj["custom"] = stringToAPIValue(m.Custom)
	obj["http_after_response_action"] = stringToAPIValue(m.HttpAfterResponseAction)
	obj["http_after_response_option"] = stringToAPIValue(m.HttpAfterResponseOption)
	obj["http_request_action"] = stringToAPIValue(m.HttpRequestAction)
	obj["http_request_option"] = stringToAPIValue(m.HttpRequestOption)
	obj["http_response_action"] = stringToAPIValue(m.HttpResponseAction)
	obj["http_response_option"] = stringToAPIValue(m.HttpResponseOption)
	obj["tcp_request_action"] = stringToAPIValue(m.TcpRequestAction)
	obj["tcp_request_option"] = stringToAPIValue(m.TcpRequestOption)
	obj["tcp_response_action"] = stringToAPIValue(m.TcpResponseAction)
	obj["tcp_response_option"] = stringToAPIValue(m.TcpResponseOption)
	obj["http_request_auth"] = stringToAPIValue(m.HttpRequestAuth)
	obj["http_request_deny_status"] = stringToAPIValue(m.HttpRequestDenyStatus)
	obj["http_request_redirect"] = stringToAPIValue(m.HttpRequestRedirect)
	obj["http_request_lua"] = stringToAPIValue(m.HttpRequestLua)
	obj["http_request_use_service"] = stringToAPIValue(m.HttpRequestUseService)
	obj["http_request_add_header_name"] = stringToAPIValue(m.HttpRequestAddHeaderName)
	obj["http_request_add_header_content"] = stringToAPIValue(m.HttpRequestAddHeaderContent)
	obj["http_request_set_header_name"] = stringToAPIValue(m.HttpRequestSetHeaderName)
	obj["http_request_set_header_content"] = stringToAPIValue(m.HttpRequestSetHeaderContent)
	obj["http_request_del_header_name"] = stringToAPIValue(m.HttpRequestDelHeaderName)
	obj["http_request_replace_header_name"] = stringToAPIValue(m.HttpRequestReplaceHeaderName)
	obj["http_request_replace_header_regex"] = stringToAPIValue(m.HttpRequestReplaceHeaderRegex)
	obj["http_request_replace_value_name"] = stringToAPIValue(m.HttpRequestReplaceValueName)
	obj["http_request_replace_value_regex"] = stringToAPIValue(m.HttpRequestReplaceValueRegex)
	obj["http_request_set_path"] = stringToAPIValue(m.HttpRequestSetPath)
	obj["http_request_set_var_scope"] = stringToAPIValue(m.HttpRequestSetVarScope)
	obj["http_request_set_var_name"] = stringToAPIValue(m.HttpRequestSetVarName)
	obj["http_request_set_var_expr"] = stringToAPIValue(m.HttpRequestSetVarExpr)
	obj["http_response_lua"] = stringToAPIValue(m.HttpResponseLua)
	obj["http_response_add_header_name"] = stringToAPIValue(m.HttpResponseAddHeaderName)
	obj["http_response_add_header_content"] = stringToAPIValue(m.HttpResponseAddHeaderContent)
	obj["http_response_set_header_name"] = stringToAPIValue(m.HttpResponseSetHeaderName)
	obj["http_response_set_header_content"] = stringToAPIValue(m.HttpResponseSetHeaderContent)
	obj["http_response_del_header_name"] = stringToAPIValue(m.HttpResponseDelHeaderName)
	obj["http_response_replace_header_name"] = stringToAPIValue(m.HttpResponseReplaceHeaderName)
	obj["http_response_replace_header_regex"] = stringToAPIValue(m.HttpResponseReplaceHeaderRegex)
	obj["http_response_replace_value_name"] = stringToAPIValue(m.HttpResponseReplaceValueName)
	obj["http_response_replace_value_regex"] = stringToAPIValue(m.HttpResponseReplaceValueRegex)
	obj["http_response_set_status_code"] = stringToAPIValue(m.HttpResponseSetStatusCode)
	obj["http_response_set_status_reason"] = stringToAPIValue(m.HttpResponseSetStatusReason)
	obj["http_response_set_var_scope"] = stringToAPIValue(m.HttpResponseSetVarScope)
	obj["http_response_set_var_name"] = stringToAPIValue(m.HttpResponseSetVarName)
	obj["http_response_set_var_expr"] = stringToAPIValue(m.HttpResponseSetVarExpr)
	obj["tcp_request_content_lua"] = stringToAPIValue(m.TcpRequestContentLua)
	obj["tcp_request_content_use_service"] = stringToAPIValue(m.TcpRequestContentUseService)
	obj["tcp_request_inspect_delay"] = stringToAPIValue(m.TcpRequestInspectDelay)
	obj["tcp_response_content_lua"] = stringToAPIValue(m.TcpResponseContentLua)
	obj["tcp_response_inspect_delay"] = stringToAPIValue(m.TcpResponseInspectDelay)
	obj["map_data_use_backend_file"] = stringToAPIValue(m.MapDataUseBackendFile)
	obj["map_data_use_backend_default"] = stringToAPIValue(m.MapDataUseBackendDefault)
	obj["map_data_use_backend_input"] = stringToAPIValue(m.MapDataUseBackendInput)
	obj["map_use_backend_file"] = stringToAPIValue(m.MapUseBackendFile)
	obj["map_use_backend_default"] = stringToAPIValue(m.MapUseBackendDefault)
	obj["compression_algo_res"] = strings.Join(tools.SetToStringSlice(m.CompressionAlgoRes), ",")
	obj["compression_algo_req"] = strings.Join(tools.SetToStringSlice(m.CompressionAlgoReq), ",")
	obj["compression_mime_res"] = strings.Join(tools.SetToStringSlice(m.CompressionMimeRes), ",")
	obj["compression_mime_req"] = strings.Join(tools.SetToStringSlice(m.CompressionMimeReq), ",")
	obj["compression_offloading"] = boolToAPIString(m.CompressionOffloading)
	obj["compression_minsize_res"] = stringToAPIValue(m.CompressionMinsizeRes)
	obj["compression_minsize_req"] = stringToAPIValue(m.CompressionMinsizeReq)
	obj["compression_direction"] = stringToAPIValue(m.CompressionDirection)
	obj["gpc_number"] = stringToAPIValue(m.GpcNumber)
	obj["gpt_number"] = stringToAPIValue(m.GptNumber)
	obj["sc_number"] = stringToAPIValue(m.ScNumber)
	obj["mapfile"] = stringToAPIValue(m.Mapfile)
	obj["map_default"] = stringToAPIValue(m.MapDefault)
	obj["sample_fetch"] = stringToAPIValue(m.SampleFetch)
	return obj
}

func actionResponseToModel(ctx context.Context, id string, obj ophaproxy.HAProxyObject) actionResourceModel {
	m := actionResourceModel{}
	m.Id = types.StringValue(id)
	m.Enabled = types.BoolValue(tools.StringToBool(apiValueToString(obj["enabled"])))
	m.Name = types.StringValue(apiValueToString(obj["name"]))
	m.Description = types.StringValue(apiValueToString(obj["description"]))
	m.TestType = types.StringValue(apiValueToString(obj["testType"]))
	m.LinkedAcls = tools.StringSliceToSet(objectSetKeys(obj["linkedAcls"]))
	m.Operator = types.StringValue(apiValueToString(obj["operator"]))
	m.Type_ = types.StringValue(apiValueToString(obj["type"]))
	m.UseBackend = types.StringValue(apiValueToString(obj["use_backend"]))
	m.UseServer = types.StringValue(apiValueToString(obj["use_server"]))
	m.FcgiPassHeader = types.StringValue(apiValueToString(obj["fcgi_pass_header"]))
	m.FcgiSetParam = types.StringValue(apiValueToString(obj["fcgi_set_param"]))
	m.MonitorFailUri = types.StringValue(apiValueToString(obj["monitor_fail_uri"]))
	m.Custom = types.StringValue(apiValueToString(obj["custom"]))
	m.HttpAfterResponseAction = types.StringValue(apiValueToString(obj["http_after_response_action"]))
	m.HttpAfterResponseOption = types.StringValue(apiValueToString(obj["http_after_response_option"]))
	m.HttpRequestAction = types.StringValue(apiValueToString(obj["http_request_action"]))
	m.HttpRequestOption = types.StringValue(apiValueToString(obj["http_request_option"]))
	m.HttpResponseAction = types.StringValue(apiValueToString(obj["http_response_action"]))
	m.HttpResponseOption = types.StringValue(apiValueToString(obj["http_response_option"]))
	m.TcpRequestAction = types.StringValue(apiValueToString(obj["tcp_request_action"]))
	m.TcpRequestOption = types.StringValue(apiValueToString(obj["tcp_request_option"]))
	m.TcpResponseAction = types.StringValue(apiValueToString(obj["tcp_response_action"]))
	m.TcpResponseOption = types.StringValue(apiValueToString(obj["tcp_response_option"]))
	m.HttpRequestAuth = types.StringValue(apiValueToString(obj["http_request_auth"]))
	m.HttpRequestDenyStatus = types.StringValue(apiValueToString(obj["http_request_deny_status"]))
	m.HttpRequestRedirect = types.StringValue(apiValueToString(obj["http_request_redirect"]))
	m.HttpRequestLua = types.StringValue(apiValueToString(obj["http_request_lua"]))
	m.HttpRequestUseService = types.StringValue(apiValueToString(obj["http_request_use_service"]))
	m.HttpRequestAddHeaderName = types.StringValue(apiValueToString(obj["http_request_add_header_name"]))
	m.HttpRequestAddHeaderContent = types.StringValue(apiValueToString(obj["http_request_add_header_content"]))
	m.HttpRequestSetHeaderName = types.StringValue(apiValueToString(obj["http_request_set_header_name"]))
	m.HttpRequestSetHeaderContent = types.StringValue(apiValueToString(obj["http_request_set_header_content"]))
	m.HttpRequestDelHeaderName = types.StringValue(apiValueToString(obj["http_request_del_header_name"]))
	m.HttpRequestReplaceHeaderName = types.StringValue(apiValueToString(obj["http_request_replace_header_name"]))
	m.HttpRequestReplaceHeaderRegex = types.StringValue(apiValueToString(obj["http_request_replace_header_regex"]))
	m.HttpRequestReplaceValueName = types.StringValue(apiValueToString(obj["http_request_replace_value_name"]))
	m.HttpRequestReplaceValueRegex = types.StringValue(apiValueToString(obj["http_request_replace_value_regex"]))
	m.HttpRequestSetPath = types.StringValue(apiValueToString(obj["http_request_set_path"]))
	m.HttpRequestSetVarScope = types.StringValue(apiValueToString(obj["http_request_set_var_scope"]))
	m.HttpRequestSetVarName = types.StringValue(apiValueToString(obj["http_request_set_var_name"]))
	m.HttpRequestSetVarExpr = types.StringValue(apiValueToString(obj["http_request_set_var_expr"]))
	m.HttpResponseLua = types.StringValue(apiValueToString(obj["http_response_lua"]))
	m.HttpResponseAddHeaderName = types.StringValue(apiValueToString(obj["http_response_add_header_name"]))
	m.HttpResponseAddHeaderContent = types.StringValue(apiValueToString(obj["http_response_add_header_content"]))
	m.HttpResponseSetHeaderName = types.StringValue(apiValueToString(obj["http_response_set_header_name"]))
	m.HttpResponseSetHeaderContent = types.StringValue(apiValueToString(obj["http_response_set_header_content"]))
	m.HttpResponseDelHeaderName = types.StringValue(apiValueToString(obj["http_response_del_header_name"]))
	m.HttpResponseReplaceHeaderName = types.StringValue(apiValueToString(obj["http_response_replace_header_name"]))
	m.HttpResponseReplaceHeaderRegex = types.StringValue(apiValueToString(obj["http_response_replace_header_regex"]))
	m.HttpResponseReplaceValueName = types.StringValue(apiValueToString(obj["http_response_replace_value_name"]))
	m.HttpResponseReplaceValueRegex = types.StringValue(apiValueToString(obj["http_response_replace_value_regex"]))
	m.HttpResponseSetStatusCode = types.StringValue(apiValueToString(obj["http_response_set_status_code"]))
	m.HttpResponseSetStatusReason = types.StringValue(apiValueToString(obj["http_response_set_status_reason"]))
	m.HttpResponseSetVarScope = types.StringValue(apiValueToString(obj["http_response_set_var_scope"]))
	m.HttpResponseSetVarName = types.StringValue(apiValueToString(obj["http_response_set_var_name"]))
	m.HttpResponseSetVarExpr = types.StringValue(apiValueToString(obj["http_response_set_var_expr"]))
	m.TcpRequestContentLua = types.StringValue(apiValueToString(obj["tcp_request_content_lua"]))
	m.TcpRequestContentUseService = types.StringValue(apiValueToString(obj["tcp_request_content_use_service"]))
	m.TcpRequestInspectDelay = types.StringValue(apiValueToString(obj["tcp_request_inspect_delay"]))
	m.TcpResponseContentLua = types.StringValue(apiValueToString(obj["tcp_response_content_lua"]))
	m.TcpResponseInspectDelay = types.StringValue(apiValueToString(obj["tcp_response_inspect_delay"]))
	m.MapDataUseBackendFile = types.StringValue(apiValueToString(obj["map_data_use_backend_file"]))
	m.MapDataUseBackendDefault = types.StringValue(apiValueToString(obj["map_data_use_backend_default"]))
	m.MapDataUseBackendInput = types.StringValue(apiValueToString(obj["map_data_use_backend_input"]))
	m.MapUseBackendFile = types.StringValue(apiValueToString(obj["map_use_backend_file"]))
	m.MapUseBackendDefault = types.StringValue(apiValueToString(obj["map_use_backend_default"]))
	m.CompressionAlgoRes = tools.StringSliceToSet(objectSetKeys(obj["compression_algo_res"]))
	m.CompressionAlgoReq = tools.StringSliceToSet(objectSetKeys(obj["compression_algo_req"]))
	m.CompressionMimeRes = tools.StringSliceToSet(objectSetKeys(obj["compression_mime_res"]))
	m.CompressionMimeReq = tools.StringSliceToSet(objectSetKeys(obj["compression_mime_req"]))
	m.CompressionOffloading = types.BoolValue(tools.StringToBool(apiValueToString(obj["compression_offloading"])))
	m.CompressionMinsizeRes = types.StringValue(apiValueToString(obj["compression_minsize_res"]))
	m.CompressionMinsizeReq = types.StringValue(apiValueToString(obj["compression_minsize_req"]))
	m.CompressionDirection = types.StringValue(apiValueToString(obj["compression_direction"]))
	m.GpcNumber = types.StringValue(apiValueToString(obj["gpc_number"]))
	m.GptNumber = types.StringValue(apiValueToString(obj["gpt_number"]))
	m.ScNumber = types.StringValue(apiValueToString(obj["sc_number"]))
	m.Mapfile = types.StringValue(apiValueToString(obj["mapfile"]))
	m.MapDefault = types.StringValue(apiValueToString(obj["map_default"]))
	m.SampleFetch = types.StringValue(apiValueToString(obj["sample_fetch"]))
	return m
}

func fetchActionModel(ctx context.Context, c *ophaproxy.Controller, id string) (actionResourceModel, error) {
	obj, err := c.HAProxyGetAction(ctx, id)
	if err != nil {
		return actionResourceModel{}, err
	}
	return actionResponseToModel(ctx, id, obj), nil
}

func findActionIDByName(ctx context.Context, c *ophaproxy.Controller, name string) (string, bool, error) {
	result, err := c.HAProxySearchActions(ctx)
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
