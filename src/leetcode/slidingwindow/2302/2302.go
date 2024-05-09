package main

/*
*
一个数组的 分数 定义为数组之和 乘以 数组的长度。

比方说，[1, 2, 3, 4, 5] 的分数为 (1 + 2 + 3 + 4 + 5) * 5 = 75 。
给你一个正整数数组 nums 和一个整数 k ，请你返回 nums 中分数 严格小于 k 的 非空整数子数组数目。

子数组 是数组中的一个连续元素序列。
*/
func countSubarrays(nums []int, k int64) int64 {

	l := 0
	s := 0
	ans := 0

	for r, v := range nums {
		s += v
		for int64(s*(r-l+1)) >= k {
			s -= nums[l]
			l++
		}
		ans += r - l + 1
	}
	return int64(ans)
}

func main() {
	println(countSubarrays([]int{2, 1, 4, 3, 5}, 10))
}
