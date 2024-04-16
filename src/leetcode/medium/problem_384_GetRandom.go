package main

import "math/rand"

type Solution []int

func Constructor(head *ListNode) (s Solution) {
	for node := head; node != nil; node = node.Next {
		s = append(s, node.Val)
	}
	return s
}

func (s Solution) GetRandom() int {
	return s[rand.Intn(len(s))]
}
