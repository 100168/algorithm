package main

import "github.com/emirpasic/gods/trees/redblacktree"

type ExamRoom struct {
	rbt   *redblacktree.Tree
	left  map[int]int
	right map[int]int
	n     int
}

func Constructor(n int) ExamRoom {
	dist := func(s []int) int {
		if s[0] == -1 || s[1] == n {
			return s[1] - s[0] - 1
		}
		return (s[1] - s[0]) >> 1
	}
	cmp := func(a, b any) int {
		x, y := a.([]int), b.([]int)
		d1, d2 := dist(x), dist(y)
		if d1 == d2 {
			return x[0] - y[0]
		}
		return d2 - d1
	}
	this := ExamRoom{redblacktree.NewWith(cmp), map[int]int{}, map[int]int{}, n}
	this.add([]int{-1, n})
	return this
}

func (e *ExamRoom) Seat() int {
	s := e.rbt.Left().Key.([]int)
	p := (s[0] + s[1]) >> 1
	if s[0] == -1 {
		p = 0
	} else if s[1] == e.n {
		p = e.n - 1
	}
	e.del(s)
	e.add([]int{s[0], p})
	e.add([]int{p, s[1]})
	return p
}

func (e *ExamRoom) Leave(p int) {
	l, _ := e.left[p]
	r, _ := e.right[p]
	e.del([]int{l, p})
	e.del([]int{p, r})
	e.add([]int{l, r})
}

func (e *ExamRoom) add(s []int) {
	e.rbt.Put(s, struct{}{})
	e.left[s[1]] = s[0]
	e.right[s[0]] = s[1]
}

func (e *ExamRoom) del(s []int) {
	e.rbt.Remove(s)
	delete(e.left, s[1])
	delete(e.right, s[0])
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
