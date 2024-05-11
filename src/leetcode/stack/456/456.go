package main

import (
	"math"
	"slices"
	"sort"
)

/*
*
给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列
由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。

如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。
*/

// 离散化+树状数组
func find132pattern2(nums []int) bool {

	n := len(nums)
	newNums := make([]int, n)

	copy(newNums, nums)

	sort.Ints(newNums)
	newNums = slices.Compact(newNums)

	sa := make(map[int]int)

	for i, v := range newNums {
		sa[v] = i
	}

	bt := new(bitTree)
	bt.sum = make([]int, n+1)
	bt.l = n + 1
	for _, v := range nums[1:] {
		bt.update(sa[v]+1, 1)
	}

	minVal := nums[0]
	for _, v := range nums[1:] {
		if v > minVal {
			diff := bt.query(sa[v]) - bt.query(sa[minVal]+1)
			if diff > 0 {
				return true
			}
		}
		bt.update(sa[v]+1, -1)
		minVal = min(v, minVal)
	}
	return false
}

type bitTree struct {
	sum []int
	l   int
}

func lowBit(index int) int {
	return index & -index
}

func (bt bitTree) query(index int) int {

	ans := 0
	for index > 0 {
		ans += bt.sum[index]
		index -= lowBit(index)
	}
	return ans
}

func (bt bitTree) update(index int, val int) {
	for index < bt.l {
		bt.sum[index] += val
		index += lowBit(index)
	}
}

func find132pattern(nums []int) bool {

	n := len(nums)
	st := []int{nums[n-1]}

	maxK := math.MinInt
	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(st) > 0 && st[len(st)-1] < nums[i] {
			maxK = max(maxK, st[len(st)-1])
			st = st[:len(st)-1]
		}
		st = append(st, nums[i])
	}
	return false
}

func main() {
	println(find132pattern2([]int{3, 5, 0, 3, 4}))
}
