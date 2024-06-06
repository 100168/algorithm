package main

import (
	"fmt"
	"math"
)

const MOD = int(1e9) + 7

func maximumSumSubsequence(nums []int, queries [][]int) int {

	return maxNonAdjacentSum(nums, queries)
}

func updateAndCalcMaxSum(nums []int, posi, xi int) int {
	// 只需要更新 posi 位置的元素以及可能影响的相邻元素
	n := len(nums)
	if posi > 0 {
		prev := nums[posi-1]
		nums[posi-1] = math.MinInt32 // 临时设为极小值，保证不选它
		if posi%2 == 0 {
			// posi 是偶数索引，需要更新 prev_odd
			if posi-2 >= 0 {
				nums[posi-2] = max(nums[posi-2], prev+nums[posi-1])
			}
		} else {
			if posi-2 >= 0 {
				nums[posi-2] = prev + nums[posi-1]
			}
			// posi 是奇数索引，需要更新 prev_even

		}
	}

	nums[posi] = xi // 更新 nums[posi]

	// 计算新的最大和
	maxSum := xi
	if posi > 0 {
		maxSum = max(maxSum, nums[posi-1])
	}
	if posi < n-1 {
		maxSum = max(maxSum, nums[posi+1])
	}

	// 使用动态规划重新计算 prev_even 和 prev_odd（从 posi 开始）
	prevEven, prevOdd := xi, 0
	for i := posi + 1; i < n; i++ {
		if i%2 == 0 {
			newPrevEven := max(prevEven, prevOdd+nums[i])
			prevOdd = prevEven
			prevEven = newPrevEven
			nums[i-1] = prevEven // 更新 nums 数组是为了下一次查询时的快速更新
		} else {
			newPrevOdd := prevEven + nums[i]
			prevEven = prevOdd
			prevOdd = newPrevOdd
			nums[i-1] = prevOdd // 更新 nums 数组是为了下一次查询时的快速更新
		}
	}

	// 返回当前查询的答案
	return max(prevEven, prevOdd)
}

func maxNonAdjacentSum(nums []int, queries [][]int) int {
	totalSum := 0
	for _, q := range queries {
		posi, xi := q[0], q[1]
		sum := updateAndCalcMaxSum(nums, posi, xi)
		totalSum = (totalSum + sum) % MOD
	}
	return totalSum
}

func main() {
	nums := []int{3, 5, 9}
	queries := [][]int{{1, -2}, {0, -3}}
	result := maxNonAdjacentSum(nums, queries)
	fmt.Println(result)
}
