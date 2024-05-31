package main

/*
*
给你一个下标从 0 开始的二维整数矩阵 grid，大小为 n * n ，其中的值在 [1, n2] 范围内。
除了 a 出现 两次，b 缺失 之外，每个整数都 恰好出现一次 。

任务是找出重复的数字a 和缺失的数字 b 。

返回一个下标从 0 开始、长度为 2 的整数数组 ans ，其中 ans[0] 等于 a ，ans[1] 等于 b 。

示例 1：

输入：grid = [[1,3],[2,2]]
输出：[2,4]
解释：数字 2 重复，数字 4 缺失，所以答案是 [2,4] 。
示例 2：

输入：grid = [[9,1,7],[8,9,2],[3,4,6]]
输出：[9,5]
解释：数字 9 重复，数字 5 缺失，所以答案是 [9,5] 。
*/
func findMissingAndRepeatedValues(grid [][]int) []int {

	axb := 0
	x := 1
	for i := range grid {
		for _, v := range grid[i] {
			axb ^= v ^ x
			x++
		}
	}
	lowBit := axb & -axb
	x = 1
	a := 0
	for i := range grid {
		for _, v := range grid[i] {
			if v&lowBit != 0 {
				a ^= v
			}
			if x&lowBit != 0 {
				a ^= x
			}
			x++
		}
	}

	b := a ^ axb
	for i := range grid {
		for _, v := range grid[i] {
			if v == b {
				b, a = a, b
			}
		}
	}

	return []int{a, b}
}
