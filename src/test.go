package main

import "math"

func minimumPossibleSum(n int, target int) int {
	//4/2 = 2
	//5/2 = 2
	mid := target / 2
	if n <= mid {
		return (1 + n) * n / 2
	}

	lSum := (1 + mid) * mid / 2
	n -= mid
	rSum := (target + target + n - 1) * n / 2
	return lSum + rSum
}
func main() {
	println(minimumPossibleSum(1000000000, 100000000))
	println(math.MaxInt32)
}
