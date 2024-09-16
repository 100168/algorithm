package main

import "fmt"

func sumScores(s string) int64 {

	n := len(s)

	z := make([]int, n)

	ans := len(s)

	for i, c, r := 1, 1, 1; i < n; i++ {

		l := min(z[i-c], max(r-i, 0))
		for i+l < n && s[i+l] == s[l] {
			l++
		}
		if i+l > r {
			r = i + l
			c = i
		}
		z[i] = l

		ans += l
	}
	return int64(ans)
}

func main() {
	fmt.Println(sumScores("babab"))
}
