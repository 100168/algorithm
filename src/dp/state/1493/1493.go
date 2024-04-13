package main

func longestSubarray(nums []int) int {
	ans := 0
	pre := 0
	cur := 0
	for _, v := range nums {
		if v == 0 {
			pre = cur
			cur = 0
		} else {
			cur++
		}
		ans = max(pre+cur, ans)
	}
	if ans == len(nums) {
		ans -= 1
	}
	return ans
}
