package main

import (
	"fmt"
	"slices"
)

/*
*给你一个下标从 0 开始、由 正整数 组成的数组 nums。

将数组分割成一个或多个 连续 子数组，如果不存在包含了相同数字的两个子数组，
则认为是一种 好分割方案 。

返回 nums 的 好分割方案 的 数目。

由于答案可能很大，请返回答案对 109 + 7 取余 的结果。
*/
/**
只统计右边区间
*/
func numberOfGoodPartitions(nums []int) int {
	r := map[int]int{}
	for i, x := range nums {
		r[x] = i
	}
	ans := 1
	maxR := 0
	for i, x := range nums[:len(nums)-1] { // 少统计最后一段区间
		maxR = max(maxR, r[x])
		if maxR == i { // 区间无法延长
			ans = ans * 2 % 1_000_000_007
		}
	}
	return ans
}

/*
*
区间合并
*/
func numberOfGoodPartitions2(nums []int) int {
	type pair struct{ l, r int }
	ps := map[int]pair{}
	for i, x := range nums {
		if p, ok := ps[x]; ok {
			p.r = i // 更新区间右端点
			ps[x] = p
		} else {
			ps[x] = pair{i, i}
		}
	}

	a := make([]pair, 0, len(ps))
	for _, p := range ps {
		a = append(a, p)
	}
	slices.SortFunc(a, func(a, b pair) int { return a.l - b.l }) // 按区间左端点排序

	ans := 1
	maxR := a[0].r
	for _, p := range a[1:] {
		if p.l > maxR { // 无法合并
			ans = ans * 2 % 1_000_000_007
		}
		maxR = max(maxR, p.r)
	}
	return ans
}

func main() {
	//fmt.Println(numberOfGoodPartitions([]int{1, 2, 1, 3}))
	fmt.Println(numberOfGoodPartitions([]int{2, 3, 3, 8, 8}))
}
