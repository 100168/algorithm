package main

func longestObstacleCourseAtEachPosition(nums []int) []int {

	n := len(nums)

	left := make([]int, n)

	st := make([]int, 0)

	for i, v := range nums {

		l, r := 0, len(st)-1

		for l <= r {
			m := (r + l) / 2
			if nums[st[m]] > v {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		if l == len(st) {
			st = append(st, i)
		} else {
			st[l] = i
		}
		left[i] = l + 1
	}

	return left
}
