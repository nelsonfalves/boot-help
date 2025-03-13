package util

import "strings"

func EmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
