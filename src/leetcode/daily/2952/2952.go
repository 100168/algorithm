package main

import "sort"

func minimumAddedCoins(coins []int, target int) int {

	l, r := 0, 0

	sort.Ints(coins)
	cnt := 0
	n := len(coins)
	for i := 0; i < n; i++ {
		x, y := coins[i]+l, coins[i]+r
		for x > r+1 {
			cnt += 1
			r = r + 1 + r
			y = r + coins[i]
		}
		r = max(r, y)
		if r >= target {
			return cnt
		}

	}
	for r < target {
		cnt++
		r = r + 1 + r
	}
	return cnt
}

func main() {
	println(minimumAddedCoins([]int{1, 4, 10}, 19))

	println(1 << 31)
}
