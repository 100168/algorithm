package main

import (
	"fmt"
	"math"
	"sort"
)

/*
*
给你一个长度为 n 的正整数数组 nums 和一个整数 k 。

一开始，你的分数为 1 。你可以进行以下操作至多 k 次，目标是使你的分数最大：

选择一个之前没有选过的 非空 子数组 nums[l, ..., r] 。
从 nums[l, ..., r] 里面选择一个 质数分数 最高的元素 x 。如果多个元素质数分数相同且最高，选择下标最小的一个。
将你的分数乘以 x 。
nums[l, ..., r] 表示 nums 中起始下标为 l ，结束下标为 r 的子数组，两个端点都包含。

一个整数的 质数分数 等于 x 不同质因子的数目。比方说， 300 的质数分数为 3 ，因为 300 = 2 * 2 * 3 * 5 * 5 。

请你返回进行至多 k 次操作后，可以得到的 最大分数 。

由于答案可能很大，请你将结果对 109 + 7 取余后返回。

1 <= k <= min(n * (n + 1) / 2, 109)

输入：nums = [8,3,9,3,8], k = 2
输出：81
解释：进行以下操作可以得到分数 81 ：
- 选择子数组 nums[2, ..., 2] 。nums[2] 是子数组中唯一的元素。所以我们将分数乘以 nums[2] ，分数变为 1 * 9 = 9 。
- 选择子数组 nums[2, ..., 3] 。nums[2] 和 nums[3] 质数分数都为 1 ，但是 nums[2] 下标更小。所以我们将分数乘以 nums[2] ，分数变为 9 * 9 = 81 。
81 是可以得到的最高得分。
*/
func maximumScore(nums []int, k int) int {

	n := len(nums)
	mod := int(1e9 + 7)

	type pair struct {
		cnt int
		x   int
	}
	var st []int
	var pairs []pair
	bonus := make([]int, n)
	for i, v := range nums {
		cur := v
		cnt := 0
		for j := 2; j*j <= cur; j++ {
			if cur%j == 0 {
				cnt++
				for cur%j == 0 {
					cur /= j
				}
			}
		}
		if cur > 1 {
			cnt++
		}
		for len(st) > 0 && bonus[st[len(st)-1]] < cnt {
			pop := st[len(st)-1]
			st = st[:len(st)-1]
			pre := -1
			if len(st) > 0 {
				pre = st[len(st)-1]
			}
			diff := (pop - pre) * (i - pop)
			pairs = append(pairs, pair{diff, nums[pop]})
		}
		bonus[i] = cnt
		st = append(st, i)
	}

	for len(st) > 0 {
		pop := st[len(st)-1]
		st = st[:len(st)-1]
		pre := -1
		if len(st) > 0 {
			pre = st[len(st)-1]
		}
		diff := (pop - pre) * (n - pop)
		pairs = append(pairs, pair{diff, nums[pop]})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].x > pairs[j].x
	})

	ans := 1
	for i := 0; i < len(pairs) && k > 0; i++ {
		curX := pairs[i].x
		cnt := pairs[i].cnt
		cnt = min(k, cnt)
		k -= cnt
		for p, t := curX, cnt; t > 0; t >>= 1 {
			if t&1 == 1 {
				ans = ans * p % mod
			}
			p = p * p % mod
		}
	}

	return ans
}

func maximumScore2(nums []int, k int) int {
	n := len(nums)
	cntPrime := make([]int, n)

	mod := int(1e9 + 7)
	for i, v := range nums {

		for j := 2; j*j <= v; j++ {

			if v%j == 0 {
				cntPrime[i]++
			}
			for v%j == 0 {
				v /= j
			}
		}
		if v > 1 {
			cntPrime[i]++
		}
	}
	cntPrime = append(cntPrime, math.MaxInt)
	var st []int

	indexLeft := make([]int, n)

	cntMap := make(map[int]int)
	for i, v := range cntPrime {
		for len(st) > 0 && cntPrime[st[len(st)-1]] < v {
			pop := st[len(st)-1]
			st = st[:len(st)-1]
			cntMap[nums[pop]] += (i - pop) * indexLeft[pop]
		}
		if len(st) > 0 {
			indexLeft[i] = i - st[len(st)-1]
		} else if i < n {
			indexLeft[i] = i + 1
		}
		st = append(st, i)
	}

	type pair struct {
		x, y int
	}
	var pairs []pair

	for t, v := range cntMap {
		pairs = append(pairs, pair{t, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].x > pairs[j].x
	})
	ans := 1
	for _, v := range pairs {
		diff := min(k, v.y)
		k -= diff
		x := v.x
		for ; diff > 0; diff >>= 1 {
			if diff&1 != 0 {
				ans = ans * x % mod
			}
			x = x * x % mod
		}
		if k == 0 {
			break
		}
	}
	return ans
}

func main() {

	//fmt.Println(maximumScore([]int{8, 3, 9, 3, 8}, 2))
	//fmt.Println(maximumScore2([]int{8, 3, 9, 3, 8}, 2))
	fmt.Println(maximumScore2([]int{2, 1, 14, 5, 18, 1, 8, 5}, 32))
	fmt.Println(maximumScore([]int{2, 1, 14, 5, 18, 1, 8, 5}, 32))

	fmt.Println(1 << 14)
}
