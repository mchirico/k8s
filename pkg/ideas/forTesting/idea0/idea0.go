package idea0

import "strings"

// alt-cmd-t to generate test
func Ex0(s string, s2 string) string {
   if strings.Contains(s , s2) {
	   return s2
   }
   return "no find"
}