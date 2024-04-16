package main

import (
	"container/heap"
	"sort"
)

type hp []int

func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Less(i, j int) bool {
	a, b := (*h)[i], (*h)[j]
	return a < b
}
func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) {
	heap.Push(h, v)
}
func (h *hp) pop() int {
	return heap.Pop(h).(int)
}

func minProcessingTime(processorTime []int, tasks []int) int {

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})

	hp := new(hp)
	for i := range processorTime {
		hp.push(processorTime[i])
	}
	newTask := make([]int, len(tasks)/4)
	for i := 0; i < len(tasks); i += 4 {
		newTask[i/4] = tasks[i]
	}
	ans := 0
	for i := 0; i < len(newTask); i++ {
		pop := hp.pop()
		ans = max(pop+newTask[i], ans)
	}
	return ans
}
func max(a, b int) int {
	if a > b {

		return a
	}
	return b
}

func main() {
	println(minProcessingTime([]int{10, 20}, []int{
		2, 3, 1, 2, 5, 8, 4, 3}))
}
