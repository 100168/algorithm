package main

/*
给定一个整数数组 A，坡是元组 (i, j)，其中  i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。
找出 A 中的坡的最大宽度，如果不存在，返回 0 。
*/
func maxWidthRamp(nums []int) int {

	n := len(nums)
	st := make([]int, n)

	ans := 0
	for i := 0; i < n; i++ {
		l, r := 0, len(st)-1
		for l <= r {
			m := (r + l) / 2
			if nums[st[m]] <= nums[i] {
				r = m - 1
			} else {
				l = m + 1
			}
		}

		if l != len(st) {
			ans = max(ans, i-st[l])
		}
		if nums[st[len(st)-1]] > nums[i] {
			st = append(st, i)
		}
	}
	return ans
}

func main() {
	println(maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
}
