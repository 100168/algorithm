package main

import (
	"fmt"
	"math/bits"
)

/**
给你一个整数数组 nums 和一个二维数组 queries，其中 queries[i] = [posi, xi]。

对于每个查询 i，首先将 nums[posi] 设置为 xi，然后计算查询 i 的答案，该答案为 nums 中 不包含相邻元素 的
子序列
 的 最大 和。

返回所有查询的答案之和。

由于最终答案可能非常大，返回其对 109 + 7 取余 的结果。

子序列 是指从另一个数组中删除一些或不删除元素而不改变剩余元素顺序得到的数组。

示例 1：

输入：nums = [3,5,9], queries = [[1,-2],[0,-3]]

输出：21

解释：
执行第 1 个查询后，nums = [3,-2,9]，不包含相邻元素的子序列的最大和为 3 + 9 = 12。
执行第 2 个查询后，nums = [-3,-2,9]，不包含相邻元素的子序列的最大和为 9 。
*/

const MOD = int(1e9) + 7

// f00 表示第一个数一定不选，最后一个数一定不选
// f01 表示第一个数一定不选，最后一个数可选可不选
// f10 表示第一个数可选可不选，最后一个数一定不选
// f11 表示第一个数可选可不选，最后一个数可选可不选，也就是没有任何限制
type data struct{ f00, f01, f10, f11 int }
type seg []data

func (t seg) maintain(o int) {
	a, b := t[o<<1], t[o<<1|1]
	t[o] = data{
		max(a.f00+b.f10, a.f01+b.f00),
		max(a.f00+b.f11, a.f01+b.f01),
		max(a.f10+b.f10, a.f11+b.f00),
		max(a.f10+b.f11, a.f11+b.f01),
	}
}

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o].f11 = max(a[l], 0)
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, i, val int) {
	if l == r {
		t[o].f11 = max(val, 0)
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(o<<1, l, m, i, val)
	} else {
		t.update(o<<1|1, m+1, r, i, val)
	}
	t.maintain(o)
}

func maximumSumSubsequence(nums []int, queries [][]int) (ans int) {
	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)
	for _, q := range queries {
		t.update(1, 0, n-1, q[0], q[1])
		ans += t[1].f11 // 注意 f11 没有任何限制，也就是整个数组的打家劫舍
	}
	return ans % 1_000_000_007
}

func main() {
	nums := []int{3, 5, 9}
	queries := [][]int{{1, -2}, {0, -3}}
	result := maximumSumSubsequence(nums, queries)
	fmt.Println(result)
}
