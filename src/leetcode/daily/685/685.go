package main

import "fmt"

/*
*
在本问题中，有根树指满足以下条件的 有向 图。该树只有一个根节点，所有其他节点都是该根节点的后继。
该树除了根节点之外的每一个节点都有且只有一个父节点，而根节点没有父节点。

输入一个有向图，该图由一个有着 n 个节点（节点值不重复，从 1 到 n）的树及一条附加的有向边构成。
附加的边包含在 1 到 n 中的两个不同顶点间，这条附加的边不属于树中已存在的边。

结果图是一个以边组成的二维数组 edges 。
每个元素是一对 [ui, vi]，用以表示 有向 图中连接顶点 ui 和顶点 vi 的边，其中 ui 是 vi 的一个父节点。

返回一条能删除的边，使得剩下的图是有 n 个节点的有根树。若有多个答案，返回最后出现在给定二维数组的答案。

存在以下两种情况：
1.节点x有两个父节点
2.节点边[x,y]加入后成为一个环
*/
func findRedundantDirectedConnection(edges [][]int) []int {

	n := len(edges)

	u := new(uf)
	u.parent = make([]int, n+1)

	for i := range u.parent {
		u.parent[i] = i
	}

	var conflictE []int

	var circleE []int
	inCnt := make([]int, n+1)
	for _, v := range edges {
		x, y := v[0], v[1]

		if inCnt[y] > 0 {
			conflictE = v
		} else {
			inCnt[y] = x
			if u.find(x) == u.find(y) {
				circleE = v
			}
			u.union(x, y)
		}

	}
	//冲突点为nil直接返回循环点
	if conflictE == nil {
		return circleE
	}
	//如果循环节点为nil直接返回第一次接触的冲突边
	//否则返回最后一次出现的冲突边
	if circleE != nil {
		return []int{inCnt[conflictE[1]], conflictE[1]}
	}

	return conflictE
}

type uf struct {
	parent []int
}

func (u uf) find(a int) int {
	p := u.parent
	for p[a] != a {
		p[a] = p[p[a]]
		a = p[a]
	}
	return a
}

func (u uf) union(a, b int) {
	u.parent[u.find(a)] = u.parent[u.find(b)]
}

func main() {

	fmt.Println(findRedundantDirectedConnection([][]int{{5, 2}, {5, 1}, {3, 1}, {3, 4}, {3, 5}}))
}
