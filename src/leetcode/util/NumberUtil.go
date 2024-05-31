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

	println(brace("[[0,0],[4,5],[7,8],[8,9],[5,6],[3,4],[1,1]]"))
}
