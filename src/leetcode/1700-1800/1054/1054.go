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

	cnt := make(map[int]int)
	for i := range barcodes {
		cnt[barcodes[i]]++
	}
	h := new(hp)
	for i := range cnt {
		heap.Push(h, pair{i, cnt[i]})
	}

	ans := make([]int, 0)
	for h.Len() > 0 {
		x1 := heap.Pop(h).(pair)
		if len(ans) > 0 && ans[len(ans)-1] == x1.key {
			x2 := heap.Pop(h).(pair)
			ans = append(ans, x2.key)
			x2.val--
			if x2.val > 0 {
				heap.Push(h, x2)
			}
		} else {
			ans = append(ans, x1.key)
			x1.val--
		}
		if x1.val > 0 {
			heap.Push(h, x1)
		}
	}
	return ans
}

func main() {
	print(rearrangeBarcodes([]int{7, 7, 8, 8}))
}
