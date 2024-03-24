package main

func maxSlidingWindow(nums []int, k int) []int {

	ans := make([]int, 0)
	w := make([]int, 0)
	n := len(nums)

	for i := 0; i < n; i++ {
		for len(w) > 0 && nums[w[len(w)-1]] <= nums[i] {
			w = w[:len(w)-1]
		}
		w = append(w, i)

		if i-w[0]+1 > k {
			w = w[1:]
		}
		if i+1 >= k {
			ans = append(ans, nums[w[0]])
		}
	}
	return ans
}
