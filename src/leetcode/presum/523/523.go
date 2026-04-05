package main

/*
给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：

子数组大小 至少为 2 ，且
子数组元素总和为 k 的倍数。
如果存在，返回 true ；否则，返回 false 。

如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终视为 k 的一个倍数。

1. (s[r] - s[l])%k==0 ==> s[r]%k=s[l]%k
*/
func checkSubarraySum(nums []int, k int) bool {

	cnt := make(map[int][]int)
	ans := 0
	s := 0
	zero := cnt[0]
	zero = append(zero, 1)
	zero = append(zero, -1)
	cnt[0] = zero

	for i, num := range nums {
		s = (num + s) % k
		if cur, ok := cnt[s]; ok {
			ans += cur[0]
			if cur[1] == i-1 {
				ans--
			}
		}
		cur := cnt[s]
		if cur == nil {
			cur = make([]int, 2)
		}
		cur[0]++
		cur[1] = i
		cnt[s] = cur
		if ans > 1 {
			return true
		}

	}
	return false

}

func main() {

	println(checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
}
