package main

func longestNiceSubarray(nums []int) int {

	l, or := 0, 0

	ans := 0
	for right, v := range nums {
		for or&v != 0 {
			or ^= nums[l]
			l++
		}
		or ^= v
		ans = max(ans, right-l+1)

	}
	return ans
}

func main() {
	println(longestNiceSubarray([]int{1, 3, 8, 48, 10}))
}
