package main

import "fmt"

/*
*
给你一个长度为 n 的整数数组 nums 和一个正整数 k 。

一个数组的 能量值 定义为：

如果 所有 元素都是依次 连续 且 上升 的，那么能量值为 最大 的元素。
否则为 -1 。
你需要求出 nums 中所有长度为 k 的
子数组的能量值。

请你返回一个长度为 n - k + 1 的整数数组 results ，其中 results[i] 是子数组 nums[i..(i + k - 1)] 的能量值。
*/
func resultsArray(nums []int, k int) []int {

	n := len(nums)

	ans := make([]int, n-k+1)

	pre := k - 1

	if k == 1 {
		return nums
	}

next:
	for i := 1; i < n; i++ {
		if nums[i-1]+1 != nums[i] {
			// 最右  i+k-1
			// 最左
			for ; pre <= min(i+k-2, n-1); pre++ {
				ans[pre-k+1] = -1
			}
			continue next
		}
		if i == pre {
			ans[i-k+1] = nums[i]
			pre++
		}
	}
	return ans
}

func main() {
	fmt.Println(resultsArray([]int{1, 2, 3, 4, 3, 2, 5}, 1))
	//fmt.Println(resultsArray([]int{2, 2, 2, 2, 2}, 4))
	//fmt.Println(resultsArray([]int{3, 2, 3, 2, 3, 2}, 2))
	//fmt.Println(resultsArray([]int{1, 3, 4}, 2))
	//fmt.Println(resultsArray([]int{1}, 1))
}
