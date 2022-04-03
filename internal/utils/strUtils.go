package utils

import (
	"strings"
)

// Fix common string issues
func FixString(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(strings.ReplaceAll(s, "-", " - ")), "\n", ""), "  ", " ")
}

func FixLink(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(s), "\n", ""), "  ", " ")
}
