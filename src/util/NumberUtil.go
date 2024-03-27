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
	println(brace("[5,11],[20,22],[1,3],[21,22],[11,11]"))
}
