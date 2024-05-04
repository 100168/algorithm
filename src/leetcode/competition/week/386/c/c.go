package main

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	m := len(changeIndices)
	l, r := 1, m

	for l <= r {
		t := (r + l) / 2
		if check(changeIndices, t, nums) {
			r = r - 1
		} else {
			l = t + 1
		}
	}
	if l > m {
		return -1
	}
	return l
}

func check(changeIndices []int, t int, nums []int) bool {

	n := len(nums)
	last := make([]int, n)

	for i := range last {
		last[i] = -1
	}

	for i, v := range changeIndices[:t] {
		last[v-1] = i
	}
	for i := range last {
		if last[i] == -1 {
			return false
		}
	}

	cnt := 0

	for i, v := range changeIndices[:t] {

		if last[v-1] == i {
			if cnt < nums[v-1] {
				return false
			}
			cnt -= nums[v-1]
		} else {
			cnt++
		}
	}
	return true
}

func main() {
	println(earliestSecondToMarkIndices([]int{2, 2, 0}, []int{2, 2, 2, 2, 3, 2, 2, 1}))
}
