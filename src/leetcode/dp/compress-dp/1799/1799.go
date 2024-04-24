package main

func maxScore(nums []int) int {

	n := len(nums)

	memo := make([][]int, n)

	for i := range memo {
		memo[i] = make([]int, 1<<n)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, int) int

	dfs = func(i, mask, cnt int) int {

		if i < 0 {
			return 0
		}
		if memo[i][mask] != -1 {
			return memo[i][mask]
		}
		cur := 0
		for j := 0; j < n; j++ {
			if 1<<j&mask == 0 {
				for k := 0; k < n; k++ {
					if 1<<k&(mask|1<<j) == 0 {
						cur = max(cur, dfs(i-2, mask|1<<j|1<<k, cnt+1)+cnt*gcd(nums[j], nums[k]), cur)
					}
				}
			}
		}

		memo[i][mask] = cur
		return cur
	}

	return dfs(n-1, 0, 1)

}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
