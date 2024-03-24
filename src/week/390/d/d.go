package main

import (
	"math"
)

type trie struct {
	root *node
}
type node struct {
	next []*node
	h    pair
}

type pair struct {
	index int
	len   int
}

func (t *trie) insert(index int, word string) {
	cur := t.root
	n := len(word)
	for i := len(word) - 1; i >= 0; i-- {
		curIndex := word[i] - 'a'
		if cur.next[curIndex] == nil {
			cur.next[curIndex] = new(node)
			cur.next[curIndex].next = make([]*node, 26)
			cur.next[curIndex].h = pair{math.MaxInt, math.MaxInt}
		}
		if cur.h.len > n {
			cur.h = pair{index, n}
		}
		cur = cur.next[curIndex]
	}
	if cur.h.len > n {
		cur.h = pair{index, n}
	}

}

func (t *trie) search(word string) int {
	cur := t.root
	for i := len(word) - 1; i >= 0; i-- {
		if cur.next[word[i]-'a'] == nil {
			break
		}
		cur = cur.next[word[i]-'a']
	}
	return cur.h.index
}
func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	t := new(trie)
	t.root = new(node)
	t.root.h = pair{math.MaxInt, math.MaxInt}
	t.root.next = make([]*node, 26)
	for i := 0; i < len(wordsContainer); i++ {
		t.insert(i, wordsContainer[i])
	}
	ans := make([]int, len(wordsQuery))
	for i, v := range wordsQuery {
		index := t.search(v)
		ans[i] = index

	}
	return ans
}

func main() {
	println(stringIndices([]string{"abcd", "bcd", "xbcd"}, []string{"bcd"}))
	//println(stringIndices([]string{"abcdefgh", "poiuygh", "ghghgh"}, []string{"acbfgh"}))
}
