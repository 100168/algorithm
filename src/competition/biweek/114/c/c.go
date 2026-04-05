package main

func maxSubarrays(nums []int) int {

	minAnd := nums[0]
	for i := range nums {
		minAnd &= nums[i]
	}
	if minAnd > 0 {
		return 1
	}

	ans := 0
	curAnd := nums[0]
	for i, v := range nums {
		curAnd &= v
		if curAnd == 0 {
			ans++
			if i+1 < len(nums) {
				curAnd = nums[i+1]
			}
		}
	}
	return ans
}
func main() {
	println(maxSubarrays([]int{30, 18, 19, 20, 11, 21, 12, 22, 26}))
}
