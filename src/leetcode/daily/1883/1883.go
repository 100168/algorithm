package main

import "math"

/*
给你一个整数 hoursBefore ，表示你要前往会议所剩下的可用小时数。要想成功抵达会议现场，你必须途经 n 条道路。
道路的长度用一个长度为 n 的整数数组 dist 表示，其中 dist[i] 表示第 i 条道路的长度（单位：千米）。
另给你一个整数 speed ，表示你在道路上前进的速度（单位：千米每小时）。

当你通过第 i 条路之后，就必须休息并等待，直到 下一个整数小时 才能开始继续通过下一条道路。
注意：你不需要在通过最后一条道路后休息，因为那时你已经抵达会议现场。

例如，如果你通过一条道路用去 1.4 小时，那你必须停下来等待，到 2 小时才可以继续通过下一条道路。
如果通过一条道路恰好用去 2 小时，就无需等待，可以直接继续。
然而，为了能准时到达，你可以选择 跳过 一些路的休息时间，这意味着你不必等待下一个整数小时。
注意，这意味着与不跳过任何休息时间相比，你可能在不同时刻到达接下来的道路。

例如，假设通过第 1 条道路用去 1.4 小时，且通过第 2 条道路用去 0.6 小时。
跳过第 1 条道路的休息时间意味着你将会在恰好 2 小时完成通过第 2 条道路，且你能够立即开始通过第 3 条道路。
返回准时抵达会议现场所需要的 最小跳过次数 ，如果 无法准时参会 ，返回 -1
*/
func minSkips(dist []int, speed int, hoursBefore int) int {

	n := len(dist)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + dist[i-1]
	}

	//将小数变成整数
	hoursBefore *= speed
	if sum[n] > hoursBefore {
		return -1
	}

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(i, j int) int {
		if j < 0 {
			return math.MaxInt / 2
		}
		if i < 0 {
			return 0
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		//表示跳过
		cur := dfs(i-1, j-1) + dist[i]
		//不跳
		no := dfs(i-1, j) + dist[i]
		//不跳要休息 有可能余数是0，所以最后要取余
		no = (no + speed - 1) / speed * speed
		cur = min(cur, no)
		cache[i][j] = cur
		return cur
	}

	for i := 0; ; i++ {
		cur := dfs(n-2, i)
		if cur+dist[n-1] <= hoursBefore {
			return i
		}
	}
}

func minSkips2(dist []int, speed int, hoursBefore int) int {

	n := len(dist)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + dist[i-1]
	}

	//将小数变成整数
	hoursBefore *= speed
	if sum[n] > hoursBefore {
		return -1
	}

	var cache [][]int

	var dfs func(int, int, int) int

	dfs = func(i int, j, maxSkip int) int {
		if j > maxSkip {
			return math.MaxInt / 2
		}
		if i < 0 {
			return 0
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		//表示跳过
		cur := dfs(i-1, j+1, maxSkip) + dist[i]
		//不跳
		no := dfs(i-1, j, maxSkip) + dist[i]
		//不跳要休息 有可能余数是0，所以最后要取余
		no = (speed-no%speed)%speed + no
		cur = min(cur, no)
		cache[i][j] = cur
		return cur
	}

	check := func(m int) bool {
		cache = make([][]int, n)
		for i := range cache {
			cache[i] = make([]int, n)
			for j := range cache[i] {
				cache[i][j] = -1
			}
		}
		cur := dfs(n-2, 0, m)
		if cur+dist[n-1] <= hoursBefore {
			return true
		}
		return false

	}
	l, r := 0, n-1

	for l <= r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l

}

func main() {

	println(minSkips([]int{38, 28, 8, 75, 95, 70, 8, 96, 41, 8, 7, 75, 62, 65, 68, 21, 8, 66, 11, 24, 9, 77, 9, 87, 31, 52, 16, 73, 32, 75, 77, 6, 80, 11, 54, 85, 75, 73, 67, 41, 34, 27, 86, 92, 41, 31, 34, 51, 17, 86, 83, 39, 59, 41, 97, 10, 2, 59, 80, 73, 13, 10}, 30, 105))
}
