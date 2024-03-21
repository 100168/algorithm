package main

import "container/heap"

//1054. 距离相等的条形码
//尝试过
//中等
//相关标签
//相关企业
//提示
//在一个仓库里，有一排条形码，其中第 i 个条形码为 barcodes[i]。
//
//请你重新排列这些条形码，使其中任意两个相邻的条形码不能相等。 你可以返回任何满足该要求的答案，此题保证存在答案。

type pair struct {
	key int
	val int
}
type hp []pair

func (h *hp) Less(i, j int) bool {
	return (*h)[i].val > (*h)[j].val
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
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
func rearrangeBarcodes(barcodes []int) []int {

	n := len(barcodes)
	cnt := make(map[int]int)
	for i := range barcodes {
		cnt[barcodes[i]]++
	}
	h := new(hp)
	for i := range cnt {
		heap.Push(h, pair{i, cnt[i]})
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		p := heap.Pop(h).(pair)
		k, v := p.key, p.val
		if i == 0 || ans[i-1] != k {
			ans[i] = k
			if v > 0 {
				heap.Push(h, pair{k, v - 1})
			}
		} else {
			p2 := heap.Pop(h).(pair)
			k2, v2 := p2.key, p2.val
			ans[i] = k2
			if v2 > 0 {
				heap.Push(h, pair{k2, v2 - 1})
			}
			heap.Push(h, p)
		}
	}
	return ans
}

func main() {
	print(rearrangeBarcodes([]int{7, 7, 8, 8}))
}
