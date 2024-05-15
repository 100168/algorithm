package main

import (
	"fmt"
	"math"
)

/*
*
给你一个正整数数组 nums 。

如果数组 nums 的子集中的元素乘积是一个 无平方因子数 ，则认为该子集是一个 无平方 子集。

无平方因子数 是无法被除 1 之外任何平方数整除的数字。

返回数组 nums 中 无平方 且 非空 的子集数目。因为答案可能很大，返回对 109 + 7 取余的结果。

nums 的 非空子集 是可以由删除 nums 中一些元素（可以不删除，但不能全部删除）得到的一个数组。
如果构成两个子集时选择删除的下标不同，则认为这两个子集不同。
*/
func squareFreeSubsets(nums []int) int {

	cntOdd := 0
	cnt := make(map[int]int)

	for i := range nums {
		if nums[i] != 1 && (nums[i]%4 == 0 || nums[i]%9 == 0 || nums[i]%25 == 0) {
			continue
		}
		cnt[nums[i]]++
		if nums[i]%2 != 0 {
			cntOdd++
		}
	}

	oddNums := make([]int, 0)
	evenNums := make([]int, 0)

	for k := range cnt {
		if k%2 != 0 {
			oddNums = append(oddNums, k)
		} else {
			evenNums = append(evenNums, k)
		}
	}
	n := len(oddNums)
	memo := make([]map[int]int, 1<<n)

	for i := range memo {
		memo[i] = make(map[int]int)
	}

	mod := int(1e9 + 7)
	ans := 0
	for i := range evenNums {
	next:
		for j := 0; j < 1<<n; j++ {
			c := evenNums[i]
			res := cnt[c]
			for k := 0; k < n; k++ {
				if 1<<k&j != 0 {
					if gcd(c, oddNums[k]) != 1 {
						continue next
					}
					c *= oddNums[k]
					res = (res * cnt[oddNums[k]]) % mod
				}
			}
			ans += res
			ans %= mod
		}
	}

next2:
	for j := 1; j < 1<<n; j++ {
		c := 1
		res := 1
		for k := 0; k < n; k++ {
			if 1<<k&j != 0 {
				if gcd(c, oddNums[k]) != 1 {
					continue next2
				}
				c *= oddNums[k]
				res = (res * cnt[oddNums[k]]) % mod
			}
		}
		ans += res
		ans %= mod
	}

	return ans
}

func isSqrt(a int) bool {
	sqrt := int(math.Sqrt(float64(a)))
	return sqrt*sqrt == a
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

	fmt.Println(squareFreeSubsets([]int{3, 4, 4, 5}))
	fmt.Println(squareFreeSubsets([]int{1, 2, 2, 11, 11, 11, 5, 6, 3, 3, 30}))
	fmt.Println(squareFreeSubsets([]int{11, 2, 19, 7, 9, 27}))
}
