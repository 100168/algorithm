package main

func countAlternatingSubarrays(nums []int) int64 {

	n := len(nums)
	ans := int64(0)
	pre := 0
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			ans += int64(i - pre + 1)
		} else {
			ans++
			pre = i
		}
	}
	return ans + 1
}

func main() {
	println(countAlternatingSubarrays([]int{0, 1, 1, 1}))
}
