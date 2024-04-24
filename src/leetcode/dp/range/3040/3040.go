package main

import "fmt"

/*给你一个整数数组 nums ，如果 nums 至少 包含 2 个元素，你可以执行以下操作中的 任意 一个：

选择 nums 中最前面两个元素并且删除它们。
选择 nums 中最后两个元素并且删除它们。
选择 nums 中第一个和最后一个元素并且删除它们。
一次操作的 分数 是被删除元素的和。

在确保 所有操作分数相同 的前提下，请你求出 最多 能进行多少次操作。

请你返回按照上述要求 最多 可以进行的操作次数。*/

func maxOperations(nums []int) int {

	//1 2 3 4 5
	n := len(nums)

	memo := make([][]map[int]int, n)

	for i := range memo {
		memo[i] = make([]map[int]int, n)
		for j := range memo[i] {
			memo[i][j] = make(map[int]int)
		}
	}
	var dfs func(int, int, int) int

	dfs = func(i, j int, score int) int {
		if i >= j {
			return 0
		}
		if _, ok := memo[i][j][score]; ok {
			return memo[i][j][score]
		}
		cur := 0

		if nums[i]+nums[i+1] == score {
			cur = dfs(i+2, j, score) + 1
		}
		if nums[i]+nums[j] == score {
			cur = max(dfs(i+1, j-1, score)+1, cur)
		}

		if nums[j]+nums[j-1] == score {
			cur = max(dfs(i, j-2, score)+1, cur)
		}
		memo[i][j][score] = cur
		return cur

	}
	score1 := nums[0] + nums[1]
	score2 := nums[n-1] + nums[n-2]
	score3 := nums[0] + nums[n-1]
	return max(dfs(2, n-1, score1), dfs(0, n-3, score2), dfs(1, n-2, score3)) + 1
}

func maxOperations2(nums []int) int {
	n := len(nums)
	res1 := helper(nums[2:], nums[0]+nums[1])       // 最前面两个
	res2 := helper(nums[:n-2], nums[n-2]+nums[n-1]) // 最后两个
	res3 := helper(nums[1:n-1], nums[0]+nums[n-1])  // 第一个和最后一个
	return max(res1, res2, res3) + 1                // 加上第一次操作
}

func helper(a []int, target int) int {
	n := len(a)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i >= j {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		if a[i]+a[i+1] == target { // 最前面两个
			res = max(res, dfs(i+2, j)+1)
		}
		if a[j-1]+a[j] == target { // 最后两个
			res = max(res, dfs(i, j-2)+1)
		}
		if a[i]+a[j] == target { // 第一个和最后一个
			res = max(res, dfs(i+1, j-1)+1)
		}
		*p = res
		return
	}
	return dfs(0, n-1)
}

func main() {
	fmt.Println(maxOperations([]int{3, 2, 1, 2, 3, 4}))
}
