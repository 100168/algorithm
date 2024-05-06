package main

import (
	"fmt"
	"sort"
)

/*
*
给你一个长度为 n 的整数数组 nums ，请你求出每个长度为 k 的子数组的 美丽值 。

一个子数组的 美丽值 定义为：如果子数组中第 x 小整数 是 负数 ，那么美丽值为第 x 小的数，否则美丽值为 0 。

请你返回一个包含 n - k + 1 个整数的数组，依次 表示数组中从第一个下标开始，每个长度为 k 的子数组的 美丽值 。

子数组指的是数组中一段连续 非空 的元素序列。
*/

// 数据范围比较小，可以直接暴力做+计数排序
// 数据范围大可以用离散化+数状数组+二分
func getSubarrayBeauty(nums []int, k int, x int) []int {

	n := len(nums)
	ans := make([]int, n-k+1)

	b := bt{}
	b.n = n + 1
	b.sum = make([]int, n+1)

	rkNums := make([]int, n)
	copy(rkNums, nums)
	sort.Ints(rkNums)
	sa := make(map[int]int)
	rk := make(map[int]int)
	for i, v := range rkNums {
		sa[v] = i
		rk[i] = v
	}

	for i := 0; i < k-1; i++ {
		rk := sa[nums[i]] + 1
		b.update(rk, 1)
	}

	for i := k - 1; i < n; i++ {
		rank := sa[nums[i]] + 1
		b.update(rank, 1)

		l, r := 1, n
		for l <= r {
			m := (r + l) / 2
			cur := b.query(m)
			if cur >= x {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		cur := rk[l-1]
		if cur < 0 {
			ans[i-k+1] = cur
		}

		b.update(sa[nums[i-k+1]]+1, -1)
	}

	return ans
}

type bt struct {
	sum []int
	n   int
}

func lowBit(index int) int {
	return index & -index
}

func (b bt) query(index int) int {

	ans := 0
	for index > 0 {
		ans += b.sum[index]
		index -= lowBit(index)
	}

	return ans
}

func (b bt) update(index int, val int) {
	for index < b.n {
		b.sum[index] += val
		index += lowBit(index)
	}
}

func main() {
	fmt.Println(getSubarrayBeauty([]int{1, -1, -3, -2, 3}, 3, 2))
}
