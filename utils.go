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
	return strings.ReplaceAll(s, "-", "")
}

func getArg(args []string, i int) (string, bool) {
	if len(args) <= i {
		return "", false
	}
	return args[i], true
}

func getNextValue(args_copy []string, i int) string {
	f_value, _ := getArg(args_copy, i+1)
	return f_value
}

func extractValues(flag string) (string, string) {
	f_parts := strings.Split(flag, "=")

	if len(f_parts) < 2 {
		return flag, ""
	}

	f_name, f_value := f_parts[0], strings.Join(f_parts[1:], "")

	if f_value == "" {
		return f_name, f_value
	}

	return f_name, f_value
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
