package main

func maxProduct(nums []int, k int, limit int) int {

	n := len(nums)
	f := make([][]map[int][map[int]bool]int, n)

	var dfs func(i int, s int, m int, f int) int

}
