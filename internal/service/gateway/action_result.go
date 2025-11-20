package gateway

import (
	"fmt"
	"sort"
	"strings"

	"github.com/browningluke/opnsense-go/pkg/api"
)

func formatActionResultFailure(operation string, res *api.ActionResult) string {
	if res == nil {
		return fmt.Sprintf("Unable to %s: action failed without a response payload.", operation)
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Unable to %s. Result: %s.", operation, res.Result))

	if len(res.Validations) > 0 {
		builder.WriteString("\nValidation errors:")

		keys := make([]string, 0, len(res.Validations))
		for k := range res.Validations {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			msg := res.Validations[key]
			if msg == "" {
				msg = "unspecified error"
			}
			builder.WriteString(fmt.Sprintf("\n  - %s: %s", key, msg))
		}
	}

	return builder.String()
}
