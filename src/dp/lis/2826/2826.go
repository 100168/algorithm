package main

func minimumOperations(nums []int) int {

	n := len(nums)
	st := make([]int, 0)

	ans := 0
	for i, v := range nums {

		l, r := 0, len(st)-1

		for l <= r {
			m := (l + r) / 2

			if nums[st[m]] > v {
				r = m - 1
			} else {
				l = m + 1
			}
		}

		if l != len(st) {
			st[l] = i
		} else {
			st = append(st, i)
		}
		ans = max(ans, len(st))
	}
	return n - ans
}
