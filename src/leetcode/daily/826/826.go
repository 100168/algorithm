package main

import (
	"sort"
)

/*
*你有 n 个工作和 m 个工人。给定三个数组： difficulty, profit 和 worker ，其中:

difficulty[i] 表示第 i 个工作的难度，profit[i] 表示第 i 个工作的收益。
worker[i] 是第 i 个工人的能力，即该工人只能完成难度小于等于 worker[i] 的工作。
每个工人 最多 只能安排 一个 工作，但是一个工作可以 完成多次 。

举个例子，如果 3 个工人都尝试完成一份报酬为 $1 的同样工作，那么总收益为 $3 。如果一个工人不能完成任何工作，他的收益为 $0 。
返回 在把工人分配到工作岗位后，我们所能获得的最大利润 。

可线性，可树状数组
*/
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	type pair struct {
		d int
		p int
	}
	var pairs []pair
	for i, d := range difficulty {
		p := profit[i]
		pairs = append(pairs, pair{d, p})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].d == pairs[j].d {
			return pairs[i].p > pairs[j].p
		}
		return pairs[i].d < pairs[j].d
	})
	rkMap := make(map[int]int)

	j := 0

	for i := 0; i < len(pairs); i++ {
		if pairs[i].d != pairs[j].d {
			pairs[j+1] = pairs[i]
			j++
		}
	}
	pairs = pairs[:j+1]
	bt := new(bitTree)
	bt.sum = make([]int, len(pairs)+1)
	bt.n = len(pairs) + 1
	for i, v := range pairs {
		rkMap[v.d] = i
		bt.update(i+1, v.p)
	}

	ans := 0

	for _, v := range worker {
		l, r := 0, len(pairs)-1

		for l <= r {
			m := (r + l) / 2
			if pairs[m].d > v {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		if r == -1 {
			continue
		}
		ans += bt.query(rkMap[pairs[r].d] + 1)

	}
	return ans
}

type bitTree struct {
	sum []int
	n   int
}

func lowBit(index int) int {
	return index & -index
}
func (bt bitTree) query(index int) int {
	ans := 0

	for index > 0 {
		ans = max(ans, bt.sum[index])
		index -= lowBit(index)
	}
	return ans
}

func (bt bitTree) update(index int, val int) {
	for index < bt.n {
		bt.sum[index] = max(bt.sum[index], val)
		index += lowBit(index)
	}
}

func main() {
	println(maxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}))
}
