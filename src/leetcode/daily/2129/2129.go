package main

import (
	"strings"
)

func capitalizeTitle(title string) string {

	lower := strings.ToLower(title)

	split := strings.Split(lower, " ")

	ans := ""
	for i := range split {
		if len(split[i]) > 2 {
			curS := []byte(split[i])
			curS[0] ^= 32
			ans += "" + string(curS)
		} else {
			ans += " " + split[i]
		}
	}
	return strings.TrimLeft(ans, " ")
}

func main() {
	println(capitalizeTitle("First leTTeR of EACH Word"))
}
