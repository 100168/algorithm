package main

import (
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"math/bits"
	"slices"
)

/*
有一条无限长的数轴，原点在 0 处，沿着 x 轴 正 方向无限延伸。

给你一个二维数组 queries ，它包含两种操作：

操作类型 1 ：queries[i] = [1, x] 。在距离原点 x 处建一个障碍物。数据保证当操作执行的时候，位置 x 处 没有 任何障碍物。
操作类型 2 ：queries[i] = [2, x, sz] 。判断在数轴范围 [0, x] 内是否可以放置一个长度为 sz 的物块，这个物块需要 完全 放置在范围 [0, x] 内。如果物块与任何障碍物有重合，那么这个物块 不能 被放置，但物块可以与障碍物刚好接触。注意，你只是进行查询，并 不是 真的放置这个物块。每个查询都是相互独立的。
请你返回一个 boolean 数组results ，如果第 i 个操作类型 2 的操作你可以放置物块，那么 results[i] 为 true ，否则为 false 。


*/
// 树状数组，为什么要倒着计算？
func getResults(queries [][]int) []bool {

	tree := rbt.NewWithIntComparator()

	maxVal := 0
	tree.Put(1, 0)
	var ans []bool
	for _, v := range queries {
		op, x := v[0], v[1]+1
		maxVal = max(x, maxVal)
		if op == 1 {
			if floor, ok := tree.Floor(x - 1); ok {
				tree.Put(x, x-floor.Key.(int))
			}
			if cell, ok := tree.Ceiling(x + 1); ok {
				tree.Put(cell.Key, cell.Key.(int)-x)
			}

		}
	}

	if floor, ok := tree.Floor(maxVal - 1); ok {
		tree.Put(maxVal, maxVal-floor.Key.(int))
	}
	bt := new(bitTree)
	bt.sum = make([]int, maxVal+1)
	bt.n = maxVal + 1
	for _, k := range tree.Keys() {
		node := tree.GetNode(k)
		bt.update(k.(int), node.Value.(int))
	}

	for i := len(queries) - 1; i >= 0; i-- {

		op, x := queries[i][0], queries[i][1]+1
		if op == 1 {
			tree.Remove(x)
			floor, _ := tree.Floor(x - 1)
			if cell, ok := tree.Ceiling(x); ok {
				tree.Put(cell.Key, cell.Key.(int)-floor.Key.(int))
				bt.update(cell.Key.(int), cell.Key.(int)-floor.Key.(int))
			}

		} else {
			se := queries[i][2]
			query := bt.query(x)
			floor, _ := tree.Floor(x - 1)
			diff := x - floor.Key.(int)
			if se <= query || diff >= se {
				ans = append(ans, true)
			} else {
				ans = append(ans, false)
			}
		}
	}
	slices.Reverse(ans)
	return ans
}

type bitTree struct {
	sum []int
	n   int
}

func lowBit(index int) int {
	return index & -index
}

func (bt bitTree) query(index int) int {
	ans := 0
	for index > 0 {
		ans = max(bt.sum[index], ans)
		index -= lowBit(index)

	}
	return ans
}

func (bt bitTree) update(index int, value int) {
	for index < bt.n {
		bt.sum[index] = max(bt.sum[index], value)
		index += lowBit(index)
	}
}

type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return res
}

func getResults2(queries [][]int) (ans []bool) {
	m := 0
	set := redblacktree.New[int, struct{}]()
	set.Put(0, struct{}{})
	pos := []int{0}
	for _, q := range queries {
		m = max(m, q[1])
		if q[0] == 1 {
			set.Put(q[1], struct{}{})
			pos = append(pos, q[1])
		}
	}
	m++
	set.Put(m, struct{}{}) // 哨兵

	t := make(fenwick, m)
	slices.Sort(pos)
	for i := 1; i < len(pos); i++ {
		t.update(pos[i], pos[i]-pos[i-1])
	}

	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		x := q[1]
		pre, _ := set.Floor(x - 1) // x 左侧最近障碍物的位置
		if q[0] == 1 {
			set.Remove(x)
			nxt, _ := set.Ceiling(x)           // x 右侧最近障碍物的位置
			t.update(nxt.Key, nxt.Key-pre.Key) // 更新 d[nxt] = nxt - pre
		} else {
			// 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
			maxGap := max(t.preMax(pre.Key), x-pre.Key)
			ans = append(ans, maxGap >= q[2])
		}
	}
	slices.Reverse(ans)
	return
}

// 动态开点
func getResults3(queries [][]int) (ans []bool) {

	rt := rbt.NewWithIntComparator()

	root := createNode(0, 1e9)
	sg := new(segmentTree)
	sg.root = root

	rt.Put(0, 0)
	rt.Put(int(1e9), 0)
	for _, v := range queries {

		op, x := v[0], v[1]
		pre, _ := rt.Floor(x - 1)
		if op == 1 {
			next, _ := rt.Ceiling(x + 1)
			sg.update(sg.root, x, x, x-pre.Key.(int))
			sg.update(sg.root, next.Key.(int), next.Key.(int), next.Key.(int)-x)
			rt.Put(x, 0)
		} else {
			st := v[2]
			query := sg.query(sg.root, 0, x)
			diff := max(query, x-pre.Key.(int))
			ans = append(ans, diff >= st)
		}
	}
	return
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
}

type seg []int

// 把 i 处的值改成 val
func (t seg) update(o, l, r, i, val int) {
	if l == r {
		t[o] = val
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(o<<1, l, m, i, val)
	} else {
		t.update(o<<1|1, m+1, r, i, val)
	}
	t[o] = max(t[o<<1], t[o<<1|1])
}

// 查询 [0,R] 中的最大值
func (t seg) query(o, l, r, R int) int {
	if r <= R {
		return t[o]
	}
	m := (l + r) >> 1
	if R <= m {
		return t.query(o<<1, l, m, R)
	}
	return max(t[o<<1], t.query(o<<1|1, m+1, r, R))
}

// 静态开点
func getResults4(queries [][]int) (ans []bool) {
	m := 0
	for _, q := range queries {
		m = max(m, q[1])
	}
	m++

	set := redblacktree.New[int, struct{}]()
	set.Put(0, struct{}{}) // 哨兵
	set.Put(m, struct{}{})
	t := make(seg, 2<<bits.Len(uint(m)))

	for _, q := range queries {
		x := q[1]
		pre, _ := set.Floor(x - 1) // x 左侧最近障碍物的位置
		if q[0] == 1 {
			nxt, _ := set.Ceiling(x) // x 右侧最近障碍物的位置
			set.Put(x, struct{}{})
			t.update(1, 0, m, x, x-pre.Key)       // 更新 d[x] = x - pre
			t.update(1, 0, m, nxt.Key, nxt.Key-x) // 更新 d[nxt] = nxt - x
		} else {
			// 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
			maxGap := max(t.query(1, 0, m, pre.Key), x-pre.Key)
			ans = append(ans, maxGap >= q[2])
		}
	}
	return
}
func main() {

	fmt.Println(getResults3([][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}}))
	fmt.Println(getResults2([][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}}))
	fmt.Println(getResults([][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}}))
}
