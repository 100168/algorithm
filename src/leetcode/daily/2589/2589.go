package main

import (
	"fmt"
	"math/rand"
	"slices"
)

/*
*
你有一台电脑，它可以 同时 运行无数个任务。给你一个二维整数数组 tasks ，
其中 tasks[i] = [starti, endi, durationi] 表示第 i 个任务需要在 闭区间
时间段 [starti, endi] 内运行 durationi 个整数时间点（但不需要连续）。

当电脑需要运行任务时，你可以打开电脑，如果空闲时，你可以将电脑关闭。

请你返回完成所有任务的情况下，电脑最少需要运行多少秒。

1 3 1,     2,3 1
*/

type seg struct {
	root *node
}

type node struct {
	l, r, m int
	left    *node
	right   *node
	cnt     int
	todo    bool
}

func newNode(l, r int) *node {
	cur := new(node)
	cur.l = l
	cur.r = r
	cur.m = (r + l) >> 1
	return cur
}

func (t seg) pushDown(cur *node) {
	if cur.left == nil {
		cur.left = newNode(cur.l, cur.m)
	}
	if cur.right == nil {
		cur.right = newNode(cur.m+1, cur.r)
	}
	if cur.todo {
		cur.todo = false
		t.do(cur.left)
		t.do(cur.right)
	}
}

func (t seg) do(cur *node) {
	cur.cnt = cur.r - cur.l + 1
	cur.todo = true
}

func (t seg) query(cur *node, l, r int) int {
	if cur.l >= l && cur.r <= r {
		return cur.cnt
	}
	t.pushDown(cur)

	res := 0
	if l <= cur.m {
		res += t.query(cur.left, l, r)
	}
	if r > cur.m {
		res += t.query(cur.right, l, r)
	}
	return res

}

func (t seg) pushUp(cur *node) {
	cur.cnt = cur.left.cnt + cur.right.cnt
}

// l <= t[o].l && t[o].r <= r && size-t[o].cnt <= *suffix { // 整个区间全部改为运行中
//		*suffix -= size - t[o].cnt
//		t.do(o)
//		return
//	}

func (t seg) update(cur *node, l, r int, val *int) {
	size := cur.r - cur.l + 1
	if cur.l >= l && cur.r <= r && size <= *val {
		*val -= size - cur.cnt
		t.do(cur)
		return
	}
	t.pushDown(cur)
	if r > cur.m {
		t.update(cur.right, l, r, val)
	}
	if *val > 0 {
		t.update(cur.left, l, r, val)
	}
	t.pushUp(cur)
}

func findMinimumTime(tasks [][]int) int {
	slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
	s := seg{newNode(1, tasks[len(tasks)-1][1])}
	for _, t := range tasks {
		start, end, d := t[0], t[1], t[2]
		d -= s.query(s.root, start, end)
		if d > 0 {
			s.update(s.root, start, end, &d)
		}
	}
	return s.root.cnt
}

func compare(tasks [][]int) (ans int) {
	slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
	do := make([]bool, tasks[len(tasks)-1][1]+1)

	for _, t := range tasks {
		start, end, d := t[0], t[1], t[2]
		for i := start; i <= end; i++ {
			if do[i] {
				d--
			}
		}
		if d > 0 {
			for i := end; d > 0; i-- {
				if !do[i] {
					do[i] = true
					d--
					ans++
				}
			}
		}
	}
	return
}

/**

type seg []struct {
	l, r, cnt int
	todo      bool
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) do(i int) {
	o := &t[i]
	o.cnt = o.r - o.l + 1
	o.todo = true
}

func (t seg) spread(o int) {
	if t[o].todo {
		t[o].todo = false
		t.do(o << 1)
		t.do(o<<1 | 1)
	}
}

// 查询区间 [l,r]   o=1
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].cnt
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1

	cur := 0
	if l <= m {
		cur += t.query(o<<1, l, r)
	}
	if m < r {
		cur += t.query(o<<1|1, l, r)
	}
	return cur
}

// 新增区间 [l,r] 后缀的 suffix 个时间点   o=1
// 相当于把线段树二分和线段树更新合并成了一个函数，时间复杂度为 O(log u)
func (t seg) update(o, l, r int, suffix *int) {
	size := t[o].r - t[o].l + 1
	if t[o].cnt == size { // 全部为运行中
		return
	}
	if l <= t[o].l && t[o].r <= r && size-t[o].cnt <= *suffix { // 整个区间全部改为运行中
		*suffix -= size - t[o].cnt
		t.do(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r > m { // 先更新右子树
		t.update(o<<1|1, l, r, suffix)
	}
	if *suffix > 0 { // 再更新左子树（如果还有需要新增的时间点）
		t.update(o<<1, l, r, suffix)
	}
	t[o].cnt = t[o<<1].cnt + t[o<<1|1].cnt
}

func findMinimumTime2(tasks [][]int) (ans int) {
	slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
	u := tasks[len(tasks)-1][1]
	st := make(seg, 2<<bits.Len(uint(u-1)))
	st.build(1, 1, u)
	for _, t := range tasks {
		start, end, d := t[0], t[1], t[2]
		d -= st.query(1, start, end) // 去掉运行中的时间点
		if d > 0 {
			st.update(1, start, end, &d) // 新增时间点
		}
	}
	return st[1].cnt
}

*/

func main() {
	//println(findMinimumTime([][]int{{2, 13, 2}, {6, 18, 5}, {2, 13, 3}}))
	//println(compare([][]int{{2, 13, 2}, {6, 18, 5}, {2, 13, 3}}))

	for i := 0; i < 1000; i++ {
		l := rand.Intn(1000)
		cur := make([][]int, 0)
		for j := 0; j <= l; j++ {
			start := rand.Intn(1000)
			end := rand.Intn(1000-start) + start
			dur := rand.Intn(end - start + 1)
			cur = append(cur, []int{start, end, dur})
		}
		if findMinimumTime(cur) != compare(cur) {
			fmt.Println("fuck")
		}
	}
}
