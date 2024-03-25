package main

import "math"

func shipWithinDays(weights []int, days int) int {

	l, r := 1, math.MaxInt

	for l <= r {
		m := (r-l)/2 + l
		if check(weights, m) <= days {
			r = m - 1
		} else {
			l = m + 1
		}

	}
	return l
}

func check(weight []int, t int) int {
	sum := 1
	cnt := 0
	for _, v := range weight {
		if cnt+v > t {
			sum++
			cnt = 0
		}
		cnt += v
	}

	return sum
}

func main() {
	println(check([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
}
