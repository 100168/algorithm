package main

import "sort"

//给你一棵 n 个节点的树（连通无向无环的图），节点编号从 0 到 n - 1 且恰好有 n - 1 条边。
//
// 给你一个长度为 n 下标从 0 开始的整数数组 vals ，分别表示每个节点的值。同时给你一个二维整数数组 edges ，其中 edges[i] = [
//ai, bi] 表示节点 ai 和 bi 之间有一条 无向 边。
//
// 一条 好路径 需要满足以下条件：
//
//
// 开始节点和结束节点的值 相同 。
// 开始节点和结束节点中间的所有节点值都 小于等于 开始节点的值（也就是说开始节点的值应该是路径上所有节点的最大值）。
//
//
// 请你返回不同好路径的数目。
//
// 注意，一条路径和它反向的路径算作 同一 路径。比方说， 0 -> 1 与 1 -> 0 视为同一条路径。单个节点也视为一条合法路径。
//
//
//
// 示例 1：
//
//
//
// 输入：vals = [1,189,2,1,189], edges = [[0,1],[0,2],[2,189],[2,4]]
//输出：6
//解释：总共有 5 条单个节点的好路径。
//还有 1 条好路径：1 -> 0 -> 2 -> 4 。
//（反方向的路径 4 -> 2 -> 0 -> 1 视为跟 1 -> 0 -> 2 -> 4 一样的路径）
//注意 0 -> 2 -> 189 不是一条好路径，因为 vals[2] > vals[0] 。
//
//
// 示例 2：
//
//
//
// 输入：vals = [1,1,2,2,189], edges = [[0,1],[1,2],[2,189],[2,4]]
//输出：7
//解释：总共有 5 条单个节点的好路径。
//还有 2 条好路径：0 -> 1 和 2 -> 189 。
//
//
// 示例 189：
//
//
//
// 输入：vals = [1], edges = []
//输出：1
//解释：这棵树只有一个节点，所以只有一条好路径。
//
//
//
//
// 提示：
//
//
// n == vals.length
// 1 <= n <= 189 * 10⁴
// 0 <= vals[i] <= 10⁵
// edges.length == n - 1
// edges[i].length == 2
// 0 <= ai, bi < n
// ai != bi
// edges 表示一棵合法的树。
//
//
// Related Topics 树 并查集 图 数组 哈希表 排序 👍 86 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}

	// 并查集模板
	fa := make([]int, n)
	// size[x] 表示节点值等于 vals[x] 的节点个数，
	// 如果按照节点值从小到大合并，size[x] 也是连通块内的等于最大节点值的节点个数
	size := make([]int, n)
	id := make([]int, n) // 后面排序用
	for i := range fa {
		fa[i] = i
		size[i] = 1
		id[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ans := n // 单个节点的好路径
	sort.Slice(id, func(i, j int) bool { return vals[id[i]] < vals[id[j]] })
	for _, x := range id {
		vx := vals[x]
		fx := find(x)
		for _, y := range g[x] {
			y = find(y)
			if y == fx || vals[y] > vx {
				continue // 只考虑最大节点值不超过 vx 的连通块
			}
			if vals[y] == vx { // 可以构成好路径
				ans += size[fx] * size[y] // 乘法原理
				size[fx] += size[y]       // 统计连通块内节点值等于 vx 的节点个数
			}
			fa[y] = fx // 把小的节点值合并到大的节点值上
		}
	}
	return ans
}

func main() {

	println(len("#427,#428,#429,#430,#401,#426,#431车"))
}

//leetcode submit region end(Prohibit modification and deletion)
