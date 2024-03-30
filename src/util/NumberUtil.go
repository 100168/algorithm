package main

import "strings"

func MAX(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
func Min(a int, b int) int {
	if a >= b {
		return b
	}
	return a
}

func brace(s string) string {
	s = strings.Replace(s, "[", "{", len(s))
	s = strings.Replace(s, "]", "}", len(s))
	return s
}
func main() {
	println(brace("[[0,0,0,1],[0,0,0,0],[0,0,0,0],[1,0,0,0]]\n"))
}
