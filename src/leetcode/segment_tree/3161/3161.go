package main

import (
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"slices"
)

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

func main() {

	fmt.Println(getResults([][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}}))
}
