package main

import "fmt"

/*
*
你被给定两个正整数 n 和 limit。
返回 在每个孩子得到不超过 limit 个糖果的情况下，将 n 个糖果分发给 3 个孩子的 总方法数。
*/
func distributeCandies(n int, limit int) int64 {

	return int64(comb(n+2) - (3*comb(n-limit+1) - 3*comb(n-2*limit) + comb(n-3*limit-1)))
}

func comb(n int) int {

	if n < 2 {
		return 0
	}
	return n * (n - 1) / 2
}

func main() {

	fmt.Println(distributeCandies(4, 2))
}
