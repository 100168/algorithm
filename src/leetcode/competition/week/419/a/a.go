package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

/*
*给你一个由 n 个整数组成的数组 nums，以及两个整数 k 和 x。

数组的 x-sum 计算按照以下步骤进行：

统计数组中所有元素的出现次数。
仅保留出现次数最多的前 x 个元素的每次出现。如果两个元素的出现次数相同，则数值 较大 的元素被认为出现次数更多。
计算结果数组的和。
注意，如果数组中的不同元素少于 x 个，则其 x-sum 是数组的元素总和。

返回一个长度为 n - k + 1 的整数数组 answer，其中 answer[i] 是
子数组

	nums[i..i + k - 1] 的 x-sum。

子数组 是数组内的一个连续 非空 的元素序列。

示例 1：

输入：nums = [1,1,2,2,3,4,2,3], k = 6, x = 2

输出：[6,10,12]

解释：

对于子数组 [1, 1, 2, 2, 3, 4]，只保留元素 1 和 2。因此，answer[0] = 1 + 1 + 2 + 2。
对于子数组 [1, 2, 2, 3, 4, 2]，只保留元素 2 和 4。因此，answer[1] = 2 + 2 + 2 + 4。注意 4 被保留是因为其数值大于出现其他出现次数相同的元素（3 和 1）。
对于子数组 [2, 2, 3, 4, 2, 3]，只保留元素 2 和 3。因此，answer[2] = 2 + 2 + 2 + 3 + 3。
示例 2：

输入：nums = [3,8,7,8,7,5], k = 2, x = 2

输出：[11,15,15,15,12]

解释：

由于 k == x，answer[i] 等于子数组 nums[i..i + k - 1] 的总和。

提示：

1 <= n == nums.length <= 50
1 <= nums[i] <= 50
1 <= x <= k <= nums.length

1. 计数cntMap  前i~i+k-1个
2. 前x个min(i+k,x) 个总和
3. 更新操作  cntMap[i]--  cntMap[i+k]++


思路：两个有序表
*/

type pair struct {
	v   int
	cnt int
}

func findXSum(nums []int, k int, x int) []int {

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

	ans := make([]int, n-k+1)

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
		ans[i-k] = sum

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
	ans[n-k] = sum
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
