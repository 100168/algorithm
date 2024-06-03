package main

import (
	"fmt"
	"math"
)

/*
*
给你一个下标从 0 开始的 正 整数数组 nums 。你可以对数组执行以下操作 任意 次：

选择一个满足 0 <= i < n - 1 的下标 i ，将 nums[i] 或者 nums[i+1] 两者之一替换成它们的最大公约数。
请你返回使数组 nums 中所有元素都等于 1 的 最少 操作次数。如果无法让数组全部变成 1 ，请你返回 -1 。

两个正整数的最大公约数指的是能整除这两个数的最大正整数。

输入：nums = [2,6,3,4]
输出：4
解释：我们可以执行以下操作：
- 选择下标 i = 2 ，将 nums[2] 替换为 gcd(3,4) = 1 ，得到 nums = [2,6,1,4] 。
- 选择下标 i = 1 ，将 nums[1] 替换为 gcd(6,1) = 1 ，得到 nums = [2,1,1,4] 。
- 选择下标 i = 0 ，将 nums[0] 替换为 gcd(2,1) = 1 ，得到 nums = [1,1,1,4] 。
- 选择下标 i = 2 ，将 nums[3] 替换为 gcd(1,4) = 1 ，得到 nums = [1,1,1,1]
*/
func minOperations(nums []int) int {

	type pair struct{ x, y int }

	var gcdList []pair
	minGcdOneL := math.MaxInt

	cntOne := 0
	for i, v := range nums {
		var newGcdList []pair
		gcdList = append(gcdList, pair{v, i})
		for _, x := range gcdList {
			cur := gcd(x.x, v)
			if cur == 1 {
				minGcdOneL = min(minGcdOneL, i-x.y)
			}
			if len(newGcdList) > 1 && newGcdList[len(newGcdList)-1].x == cur {
				newGcdList[len(newGcdList)-1].y = x.y
			} else {
				newGcdList = append(newGcdList, pair{cur, x.y})
			}
		}
		if v == 1 {
			cntOne++
		}
		gcdList = newGcdList

	}

	if minGcdOneL == math.MaxInt {
		return -1
	}
	if cntOne > 0 {
		return len(nums) - cntOne
	}

	return len(nums) - 1 + minGcdOneL

}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(minOperations([]int{6, 10, 15}))
}
