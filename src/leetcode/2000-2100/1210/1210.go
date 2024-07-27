package main

import "fmt"

/*
*
你还记得那条风靡全球的贪吃蛇吗？

我们在一个 n*n 的网格上构建了新的迷宫地图，蛇的长度为 2，也就是说它会占去两个单元格。蛇会从左上角（(0, 0) 和 (0, 1)）开始移动。我们用 0 表示空单元格，用 1 表示障碍物。蛇需要移动到迷宫的右下角（(n-1, n-2) 和 (n-1, n-1)）。

每次移动，蛇可以这样走：

如果没有障碍，则向右移动一个单元格。并仍然保持身体的水平／竖直状态。
如果没有障碍，则向下移动一个单元格。并仍然保持身体的水平／竖直状态。
如果它处于水平状态并且其下面的两个单元都是空的，就顺时针旋转 90 度。蛇从（(r, c)、(r, c+1)）移动到 （(r, c)、(r+1, c)）。

如果它处于竖直状态并且其右面的两个单元都是空的，就逆时针旋转 90 度。蛇从（(r, c)、(r+1, c)）移动到（(r, c)、(r, c+1)）。

返回蛇抵达目的地所需的最少移动次数。
*/
func minimumMoves(grid [][]int) int {

	n := len(grid)

	type pair struct {
		tail int
		head int
	}

	st := []pair{{0, 1}}

	visited := make(map[pair]bool)
	visited[st[0]] = true

	end := pair{(n-1)*n + n - 2, (n-1)*n + n - 1}
	ans := 0
	for len(st) > 0 {

		temp := st
		st = nil

		for _, v := range temp {

			if v == end {
				return ans
			}

			tail, head := v.tail, v.head

			x1 := tail / n
			x2 := head / n

			tailY := tail % n
			headY := head % n

			if x1+1 < n && x2+1 < n {
				next := pair{tail + n, head + n}
				if !visited[next] && grid[x1+1][tailY] == 0 && grid[x2+1][headY] == 0 {
					visited[next] = true
					st = append(st, next)
				}
			}

			if tailY+1 < n && headY+1 < n {
				next := pair{tail + 1, head + 1}
				if !visited[next] && grid[x1][tailY+1] == 0 && grid[x2][headY+1] == 0 {
					visited[next] = true
					st = append(st, next)
				}
			}

			if x1 == x2 && x1+1 < n {
				downX1 := x1 + 1
				downY1 := tailY
				downX2 := x2 + 1
				downY2 := headY
				next := pair{tail, tail + n}
				if grid[downX1][downY1] == 0 && grid[downX2][downY2] == 0 && !visited[next] {
					visited[next] = true
					st = append(st, next)
				}

			}
			if x1 != x2 && tailY+1 < n {
				downX1 := x1
				downY1 := tailY + 1
				downX2 := x2
				downY2 := tailY + 1
				next := pair{tail, tail + 1}
				if grid[downX1][downY1] == 0 && grid[downX2][downY2] == 0 && !visited[next] {
					visited[next] = true
					st = append(st, next)
				}
			}

		}
		ans++

	}
	return -1

}

func main() {
	fmt.Println(minimumMoves([][]int{{0, 0, 0, 0, 0, 1}, {1, 1, 0, 0, 1, 0}, {0, 0, 0, 0, 1, 1}, {0, 0, 1, 0, 1, 0}, {0, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 0, 0}}))
}
