package main

import "fmt"

/*
*
给你一棵 n 个节点的树（一个无向、连通、无环图），每个节点表示一个城市，编号从 0 到 n - 1 ，
且恰好有 n - 1 条路。0 是首都。给你一个二维整数数组 roads ，其中 roads[i] = [ai, bi] ，表示城市 ai 和 bi 之间有一条 双向路 。

每个城市里有一个代表，他们都要去首都参加一个会议。

每座城市里有一辆车。给你一个整数 seats 表示每辆车里面座位的数目。

城市里的代表可以选择乘坐所在城市的车，或者乘坐其他城市的车。相邻城市之间一辆车的油耗是一升汽油。

请你返回到达首都最少需要多少升汽油。

思路：树形dp
1.只需要考虑到达每个地区有多少个人，然后从该地区派多少量车
*/
func minimumFuelCost(roads [][]int, seats int) int64 {

	n := len(roads) + 1
	g := make([][]int, n)
	for _, v := range roads {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(x, fa int) (int, int, int)
	dfs = func(x, fa int) (int, int, int) {

		curCost := 0
		curCars := 0
		curSeatCost := 0
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			gasCost, cars, seatCost := dfs(y, x)
			curCost += gasCost
			curSeatCost += seatCost
			curCars += cars
		}
		curSeatCost++
		return curCost + curCars, ((curSeatCost - 1) / seats) + 1, curSeatCost
	}
	cost, _, _ := dfs(0, -1)
	return int64(cost)
}

func main() {
	//fmt.Println(minimumFuelCost([][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}}, 5))
	fmt.Println(minimumFuelCost([][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2))
}
