package main

/*
给你一个下标从 0 开始的整数数组 nums ，以及整数 modulo 和整数 k 。

请你找出并统计数组中 趣味子数组 的数目。

如果 子数组 nums[l..r] 满足下述条件，则称其为 趣味子数组 ：

在范围 [l, r] 内，设 cnt 为满足 nums[i] % modulo == k 的索引 i 的数量。并且 cnt % modulo == k 。
以整数形式表示并返回趣味子数组的数目。

注意：子数组是数组中的一个连续非空的元素序列。

(sum[r]%mod-sum[l]%mod+mod)%mod = k

(sum[r] - sum[l]+mod)%mod = k
*/
func countInterestingSubarrays(nums []int, modulo int, k int) int64 {

	ans := int64(0)
	cnt := make(map[int]int64)

	cnt[0] = 1
	s := 0
	for _, num := range nums {
		if num%modulo == k {
			s = (s + 1) % modulo
		}
		ans += cnt[((s-k)+modulo)%modulo]
		cnt[s]++
	}
	return ans
}

func main() {

	nums := []int{7, 5, 2, 1}
	println(countInterestingSubarrays(nums, 1, 0))
}
