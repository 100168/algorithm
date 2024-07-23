package main

import (
	"sort"
)

/*
假设你是球队的经理。对于即将到来的锦标赛，你想组合一支总体得分最高的球队。球队的得分是球队中所有球员的分数 总和 。

然而，球队中的矛盾会限制球员的发挥，所以必须选出一支 没有矛盾 的球队。如果一名年龄较小球员的分数 严格大于 一名年龄较大的球员，则存在矛盾。同龄球员之间不会发生矛盾。

给你两个列表 scores 和 ages，其中每组 scores[i] 和 ages[i] 表示第 i 名球员的分数和年龄。请你返回 所有可能的无矛盾球队中得分最高那支的分数 。

示例 1：

输入：scores = [1,3,5,10,15], ages = [1,2,3,4,5]
输出：34
解释：你可以选中所有球员。
示例 2：

输入：scores = [4,5,6,5], ages = [2,1,2,1]
输出：16
解释：最佳的选择是后 3 名球员。注意，你可以选中多个同龄球员。
示例 3：

输入：scores = [1,2,3,5], ages = [8,9,10,1]
输出：6
解释：最佳的选择是前 3 名球员。
*/
func bestTeamScore(scores []int, ages []int) int {
	n := len(scores)

	type pair struct{ x, y int }

	pairs := make([]pair, n)
	for i := 0; i < n; i++ {
		pairs[i] = pair{scores[i], ages[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].x == pairs[j].x {
			return pairs[i].y < pairs[j].y
		}
		return pairs[i].x < pairs[j].x
	})

	st := make([]int, 0)
	sum := 0
	ans := 0
	for i, p := range pairs {
		l, r := 0, len(st)-1

		for l <= r {
			m := (l + r) / 2
			if pairs[st[m]].y > p.y {
				r = m - 1
			} else {
				l = m + 1
			}
		}

		if l < len(st) {
			sum -= pairs[st[l]].x

			st[l] = i
			curSum := 0
			for _, v := range st[:l+1] {
				curSum += pairs[v].x
			}
			sum += p.x
			ans = max(ans, curSum)
		} else {
			st = append(st, i)
			sum += p.x
			ans = max(ans, sum)
		}
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
		s += b.sum[index]
		index -= lowBit(index)
	}
	return s
}

func (b *bitSet) update(index int, val int) {
	for b.len > index {
		b.sum[index] += val
		index += lowBit(index)
	}
}
func main() {

	println(bestTeamScore([]int{4, 5, 6, 5}, []int{2, 1, 2, 1}))
}
