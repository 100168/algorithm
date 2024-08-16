package main

import (
	"fmt"
	"math"
	"sort"
)

/*
*
给你 n 个包裹，你需要把它们装在箱子里，每个箱子装一个包裹。总共有 m 个供应商提供 不同尺寸 的箱子（每个规格都有无数个箱子）。
如果一个包裹的尺寸 小于等于 一个箱子的尺寸，那么这个包裹就可以放入这个箱子之中。

包裹的尺寸用一个整数数组 packages 表示，其中 packages[i] 是第 i 个包裹的尺寸。
供应商用二维数组 boxes 表示，其中 boxes[j] 是第 j 个供应商提供的所有箱子尺寸的数组。

你想要选择 一个供应商 并只使用该供应商提供的箱子，使得 总浪费空间最小 。
对于每个装了包裹的箱子，我们定义 浪费的 空间等于 箱子的尺寸 - 包裹的尺寸 。总浪费空间 为 所有 箱子中浪费空间的总和。

比方说，如果你想要用尺寸数组为 [4,8] 的箱子装下尺寸为 [2,3,5] 的包裹，
你可以将尺寸为 2 和 3 的两个包裹装入两个尺寸为 4 的箱子中，
同时把尺寸为 5 的包裹装入尺寸为 8 的箱子中。总浪费空间为 (4-2) + (4-3) + (8-5) = 6 。
请你选择 最优 箱子供应商，使得 总浪费空间最小 。如果 无法 将所有包裹放入箱子中，请你返回 -1 。
由于答案可能会 很大 ，请返回它对 109 + 7 取余 的结果。

示例 1：

输入：packages = [2,3,5], boxes = [[4,8],[2,8]]
输出：6
解释：选择第一个供应商最优，用两个尺寸为 4 的箱子和一个尺寸为 8 的箱子。
总浪费空间为 (4-2) + (4-3) + (8-5) = 6 。
*/
func minWastedSpace(packages []int, boxes [][]int) int {

	mod := int(1e9 + 7)
	sort.Ints(packages)
	for i := range boxes {
		sort.Ints(boxes[i])
	}

	n := len(packages)
	sum := make([]int, n+1)

	for i, v := range packages {
		sum[i+1] = sum[i] + v
	}

	find := func(t int) int {
		l, r := 0, n-1
		for l <= r {
			m := (l + r) / 2
			if packages[m] > t {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		return r
	}

	ans := math.MaxInt
	for _, box := range boxes {
		if box[len(box)-1] < packages[n-1] {
			continue
		}
		s := 0
		pre := -1
		for _, v := range box {
			end := find(v)
			if end < 0 {
				continue
			}
			if pre == 0 {

			}
			s += v*(end-pre) - (sum[end+1] - sum[pre+1])
			pre = end
			if pre == n {
				break
			}

		}
		ans = min(ans, s)

	}
	if ans == math.MaxInt {
		return -1
	}
	return ans % mod
}

func main() {
	fmt.Println(minWastedSpace([]int{2, 3, 5}, [][]int{{4, 8}, {2, 8}}))
	fmt.Println(minWastedSpace([]int{2, 3, 5}, [][]int{{1, 4}, {2, 3}, {3, 4}}))
}
