package day87

import (
	"errors"
	"strings"
)


func example1(s string, sub string) (string, error) {


    if s == "" || sub == "" {
		return "",errors.New("Empty string")
	}

	if strings.Contains(s, sub) {
		return "match", nil
	}
	return "no match", nil
}