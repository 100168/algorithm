package main

import "fmt"

/*
*
你正在维护一个项目，该项目有 n 个方法，编号从 0 到 n - 1。

给你两个整数 n 和 k，以及一个二维整数数组 invocations，其中 invocations[i] = [ai, bi] 表示方法 ai 调用了方法 bi。

已知如果方法 k 存在一个已知的 bug。那么方法 k 以及它直接或间接调用的任何方法都被视为 可疑方法 ，我们需要从项目中移除这些方法。

只有当一组方法没有被这组之外的任何方法调用时，这组方法才能被移除。

返回一个数组，包含移除所有 可疑方法 后剩下的所有方法。你可以以任意顺序返回答案。如果无法移除 所有 可疑方法，则 不 移除任何方法。
*/
func remainingMethods(n int, k int, invocations [][]int) []int {

	g := make(map[int][]int)

	for _, inv := range invocations {
		x, y := inv[0], inv[1]
		g[x] = append(g[x], y)
	}

	delMap := make(map[int]bool)

	var dfs func(x int)

	dfs = func(x int) {
		if delMap[x] {
			return
		}
		delMap[x] = true
		for _, v := range g[x] {
			dfs(v)
		}
	}

	ans1 := make([]int, n)

	for i := range ans1 {
		ans1[i] = i
	}
	dfs(k)
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if !delMap[i] {
			for _, v := range g[i] {
				if delMap[v] {
					return ans1
				}
			}
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Println(remainingMethods(5, 0, [][]int{{1, 2}, {0, 2}, {0, 1}, {3, 4}}))
}
