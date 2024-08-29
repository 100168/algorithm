package main

import (
	"fmt"
	"sort"
)

/*
*
给你一个二维整数数组 intervals ，其中 intervals[i] = [starti, endi] 表示从 starti 到 endi 的所有整数，包括 starti 和 endi 。

包含集合 是一个名为 nums 的数组，并满足 intervals 中的每个区间都 至少 有 两个 整数在 nums 中。

例如，如果 intervals = [[1,3], [3,7], [8,9]] ，那么 [1,2,4,7,8,9] 和 [2,3,4,8,9] 都符合 包含集合 的定义。
返回包含集合可能的最小大小。

示例 1：

输入：intervals = [[1,3],[3,7],[8,9]]
输出：5
解释：nums = [2, 3, 4, 8, 9].
可以证明不存在元素数量为 4 的包含集合。
*/
func intersectionSizeTwo(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	cnt := 2
	mostR := intervals[0][1]

	//
	for _, v := range intervals[1:] {
		s, e := v[0], v[1]
		if s >= mostR {
			cnt += 2
			mostR = e
		} else {
			mostR = min(mostR, e)
		}
	}
	return cnt

}

func main() {
	//fmt.Println(intersectionSizeTwo([][]int{{1, 3}, {3, 7}, {8, 9}}))
	fmt.Println(intersectionSizeTwo([][]int{{1, 3}, {3, 7}, {5, 7}, {7, 8}}))
}
