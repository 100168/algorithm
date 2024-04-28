package main

/*
*

给你 3 个正整数 zero ，one 和 limit 。

一个 二进制数组 arr 如果满足以下条件，那么我们称它是 稳定的 ：

0 在 arr 中出现次数 恰好 为 zero 。
1 在 arr 中出现次数 恰好 为 one 。
arr 中每个长度超过 limit 的 子数组 都 同时 包含 0 和 1 。
请你返回 稳定 二进制数组的 总 数目。

由于答案可能很大，将它对 109 + 7 取余 后返回。
*/
func numberOfStableArrays(zero int, one int, limit int) int {

	mod := int(1e9 + 7)
	n := zero + one

	memo := make([][][]int, n)

	for i := range memo {
		memo[i] = make([][]int, one+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, n+1)
			for k := 0; k <= n; k++ {
				memo[i][j][k] = -1
			}
		}

	}
	var dfs func(int, int, int) int

	dfs = func(i, rest int, preOne int) int {

		if rest < 0 {
			return 0
		}
		if i < 0 {
			if rest > 0 {
				return 0
			}
			return 1
		}
		if preOne != -1 && memo[i][rest][preOne] != -1 {
			return memo[i][rest][preOne]
		}
		cur := 0

		if preOne-i < limit {
			cur += dfs(i-1, rest, preOne)
		}

		//fmt.Println("select", i, rest, one)
		cur += dfs(i-1, rest-1, i)

		cur %= mod
		if preOne != -1 {
			memo[i][rest][preOne] = cur
		}

		return cur
	}
	return dfs(n-1, one, n-1)
}

func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

func main() {
	println(numberOfStableArrays(1, 2, 1))
	println(numberOfStableArrays(2, 1, 2))
	println(numberOfStableArrays(3, 3, 2))
}
