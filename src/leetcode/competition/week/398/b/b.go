package main

func isArraySpecial(nums []int, queries [][]int) []bool {

	n := len(nums)
	sp := make([]int, n-1)
	for i := 1; i < n; i++ {
		if nums[i]%2 == nums[i-1]%2 {
			sp[i-1] = 1
		}
	}
	s := make([]int, n)
	for i := 1; i < n; i++ {
		s[i] = s[i-1] + sp[i-1]
	}
	ans := make([]bool, len(queries))
	for i, v := range queries {
		start := v[0]
		end := v[1]
		if s[end]-s[start] == 0 {
			ans[i] = true
		}
	}
	return ans
}
