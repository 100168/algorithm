package main

/*
给你一个下标从 0 开始的整数数组 nums ，它包含 n 个 互不相同 的正整数。如果 nums 的一个排列满足以下条件，我们称它是一个特别的排列：

对于 0 <= i < n - 1 的下标 i ，要么 nums[i] % nums[i+1] == 0 ，要么 nums[i+1] % nums[i] == 0 。
请你返回特别排列的总数目，由于答案可能很大，请将它对 109 + 7 取余 后返回。
*/
func specialPerm(nums []int) int {

	mod := int(1e9 + 7)

	n := len(nums)

	memo := make([][]int, 1<<n)
	for i := range memo {
		memo[i] = make([]int, n+1)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(mask, pre int) int {
		if mask == 1<<n-1 {
			return 1
		}
		if memo[mask][pre] != -1 {
			return memo[mask][pre]
		}
		cur := 0
		for j := 0; j < n; j++ {
			if 1<<j&mask == 0 && (pre == 0 || nums[pre-1]%nums[j] == 0 || nums[j]%nums[pre-1] == 0) {
				cur += dfs(mask|1<<j, j+1)
				cur %= mod
			}
		}
		memo[mask][pre] = cur
		return cur
	}
	return dfs(0, 0)
}

func specialPerm2(nums []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(nums)
	m := 1 << n
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == 0 {
			return 1 // 找到一个特别排列
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		for k, x := range nums {
			if i>>k&1 > 0 && (nums[j]%x == 0 || x%nums[j] == 0) {
				res = (res + dfs(i^(1<<k), k)) % mod
			}
		}
		*p = res
		return
	}
	for j := range nums {
		ans = (ans + dfs((m-1)^(1<<j), j)) % mod
	}
	return
}

func main() {
	println(specialPerm([]int{1, 4, 3}))
}
