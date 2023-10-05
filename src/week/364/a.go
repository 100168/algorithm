package main

import (
	"strings"
)

func maximumOddBinaryNumber(s string) string {
	count := strings.Count(s, "1")
	ans := make([]byte, len(s))
	for i := range ans {
		ans[i] = '0'
	}
	index := 0
	for i := 1; i < count; i++ {
		ans[index] = '1'
		index++
	}

	ans[len(s)-1] = '1'
	return string(ans)
}

func main() {
	println(maximumOddBinaryNumber("010101"))
}
