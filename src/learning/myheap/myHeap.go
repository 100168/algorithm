package main

import "container/heap"

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func maxKElements(nums []int, k int) int64 {

	hp := new(myHeap)
	for i := range nums {
		heap.Push(hp, nums[i])
	}
	ans := int64(0)
	for hp.Len() > 0 && k > 0 {
		cur := heap.Pop(hp).(int)
		ans += int64(cur)
		heap.Push(hp, (cur+2)/3)
		k--
	}
	return ans
}

func main() {
	println(maxKElements([]int{1, 10, 3, 3, 3}, 3))

}
