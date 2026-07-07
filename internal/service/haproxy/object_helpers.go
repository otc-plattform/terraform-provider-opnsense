package haproxy

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/browningluke/opnsense-go/pkg/api"
	ophaproxy "github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type haproxyObjectKind struct {
	typeSuffix     string
	summaryName    string
	resourceDesc   string
	dataSourceDesc string
	search         func(context.Context, *ophaproxy.Controller) (*ophaproxy.HAProxyObjectSearchResult, error)
	get            func(context.Context, *ophaproxy.Controller, string) (ophaproxy.HAProxyObject, error)
	add            func(context.Context, *ophaproxy.Controller, ophaproxy.HAProxyObject) (*api.ActionResult, error)
	edit           func(context.Context, *ophaproxy.Controller, string, ophaproxy.HAProxyObject) (*api.ActionResult, error)
	delete         func(context.Context, *ophaproxy.Controller, string) (*api.ActionResult, error)
}

func objectModelFromAPI(ctx context.Context, id string, object ophaproxy.HAProxyObject) haproxyObjectResourceModel {
	config := map[string]string{}
	for key, value := range object {
		if strings.HasPrefix(key, "%") || key == "uuid" || key == "id" {
			continue
		}
		config[key] = apiValueToString(value)
	}

	configValue, _ := types.MapValueFrom(ctx, types.StringType, config)
	return haproxyObjectResourceModel{
		Id:     types.StringValue(id),
		Config: configValue,
	}
}

func objectFromModel(ctx context.Context, model *haproxyObjectResourceModel) (ophaproxy.HAProxyObject, error) {
	config := map[string]string{}
	diags := model.Config.ElementsAs(ctx, &config, false)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to decode config map")
	}

	object := ophaproxy.HAProxyObject{}
	for key, value := range config {
		object[key] = value
	}
	return object, nil
}

func apiValueToString(value interface{}) string {
	switch typed := value.(type) {
	case nil:
		return ""
	case string:
		return typed
	case bool:
		if typed {
			return "1"
		}
		return "0"
	case float64:
		if typed == float64(int64(typed)) {
			return strconv.FormatInt(int64(typed), 10)
		}
		return strconv.FormatFloat(typed, 'f', -1, 64)
	case []interface{}:
		parts := make([]string, 0, len(typed))
		for _, item := range typed {
			parts = append(parts, apiValueToString(item))
		}
		return strings.Join(parts, ",")
	case map[string]interface{}:
		selected := selectedKeysFromRawOptionMap(typed)
		if len(selected) > 0 {
			return strings.Join(selected, ",")
		}
		return ""
	default:
		return fmt.Sprint(typed)
	}
}

func selectedKeysFromRawOptionMap(options map[string]interface{}) []string {
	keys := make([]string, 0, len(options))
	for key, value := range options {
		option, ok := value.(map[string]interface{})
		if !ok {
			continue
		}
		selected, ok := option["selected"]
		if !ok {
			continue
		}
		if apiValueToString(selected) == "1" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	return keys
}

func emptyConfigMap() types.Map {
	value, _ := types.MapValue(types.StringType, map[string]attr.Value{})
	return value
}

func findObjectIDByName(ctx context.Context, kind haproxyObjectKind, controller *ophaproxy.Controller, name string) (string, bool, error) {
	result, err := kind.search(ctx, controller)
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
