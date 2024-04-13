package main

func wiggleMaxLength(nums []int) int {

	m := len(nums)
	cache := make([][][]int, m)
	for i := range cache {
		cache[i] = make([][]int, m)
		for j := range cache[i] {
			cache[i][j] = make([]int, 2)
			for k := range cache[i][j] {
				cache[i][j][k] = -1
			}

		}
	}
	var dfs func(int, int, int) int

	dfs = func(i, j, k int) int {
		if i < 0 {
			return 1
		}
		if cache[i][j][k] != -1 {
			return cache[i][j][k]
		}
		cur := 0
		if k == 0 {
			if nums[i] < nums[j] {
				cur = dfs(i-1, i, 1) + 1
			}

		} else {
			if nums[i] > nums[j] {
				cur = dfs(i-1, i, 0) + 1
			}
		}
		cur = max(dfs(i-1, j, k), cur)
		cache[i][j][k] = cur
		return cur
	}
	return max(dfs(m-2, m-1, 0), dfs(m-2, m-1, 1))
}

func wiggleMaxLength2(nums []int) int {

	m := len(nums)

	up := make([]int, m)
	down := make([]int, m)
	up[0] = 1
	down[0] = 1

	for i := 1; i < m; i++ {
		if nums[i] < nums[i-1] {
			down[i] = max(down[i-1], up[i-1]+1)
			up[i] = up[i-1]
		} else if nums[i] > nums[i-1] {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return max(down[m-1], up[m-1])

}
