package main

func minimumMountainRemovals(nums []int) int {

	n := len(nums)
	maxLeft := getMaxLen(nums, 0, n, 1)
	maxRight := getMaxLen(nums, n-1, -1, -1)
	ans := n
	for i := 0; i < n; i++ {
		if maxLeft[i] > 1 && maxRight[i] > 1 {
			if n-maxLeft[i]-maxRight[i]+1 < ans {
				ans = n - maxLeft[i] - maxRight[i] + 1
			}
		}
	}
	return ans

}

func getMaxLen(nums []int, start int, end int, add int) []int {
	n := len(nums)
	maxLen := make([]int, n)
	maxItem := make([]int, n)

	le := 0

	for i := start; i != end; i += add {
		if le == 0 || nums[i] > maxItem[le-1] {
			maxItem[le] = nums[i]
			le++
			maxLen[i] = le
		} else {
			l := 0
			r := le - 1
			for l < r {
				m := l + r>>1
				if nums[i] > maxItem[m] {
					r = m
				} else {
					l = m + 1
				}
			}

			maxLen[i] = l + 1
			maxItem[l] = nums[i]

		}
	}

	return maxLen

}
