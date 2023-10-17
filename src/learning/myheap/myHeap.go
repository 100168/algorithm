package main

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
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

func main() {
	heap := new(myHeap)
	println(heap.Len())
	heap.Push(1)
	heap.Push(5)
	heap.Push(2)
	heap.Push(12)
	for heap.Len() > 0 {
		println(heap.Pop().(int))
	}

}
