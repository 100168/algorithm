package main

import (
	"fmt"
	"math"
)

/*
*
给定一个由正整数组成的数组 nums。

将数组拆分为 一个或多个 互相不覆盖的子数组，如下所示:

数组中的每个元素都 只属于一个 子数组，并且
每个子数组元素的 最大公约数 严格大于 1。
返回拆分后可获得的子数组的最小数目。

注意:

子数组的 最大公约数 是能将子数组中所有元素整除的最大正整数。
子数组 是数组的连续部分。
*/
func minimumSplits(nums []int) int {
	ans := 1
	g := 0
	for _, v := range nums {
		if g = gcd(v, g); g == 1 {
			ans++
			g = v
		}
	}
	return ans
}

func minimumSplits2(nums []int) int {

	memo := make([]int, len(nums))

	for i := range memo {
		memo[i] = -1
	}

	var dfs func(index int) int

	dfs = func(index int) int {

		if index < 0 {
			return 0
		}
		if memo[index] != -1 {
			return memo[index]
		}

		cur := math.MaxInt / 2

		g := nums[index]

		for j := index; j >= 0; j-- {
			g = gcd(nums[j], g)
			if g == 1 {
				break
			}
			cur = min(cur, dfs(j-1)+1)
		}
		memo[index] = cur
		return memo[index]
	}

	return dfs(len(nums) - 1)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(minimumSplits2([]int{12, 6, 3, 14, 8}))
}
