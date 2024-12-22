package main

import (
	"fmt"
	"slices"
)

/*
*
有一个 m x n 大小的矩形蛋糕，需要切成 1 x 1 的小块。

给你整数 m ，n 和两个数组：

horizontalCut 的大小为 m - 1 ，其中 horizontalCut[i] 表示沿着水平线 i 切蛋糕的开销。
verticalCut 的大小为 n - 1 ，其中 verticalCut[j] 表示沿着垂直线 j 切蛋糕的开销。
一次操作中，你可以选择任意不是 1 x 1 大小的矩形蛋糕并执行以下操作之一：

沿着水平线 i 切开蛋糕，开销为 horizontalCut[i] 。
沿着垂直线 j 切开蛋糕，开销为 verticalCut[j] 。
每次操作后，这块蛋糕都被切成两个独立的小蛋糕。

每次操作的开销都为最开始对应切割线的开销，并且不会改变。

请你返回将蛋糕全部切成 1 x 1 的蛋糕块的 最小 总开销。

输入：m = 3, n = 2, horizontalCut = [1,3], verticalCut = [5]

输出：13

思路：贪心
*/
func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {

	slices.SortFunc(horizontalCut, func(a, b int) int { return b - a })
	slices.SortFunc(verticalCut, func(a, b int) int { return b - a })
	ans := 0
	cntH, cntV := 1, 1
	i, j := 0, 0
	for i < m-1 || j < n-1 {
		if j == n-1 || i < m-1 && horizontalCut[i] > verticalCut[j] {
			ans += horizontalCut[i] * cntH // 横切
			i++
			cntV++ // 需要竖切的蛋糕块增加
		} else {
			ans += verticalCut[j] * cntV // 竖切
			j++
			cntH++ // 需要横切的蛋糕块增加
		}
	}
	return ans
}

func main() {
	//fmt.Println(minimumCost(3, 2, []int{1, 3}, []int{5}))
	fmt.Println(minimumCost(2, 2, []int{7}, []int{4}))
}
