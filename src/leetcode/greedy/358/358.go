package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/**
给你一个非空的字符串 s 和一个整数 k ，你要将这个字符串 s 中的字母进行重新排列，使得重排后的字符串中相同字母的位置间隔距离 至少 为 k 。
如果无法做到，请返回一个空字符串 ""。



示例 1：

输入: s = "aabbcc", k = 3
输出: "abcabc"
解释: 相同的字母在新的字符串中间隔至少 3 个单位距离。
示例 2:

输入: s = "aaabc", k = 3
输出: ""
解释: 没有办法找到可能的重排结果。
示例 3:

输入: s = "aaadbbcc", k = 2
输出: "abacabcd"
解释: 相同的字母在新的字符串中间隔至少 2 个单位距离。


*/

func rearrangeString(s string, k int) string {

	if k == 0 || k == 1 {
		return s
	}

	cnt := make(map[rune]int)

	type pair struct {
		x rune
		v int
	}

	n := len(s)

	for _, v := range s {
		cnt[v]++
	}
	pairs := make([]pair, 0)

	for i, v := range cnt {

		pairs = append(pairs, pair{i, v})

	}
	ans := make([]byte, n)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].v > pairs[j].v
	})

	p := make([]int, k)

	for i := range p {
		p[i] = i
	}

	cur := 0
	for _, v := range pairs {
		for ; v.v > 0; v.v-- {
			if p[cur] >= n {
				cur = (cur + 1) % k
			}
			ans[p[cur]] = byte(v.x)
			p[cur] += k
		}
	}

	wc := make(map[byte]int)
	l := 0
	for i, v := range ans {

		wc[v]++
		if wc[v] > 1 {
			return ""
		}
		if i-l+1 >= k {
			wc[ans[l]]--
			l++
		}

	}

	return string(ans)
}

func rearrangeString2(s string, k int) string {

	if k == 0 || k == 1 {
		return s
	}

	cnt := make(map[rune]int)

	for _, v := range s {
		cnt[v]++
	}

	h := new(hp)

	for i, v := range cnt {
		heap.Push(h, pair{i, v})
	}
	ans := make([]byte, 0)

	for h.Len() > 0 {
		first := heap.Pop(h).(pair)
		if first.v > 1 && h.Len() < k-1 {
			return ""
		}
		first.v--

		ans = append(ans, byte(first.x))

		curPair := make([]pair, 0)

		if first.v > 0 {
			curPair = append(curPair, first)
		}
		for i := 0; i < k-1 && h.Len() > 0; i++ {
			c := heap.Pop(h).(pair)
			c.v--
			if c.v > 0 {
				curPair = append(curPair, c)
			}
			ans = append(ans, byte(c.x))
		}

		for _, v := range curPair {
			heap.Push(h, v)
		}
	}

	return string(ans)
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
	//fmt.Println(rearrangeString2("aabbcc", 3))
	//fmt.Println(rearrangeString2("aaabc", 3))
	//fmt.Println(rearrangeString2("aaadbbcc", 2))
	//fmt.Println(rearrangeString2("abb", 2))
	//fmt.Println(rearrangeString2("aabbcc", 4))
	//fmt.Println(rearrangeString2("abcabcabcd", 4))
	//fmt.Println(rearrangeString2("aa", 2))
	fmt.Println(rearrangeString2("abcdabcdabdeac", 4))
}
