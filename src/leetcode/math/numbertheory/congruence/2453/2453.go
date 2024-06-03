package main

import (
	"fmt"
	"math"
)

/*
*
给你一个下标从 0 开始的数组 nums ，它包含若干正整数，
表示数轴上你需要摧毁的目标所在的位置。同时给你一个整数 space 。

你有一台机器可以摧毁目标。给机器 输入 nums[i] ，
这台机器会摧毁所有位置在 nums[i] + c * space 的目标，其中 c 是任意非负整数。
你想摧毁 nums 中 尽可能多 的目标。

请你返回在摧毁数目最多的前提下，nums[i] 的 最小值 。
*/
func destroyTargets(nums []int, space int) int {

	reminderMap := make(map[int]int)
	maxCnt := 0
	for _, v := range nums {
		reminderMap[v%space]++
		maxCnt = max(maxCnt, reminderMap[v%space])
	}
	ans := math.MaxInt
	for _, v := range nums {
		if reminderMap[v%space] == maxCnt {
			ans = min(ans, v)
		}
	}
	return ans
}

func main() {

	fmt.Println(destroyTargets([]int{3, 7, 8, 1, 1, 5}, 2))
}
