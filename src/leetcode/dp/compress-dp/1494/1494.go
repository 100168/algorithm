package main

import (
	"math"
	"math/bits"
)

/*
*
给你一个整数 n 表示某所大学里课程的数目，编号为 1 到 n ，数组 relations 中，
relations[i] = [xi, yi]  表示一个先修课的关系，也就是课程 xi 必须在课程 yi 之前上。同时你还有一个整数 k 。

在一个学期中，你 最多 可以同时上 k 门课，前提是这些课的先修课在之前的学期里已经上过了。

请你返回上完所有课最少需要多少个学期。题目保证一定存在一种上完所有课的方式。
*/
func minNumberOfSemesters(n int, relations [][]int, k int) int {

	g := make([]int, n)
	for _, v := range relations {
		x, y := v[0]-1, v[1]-1
		//表示y必须先修的课程
		g[y] |= 1 << x
	}

	pre := make([]int, 1<<n)

	//计算所有必须先修课程的并集
	for i := 0; i < n; i++ {
		for j, k := 0, 1<<i; j < k; j++ {
			pre[j|k] = pre[j] | g[i]
		}
	}
	f := make([]int, 1<<n)

	for i := range f {
		f[i] = math.MaxInt / 2
	}

	f[0] = 0
	for s := range f {
		for sub := s; sub > 0; sub = (sub - 1) & s {
			//前置课程必须先修完并且不能同时修前置课程跟当前课程
			if pre[sub]&s != pre[sub] || sub&pre[sub] != 0 {
				continue
			}
			if bits.OnesCount(uint(sub)) <= k {
				f[s] = min(f[s], f[s^sub]+1)
			}
		}
	}

	return f[1<<n-1]

}

func main() {
	println(minNumberOfSemesters(15, [][]int{{9, 6}, {9, 5}, {4, 15}, {7, 5}, {13, 5}, {1, 10}, {13, 6}, {14, 3}, {2, 13}, {14, 10}, {8, 1}, {8, 9}, {14, 6}, {14, 1}, {4, 12}, {9, 10}, {4, 6}, {2, 6}, {2, 14}, {15, 9}, {13, 9}, {12, 9}, {11, 5}, {8, 3}, {4, 9}, {13, 10}, {6, 10}, {11, 1}, {13, 7}, {2, 15}, {12, 10}, {4, 8}, {15, 1}, {7, 10}, {14, 13}, {4, 11}, {7, 9}, {1, 3}, {15, 10}, {2, 11}, {2, 7}, {4, 1}, {4, 3}, {11, 3}, {11, 9}}, 1))
}
