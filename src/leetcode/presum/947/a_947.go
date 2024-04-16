package main

/*
974. 和可被 K 整除的子数组
相关企业
给定一个整数数组 nums 和一个整数 k ，返回其中元素之和可被 k 整除的（连续、非空） 子数组 的数目。
子数组 是数组的 连续 部分。
示例 1：
输入：nums = [4,5,0,-2,-3,1], k = 5
输出：7
解释：
有 7 个子数组满足其元素之和可被 k = 5 整除：
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
示例 2:

输入: nums = [5], k = 9
输出: 0

1.(sum[r] -sum[l])%k = 0
2.(sum[r]%k - sum[l]%k = 0) ==> sum[r]%k = sum[l]%k
*/
func subarraysDivByK(nums []int, k int) int {

	cnt := make(map[int]int)
	s := 0
	cnt[0] = 1

	ans := 0
	for _, num := range nums {
		s = (((s + num) % k) + k) % k
		ans += cnt[s]
		cnt[s]++

	}
	return ans
}

func main() {

}
