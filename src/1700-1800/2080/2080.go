package main

import "sort"

// RangeFreqQuery 2080. 区间内查询数字的频率
// 中等
// 相关标签
// 相关企业
// 提示
// 请你设计一个数据结构，它能求出给定子数组内一个给定值的 频率 。
//
// 子数组中一个值的 频率 指的是这个子数组中这个值的出现次数。
//
// 请你实现 RangeFreqQuery 类：
//
// RangeFreqQuery(int[] arr) 用下标从 0 开始的整数数组 arr 构造一个类的实例。
// int query(int left, int right, int value) 返回子数组 arr[left...right] 中 value 的 频率 。
// 一个 子数组 指的是数组中一段连续的元素。arr[left...right] 指的是 nums 中包含下标 left 和 right 在内 的中间一段连续元素。/***/
type RangeFreqQuery struct {
	m map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	r := new(RangeFreqQuery)
	r.m = make(map[int][]int)

	for i, v := range arr {
		r.m[v] = append(r.m[v], i)
	}
	return *r
}

func (r *RangeFreqQuery) Query(left int, right int, value int) int {
	left = sort.SearchInts(r.m[value], left)
	right = sort.SearchInts(r.m[value], right+1)

	return right - left
}

func search(nums []int, target int) int {

	l := -1
	r := len(nums)
	for l+1 < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m
		} else {
			l = m
		}
	}
	return r
}
func main() {
	freqQuery := Constructor([]int{4, 2, 9, 7, 5, 5, 1, 9, 8, 6})
	println(freqQuery.Query(2, 6, 10))
}
