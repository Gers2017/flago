package flago

import (
	"strconv"
	"strings"
)

func removeFlagPrefix(s string) string {
	for strings.HasPrefix(s, "-") {
		s = strings.TrimPrefix(s, "-")
	}
	return s
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
