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

	println(brace("[[1,7],[2,7,6],[1,2],[2,7,5],[2,7,6]]"))
}
