package main

import (
	"strings"
)

func Copy[T any](slice []T) []T {
	cp := make([]T, len(slice))
	for i, v := range slice {
		cp[i] = v
	}
	return cp
}

func Clean(s string) string {
	s = strings.Trim(s, " ")
	return strings.ReplaceAll(s, "-", "")
}

func GetArg(args []string, i int) (string, bool) {
	if len(args) <= i {
		return "", false
	}
	return args[i], true
}

func IsFlag(s string) bool {
	return strings.HasPrefix(s, "-") && len(s) > 1
}

func IsShortFlag(s string) bool {
	return IsFlag(s) && strings.Count(s, "-") == 1
}

func ExtractValues(s string) (string, string) {
	parts := strings.Split(s, "=")

	if len(parts) <= 1 {
		return s, ""
	}

	return parts[0], parts[1]
}
