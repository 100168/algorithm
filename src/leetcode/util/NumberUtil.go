package main

import (
	"fmt"
	"strings"
)

func brace(s string) string {
	s = strings.Replace(s, "[", "{", len(s))
	s = strings.Replace(s, "]", "}", len(s))
	return s
}
func main() {

	//{{1,5},
	//{1,5},
	//{1,5},
	//{2,3},
	//{2,3}}
	fmt.Println(brace("[[1,5],[1,5],[1,5],[2,3],[2,3]]"))
}
