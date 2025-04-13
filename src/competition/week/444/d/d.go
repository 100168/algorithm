package main

import (
	"container/heap"
)

type Node struct {
	sum   int
	left  *Node
	right *Node
	start int
	end   int
}

type Pair struct {
	sum   int
	left  *Node
	right *Node
}

type hp []*Pair

func (pq hp) Len() int { return len(pq) }

func (pq hp) Less(i, j int) bool {
	if pq[i].sum != pq[j].sum {
		return pq[i].sum < pq[j].sum
	}
	return pq[i].left.start < pq[j].left.start
}

func (pq hp) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *hp) Push(x interface{}) {
	*pq = append(*pq, x.(*Pair))
}

func (pq *hp) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func minimumPairRemoval(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	nodes := make([]*Node, n)
	for i := range nums {
		nodes[i] = &Node{
			sum:   nums[i],
			start: i,
			end:   i,
		}
	}

	for i := 0; i < n-1; i++ {
		nodes[i].right = nodes[i+1]
		nodes[i+1].left = nodes[i]
	}

	count := 0
	for i := 0; i < n-1; i++ {
		if nodes[i].sum > nodes[i+1].sum {
			count++
		}
	}
	if count == 0 {
		return 0
	}

	pq := make(hp, 0)
	heap.Init(&pq)
	for i := 0; i < n-1; i++ {
		left := nodes[i]
		right := nodes[i+1]
		heap.Push(&pq, &Pair{
			sum:   left.sum + right.sum,
			left:  left,
			right: right,
		})
	}

	ops := 0

	for pq.Len() > 0 && count > 0 {
		pair := heap.Pop(&pq).(*Pair)
		left := pair.left
		right := pair.right

		if left.right != right || right.left != left {
			continue
		}

		ops++

		if left.left != nil {
			if left.left.sum > left.sum {
				count--
			}
		}

		if left.sum > right.sum {
			count--
		}

		if right.right != nil {
			if right.sum > right.right.sum {
				count--
			}
		}

		newNode := &Node{
			sum:   left.sum + right.sum,
			left:  left.left,
			right: right.right,
			start: left.start,
			end:   right.end,
		}

		if newNode.left != nil {
			newNode.left.right = newNode
		}
		if newNode.right != nil {
			newNode.right.left = newNode
		}

		if newNode.left != nil {
			if newNode.left.sum > newNode.sum {
				count++
			}
		}
		if newNode.right != nil {
			if newNode.sum > newNode.right.sum {
				count++
			}
		}

		if newNode.left != nil {
			newLeft := newNode.left
			heap.Push(&pq, &Pair{
				sum:   newLeft.sum + newNode.sum,
				left:  newLeft,
				right: newNode,
			})
		}
		if newNode.right != nil {
			newRight := newNode.right
			heap.Push(&pq, &Pair{
				sum:   newNode.sum + newRight.sum,
				left:  newNode,
				right: newRight,
			})
		}
	}

	return ops
}
