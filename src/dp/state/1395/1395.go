package main

import "sort"

func numTeams(rating []int) int {

	m := len(rating)
	ans := 0

	dpMax := make([]int, m)
	dpMin := make([]int, m)

	sa := make([]int, m)
	for i := range sa {
		sa[i] = i
	}
	sort.Slice(sa, func(i, j int) bool {
		return rating[i] < rating[j]
	})

	rk := make([]int, m)
	for i := range sa {
		rk[sa[i]] = i
	}

	b := new(bs)
	b.sum = make([]int, m+1)
	b.len = m + 1

	for i := m - 1; i >= 0; i-- {
		maxR := b.query(rk[i])
		minR := b.query(m) - maxR
		dpMax[i] = maxR
		dpMin[i] = minR
		b.update(rk[i] + 1)
	}

	clear(b.sum)

	for i := 0; i < m; i++ {
		maxL := b.query(rk[i])
		minL := b.query(m) - maxL
		b.update(rk[i] + 1)
		ans += maxL*dpMin[i] + minL*dpMax[i]
	}
	return ans
}

type bs struct {
	sum []int
	len int
}

func lowBit(index int) int {
	return index & -index
}

func (b bs) query(index int) int {

	s := 0
	for index > 0 {
		s += b.sum[index]
		index -= lowBit(index)
	}

	return s
}

func (b bs) update(index int) {
	for index < b.len {
		b.sum[index] += 1
		index += lowBit(index)
	}
}
