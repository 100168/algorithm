package main

/*
*
给你两个整数 n 和 k。

最初，你有一个长度为 n 的整数数组 a，对所有 0 <= i <= n - 1，都有 a[i] = 1 。
每过一秒，你会同时更新每个元素为其前面所有元素的和加上该元素本身。
例如，一秒后，a[0] 保持不变，a[1] 变为 a[0] + a[1]，a[2] 变为 a[0] + a[1] + a[2]，以此类推。

返回 k 秒后 a[n - 1] 的值。

由于答案可能非常大，返回其对 109 + 7 取余 后的结果。
*/
func valueAfterKSeconds(n int, k int) int {

	mod := int(1e9 + 7)
	s := make([]int, n)

	for i := range s {
		s[i] = 1
	}
	for i := 0; i < k; i++ {
		for j := 1; j < n; j++ {

			s[j] = (s[j-1] + s[j]) % mod
		}
	}
	return s[n-1]

}
