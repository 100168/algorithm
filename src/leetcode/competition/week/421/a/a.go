package main

import "fmt"

/**
给你一个整数数组 nums。

因子得分 定义为数组所有元素的最小公倍数（LCM）与最大公约数（GCD）的 乘积。

在 最多 移除一个元素的情况下，返回 nums 的 最大因子得分。

注意，单个数字的
LCM
 和
GCD
 都是其本身，而 空数组 的因子得分为 0。



示例 1：

输入： nums = [2,4,8,16]

输出： 64

解释：

移除数字 2 后，剩余元素的 GCD 为 4，LCM 为 16，因此最大因子得分为 4 * 16 = 64。

*/

func maxScore(nums []int) int64 {

	ans := 0
	for i := range nums {

		g := 0
		lcm := 1
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			g = gcd(nums[j], g)
			lcm = lcm * nums[j] / gcd(lcm, nums[j])
		}

		ans = max(ans, lcm*g)
	}

	g := 0
	lcm := nums[0]
	for j := 0; j < len(nums); j++ {
		g = gcd(nums[j], g)
		lcm = lcm * nums[j] / gcd(lcm, nums[j])
	}
	ans = max(ans, lcm*g)
	return int64(ans)
}

func gcd(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(maxScore([]int{2, 4, 8, 16}))
}
