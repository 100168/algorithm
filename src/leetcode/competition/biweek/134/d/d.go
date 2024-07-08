package main

/**
给你一个整数数组 nums 和一个整数 k ，请你返回 nums 中有多少个
子数组
满足：子数组中所有元素按位 AND 的结果为 k 。
*/

func countSubarrays(nums []int, k int) int64 {

	type pair struct{ x, cnt int }
	ls := make([]pair, 0)

	ans := 0
	for _, v := range nums {

		ls = append(ls, pair{v, 1})

		nls := ls
		ls = nil

		for _, t := range nls {

			cur := t.x & v
			if cur == k {
				ans += t.cnt
			}
			if len(ls) > 0 && ls[len(ls)-1].x == cur {
				ls[len(ls)-1].cnt += t.cnt
			} else {
				ls = append(ls, pair{cur, t.cnt})
			}
		}

	}
	return int64(ans)

}
