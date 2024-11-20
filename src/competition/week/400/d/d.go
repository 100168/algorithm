package main

import (
	"math"
)

/*

给你一个数组 nums 和一个整数 k 。你需要找到 nums 的一个
子数组
 ，满足子数组中所有元素按位与运算 AND 的值与 k 的 绝对差 尽可能 小 。换言之，
你需要选择一个子数组 nums[l..r] 满足 |k - (nums[l] AND nums[l + 1] ... AND nums[r])| 最小。

请你返回 最小 的绝对差值。

子数组是数组中连续的 非空 元素序列。
*/

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minimumDifference(arr []int, target int) int {

	ans := math.MaxInt

	pres := make([]int, 0)
	for _, v := range arr {

		pres = append(pres, v)

		newPres := make([]int, 0)

		for _, e := range pres {

			cur := e & v
			if len(newPres) > 0 && newPres[len(newPres)-1] == cur {
				continue
			}
			newPres = append(newPres, cur)
			ans = min(ans, abs(cur-target))
		}
		pres = newPres
	}
	return ans
}

func main() {
	minimumDifference([]int{1, 2, 1, 2}, 9)
}
