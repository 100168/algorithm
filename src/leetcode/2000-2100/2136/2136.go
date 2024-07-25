package main

import "sort"

/*
*
你有 n 枚花的种子。每枚种子必须先种下，才能开始生长、开花。播种需要时间，种子的生长也是如此。
给你两个下标从 0 开始的整数数组 plantTime 和 growTime ，每个数组的长度都是 n ：

plantTime[i] 是 播种 第 i 枚种子所需的 完整天数 。每天，你只能为播种某一枚种子而劳作。
无须 连续几天都在种同一枚种子，但是种子播种必须在你工作的天数达到 plantTime[i] 之后才算完成。
growTime[i] 是第 i 枚种子完全种下后生长所需的 完整天数 。在它生长的最后一天 之后 ，将会开花并且永远 绽放 。
从第 0 开始，你可以按 任意 顺序播种种子。

返回所有种子都开花的 最早 一天是第几天。

思路
1.贪心
2.为啥贪心是对的，因为总的播种时间是固定的，所以开花越晚播种越早
*/
func earliestFullBloom(plantTime []int, growTime []int) int {

	type pair struct {
		p int
		g int
	}
	n := len(plantTime)

	pairs := make([]pair, n)
	for i := range plantTime {
		pairs[i] = pair{plantTime[i], growTime[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].g > pairs[j].g
	})

	pt := 0
	maxGt := 0
	for _, v := range pairs {
		pt += v.p
		maxGt = max(maxGt, pt+v.g)

	}
	return maxGt
}
