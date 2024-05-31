package main

import (
	"fmt"
)

/*
*给你一个由正整数组成的数组 nums 。

数字序列的 最大公约数 定义为序列中所有整数的共有约数中的最大整数。

例如，序列 [4,6,16] 的最大公约数是 2 。
数组的一个 子序列 本质是一个序列，可以通过删除数组中的某些元素（或者不删除）得到。

例如，[2,5,10] 是 [1,2,1,2,4,1,5,10] 的一个子序列。
计算并返回 nums 的所有 非空 子序列中 不同 最大公约数的 数目 。

思路：
1.先按分解因子，然后按因子分组
2.然后再计算单前数跟因子数的最大公约数是否是当前因子
*/
func countDifferentSubsequenceGCDs(nums []int) int {
	cnt := make(map[int]int)

	gcdMap := make(map[int]bool)
	allG := 0
	for _, v := range nums {
		allG = gcd(allG, v)
		for i := 2; i*i <= v; i++ {
			if v%i == 0 {
				if !gcdMap[i] {
					cc := cnt[i]
					g := gcd(cc, v)
					cnt[i] = g
					if g == i {
						gcdMap[i] = true
					}
				}
				if !gcdMap[v/i] && v/i != i {
					cc := cnt[v/i]
					g := gcd(cc, v)
					cnt[v/i] = g
					if g == v/i {
						gcdMap[v/i] = true
					}
				}
			}
		}
		gcdMap[v] = true
	}
	if allG == 1 {
		gcdMap[1] = true
	}
	return len(gcdMap)
}

func gcd(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(countDifferentSubsequenceGCDs([]int{5, 15, 40, 5, 6}))
}
