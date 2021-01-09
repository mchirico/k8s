package idea0

import (
	"errors"
	"strings"
)

func fun(s string, sub string) string {
	if strings.Contains(s, sub) {
		return "match"
	}
	return "no match"
}

func fun2(s string, sub string) (string, error) {

	if s == "" || sub == "" {
		return "", errors.New("Empty")
	}

	if strings.Contains(s, sub) {
		return "match", nil
	}
	return "no match", nil
}
