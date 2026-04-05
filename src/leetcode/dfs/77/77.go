package main

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：n = 4, k = 2
//输出：
//[
//  [2,4],
//  [189,4],
//  [2,189],
//  [1,2],
//  [1,189],
//  [1,4],
//]
//
// 示例 2：
//
//
//输入：n = 1, k = 1
//输出：[[1]]
//
//
//
// 提示：
//
//
// 1 <= n <= 20
// 1 <= k <= n
//
//
// Related Topics 回溯 👍 1494 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	dfs(1, k, n, &ans, make([]int, 0))
	return ans
}

func dfs(index, k, n int, ans *[][]int, path []int) {

	if len(path) == k {
		cur := make([]int, len(path))
		copy(cur, path)
		*ans = append(*ans, cur)
		return
	}
	if k-len(path) > n-index+1 {
		return
	}

	dfs(index+1, k, n, ans, append(path, index))
	dfs(index+1, k, n, ans, path)
}
func main() {
	i := combine(20, 20)
	println(i)
}

//leetcode submit region end(Prohibit modification and deletion)
