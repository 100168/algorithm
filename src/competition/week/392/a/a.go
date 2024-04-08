package main

func longestMonotonicSubarray(nums []int) int {

	ans := 1
	n := len(nums)
	preMax := 1
	preMin := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			preMax += 1
			preMin = 1
		} else if nums[i] < nums[i-1] {
			preMax = 1
			preMin += 1
		} else {
			preMin = 1
			preMax = 1
		}
		ans = max(ans, preMax, preMin)
	}
	return ans
}
func main() {
	longestMonotonicSubarray([]int{1})
}
