package main

import "fmt"

/*用以太网线缆将n台计算机连接成一个网络，计算机的编号从0到n-1。线缆用connections表示，其中 connections[i] = [a, b] 连接了计算机 a 和 b。

网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。

给你这个计算机网络的初始布线 connections，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回 -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-operations-to-make-network-connected
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type UnionSet struct {
	parent   []int
	size     []int
	setCount int
}

func newUnionSet(n int) *UnionSet {

	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &UnionSet{parent, size, n}

}
func (unionSet *UnionSet) findFather(cur int) int {
	var stack []int
	for cur != unionSet.parent[cur] {
		cur = unionSet.parent[cur]
		stack = append(stack, cur)
	}

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		unionSet.parent[pop] = cur
	}
	return cur
}

func (unionSet *UnionSet) isSameSet(v1 int, v2 int) bool {
	return unionSet.findFather(v1) == unionSet.findFather(v2)
}

func (unionSet *UnionSet) union(v1 int, v2 int) {

	fx := unionSet.findFather(v1)
	fy := unionSet.findFather(v2)

	newSize := unionSet.size[fx] + unionSet.size[fy]
	//相同就说名连通分量同源，可以把线让出来
	if fx != fy {
		if unionSet.size[fy] > unionSet.size[fx] {
			fx, fy = fy, fx
		}
		unionSet.parent[fy] = fx
		unionSet.size[fx] = newSize
		unionSet.setCount--
	}
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	uf := newUnionSet(n)
	for _, c := range connections {
		uf.union(c[0], c[1])
	}

	return uf.setCount - 1
}
func main() {

	co := [][]int{{0, 1}, {0, 2}, {1, 2}}
	connected := makeConnected(4, co)
	fmt.Println(connected)
}
