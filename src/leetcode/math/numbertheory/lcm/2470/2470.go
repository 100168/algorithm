package main

import "fmt"

/*
*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 nums 的 子数组 中满足 元素最小公倍数为 k 的子数组数目。

子数组 是数组中一个连续非空的元素序列。

数组的最小公倍数 是可被所有数组元素整除的最小正整数。
*/
func subarrayLCM(nums []int, k int) int {

	var pre []int
	ans := 0

	for _, v := range nums {

		pre = append(pre, v)
		temp := pre
		pre = nil
		for _, lcm := range temp {
			g := gcd(lcm, v)
			curLcm := lcm * v / g
			if curLcm == k {
				ans++
			}
			if curLcm <= k {
				pre = append(pre, curLcm)
			}
		}
	}
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

	fmt.Println(subarrayLCM([]int{3, 6, 2, 7, 1}, 6))
}
