package main

import (
	"container/heap"
	"fmt"
)

/*
*
给定一个字符串 s ，检查是否能重新排布其中的字母，使得两相邻的字符不同。

返回 s 的任意可能的重新排列。若不可行，返回空字符串 "" 。

示例 1:

输入: s = "aab"
输出: "aba"
示例 2:

输入: s = "aaab"
输出: ""
*/
func reorganizeString(s string) string {

	cnt := make(map[rune]int)

	maxV := 0
	for _, v := range s {
		cnt[v]++
	}

	h := new(hp)
	for i, v := range cnt {

		maxV = max(maxV, v)
		if v > 0 {
			heap.Push(h, pair{i, v})
		}
	}
	if maxV > len(s)-maxV+1 {
		return ""
	}
	ans := ""
	for h.Len() > 0 {
		first := heap.Pop(h).(pair)
		ans += string(first.x)
		if h.Len() > 0 {
			second := heap.Pop(h).(pair)
			if second.v--; second.v > 0 {
				heap.Push(h, second)
			}
			ans += string(second.x)
		}
		if first.v--; first.v > 0 {
			heap.Push(h, first)
		}

	}
	return ans

}

type pair struct {
	x rune
	v int
}
type hp []pair

func (h *hp) Less(i, j int) bool {

	return (*h)[i].v > (*h)[j].v || (*h)[i].v == (*h)[j].v && (*h)[i].x < (*h)[j].x
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
	fmt.Println(reorganizeString("aab"))
}
