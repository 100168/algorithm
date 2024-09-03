package main

import (
	"fmt"
	"sort"
)

/**
有一个无限大的二维平面。

给你一个正整数 k ，同时给你一个二维数组 queries ，包含一系列查询：

queries[i] = [x, y] ：在平面上坐标 (x, y) 处建一个障碍物，数据保证之前的查询 不会 在这个坐标处建立任何障碍物。
每次查询后，你需要找到离原点第 k 近 障碍物到原点的 距离 。

请你返回一个整数数组 results ，其中 results[i] 表示建立第 i 个障碍物以后，离原地第 k 近障碍物距离原点的距离。如果少于 k 个障碍物，results[i] == -1 。

注意，一开始 没有 任何障碍物。

坐标在 (x, y) 处的点距离原点的距离定义为 |x| + |y| 。



示例 1：

输入：queries = [[1,2],[3,4],[2,3],[-3,0]], k = 2

输出：[-1,7,5,3]


*/

func resultsArray(queries [][]int, k int) []int {

	n := len(queries)

	bt := new(bitTree)

	bt.n = n + 1
	bt.sum = make([]int, n+1)

	rk := make([]int, n)

	for i, v := range queries {
		x, y := v[0], v[1]
		rk[i] = abs(x) + abs(y)

	}

	sort.Ints(rk)

	rkMap := make(map[int]int)

	vMap := make(map[int]int)

	for i, v := range rk {
		rkMap[v] = i + 1
		vMap[i+1] = v
	}

	ans := make([]int, len(queries))

	cnt := 0
	for i, v := range queries {
		x, y := v[0], v[1]

		rk := rkMap[abs(x)+abs(y)]
		bt.update(rk, 1)

		cnt++
		if cnt < k {
			ans[i] = -1
			continue
		}
		l, r := 1, n
		for l <= r {
			m := (r + l) / 2

			if bt.query(m) >= k {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		ans[i] = vMap[l]
	}

	return ans

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

func (b bitTree) query(index int) int {
	ans := 0
	for index > 0 {
		ans += b.sum[index]
		index -= index & -index
	}
	return ans
}

func (b bitTree) update(index int, value int) {
	for index < b.n {
		b.sum[index] += value
		index += index & -index
	}
}

func main() {

	fmt.Println(resultsArray([][]int{{-6, 4}, {7, 8}, {-2, -1}, {1, -9}, {-9, 4}}, 1))
}
