package main

import (
	"strings"
)

func brace(s string) string {
	s = strings.Replace(s, "[", "{", len(s))
	s = strings.Replace(s, "]", "}", len(s))
	return s
}
func main() {

	println(brace("[[3,1,6],[-9,5,7]]"))
}
