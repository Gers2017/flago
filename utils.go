package flago

import (
	"strconv"
	"strings"
)

func copy[T any](slice []T) []T {
	cp := make([]T, len(slice))
	for i, v := range slice {
		cp[i] = v
	}
	return cp
}

func clean(s string) string {
	s = strings.Trim(s, " ")
	return strings.TrimPrefix(s, "-")
}

func getArg(args []string, i int) (string, bool) {
	if len(args) <= i {
		return "", false
	}
	return args[i], true
}

func isHelpValue(arg string) bool {
	return arg == "help" || arg == "h"
}

func extractValues(flag string) (string, string) {
	flag_name, flag_value, ok := strings.Cut(flag, "=")
	if !ok {
		return flag, ""
	}

	return flag_name, flag_value
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
