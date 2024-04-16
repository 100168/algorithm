package main

func maxSubarrays(nums []int) int {

	n := len(nums)

	minValue := nums[0]
	for i := 0; i < n; i++ {
		minValue = minValue & nums[i]
	}
	ans := 0

	if minValue > 0 {
		return 1
	}

	cur := nums[0]
	for i := 0; i < n; i++ {
		if cur&nums[i] == minValue {
			ans++
			cur = nums[min(i+1, n-1)]
		} else {
			cur = cur & nums[i]
		}

	}
	return ans
}

func main() {
	println(maxSubarrays([]int{22, 21, 29, 22}))
}
