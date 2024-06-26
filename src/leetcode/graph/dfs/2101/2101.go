package main

import "fmt"

/*
*
给你一个炸弹列表。一个炸弹的 爆炸范围 定义为以炸弹为圆心的一个圆。

炸弹用一个下标从 0 开始的二维整数数组 bombs 表示，其中 bombs[i] = [xi, yi, ri] 。
xi 和 yi 表示第 i 个炸弹的 X 和 Y 坐标，ri 表示爆炸范围的 半径 。

你需要选择引爆 一个 炸弹。当这个炸弹被引爆时，所有 在它爆炸范围内的炸弹都会被引爆，这些炸弹会进一步将它们爆炸范围内的其他炸弹引爆。

给你数组 bombs ，请你返回在引爆 一个 炸弹的前提下，最多 能引爆的炸弹数目。
*/
func maximumDetonation(bombs [][]int) int {

	n := len(bombs)
	g := make([][]int, n)

	for i := range bombs {
		x1, y1, r1 := bombs[i][0], bombs[i][1], bombs[i][2]
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			x2, y2, _ := bombs[j][0], bombs[j][1], bombs[j][2]
			dis := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
			if dis <= r1*r1 {
				g[i] = append(g[i], j)
			}
		}
	}

	visited := make([]bool, n)

	cnt := 0
	var dfs func(int)

	dfs = func(x int) {
		cnt++
		visited[x] = true
		for _, y := range g[x] {
			if !visited[y] {
				dfs(y)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		visited = make([]bool, n)
		cnt = 0
		dfs(i)
		ans = max(ans, cnt)
	}
	return ans
}

func main() {
	fmt.Println(maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}}))
}
