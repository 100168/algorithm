package main

import "fmt"

/*
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。



示例 1：

输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
示例 2:

输入: nums = [1,2,3,4], k = 3
输出: false

*/

func canPartitionKSubsets(nums []int, k int) bool {

	s := 0
	for _, v := range nums {
		s += v
	}
	if s%k != 0 {
		return false
	}

	n := len(nums)
	f := make([]int, 1<<n)

	mul := s / k

	var dfs func(int, int) bool

	dfs = func(mask int, rest int) bool {
		if mask == 1<<n-1 {
			return true
		}
		if f[mask] != 0 {
			return f[mask] == 1
		}
		cur := false
		for i := 0; i < n; i++ {
			if 1<<i&mask == 0 && rest+nums[i] <= mul {
				cur = cur || dfs(1<<i|mask, (rest+nums[i])%mul)
			}
		}
		if cur {
			f[mask] = 1
		} else {
			f[mask] = -1
		}
		return cur
	}
	return dfs(0, 0)
}

func main() {

	//fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 2))
}
