package main

func validPartition(nums []int) bool {

	n := len(nums)

	memo := make([][]map[bool]bool, n)

	for i := range memo {
		memo[i] = make([]map[bool]bool, 4)

		for j := range memo[i] {
			memo[i][j] = make(map[bool]bool)
		}
	}
	var dfs func(int, int, bool) bool

	dfs = func(i, size int, unique bool) bool {
		if i < 0 {
			if unique {
				return size == 3
			}
			return size >= 2
		}
		if _, ok := memo[i][size][unique]; ok {
			return memo[i][size][unique]
		}
		cur := false
		if unique {
			if size <= 2 && nums[i]+1 == nums[i+1] {
				cur = cur || dfs(i+-1, size+1, unique)
			}
			if size == 3 {
				cur = cur || dfs(i-1, 1, true) || dfs(i-1, 1, false)
			}
		}
		if !unique {
			if size <= 2 && nums[i] == nums[i+1] {
				cur = cur || dfs(i-1, size+1, false)
			} else if size == 3 {
				cur = cur || dfs(i-1, 1, false) || dfs(i-1, 1, true)
			}
		}
		memo[i][size][unique] = cur
		return cur

	}
	return dfs(n-1, 3, false)
}

func validPartition2(nums []int) bool {

	n := len(nums)

	memo := make(map[int]bool, n)
	var dfs func(int) bool
	dfs = func(i int) bool {
		if i < 0 {
			return true
		}
		if _, ok := memo[i]; ok {
			return memo[i]
		}
		cur := false
		if i > 1 {
			cur = dfs(i-3) && nums[i-2]+1 == nums[i-1] && nums[i-1]+1 == nums[i]
			cur = cur || dfs(i-3) && nums[i-2] == nums[i-1] && nums[i-1] == nums[i]
		}
		if i > 0 {
			cur = cur || dfs(i-2) && nums[i-1] == nums[i]
		}
		memo[i] = cur
		return cur
	}
	return dfs(n - 1)
}

func main() {

	println(validPartition([]int{4, 4, 4, 5, 6}))
}
