package util

import "strings"

func ToCacheKey(str ...string) string {
	return strings.Join(str, ":")
}
