package main

import (
	"slices"
)

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

	minVal := slices.Min(baseCosts)
	if minVal > target {
		return minVal
	}
	n := len(toppingCosts)
	ans := minVal
	var dfs func(int, int)
	dfs = func(index, cost int) {
		if cost < target {
			if abs(ans-target) >= target-cost {
				ans = cost
			}
		} else if cost >= target {
			if cost == target {
				ans = cost
			} else {
				if abs(ans-target) > cost-target {
					ans = cost
				}
			}
			return
		}
		if index == n {
			return
		}
		dfs(index+1, cost)
		dfs(index+1, cost+toppingCosts[index])
		dfs(index+1, cost+toppingCosts[index]*2)

	}
	for _, cost := range baseCosts {
		dfs(0, cost)
	}
	return ans
}

func closestCost2(baseCosts []int, toppingCosts []int, target int) int {

	minVal := slices.Min(baseCosts)
	if minVal > target {
		return minVal
	}
	//abs(y-target)<=abs(target - x)
	//y - target <=target-x ==>y<2*target - x
	n := len(toppingCosts)
	cache := make([][]int, n+1)
	maxValue := 2*target - minVal
	for i := range cache {
		cache[i] = make([]int, maxValue+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	valuesMap := make(map[int]bool)
	for _, cost := range baseCosts {
		valuesMap[cost] = true
	}
	var dfs func(int, int) bool

	ans := minVal
	dfs = func(index, value int) bool {
		if value < 0 {
			return false
		}
		if n == index {
			return valuesMap[value]
		}
		if cache[index][value] != -1 {
			return cache[index][value] == 1
		}
		cur := dfs(index+1, value) || dfs(index+1, value-toppingCosts[index]) || dfs(index+1, value-2*toppingCosts[index])
		if cur {
			cache[index][value] = 1
			if abs(target-value) < abs(ans-target) {
				ans = value
			} else if abs(target-value) == abs(ans-target) {
				ans = min(ans, value)
			}
		} else {
			cache[index][value] = 0
		}
		return cur
	}

	for i := 0; i <= maxValue; i++ {
		dfs(0, i)
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	//println(closestCost([]int{1}, []int{8, 10}, 10))
	println(closestCost2([]int{3, 10}, []int{5}, 9))
}
