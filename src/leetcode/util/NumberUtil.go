package main

import (
	"strings"
)

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

	println(brace("[[9,6],[9,5],[4,15],[7,5],[13,5],[1,10],[13,6],[14,3],[2,13],[14,10],[8,1],[8,9],[14,6],[14,1],[4,12],[9,10],[4,6],[2,6],[2,14],[15,9],[13,9],[12,9],[11,5],[8,3],[4,9],[13,10],[6,10],[11,1],[13,7],[2,15],[12,10],[4,8],[15,1],[7,10],[14,13],[4,11],[7,9],[1,3],[15,10],[2,11],[2,7],[4,1],[4,3],[11,3],[11,9]]"))
}
