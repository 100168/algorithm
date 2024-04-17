package main

func countSpecialSubsequences(nums []int) int {

	n := len(nums)

	mod := 1_000_000_007
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = []int{-1, -1, -1}
	}
	var dfs func(int, int, bool) int
	dfs = func(i, j int, isNum bool) int {
		if i < 0 {
			if j == 0 && isNum {
				return 1
			}
			return 0
		}
		if memo[i][j] != -1 && isNum {
			return memo[i][j]
		}
		res := dfs(i-1, j, isNum)
		if j-nums[i] == 1 && isNum {
			res += dfs(i-1, j-1, isNum)
		}
		if j == nums[i] {
			res += dfs(i-1, j, true)
		}
		res %= mod
		if isNum {
			memo[i][j] = res
		}
		return res
	}
	return dfs(n-1, 2, false)
}

func main() {
	println(countSpecialSubsequences([]int{0, 1, 2}))
}
