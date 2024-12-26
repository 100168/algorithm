package main

import "math/bits"

/*
*
给你一个 m * n 的矩阵 seats 表示教室中的座位分布。如果座位是坏的（不可用），就用 '#' 表示；否则，用 '.' 表示。

学生可以看到左侧、右侧、左上、右上这四个方向上紧邻他的学生的答卷，但是看不到直接坐在他前面或者后面的学生的答卷。请你计算并返回该考场可以容纳的同时参加考试且无法作弊的 最大 学生人数。

学生必须坐在状况良好的座位上。

tag：按行状压dp
*/
func maxStudents(seats [][]byte) int {

	n := len(seats[0])

	m := len(seats)
	f := make([][]int, m)

	for i := range f {

		f[i] = make([]int, 1<<n)

		for j := range f[i] {
			f[i][j] = -1
		}
	}

	masks := make([]int, m)

	for i, r := range seats {

		c := 0
		for j, v := range r {

			if v == '.' {
				c |= 1 << j
			}
		}
		masks[i] = c

	}

	var dfs func(int, int) int

	dfs = func(i, pre int) int {

		if i == m {
			return 0
		}

		if f[i][pre] != -1 {
			return f[i][pre]
		}
		mask := masks[i]

		t := pre<<1 | pre>>1
		mask = ((1<<n - 1) ^ t) & mask

		cur := 0
	next:
		for sub := mask; sub > 0; sub = (sub - 1) & mask {

			one := bits.OnesCount(uint(sub))
			for j := 0; j < n; j++ {

				if sub>>j&1 != 0 && sub&(sub>>1) != 0 {
					continue next
				}
			}
			cur = max(cur, one+dfs(i+1, sub))
		}
		cur = max(cur, dfs(i+1, 0))
		f[i][pre] = cur
		return cur

	}

	return dfs(0, 0)

}
