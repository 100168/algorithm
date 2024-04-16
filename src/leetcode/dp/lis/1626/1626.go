package main

import (
	"sort"
)

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
