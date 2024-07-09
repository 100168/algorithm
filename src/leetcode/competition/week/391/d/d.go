package main

import (
	"fmt"
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"math"
)

func minimumDistance(points [][]int) int {

	lx := redblacktree.New[int, int]()
	ly := redblacktree.New[int, int]()

	for _, v := range points {
		x, y := v[0]+v[1], v[0]-v[1]
		put(lx, x)
		put(ly, y)

	}
	ans := math.MaxInt

	for _, v := range points {
		x, y := v[0]+v[1], v[0]-v[1]
		remove(lx, x)
		remove(ly, y)
		ans = min(ans, max(lx.Right().Key-lx.Left().Key, ly.Right().Key-ly.Left().Key))
		put(lx, x)
		put(ly, y)
	}
	return ans
}

func put(t *redblacktree.Tree[int, int], k int) {
	value, _ := t.Get(k)
	t.Put(k, value+1)

}

func remove(t *redblacktree.Tree[int, int], k int) {

	value, _ := t.Get(k)
	if value == 1 {
		t.Remove(k)
	} else {
		t.Put(k, value-1)
	}
}

func main() {
	fmt.Println(minimumDistance([][]int{{1, 1}, {1, 1}, {1, 1}}))
}
