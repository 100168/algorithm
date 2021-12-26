package main

import "fmt"

/*你被给定一个m × n的二维网格 rooms ，网格中有以下三种可能的初始化值：

-1表示墙或是障碍物
0表示一扇门
INF无限表示一个空的房间。然后，我们用231 - 1 = 2147483647代表INF。你可以认为通往门的距离总是小于2147483647的。
你要给每个空房间位上填上该房间到最近门的距离 ，如果无法到达门，则填INF即可

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/walls-and-gates
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func wallsAndGates(rooms [][]int) {

	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	m := len(rooms)
	n := len(rooms[0])

	gate := 0
	inf := 1<<31 - 1
	var q [][]int
	//将所有门找出并放入队列里
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rooms[i][j] == gate {
				q = append(q, []int{i, j})
			}

		}
	}

	for len(q) > 0 {
		var x []int
		x, q = q[0], q[1:]
		row := x[0]
		col := x[1]
		for _, v := range directions {
			r := v[0] + row
			c := v[1] + col
			// 第一次被修改时就是最近的门的距离，之后因为不是INF了，就不会被后来的修改
			if r < 0 || c < 0 || r >= m || c >= n || rooms[r][c] != inf {
				continue
			}

			rooms[r][c] = rooms[row][col] + 1
			q = append(q, []int{r, c})
		}
	}

}

func main() {
	fmt.Println(1<<31 - 1)
}
