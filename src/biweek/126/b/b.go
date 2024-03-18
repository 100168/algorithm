package main

import "container/heap"

/**
 */
func unmarkedSumArray(nums []int, queries [][]int) []int64 {

	sum := int64(0)

	h := new(hp)
	for i := range nums {
		sum += int64(nums[i])
		heap.Push(h, pair{i, nums[i]})
	}
	ans := make([]int64, len(queries))
	for j, v := range queries {
		index, k := v[0], v[1]
		sum -= int64(nums[index])
		nums[index] = 0
		for i := 1; i <= k && h.Len() > 0; i++ {
			cur := heap.Pop(h).(pair)
			for h.Len() > 0 && nums[cur.index] == 0 {
				cur = heap.Pop(h).(pair)
			}
			sum -= int64(nums[cur.index])
			nums[cur.index] = 0
		}
		ans[j] = sum
	}
	return ans
}

type pair struct {
	index int
	val   int
}
type hp []pair

func (h *hp) Less(i, j int) bool {
	if (*h)[i].val == (*h)[j].val {
		return (*h)[i].index < (*h)[j].index
	}
	return (*h)[i].val < (*h)[j].val
}
func (h *hp) Swap(i, j int) {

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}

func main() {
	println(unmarkedSumArray([]int{1, 2, 2, 1, 2, 3, 1}, [][]int{{1, 2}, {3, 3}, {4, 2}}))
}
