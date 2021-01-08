package idea0

import (
	"errors"
	"strings"
)

func Fun(s string, sub string) string {
	if strings.Contains(s, sub) {
		return "match"
	}
	return "no match"
}

func Fun2(s string, sub string) (string, error) {

	if s == "" || sub == "" {
		return "", errors.New("Empty")
	}

	if strings.Contains(s, sub) {
		return "match", nil
	}
	return "no match", nil
}
