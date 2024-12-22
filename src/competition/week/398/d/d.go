package main

/*
*
给你有一个 非负 整数 k 。有一个无限长度的台阶，最低 一层编号为 0 。

虎老师有一个整数 jump ，一开始值为 0 。虎老师从台阶 1 开始，虎老师可以使用 任意 次操作，目标是到达第 k 级台阶。假设虎老师位于台阶 i ，一次 操作 中，虎老师可以：

向下走一级到 i - 1 ，但该操作 不能 连续使用，如果在台阶第 0 级也不能使用。
向上走到台阶 i + 2jump 处，然后 jump 变为 jump + 1 。
请你返回虎老师到达台阶 k 处的总方案数。

注意 ，虎老师可能到达台阶 k 处后，通过一些操作重新回到台阶 k 处，这视为不同的方案。
*/
func waysToReachStair(k int) int {
	memo := make([][]map[int]int, 31)
	for i := range memo {
		memo[i] = make([]map[int]int, 2)
		for j := range memo[i] {
			memo[i][j] = make(map[int]int)
		}

	}
	var dfs func(int, int, int) int

	dfs = func(jump int, p int, isUp int) int {

		if p > k+1 {
			return 0
		}
		cur := 0
		if p == k {
			cur++
		}
		if _, ok := memo[jump][isUp][p]; ok {
			return memo[jump][isUp][p]
		}
		if isUp == 0 {
			cur += dfs(jump, p-1, 1)
		}
		cur += dfs(jump+1, p+1<<jump, 0)
		memo[jump][isUp][p] = cur
		return cur
	}

	return dfs(0, 1, 0)
}

func main() {
	println(waysToReachStair(3))
}
