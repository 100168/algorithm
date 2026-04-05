package main

import (
	"math"
)

// 给你一个整数数组，返回它的某个 非空 子数组（连续元素）在执行一次可选的删除操作后，
// 所能得到的最大元素总和。换句话说，你可以从原数组中选出一个子数组，并可以
// 决定要不要从中删除一个元素（只能删一次哦），（删除后）子数组中至少应当有一个元素，
// 然后该子数组（剩下）的元素总和是所有子数组之中最大的。
//
// 注意，删除一个元素后，子数组 不能为空。
var dp [][]int
var nums []int

/*
1.子数组不能为空 --》可以设置一个虚拟数无穷小
*/
func maximumSum(arr []int) int {
	n := len(arr)
	nums = arr
	dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
		for j := range dp[i] {
			dp[i][j] = math.MinInt
		}
	}
	ans := math.MinInt
	for i := 0; i < n; i++ {
		cur := max(dfs(i, 1), dfs(i, 0))
		ans = max(ans, cur)
	}
	return ans
}

func dfs(i, j int) int {
	if i < 0 {
		return math.MinInt / 2
	}
	if dp[i][j] != math.MinInt {
		return dp[i][j]
	}
	if j == 0 {
		dp[i][j] = max(dfs(i-1, j), 0) + nums[i]
		return dp[i][j]
	}
	dp[i][j] = max(dfs(i-1, j)+nums[i], dfs(i-1, j-1))

	return dp[i][j]
}

func maximumSum2(arr []int) int {

	n := len(arr)
	f0 := math.MinInt / 2
	f1 := math.MinInt / 2
	ans := math.MinInt

	for i := 1; i <= n; i++ {
		f1 = max(f1+arr[i-1], f0)
		f0 = max(f0, 0) + arr[i-1]
		ans = max(ans, max(f0, f1))
	}
	return ans

}

func main() {
	nums := []int{1, -2, 0, 3}
	println(maximumSum(nums))
	println(maximumSum2(nums))
}
