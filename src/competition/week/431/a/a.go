package main

import "fmt"

func maxLength(nums []int) int {

	n := len(nums)

	ans := 0

	for i := 0; i < n; i++ {

		s := nums[i]
		g := nums[i]
		lcm := nums[i]

		for j := i + 1; j < n; j++ {
			g = gcd(g, nums[j])

			g1 := gcd(lcm, nums[j])

			lcm = lcm * nums[j] / g1

			s *= nums[j]

			if s == lcm*g {
				ans = max(ans, j-i+1)
			}

		}
	}
	return ans
}

func gcd(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(maxLength([]int{2, 6}))
}
