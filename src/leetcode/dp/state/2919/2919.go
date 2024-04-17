package main

func minIncrementOperations(nums []int, k int) int64 {

	n := len(nums)

	memo := make([][3]int64, n)
	for i := range memo {
		memo[i] = [3]int64{-1, -1, -1}
	}

	var dfs func(int, int) int64

	dfs = func(i int, rest int) int64 {
		if i < 0 {
			return 0
		}

		cur := int64(max(k-nums[i], 0))

		if memo[i][rest] != -1 {
			return memo[i][rest]
		}
		if rest == 0 {
			cur += dfs(i-1, 2)
		}
		if rest == 1 {
			cur = min(dfs(i-1, 0), dfs(i-1, 2)+cur)
		}
		if rest == 2 {
			cur = min(dfs(i-1, 2)+cur, dfs(i-1, 0), dfs(i-1, 1))
		}
		memo[i][rest] = cur
		return cur

	}

	return min(dfs(n-1, 0), dfs(n-1, 1), dfs(n-1, 2))
}

func main() {
	println(minIncrementOperations([]int{1, 1, 2}, 1))
}
