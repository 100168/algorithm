package main

/**
给你一个整数数组 nums 。数组 nums 的 唯一性数组 是一个按元素从小到大排序的数组，包含了 nums 的所有非空子数组中不同元素的个数。

换句话说，这是由所有 0 <= i <= j < nums.length 的 distinct(nums[i..j]) 组成的递增数组。

其中，distinct(nums[i..j]) 表示从下标 i 到下标 j 的子数组中不同元素的数量。

返回 nums 唯一性数组 的 中位数 。

注意，数组的 中位数 定义为有序数组的中间元素。如果有两个中间元素，则取值较小的那个。*/

func medianOfUniquenessArray(nums []int) int {

	n := len(nums)

	s := (n + 1) * n / 2

	k := (s-1)/2 + 1

	l, r := 0, n

	for l <= r {
		m := (r + l) / 2
		if check(nums, m, k) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func check(nums []int, m, k int) bool {

	n := len(nums)
	cnt := 0

	w := make(map[int]int)

	l := 0

	for r := 0; r < n; r++ {
		w[nums[r]]++
		for len(w) > m {
			w[nums[l]]--
			if w[nums[l]] == 0 {
				delete(w, nums[l])
			}
			l++
		}
		cnt += r - l + 1
	}

	return cnt >= k

}
