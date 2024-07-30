package main

/*
*
给你一个下标从 0 开始的二维数组 variables ，其中 variables[i] = [ai, bi, ci, mi]，以及一个整数 target 。

如果满足以下公式，则下标 i 是 好下标：

0 <= i < variables.length
((aibi % 10)ci) % mi == target
返回一个由 好下标 组成的数组，顺序不限 。

示例 1：
输入：variables = [[2,3,3,10],[3,3,3,1],[6,1,1,4]], target = 2
输出：[0,2]
解释：对于 variables 数组中的每个下标 i ：
1) 对于下标 0 ，variables[0] = [2,3,3,10] ，(23 % 10)3 % 10 = 2 。
2) 对于下标 1 ，variables[1] = [3,3,3,1] ，(33 % 10)3 % 1 = 0 。
3) 对于下标 2 ，variables[2] = [6,1,1,4] ，(61 % 10)1 % 4 = 2 。
因此，返回 [0,2] 作为答案。
*/
func getGoodIndices(variables [][]int, target int) []int {

	var ans []int
	for i, v := range variables {

		ai, bi, ci, mi := v[0], v[1], v[2], v[3]

		if pow(pow(ai, bi, 10), ci, mi) == target {
			ans = append(ans, i)
		}
	}
	return ans

}

func pow(p, n, mod int) int {

	ans := 1

	for ; n > 0; n >>= 1 {

		if n&1 == 1 {
			ans = ans * p % mod
		}
		p = p * p % mod
	}
	return ans
}
