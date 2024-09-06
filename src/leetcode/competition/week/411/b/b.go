package main

import "fmt"

/*
*
来自未来的体育科学家给你两个整数数组 energyDrinkA 和 energyDrinkB，数组长度都等于 n。
这两个数组分别代表 A、B 两种不同能量饮料每小时所能提供的强化能量。

你需要每小时饮用一种能量饮料来 最大化 你的总强化能量。然而，如果从一种能量饮料切换到另一种，
你需要等待一小时来梳理身体的能量体系（在那个小时里你将不会获得任何强化能量）。

返回在接下来的 n 小时内你能获得的 最大 总强化能量。

注意 你可以选择从饮用任意一种能量饮料开始。

示例 1：

输入：energyDrinkA = [1,3,1], energyDrinkB = [3,1,1]

输出：5

解释：

要想获得 5 点强化能量，需要选择只饮用能量饮料 A（或者只饮用 B）。
*/
func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {

	n := len(energyDrinkA)

	fa := make([]int, n+1)
	fb := make([]int, n+1)
	fa[1] = energyDrinkA[0]
	fb[1] = energyDrinkB[0]
	for i := 1; i < n; i++ {
		fa[i+1] = max(fa[i], fb[i-1]) + energyDrinkA[i]
		fb[i+1] = max(fb[i], fa[i-1]) + energyDrinkB[i]
	}
	return int64(max(fa[n], fb[n]))
}

func main() {
	fmt.Println(maxEnergyBoost([]int{1, 3, 1}, []int{3, 1, 1}))
}
