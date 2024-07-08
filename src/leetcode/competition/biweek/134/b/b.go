package main

import "sort"

/*
*
给你一个下标从 0 开始的整数数组 enemyEnergies ，它表示一个下标从 0 开始的敌人能量数组。

同时给你一个整数 currentEnergy ，它表示你一开始拥有的能量值总量。

你一开始的分数为 0 ，且一开始所有的敌人都未标记。

你可以通过以下操作 之一 任意次（也可以 0 次）来得分：

选择一个 未标记 且满足 currentEnergy >= enemyEnergies[i] 的敌人 i 。在这个操作中：
你会获得 1 分。
你的能量值减少 enemyEnergies[i] ，也就是说 currentEnergy = currentEnergy - enemyEnergies[i] 。
如果你目前 至少 有 1 分，你可以选择一个 未标记 的敌人 i 。在这个操作中：
你的能量值增加 enemyEnergies[i] ，也就是说 currentEnergy = currentEnergy + enemyEnergies[i] 。
敌人 i 被标记 。
请你返回通过以上操作，最多 可以获得多少分。
*/
func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {

	sort.Ints(enemyEnergies)

	if currentEnergy < enemyEnergies[0] {
		return 0
	}
	ans := 0

	s := 0

	for i := 1; i < len(enemyEnergies); i++ {
		s += enemyEnergies[i]
	}
	return int64((s + currentEnergy) / enemyEnergies[0])
}
