package main

import (
	"math"
	"math/bits"
)

/*
*
给你一个整数数组 nums 和一个整数 k 。你需要将这个数组划分到 k 个相同大小的子集中，使得同一个子集里面没有两个相同的元素。

一个子集的 不兼容性 是该子集里面最大值和最小值的差。

请你返回将数组分成 k 个子集后，各子集 不兼容性 的 和 的 最小值 ，如果无法分成分成 k 个子集，返回 -1 。

子集的定义是数组中一些数字的集合，对数字顺序没有要求。
*/
func minimumIncompatibility(nums []int, k int) int {

	n := len(nums)
	m := n / k

	memo := make([]int, 1<<n)
	for i := range memo {
		memo[i] = -1

	}

	values := make(map[int]int)

	//先预处理
	for j := 0; j < 1<<n; j++ {
		if bits.OnesCount(uint(j)) < m {
			continue
		}
		check := make(map[int]bool)
		s := 0
		curMax := 0
		curMin := math.MaxInt / 2
		for i := 0; i < n; i++ {
			if 1<<i&j != 0 && !check[nums[i]] {
				s++
				check[nums[i]] = true
				curMax = max(curMax, nums[i])
				curMin = min(curMin, nums[i])
			}
		}
		if s == m {
			values[j] = curMax - curMin
		}
	}
	//自上到下递归
	var dfs func(int) int
	dfs = func(mask int) int {

		if mask == 0 {
			return 0
		}

		if memo[mask] != -1 {
			return memo[mask]
		}
		cur := math.MaxInt / 2

		seen := make(map[int]int)
		for i := 0; i < n; i++ {
			if (mask & (1 << i)) != 0 {
				seen[nums[i]] = i
			}
		}
		sub := 0
		for _, v := range seen {
			sub |= 1 << v
		}
		//枚举子集的子集妙啊
		for nxt := sub; nxt > 0; nxt = (nxt - 1) & sub {
			if val, ok := values[nxt]; ok {
				cur = min(cur, dfs(mask^nxt)+val)
			}
		}
		memo[mask] = cur
		return cur
	}
	ans := dfs(1<<n - 1)
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}

func minimumIncompatibility2(nums []int, k int) int {
	n := len(nums)
	group := n / k
	inf := math.MaxInt32
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	values := make(map[int]int)

	for mask := 1; mask < (1 << n); mask++ {
		if bits.OnesCount(uint(mask)) != group {
			continue
		}
		minVal := 20
		maxVal := 0
		cur := make(map[int]bool)
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				if cur[nums[i]] {
					break
				}
				cur[nums[i]] = true
				minVal = min(minVal, nums[i])
				maxVal = max(maxVal, nums[i])
			}
		}
		if len(cur) == group {
			values[mask] = maxVal - minVal
		}
	}

	for mask := 0; mask < (1 << n); mask++ {
		if dp[mask] == inf {
			continue
		}
		seen := make(map[int]int)
		for i := 0; i < n; i++ {
			if (mask & (1 << i)) == 0 {
				seen[nums[i]] = i
			}
		}
		if len(seen) < group {
			continue
		}
		sub := 0
		for _, v := range seen {
			sub |= 1 << v
		}
		nxt := sub
		for nxt > 0 {
			if val, ok := values[nxt]; ok {
				dp[mask|nxt] = min(dp[mask|nxt], dp[mask]+val)
			}
			nxt = (nxt - 1) & sub
		}
	}
	if dp[(1<<n)-1] < inf {
		return dp[(1<<n)-1]
	}
	return -1
}

func main() {
	println(minimumIncompatibility([]int{1, 1, 1, 0}, 2))
	println(minimumIncompatibility([]int{6, 3, 8, 1, 3, 1, 2, 2}, 4))
}
