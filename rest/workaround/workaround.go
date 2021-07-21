package workaround

import "strings"

func FormatContentTypes(contentTypes []string) string {
	return strings.Join(contentTypes, ",")
}
