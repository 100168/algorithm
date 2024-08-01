package main

import "fmt"

func canReachCorner(X int, Y int, circles [][]int) bool {

	n := len(circles)
	g := make([][]int, n)

	visited := make([]bool, n)
	for i, c := range circles {

		x1, y1, r1 := c[0], c[1], c[2]

		x1Min := x1 - r1
		x1Max := x1 + r1
		y1Min := y1 - r1
		y1Max := y1 + r1
		if abs(x1-X)*abs(x1-X)+abs(y1-Y)*abs(y1-Y) <= r1*r1 {
			return false
		}

		if abs(x1-0)*abs(x1-0)+abs(y1-0)*abs(y1-0) <= r1*r1 {
			return false
		}
		if y1Max < 0 || y1Min > Y || x1Min > X || x1Max < 0 {
			continue
		}
		for j, v := range circles {
			if i == j {
				continue
			}
			x2, y2, r2 := v[0], v[1], v[2]
			x2Min := x2 - r2
			x2Max := x2 + r2
			y2Min := y2 - r2
			y2Max := y2 + r2
			if y2Max < 0 || y2Min > Y || x2Min > X || x2Max < 0 {
				continue
			}
			dis := abs(x1-x2)*abs(x1-x2) + abs(y1-y2)*abs(y1-y2)
			if (r1+r2)*(r1+r2) >= dis {
				g[i] = append(g[i], j)
			}
		}
	}

	var dfs func(x int) (int, int, int, int)

	dfs = func(x int) (int, int, int, int) {
		visited[x] = true
		x1, y1, r1 := circles[x][0], circles[x][1], circles[x][2]
		x1Min := x1 - r1
		x1Max := x1 + r1
		y1Min := y1 - r1
		y1Max := y1 + r1
		for _, y := range g[x] {
			if visited[y] {
				continue
			}
			x2Min, x2Max, y2Min, y2Max := dfs(y)
			x1Min = min(x2Min, x1Min)
			x1Max = max(x1Max, x2Max)
			y1Min = min(y2Min, y1Min)
			y1Max = max(y1Max, y2Max)
		}
		return x1Min, x1Max, y1Min, y1Max
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			xMin, xMax, yMin, yMax := dfs(i)
			if xMin <= 0 && xMax >= X || yMin <= 0 && yMax >= Y {
				return false
			}
			if xMin <= 0 && xMax >= 0 && yMin <= 0 && yMax >= 0 {
				return false
			}

			if xMin <= X && xMax >= X && yMin <= Y && yMax >= Y {
				return false
			}
		}
	}
	return true
}
func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(canReachCorner(3, 3, [][]int{{2, 1000, 997}, {1000, 2, 997}}))
}
