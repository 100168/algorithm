package main

import "container/heap"

type pair struct {
	id      int
	cnt     int64
	version int
}

type hp []pair

func (h *hp) Less(i, j int) bool {
	return (*h)[i].cnt > (*h)[j].cnt
}
func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}
func (h *hp) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[len(*h)-1]
	return
}

func mostFrequentIDs(nums []int, freq []int) []int64 {

	version := make(map[int]int)
	cnt := make(map[int]int64)
	h := new(hp)

	ans := make([]int64, len(freq))
	for i, v := range freq {
		cnt[nums[i]] += int64(v)
		version[nums[i]]++
		heap.Push(h, pair{id: nums[i], cnt: cnt[nums[i]], version: version[nums[i]]})
		cur := heap.Pop(h).(pair)
		for cur.version < version[cur.id] {
			cur = heap.Pop(h).(pair)
		}
		ans[i] = cur.cnt
		heap.Push(h, cur)
	}
	return ans
}
