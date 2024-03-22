package main

//1774. 最接近目标价格的甜点成本
//中等
//相关标签
//相关企业
//提示
//你打算做甜点，现在需要购买配料。目前共有 n 种冰激凌基料和 m 种配料可供选购。而制作甜点需要遵循以下几条规则：
//
//必须选择 一种 冰激凌基料。
//可以添加 一种或多种 配料，也可以不添加任何配料。
//每种类型的配料 最多两份 。
//给你以下三个输入：
//
//baseCosts ，一个长度为 n 的整数数组，其中每个 baseCosts[i] 表示第 i 种冰激凌基料的价格。
//toppingCosts，一个长度为 m 的整数数组，其中每个 toppingCosts[i] 表示 一份 第 i 种冰激凌配料的价格。
//target ，一个整数，表示你制作甜点的目标价格。
//你希望自己做的甜点总成本尽可能接近目标价格 target 。
//
//返回最接近 target 的甜点成本。如果有多种方案，返回 成本相对较低 的一种。

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	x := baseCosts[0]
	for _, c := range baseCosts {
		x = min(x, c)
	}
	if x > target {
		return x
	}
	can := make([]bool, target+1)
	ans := 2*target - x
	for _, c := range baseCosts {
		if c <= target {
			can[c] = true
		} else {
			ans = min(ans, c)
		}
	}
	for _, c := range toppingCosts {
		for count := 0; count < 2; count++ {
			for i := target; i > 0; i-- {
				if can[i] && i+c > target {
					ans = min(ans, i+c)
				}
				if i-c > 0 {
					can[i] = can[i] || can[i-c]
				}
			}
		}
	}
	for i := 0; i <= ans-target; i++ {
		if can[target-i] {
			return target - i
		}
	}
	return ans
}
