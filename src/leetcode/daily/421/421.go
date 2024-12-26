package main

import (
	"fmt"
	"math/big"
	"math/bits"
)

func findMaximumXOR(nums []int) int {

	trie := &Trie{}
	trie.root = new(Node)
	maxBit := 0

	for i := 0; i < len(nums); i++ {

		maxBit = max(maxBit, bits.Len(uint(nums[i]))-1)
	}

	for _, v := range nums {
		trie.Insert(v, maxBit)
	}

	ans := 0
	for _, v := range nums {
		ans = max(ans, trie.Search(v, maxBit))

	}
	return ans
}

type Trie struct {
	root *Node
}
type Node struct {
	next [2]*Node
}

func (t *Trie) Insert(n int, l int) {

	cur := t.root

	for i := l; i >= 0; i-- {
		c := n >> i & 1
		if cur.next[c] == nil {
			cur.next[c] = &Node{}
		}
		cur = cur.next[c]

	}

}

func (t *Trie) Search(n int, l int) int {
	cur := t.root

	path := 0
	for i := l; i >= 0; i-- {
		c := ((n >> i) & 1) ^ 1
		if cur.next[c] != nil {
			cur = cur.next[c]
			path |= 1 << i
		} else {
			cur = cur.next[c^1]
		}
	}

	big.NewInt(0)
	return path

}

func main() {
	fmt.Println(findMaximumXOR([]int{2, 4, 7, 8}))

	fmt.Println(8 ^ 7)
}
