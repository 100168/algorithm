package main

import "sort"

func maxRectangleArea(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {

		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] < points[j][1]
	})
	g := make([][]int, 101)
	for _, v := range points {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
	}

	for i, y1 := range g {

		bt := new(bitTree)
		bt.sum = make([]int, 101)
		for h, y2 := range g[i:] {

			p1 := 1
			p2 := 1

			for (p1 < len(y1) && p2 < len(y2)) && (y1[p1-1] != y2[p2-1] || y1[p2] != y1[p1]) {
				if y2[p2] < y1[p1] || y2[p2-1] < y2[p2-1] {
					p2++
					if p2 < len(y2) {
						bt.update(y2[p2])
					}
				} else {
					p1++
				}
			}
		}
	}

}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type bitTree struct {
	sum []int
	n   int
}

func (bt *bitTree) lowBit(index int) int {
	return index & -index

}
func (bt *bitTree) query(index int) int {

	c := 0
	for index > 0 {
		c += bt.sum[index]
		index -= bt.lowBit(index)
	}
	return c
}

func (bt *bitTree) update(index int) {
	for index < bt.n {

		bt.sum[index] += 1
	}

}
