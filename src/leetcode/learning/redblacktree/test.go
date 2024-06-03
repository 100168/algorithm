package main

import (
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

type TreeNode struct {
	Val int
}

func main() {

	tree := rbt.NewWithIntComparator()
	tree.Put(1, 2)
	tree.Put(2, 3)
	left := tree.Left()
	right := tree.Right()
	fmt.Println(left.Value)
	fmt.Println(right.Value)
}
