package main

import "fmt"

func sumScores(s string) int64 {

	n := len(s)

	z := make([]int, n)

	ans := len(s)
	// 10  c = 18  25 -18 = 7
	//18 19 20 21 22 23 24 25
	//0  1  2  3  4  5  6  7

	//25 26 27 28
	for i, c, r := 1, 1, 1; i < n; i++ {

		//当前长度
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
