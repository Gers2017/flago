package flago

import (
	"strconv"
	"strings"
)

func copySlice[T any](slice []T) []T {
	cp := make([]T, len(slice))
	for i, v := range slice {
		cp[i] = v
	}
	return cp
}

func removeFlagPrefix(s string) string {
	for strings.HasPrefix(s, "-") {
		s = strings.TrimPrefix(s, "-")
	}
	return s
}

func getArg(args []string, i int) string {
	if len(args) <= i {
		return ""
	}
	return args[i]
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
