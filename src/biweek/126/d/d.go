package main

/*
3082. 求出所有子序列的能量和
困难
相关企业

提示
给你一个长度为 n 的整数数组 nums 和一个 正 整数 k 。

一个整数数组的 能量 定义为和 等于 k 的子序列的数目。

请你返回 nums 中所有子序列的 能量和 。

由于答案可能很大，请你将它对 109 + 7 取余 后返回。


观察

如果和为 kkk 的子序列 SSS 的长度是 ccc，那么有多少个子序列 TTT，会包含 SSS 呢？即 SSS 是 TTT 的子序列。

例如示例 1，子序列 S=[3]S=[3]S=[3]，出现在子序列 [1,2,3],[1,3],[2,3],[3][1,2,3],[1,3],[2,3],[3][1,2,3],[1,3],[2,3],[3] 中，这 444 个子序列都可以是 TTT。

除了 333 以外的每个数，都可以选择在/不在包含 [3][3][3] 的子序列 TTT 中。

所以有 2n−c2^{n-c}2
n−c
  个子序列 TTT。

这意味着 SSS 对答案的贡献是 2n−c2^{n-c}2
n−c
 。

作者：灵茶山艾府
链接：https://leetcode.cn/problems/find-the-sum-of-the-power-of-all-subsequences/solutions/2691661/liang-chong-fang-fa-er-wei-yi-wei-0-1-be-2e47/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

func sumOfPower(nums []int, k int) int {

	n := len(nums)
	type pair struct {
		val int
		cnt int
	}
	cntMap := make(map[int]int)

	for i := range nums {
		cntMap[nums[i]]++
	}
	pairs := make([]pair, 0)
	for k, v := range cntMap {
		pairs = append(pairs, pair{k, v})
	}
	var dfs func(int, int) int

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dfs = func(index int, rest int) int {
		if rest < 0 || index == n {
			return 0
		}
		if rest == 0 {
			return 1
		}
		if dp[index][rest] != -1 {
			return dp[index][rest]
		}
		cur := int64(0)
		val, p := pairs[index].val, pairs[index].cnt
		for j := 0; j <= p; j++ {
			cur += int64(dfs(index+1, rest-val*p)) * multiply(n, j)
			cur %= 1_000_000_007
		}
		dp[index][rest] = int(cur)
		return int(cur)

	}
	return dfs(0, k)
}
func multiply(a, b int) int64 {
	if b == 0 {
		return 1
	}
	ans := int64(1)
	for i := 1; i <= a; i++ {
		ans *= int64(i)
	}
	div := int64(1)
	for i := 1; i <= b; i++ {
		div *= int64(i)
	}
	return ans / div % int64(1_000_000_007)
}

func main() {
	println(sumOfPower([]int{1, 2, 3}, 3))
}
