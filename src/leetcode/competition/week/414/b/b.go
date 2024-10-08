package main

import (
	"fmt"
	"math"
	"sort"
)

/**

给你一个整数数组 start 和一个整数 d，代表 n 个区间 [start[i], start[i] + d]。

你需要选择 n 个整数，其中第 i 个整数必须属于第 i 个区间。所选整数的 得分 定义为所选整数两两之间的 最小 绝对差。

返回所选整数的 最大可能得分 。



示例 1：

输入： start = [6,0,3], d = 2

输出： 4

解释：

可以选择整数 8, 0 和 4 获得最大可能得分，得分为 min(|8 - 0|, |8 - 4|, |0 - 4|)，等于 4。

示例 2：

输入： start = [2,6,13,13], d = 5

输出： 5

解释：

可以选择整数 2, 7, 13 和 18 获得最大可能得分，得分为 min(|2 - 7|, |2 - 13|, |2 - 18|, |7 - 13|, |7 - 18|, |13 - 18|)，等于 5。
*/

func maxPossibleScore(start []int, d int) int {

	sort.Ints(start)

	l, r := 0, math.MaxInt

	check := func(t int) bool {

		pre := start[0]
		for _, v := range start[1:] {
			minV := max(v, pre+t)
			if minV > v+d {
				return false
			}
			pre = minV
		}
		return true
	}
	for l <= r {

		m := (r-l)/2 + l
		if check(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return r

}

func main() {
	fmt.Println(maxPossibleScore([]int{6, 0, 3}, 2))
}
