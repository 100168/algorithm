package main

import "fmt"

func minimumMountainRemovals(nums []int) int {

	n := len(nums)

	left := make([]int, n)

	right := make([]int, n)

	st := make([]int, 0)

	for i, v := range nums {

		l, r := 0, len(st)-1

		for l <= r {
			m := (r + l) / 2
			if nums[st[m]] >= v {
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
		left[i] = l
	}
	st = make([]int, 0)

	for i := n - 1; i > 0; i-- {

		l, r := 0, len(st)-1
		for l <= r {
			m := (r + l) / 2
			if nums[st[m]] >= nums[i] {
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
		right[i] = l
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		if left[i] > 0 && right[i] > 0 {
			ans = max(left[i]+right[i]+1, ans)
		}

	}
	return n - ans
}

func main() {

	fmt.Println(minimumMountainRemovals([]int{100, 92, 89, 77, 74, 66, 64, 66, 64}))
}
