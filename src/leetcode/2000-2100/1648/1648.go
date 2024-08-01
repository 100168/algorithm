package main

import (
	"fmt"
	"slices"
	"sort"
)

/**
你有一些球的库存 inventory ，里面包含着不同颜色的球。一个顾客想要 任意颜色 总数为 orders 的球。

这位顾客有一种特殊的方式衡量球的价值：每个球的价值是目前剩下的 同色球 的数目。比方说还剩下 6 个黄球，
那么顾客买第一个黄球的时候该黄球的价值为 6 。这笔交易以后，只剩下 5 个黄球了，
所以下一个黄球的价值为 5 （也就是球的价值随着顾客购买同色球是递减的）

给你整数数组 inventory ，其中 inventory[i] 表示第 i 种颜色球一开始的数目。同时给你整数 orders ，
表示顾客总共想买的球数目。你可以按照 任意顺序 卖球。

请你返回卖了 orders 个球以后 最大 总价值之和。由于答案可能会很大，请你返回答案对 109 + 7 取余数 的结果。


*/

func maxProfit(inventory []int, orders int) int {

	mod := int(1e9 + 7)

	n := len(inventory)
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i] < inventory[j]
	})
	check := func(t int) bool {
		cnt := 0
		for _, v := range inventory {
			cnt += max(0, v-t)
		}
		return cnt >= orders
	}
	l, r := 0, slices.Max(inventory)

	for l <= r {
		m := (r + l) / 2
		if check(m) {
			l = m + 1
		} else {
			r = m - 1
		}

	}

	firstBig := 0
	for i, v := range inventory {
		if v > r {
			firstBig = i
			break
		}
	}

	ans := 0
	for i := firstBig + 1; i < n; i++ {
		diff := inventory[i] - inventory[firstBig]
		ans += inventory[firstBig]*diff + (1+diff)*diff/2
		ans %= mod
		orders -= diff
	}

	restL := n - firstBig

	mul := orders / restL
	rest := orders % restL

	ans += (inventory[firstBig]-mul+1+inventory[firstBig])*mul/2*(restL) + (inventory[firstBig]-mul)*rest
	ans %= mod
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{2, 8, 4, 10, 6}, 20))
}
