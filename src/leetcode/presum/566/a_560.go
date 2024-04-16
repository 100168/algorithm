package main

/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。

示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2
1。sum[r] - sum[l] = k

2. 核心就是当前位置 i 为子数组最后位置和为k的数量  ==》ans += cnt[(k - sum[i+1])]
*/
func subarraySum(nums []int, k int) int {
	cnt := make(map[int]int)
	ans := 0
	s := 0
	cnt[0] = 1
	for _, num := range nums {
		s += num
		ans += cnt[s-k]
		cnt[s]++
	}
	return ans
}
func main() {

	println(subarraySum([]int{1, 1, 1}, 2))
}
