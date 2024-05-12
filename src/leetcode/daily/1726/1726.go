package main

/*
*

给你一个由 不同 正整数组成的数组 nums ，请你返回满足 a * b = c * d 的元组 (a, b, c, d) 的数量。
其中 a、b、c 和 d 都是 nums 中的元素，且 a != b != c != d 。
*/
func tupleSameProduct(nums []int) int {

	cntMap := make(map[int]int)
	for i, v := range nums {
		for _, v2 := range nums[i+1:] {
			cntMap[v2*v]++
		}
	}

	ans := 0
	for _, v := range cntMap {
		ans += v * (v - 1) / 2
	}
	return ans * 8
}
