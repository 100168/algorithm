package main

import (
	"fmt"
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

func main() {

	fmt.Println(maximumScore([]int{8, 3, 9, 3, 8}, 2))
}
