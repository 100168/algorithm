package main

/*
*
给你一个整数数组 nums 和一个整数 k ，请你返回 nums 中 好 子数组的数目。

一个子数组 arr 如果有 至少 k 对下标 (i, j) 满足 i < j 且 arr[i] == arr[j] ，那么称它是一个 好 子数组。

子数组 是原数组中一段连续 非空 的元素序列。
*/
func countGood(nums []int, k int) int64 {

	cnt := make(map[int]int)

	w := 0

	l := 0
	ans := 0
	for _, v := range nums {
		w += cnt[v]
		cnt[v]++
		for w >= k {
			cnt[nums[l]]--
			w -= cnt[nums[l]]
			l++
		}
		ans += l

	}
	return int64(ans)
}
