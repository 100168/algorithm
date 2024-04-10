package main

import (
	"fmt"
	"slices"
)

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
//序列。
//
// 示例 1：
//
//
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,3,2,3]
//输出：4
//
//
// 示例 3：
//
//
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2500
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：
//
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
//
//
// Related Topics 数组 二分查找 动态规划 👍 3615 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {

	n := len(nums)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = 1
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n]
}
func lengthOfLIS2(nums []int) int {

	n := len(nums)
	st := make([]int, 0)
	ans := 0
	for i := 0; i < n; i++ {
		l, r := 0, len(st)-1
		for l <= r {
			m := (r + l) / 2
			if nums[st[m]] >= nums[i] {
				r = m - 1
			} else {
				l = m + 1
			}
		}

		if l != len(st) {
			st[l] = i
		} else {
			st = append(st, i)
		}
		ans = max(ans, len(st))
	}
	return ans
}

type bitSet struct {
	sum []int
	len int
}

func lowBit(index int) int {
	return index & -index
}
func (b *bitSet) query(index int) int {

	s := 0
	for index > 0 {
		s = max(b.sum[index], s)
		index -= lowBit(index)
	}
	return s
}

func (b *bitSet) update(index int, val int) {
	for b.len > index {
		b.sum[index] = max(b.sum[index], val)
		index += lowBit(index)
	}
}

func lengthOfLIS3(nums []int) int {
	n := len(nums)
	b := new(bitSet)
	b.sum = make([]int, n+1)
	b.len = n + 1

	newNums := make([]int, n)
	copy(newNums, nums)
	compact := slices.Compact(newNums)

	sa := make([]int, len(compact))

	rkMap := make(map[int]int)
	for i := range compact {
		sa[i] = i
	}
	slices.SortStableFunc(sa, func(a, b int) int {
		return compact[a] - compact[b]
	})
	for i := range sa {
		rkMap[compact[sa[i]]] = i
	}
	addMap := make(map[int]bool)
	ans := 0
	for _, v := range nums {
		cur := b.query(rkMap[v])
		ans = max(ans, cur+1)
		if !addMap[rkMap[v]] {
			b.update(rkMap[v]+1, ans)
		}
		addMap[rkMap[v]] = true
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLIS3([]int{10, 9, 2, 5, 3, 7, 101, 1}))
}

//leetcode submit region end(Prohibit modification and deletion)
