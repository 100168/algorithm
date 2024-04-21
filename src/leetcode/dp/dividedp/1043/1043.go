package main

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)

	memo := make([]int, n)

	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int

	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		cur := 0
		maxVal := 0
		for j := i; i-j+1 <= k && j >= 0; j-- {
			maxVal = max(maxVal, arr[j])
			cur = max(dfs(j-1)+maxVal*(i-j+1), cur)

		}

		memo[i] = cur
		return cur
	}
	return dfs(n - 1)
}
