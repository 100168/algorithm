package main

import (
	"fmt"
	"strings"
)

func hasMatch(s string, p string) bool {

	split := strings.Split(p, "*")

	f := strings.Index(s, split[0])
	f2 := strings.LastIndex(s, split[1])

	if len(split[0]) == 0 {
		return f2 >= 0
	}

	if len(split[1]) == 0 {
		return f >= 0
	}

	if f >= 0 && f+len(split[0]) <= f2 {
		return true
	}
	return false

}

func main() {
	fmt.Println(hasMatch("leetcode", "ee*e"))
}
