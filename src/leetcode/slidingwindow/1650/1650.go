package main

import "fmt"

/*
给你一个整数数组 nums 和一个整数 x 。
每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。
请注意，需要 修改 数组以供接下来的操作使用。

如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。
*/
func minOperations(nums []int, x int) int {

	n := len(nums)

	l := 0
	s := 0

	ans := -1
	w := 0
	for _, v := range nums {
		s += v
	}
	rest := s - x

	if rest < 0 {
		return -1
	}

	for r := 0; r < n; r++ {
		w += nums[r]
		for w > rest {
			w -= nums[l]
			l++
		}
		if w == rest {
			ans = max(ans, r-l+1)
		}
	}
	if ans == -1 {
		return ans
	}
	return n - ans
}
func main() {
	fmt.Println(minOperations([]int{1, 1, 1, 1, 2, 3, 1}, 5))
	fmt.Println(minOperations([]int{1, 1, 1, 1, 2, 3, 1}, 7))
	fmt.Println(minOperations([]int{1, 1, 4, 2, 3}, 5))
	fmt.Println(minOperations([]int{1, 1}, 3))
	fmt.Println(minOperations([]int{8828, 9581, 49, 9818, 9974, 9869, 9991, 10000, 10000, 10000, 9999, 9993, 9904, 8819, 1231, 6309}, 134365))
}
