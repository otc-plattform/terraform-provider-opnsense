package haproxy

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// apiValueToString converts a raw HAProxy API value (string, bool, float64,
// []interface{}, or map[string]interface{} of {value, selected} options) into
// the string form used by the provider. Selectable option maps collapse to a
// comma-joined list of the selected option keys.
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

// selectedKeysFromRawOptionMap returns the sorted keys of the options whose
// "selected" flag is set to "1" inside a raw HAProxy option map.
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
