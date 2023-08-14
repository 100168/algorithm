package main

import "fmt"

func longestAwesome(s string) int {

	n := len(s)
	status := 0
	pos := make([]int, 1<<10)
	for i := 1; i < 1<<10; i++ {
		pos[i] = -1
	}
	ans := 0

	for i := 0; i < n; i++ {
		pre := status
		status ^= 1 << (s[i] - '0')
		if pos[status] >= 0 {
			ans = Max(Max(ans, i+1-pos[status]), i+1-pos[pre])
		} else {
			pos[status] = i + 1
		}
		for k := 0; k < 10; k++ {
			if pos[status^(1<<k)] != -1 {
				ans = Max(ans, i+1-pos[status^(1<<k)])
			}
		}
	}
	return ans
}

func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
func main() {
	awesome := longestAwesome("3242415")

	fmt.Println(awesome)
}
