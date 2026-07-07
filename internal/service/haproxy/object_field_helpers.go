package haproxy

// objectSetKeys extracts the selected string values from a raw HAProxy API
// field that may be either a map[string]interface{} of {value, selected}
// options (multi-select dropdown) or a plain []interface{} of strings.
func objectSetKeys(value interface{}) []string {
	switch typed := value.(type) {
	case nil:
		return []string{}
	case map[string]interface{}:
		return selectedKeysFromRawOptionMap(typed)
	case []interface{}:
		out := make([]string, 0, len(typed))
		for _, item := range typed {
			switch v := item.(type) {
			case string:
				if v != "" {
					out = append(out, v)
				}
			case map[string]interface{}:
				if sel, ok := v["selected"]; ok && apiValueToString(sel) == "1" {
					if key, ok := v["value"].(string); ok && key != "" {
						out = append(out, key)
					}
				}
			}
		}
		return out
	default:
		return []string{}
	}
}
