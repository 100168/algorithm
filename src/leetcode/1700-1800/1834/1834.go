package main

import (
	"container/heap"
	"sort"
)

// 给你一个二维数组 tasks ，用于表示 n 项从 0 到 n - 1 编号的任务。其中 tasks[i] = [enqueueTimei, processingTimei]
// 意味着第 i 项任务将会于 enqueueTimei 时进入任务队列，需要 processingTimei 的时长完成执行。
//
// 现有一个单线程 CPU ，同一时间只能执行 最多一项 任务，该 CPU 将会按照下述方式运行：
//
// 如果 CPU 空闲，且任务队列中没有需要执行的任务，则 CPU 保持空闲状态。
// 如果 CPU 空闲，但任务队列中有需要执行的任务，则 CPU 将会选择 执行时间最短 的任务开始执行。如果多个任务具有同样的最短执行时间，则选择下标最小的
// 任务开始执行。
// 一旦某项任务开始执行，CPU 在 执行完整个任务 前都不会停止。
// CPU 可以在完成一项任务后，立即开始执行一项新任务。
//
// 返回 CPU 处理任务的顺序。
type pair struct{ t, i int }
type hp []pair

func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Less(i, j int) bool {
	a, b := (*h)[i], (*h)[j]
	return a.t < b.t || a.t == b.t && a.i < b.i
}
func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func (h *hp) push(v pair) {
	heap.Push(h, v)
}
func (h *hp) pop() pair {
	return heap.Pop(h).(pair)
}

func getOrder(a [][]int) (ans []int) {
	for i := range a {
		a[i] = append(a[i], i)
	}
	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] })
	h := &hp{}
	for i, cur, n := 0, 0, len(a); i < n; {
		if h.Len() > 0 {
			p := h.pop()
			ans = append(ans, p.i)
			cur += p.t
		}
		if h.Len() == 0 && cur < a[i][0] {
			cur = a[i][0]
		}
		for ; i < n && a[i][0] <= cur; i++ {
			h.push(pair{a[i][1], a[i][2]})
		}
	}
	for h.Len() > 0 {
		ans = append(ans, h.pop().i)
	}
	return
}
func main() {

}
