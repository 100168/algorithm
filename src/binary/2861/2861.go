package main

import "math"

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {

	var check func(int) bool

	check = func(m int) bool {
		minCost := math.MaxInt
		for i := 0; i < k; i++ {
			curCost := 0
			for j, v := range composition[i] {
				curCost += v * (cost[j]*m - stock[j])
			}

			minCost = min(curCost, minCost)
		}

		return minCost <= budget
	}
	l, r := 1, budget
	for l <= r {
		m := (r-l)/2 + l
		if check(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}
func main() {
	maxNumberOfAlloys(2, 3, 10, [][]int{{2, 1}, {1, 2}, {1, 1}}, []int{1, 1}, []int{5, 5})
}
