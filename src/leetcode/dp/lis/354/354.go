package main

import (
	"sort"
)

//给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
//
// 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
//
// 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
//
// 注意：不允许旋转信封。
//
// 示例 1：
//
//
//输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
//输出：3
//解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
//
// 示例 2：
//
//
//输入：envelopes = [[1,1],[1,1],[1,1]]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= envelopes.length <= 10⁵
// envelopes[i].length == 2
// 1 <= wi, hi <= 10⁵
//
//
// Related Topics 数组 二分查找 动态规划 排序 👍 1002 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	hMap := make(map[int]bool)

	b := new(bitSet)
	maxH := 0
	for _, v := range envelopes {
		maxH = max(maxH, v[1])
		hMap[v[1]] = true
	}

	distH := make([]int, 0)

	for k := range hMap {
		distH = append(distH, k)
	}
	sort.Ints(distH)
	rkMap := make(map[int]int)
	for i, v := range distH {
		rkMap[v] = i + 1
	}
	b.sum = make([]int, maxH+1)
	b.len = maxH + 1
	ans := 0
	for _, v := range envelopes {
		rk := rkMap[v[1]]
		cur := b.query(rk - 1)
		ans = max(ans, cur+1)
		b.update(rk, cur+1)
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

func (b *bitSet) update(index, val int) {
	for index < b.len {
		b.sum[index] = max(b.sum[index], val)
		index += lowBit(index)
	}
}

func (b *bitSet) query(index int) int {
	ans := 0
	for index > 0 {
		ans = max(ans, b.sum[index])
		index -= lowBit(index)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
