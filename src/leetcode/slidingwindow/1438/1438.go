package main

func longestSubarray(nums []int, limit int) int {
	ans := 0
	l := 0

	decreaseSt := make([]int, 0)
	increaseSt := make([]int, 0)
	for r, v := range nums {
		for len(increaseSt) > 0 && nums[increaseSt[len(increaseSt)-1]] <= v {
			increaseSt = increaseSt[:len(increaseSt)-1]
		}
		for len(decreaseSt) > 0 && nums[decreaseSt[len(decreaseSt)-1]] >= v {
			decreaseSt = decreaseSt[:len(decreaseSt)-1]
		}
		increaseSt = append(increaseSt, r)
		decreaseSt = append(decreaseSt, r)
		for len(decreaseSt) > 0 && nums[increaseSt[0]]-nums[decreaseSt[0]] > limit {
			if decreaseSt[0] == l {
				decreaseSt = decreaseSt[1:]
			}
			if increaseSt[0] == l {
				increaseSt = increaseSt[1:]
			}
			l++
		}

		ans = max(ans, r-l+1)
	}
	return ans
}
