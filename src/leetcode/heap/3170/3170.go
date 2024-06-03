package main

import "container/heap"

/*
*
给你一个字符串 s 。它可能包含任意数量的 '*' 字符。你的任务是删除所有的 '*' 字符。

当字符串还存在至少一个 '*' 字符时，你可以执行以下操作：

删除最左边的 '*' 字符，同时删除该星号字符左边一个字典序 最小 的字符。如果有多个字典序最小的字符，你可以删除它们中的任意一个。
请你返回删除所有 '*' 字符以后，剩余字符连接而成的
字典序最小

	的字符串。
*/
func clearStars(s string) string {
	h := new(hp)
	del := make(map[int]bool)

	for i := 0; i < len(s); i++ {
		cur := s[i]
		if cur == '*' {
			pop := heap.Pop(h).(pair)
			del[pop.index] = true
			del[i] = true
		} else {
			heap.Push(h, pair{cur, i})
		}
	}
	ans := ""
	for i := 0; i < len(s); i++ {
		if !del[i] {
			ans += string(s[i])
		}
	}
	return ans
}

type pair struct {
	x     uint8
	index int
}

type hp []pair

func (h *hp) Less(i, j int) bool {
	if (*h)[i].x == (*h)[j].x {
		return (*h)[i].index > (*h)[j].index
	}
	return (*h)[i].x < (*h)[j].x
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
