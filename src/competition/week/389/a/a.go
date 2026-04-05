package main

import (
	"strings"
)

func isSubstringPresent(s string) bool {

	n := len(s)
	re := reverse(s)
	for i := 2; i <= n; i++ {

		curs := s[i-2 : i]
		if strings.Contains(re, curs) {
			return true
		}
	}
	return false
}

func reverse(s string) string {

	ans := ""
	for _, v := range s {
		ans = string(v) + ans
	}
	return ans
}
