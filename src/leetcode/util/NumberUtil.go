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

	fmt.Println(brace("[[0,0,2],[0,1,5],[2,2,5],[0,2,4]]"))
}
