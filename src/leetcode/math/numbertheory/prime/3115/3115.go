package main

import "fmt"

func maximumPrimeDifference(nums []int) int {

	ans := 0
	p := -1
	for i, v := range nums {

		f := true
		if v == 1 {
			f = false
		}
		for c := 2; c*c <= v; c++ {
			if v%c == 0 {
				f = false
				break
			}
		}
		if f && p >= 0 {
			ans = max(ans, i-p)
		}
		if f && p == -1 {
			p = i
		}

	}
	return ans
}

func main() {
	fmt.Println(maximumPrimeDifference([]int{4, 2, 9, 5, 3}))
}
