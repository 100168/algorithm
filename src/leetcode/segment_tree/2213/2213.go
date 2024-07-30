package main

import (
	"fmt"
	"math/bits"
)

/**
给你一个下标从 0 开始的字符串 s 。另给你一个下标从 0 开始、长度为 k 的字符串 queryCharacters ，一个下标从 0 开始、长度也是 k 的整数 下标 数组 queryIndices ，这两个都用来描述 k 个查询。

第 i 个查询会将 s 中位于下标 queryIndices[i] 的字符更新为 queryCharacters[i] 。

返回一个长度为 k 的数组 lengths ，其中 lengths[i] 是在执行第 i 个查询 之后 s 中仅由 单个字符重复 组成的 最长子字符串 的 长度 。



示例 1：

输入：s = "babacc", queryCharacters = "bcb", queryIndices = [1,3,3]
输出：[3,3,4]
解释：
- 第 1 次查询更新后 s = "bbbacc" 。由单个字符重复组成的最长子字符串是 "bbb" ，长度为 3 。
- 第 2 次查询更新后 s = "bbbccc" 。由单个字符重复组成的最长子字符串是 "bbb" 或 "ccc"，长度为 3 。
- 第 3 次查询更新后 s = "bbbbcc" 。由单个字符重复组成的最长子字符串是 "bbbb" ，长度为 4 。
因此，返回 [3,3,4] 。
*/

var bytes []byte

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {

	n := len(s)
	m := len(queryCharacters)
	bytes = []byte(s)
	ans := make([]int, m)
	t := make(seg, 2<<bits.Len(uint(n-1)))

	t.build(1, 0, n-1)

	for i, index := range queryIndices {
		bytes[index] = queryCharacters[i]
		t.update(1, index)
		ans[i] = t[1].max
	}
	return ans
}

type data struct {
}
type seg []struct {
	pre int
	suf int
	max int
	l   int
	r   int
}

func (t seg) maintain(o int) {

	lo, ro := t[o<<1], t[o<<1|1]
	t[o].pre = lo.pre
	t[o].suf = ro.suf
	t[o].max = max(lo.max, ro.max)
	if bytes[lo.r] == bytes[ro.l] {
		if lo.suf == lo.r-lo.l+1 {
			t[o].pre += ro.pre
		}
		if ro.pre == ro.r-ro.l+1 {
			t[o].suf += lo.suf
		}
		t[o].max = max(t[o].max, lo.suf+ro.pre)
	}

}

func (t seg) build(o, l, r int) {
	t[o].l = l
	t[o].r = r
	if l == r {
		t[o].max = 1
		t[o].pre = 1
		t[o].suf = 1
		return
	}
	m := (l + r) >> 1

	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i int) {

	if t[o].l == t[o].r {
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i)
	} else {
		t.update(o<<1|1, i)
	}
	t.maintain(o)
}

func main() {
	fmt.Println(longestRepeating("babacc", "bcb", []int{1, 3, 3}))
}
