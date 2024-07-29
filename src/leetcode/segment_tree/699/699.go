package main

import "fmt"

/**
在二维平面上的 x 轴上，放置着一些方块。

给你一个二维整数数组 positions ，其中 positions[i] = [lefti, sideLengthi] 表示：第 i 个方块边长为 sideLengthi ，其左侧边与 x 轴上坐标点 lefti 对齐。

每个方块都从一个比目前所有的落地方块更高的高度掉落而下。方块沿 y 轴负方向下落，直到着陆到 另一个正方形的顶边 或者是 x 轴上 。一个方块仅仅是擦过另一个方块的左侧边或右侧边不算着陆。一旦着陆，它就会固定在原地，无法移动。

在每个方块掉落后，你必须记录目前所有已经落稳的 方块堆叠的最高高度 。

返回一个整数数组 ans ，其中 ans[i] 表示在第 i 块方块掉落后堆叠的最高高度。
*/

func fallingSquares(positions [][]int) []int {

	sg := new(segmentTree)

	sg.root = createNode(1, 1e9)

	ans := make([]int, len(positions))
	for i, pos := range positions {

		l, r := pos[0], pos[1]
		r = l + r
		query := sg.query(sg.root, l, r-1)
		sg.update(sg.root, l, r-1, query+pos[1])
		ans[i] = sg.root.val

	}
	return ans
}

type segmentTree struct {
	root *node
}
type node struct {
	l     int
	r     int
	m     int
	left  *node
	right *node
	val   int
	add   bool
}

func createNode(l, r int) *node {
	cur := new(node)
	cur.l = l
	cur.r = r
	cur.m = (l + r) >> 1
	return cur
}

func (sg segmentTree) query(cur *node, l, r int) int {
	if cur.l >= l && cur.r <= r {
		return cur.val
	}
	sg.pushDown(cur)
	ans := 0
	if l <= cur.m {
		ans = sg.query(cur.left, l, r)
	}
	if r > cur.m {
		ans = max(ans, sg.query(cur.right, l, r))
	}
	return ans
}

func (sg segmentTree) update(cur *node, l, r, val int) {

	if cur.l >= l && cur.r <= r {
		cur.val = val
		cur.add = true
		return
	}
	sg.pushDown(cur)
	if l <= cur.m {
		sg.update(cur.left, l, r, val)
	}
	if r > cur.m {
		sg.update(cur.right, l, r, val)
	}
	sg.pushUp(cur)
}
func (sg segmentTree) pushUp(cur *node) {
	cur.val = max(cur.left.val, cur.right.val)
}

func (sg segmentTree) pushDown(cur *node) {
	if cur.left == nil {
		cur.left = createNode(cur.l, cur.m)
	}
	if cur.right == nil {
		cur.right = createNode(cur.m+1, cur.r)
	}
	if cur.add {
		cur.left.val = max(cur.val, cur.left.val)
		cur.right.val = max(cur.val, cur.right.val)
		cur.add = false
		cur.left.add = true
		cur.right.add = true
	}

}

func main() {
	fmt.Println(fallingSquares([][]int{{1, 5}, {2, 2}, {7, 5}}))
}
