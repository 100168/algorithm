package main

import (
	"container/heap"
	"sort"
)

func smallestDistancePair(nums []int, k int) int {

	n := len(nums)
	h := new(hp)
	h.nums = nums
	sort.Ints(nums)

	for i := 1; i <= min(k, n-1); i++ {
		heap.Push(h, pair{i - 1, i})
	}

	for k > 1 {
		cur := heap.Pop(h).(pair)
		if cur.j+1 < n {
			heap.Push(h, pair{cur.i, cur.j + 1})
		}
		k--

	}
	ans := heap.Pop(h).(pair)

	return nums[ans.j] - nums[ans.i]

}

func smallestDistancePair2(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	l, r := 0, nums[n-1]-nums[0]

	for l <= r {
		m := (r + l) / 2
		if check(m, nums) <= k {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}
func check(m int, nums []int) int {
	n := len(nums)
	i := n - 2
	j := n - 1

	cnt := 1
	for i >= 0 && j >= 0 {
		if nums[j]-nums[i] >= m {
			j--
		} else {
			cnt += j - i
			i--
		}
	}
	return cnt
}

type pair struct{ i, j int }
type hp struct {
	data []pair
	nums []int
}

func (h *hp) Len() int {
	return len(h.data)
}
func (h *hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums[a.j]-h.nums[a.i] < h.nums[b.j]-h.nums[b.i]
}
func (h *hp) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *hp) Push(v any) {
	h.data = append(h.data, v.(pair))
}
func (h *hp) Pop() (v any) {
	a := h.data
	h.data, v = a[:len(a)-1], a[h.Len()-1]
	return v
}

func main() {

	println(smallestDistancePair2([]int{62, 100, 4}, 2))
}
