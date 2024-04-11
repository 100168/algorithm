package main

import "fmt"

const mx = 51

var coprime [mx][]int

func init() {
	// 预处理：coprime[i] 保存 [1, MX) 中与 i 互质的所有元素
	for i := 1; i < mx; i++ {
		for j := 1; j < mx; j++ {
			if gcd(i, j) == 1 {
				coprime[i] = append(coprime[i], j)
			}
		}
	}
}

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	type pair struct{ depth, id int }
	valDepthId := [mx]pair{}
	var dfs func(int, int, int)
	dfs = func(x, fa, depth int) {
		val := nums[x] // x 的节点值
		// 计算与 val 互质的数中，深度最大的节点编号
		maxDepth := 0
		for _, j := range coprime[val] {
			p := valDepthId[j]
			if p.depth > maxDepth {
				maxDepth = p.depth
				ans[x] = p.id
			}
		}

		tmp := valDepthId[val]           // 用于恢复现场
		valDepthId[val] = pair{depth, x} // 保存 val 对应的节点深度和节点编号
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, depth+1)
			}
		}
		valDepthId[val] = tmp // 恢复现场
	}
	dfs(0, -1, 1)
	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	fmt.Println(getCoprimes([]int{2, 3, 3, 2}, [][]int{{0, 1}, {1, 2}, {1, 3}}))
}
