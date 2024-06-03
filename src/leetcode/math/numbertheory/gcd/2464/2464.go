package main

import (
	"fmt"
	"math"
)

/*
*
给定一个整数数组 nums。

如果要将整数数组 nums 拆分为 子数组 后是 有效的，则必须满足:

每个子数组的第一个和最后一个元素的最大公约数 大于 1，且
nums 的每个元素只属于一个子数组。
返回 nums 的 有效 子数组拆分中的 最少 子数组数目。如果不能进行有效的子数组拆分，则返回 -1。

注意:

两个数的 最大公约数 是能整除两个数的最大正整数。
子数组 是数组中连续的非空部分。
*/
func validSubarraySplit(nums []int) int {

	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int

	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		cur := math.MaxInt / 2
		for j := i; j >= 0; j-- {
			if gcd(nums[i], nums[j]) > 1 {
				cur = min(cur, dfs(j-1)+1)
			}
		}
		memo[i] = cur
		return cur
	}

	ans := dfs(n - 1)
	if ans == math.MaxInt/2 {
		return -1
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

	fmt.Println(validSubarraySplit([]int{1, 2, 1}))
	fmt.Println(validSubarraySplit([]int{3, 5}))
}
