package main

func maximumPrimeDifference(nums []int) int {

	pre := -1
	ans := 0
	for i := 0; i < len(nums); i++ {
		if isF(nums[i]) {
			if pre == -1 {
				pre = i
			} else {
				ans = max(ans, i-pre)
			}
		}
	}
	return ans
}

func isF(x int) bool {

	if x == 1 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	println(maximumPrimeDifference([]int{4, 2, 9, 5, 3}))
}
