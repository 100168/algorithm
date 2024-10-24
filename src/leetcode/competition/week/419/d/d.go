package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

type pair struct {
	v   int
	cnt int
}

func findXSum(nums []int, k int, x int) []int64 {

	n := len(nums)
	left := redblacktree.NewWith(func(x, y interface{}) int {
		a, b := x.(pair), y.(pair)
		if a.cnt != b.cnt {
			return utils.IntComparator(a.cnt, b.cnt)
		}
		return utils.IntComparator(a.v, b.v)
	})

	right := redblacktree.NewWith(func(x, y interface{}) int {
		a, b := x.(pair), y.(pair)
		if a.cnt != b.cnt {
			return utils.IntComparator(a.cnt, b.cnt)
		}
		return utils.IntComparator(a.v, b.v)
	})

	ans := make([]int64, n-k+1)

	cntMap := make(map[int]int)
	for i := 0; i < k; i++ {
		right.Remove(pair{nums[i], cntMap[nums[i]]})
		cntMap[nums[i]]++
		curPair := pair{nums[i], cntMap[nums[i]]}
		right.Put(curPair, pair{})
	}

	sum := 0
	for left.Size() < x && right.Size() > 0 {
		node := right.Right()
		p := node.Key.(pair)
		sum += p.v * p.cnt
		left.Put(p, pair{})
		right.Remove(p)
	}

	for i := k; i < n; i++ {

		//分两步
		//先删再加
		//
		ans[i-k] = int64(sum)

		if nums[i] == nums[i-k] {
			continue
		}

		delPre := pair{nums[i-k], cntMap[nums[i-k]]}
		cntMap[nums[i-k]]--

		addPre := pair{nums[i], cntMap[nums[i]]}
		cntMap[nums[i]]++

		_, foundDel := right.Get(delPre)
		_, foundAdd := left.Get(addPre)
		if foundDel { //在右边
			put(right, delPre, -1)
		} else { //在左边
			sum -= delPre.v
			put(left, delPre, -1)
		}
		if !foundAdd {
			put(right, addPre, 1)
		} else {
			sum += addPre.v
			put(left, addPre, 1)
		}
		for left.Size() < x && right.Size() > 0 {
			node := right.Right()
			p := node.Key.(pair)
			sum += p.v * p.cnt
			left.Put(p, pair{})
			right.Remove(p)
		}

		if right.Size() > 0 {
			l := left.Left()
			r := right.Right()
			lk := l.Key.(pair)
			rk := r.Key.(pair)
			if cmp(lk, rk) {
				sum -= lk.v * lk.cnt
				sum += rk.v * rk.cnt
				left.Remove(lk)
				right.Remove(rk)
				left.Put(rk, pair{})
				right.Put(lk, pair{})
			}

		}

	}
	ans[n-k] = int64(sum)
	return ans

}

func put(tree *redblacktree.Tree, p pair, v int) {
	tree.Remove(p)
	if p.cnt+v != 0 {
		tree.Put(pair{p.v, p.cnt + v}, pair{})
	}
}

func cmp(pair1, pair2 pair) bool {
	return pair1.cnt < pair2.cnt || pair1.cnt == pair2.cnt && pair1.v < pair2.v
}

func main() {
	//fmt.Println(findXSum([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2))
	//fmt.Println(findXSum([]int{2, 4, 7, 10, 10}, 3, 1))
	//fmt.Println(findXSum([]int{3, 8, 7, 8, 7, 5}, 2, 2))
	fmt.Println(findXSum([]int{2, 5, 3, 5, 3, 5}, 4, 4))
}
