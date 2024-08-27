package main

import (
	"math"
	"sort"
)

/*
*
X 轴上有一些机器人和工厂。给你一个整数数组 robot ，其中 robot[i] 是第 i 个机器人的位置。
再给你一个二维整数数组 factory ，其中 factory[j] = [positionj, limitj] ，表示第 j 个工厂的位置在 positionj ，
且第 j 个工厂最多可以修理 limitj 个机器人。

每个机器人所在的位置 互不相同 。每个工厂所在的位置也 互不相同 。注意一个机器人可能一开始跟一个工厂在 相同的位置 。

所有机器人一开始都是坏的，他们会沿着设定的方向一直移动。设定的方向要么是 X 轴的正方向，要么是 X 轴的负方向。当一个机器人经过一个没达到上限的工厂时，这个工厂会维修这个机器人，且机器人停止移动。

任何时刻，你都可以设置 部分 机器人的移动方向。你的目标是最小化所有机器人总的移动距离。

请你返回所有机器人移动的最小总距离。测试数据保证所有机器人都可以被维修。

注意：

所有机器人移动速度相同。
如果两个机器人移动方向相同，它们永远不会碰撞。
如果两个机器人迎面相遇，它们也不会碰撞，它们彼此之间会擦肩而过。
如果一个机器人经过了一个已经达到上限的工厂，机器人会当作工厂不存在，继续移动。
机器人从位置 x 到位置 y 的移动距离为 |y - x| 。

思路：
1.先将机器人和工厂按位置排序
2.确定位置后再枚举当前位置可以放多少个机器人
*/
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	m := len(robot)
	n := len(factory)
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	f := make([][]int, n)

	for i := range f {
		f[i] = make([]int, m)

		for j := range f[i] {
			f[i][j] = -1
		}
	}
	var dfs func(int, int) int

	dfs = func(i int, j int) int {

		if j < 0 {
			return 0
		}
		if i < 0 {
			return math.MaxInt / 2
		}

		if f[i][j] != -1 {
			return f[i][j]
		}

		cur := math.MaxInt / 2
		s := 0
		for k := j; k >= 0 && j-k+1 <= factory[i][1]; k-- {
			s += abs(robot[k] - factory[i][0])
			cur = min(cur, dfs(i-1, k-1)+s)
		}
		cur = min(cur, dfs(i-1, j))
		f[i][j] = cur
		return cur

	}
	return int64(dfs(n-1, m-1))

}

func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}
