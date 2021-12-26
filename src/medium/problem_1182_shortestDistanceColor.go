package main

import (
	"fmt"
	"math"
)

/*给你一个数组colors，里面有1、2、3 三种颜色。

我们需要在colors 上进行一些查询操作 queries，其中每个待查项都由两个整数 i 和 c 组成。

现在请你帮忙设计一个算法，查找从索引i到具有目标颜色c的元素之间的最短距离。

如果不存在解决方案，请返回-1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-distance-to-target-color
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func shortestDistanceColor(colors []int, queries [][]int) []int {

	l := len(colors)
	color1l := make([]int, l)
	color2l := make([]int, l)
	color3l := make([]int, l)
	color1r := make([]int, l)
	color2r := make([]int, l)
	color3r := make([]int, l)

	max := math.MaxInt
	for i := 0; i < l; i++ {
		color1l[i] = max
		color2l[i] = max
		color3l[i] = max

		if colors[i] == 1 {
			color1l[i] = 0
		}
		if colors[i] == 2 {
			color2l[i] = 0
		}
		if colors[i] == 3 {
			color3l[i] = 0
		}
		if i >= 1 {
			color1l[i] = min(color1l[i-1]+1, color1l[i])
			color2l[i] = min(color2l[i-1]+1, color2l[i])
			color3l[i] = min(color3l[i-1]+1, color3l[i])
		}
	}
	for i := l - 1; i >= 0; i-- {
		color1r[i] = max
		color2r[i] = max
		color3r[i] = max

		if colors[i] == 1 {
			color1r[i] = 0
		}
		if colors[i] == 2 {
			color2r[i] = 0
		}
		if colors[i] == 3 {
			color3r[i] = 0
		}
		if i < l-1 {
			color1r[i] = min(color1r[i+1]+1, color1r[i])
			color2r[i] = min(color2r[i+1]+1, color2r[i])
			color3r[i] = min(color3r[i+1]+1, color3r[i])
		}
	}

	ans := make([]int, len(queries))

	for i, v := range queries {
		if v[1] == 1 {
			ans[i] = min(color1l[v[0]], color1r[v[0]])

		}

		if v[1] == 2 {
			ans[i] = min(color2l[v[0]], color2r[v[0]])
		}
		if v[1] == 3 {
			ans[i] = min(color3l[v[0]], color3r[v[0]])
		}

		if ans[i] == max {
			ans[i] = -1
		}

	}

	return ans
}

func min(a int, b int) int {

	if a <= b {
		return a
	} else {
		return b
	}
}
func main() {
	fmt.Println(5*10 ^ 4 + 1)
}
