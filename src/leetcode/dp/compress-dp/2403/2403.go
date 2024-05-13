package main

import (
	"math"
	"math/bits"
)

/*
*
你有一个整数数组 power，其中  power[i] 是第 i 个怪物的力量。

你从 0 点法力值开始，每天获取 gain 点法力值，最初 gain 等于 1。

每天，在获得 gain 点法力值后，如果你的法力值大于或等于怪物的力量，你就可以打败怪物。当你打败怪物时:

你的法力值会被重置为 0，并且

gain 的值增加 1。

返回打败所有怪物所需的 最少 天数。
*/
func minimumTime(power []int) int64 {

	n := len(power)
	memo := make([]int, 1<<n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(mask int) int {

		gain := bits.OnesCount(uint(mask))

		if gain == n {
			return 0
		}

		if memo[mask] != -1 {
			return memo[mask]
		}

		cur := math.MaxInt

		for j := 0; j < n; j++ {
			if 1<<j&mask == 0 {
				cur = min(cur, dfs(1<<j|mask)+getDay(power[j], gain+1))
			}
		}
		memo[mask] = cur
		return cur
	}
	return int64(dfs(0))

}

// 向上取整
func getDay(pow int, gain int) int {
	return (pow-1)/gain + 1
}

func main() {

	println(minimumTime([]int{3}))
	println(getDay(1, 1))
}
