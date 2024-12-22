package main

import (
	"fmt"
	"slices"
)

/**
给你一个整数数组 nums 和一个 正 整数 k 。

定义长度为 2 * x 的序列 seq 的 值 为：

(seq[0] OR seq[1] OR ... OR seq[x - 1]) XOR (seq[x] OR seq[x + 1] OR ... OR seq[2 * x - 1]).
请你求出 nums 中所有长度为 2 * k 的
子序列
 的 最大值 。



示例 1：

输入：nums = [2,6,7], k = 1

输出：5

解释：

子序列 [2, 7] 的值最大，为 2 XOR 7 = 5 。




*/

func maxValue(nums []int, k int) int {

	n := len(nums)

	f1 := make([][][]bool, n)
	f2 := make([][][]bool, n)

	for i := range f1 {

		f1[i] = make([][]bool, k+1)
		for j := range f1[i] {
			f1[i][j] = make([]bool, 1<<7)
		}
	}

	f1[0][1][nums[0]] = true
	for i := range f2 {
		f2[i] = make([][]bool, k+1)
		for j := range f2[i] {
			f2[i][j] = make([]bool, 1<<7)
		}
	}

	f2[n-1][1][nums[n-1]] = true

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			f1[i][j] = slices.Clone(f1[i-1][j])
			for v, c := range f1[i-1][j-1] {
				if c {
					f1[i][j][v|nums[i]] = true
				}
			}
			if j == 1 {
				f1[i][j][nums[i]] = true
			}
		}
	}

	for i := n - 2; i >= 0; i-- {
		for j := 1; j <= k; j++ {
			f2[i][j] = slices.Clone(f2[i+1][j])
			for v, c := range f2[i+1][j-1] {
				if c {
					f2[i][j][v|nums[i]] = true
				}

			}
			if j == 1 {
				f2[i][j][nums[i]] = true
			}
		}
	}

	ans := 0
	for i := k - 1; i < n-k; i++ {
		for l, lv := range f1[i][k] {
			for r, rv := range f2[i+1][k] {
				if lv && rv {
					ans = max(ans, l^r)
				}

			}
		}
	}
	return ans
}

func main() {
	//fmt.Println(maxValue([]int{2, 42}, 1))
	//fmt.Println(maxValue([]int{1, 89, 11, 90}, 2))
	//fmt.Println(maxValue([]int{2, 6, 7}, 1))
	fmt.Println(maxValue([]int{8, 114, 123, 82}, 1))
}
