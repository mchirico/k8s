package idea0

import "strings"


func Fun(s string, sub string) string {
	if strings.Contains(s, sub) {
		return "match"
	}
	return "no match"
}