package main

import "fmt"

/*
*

给你一个整数数组 nums。

nums 的子序列 sub 的长度为 x ，如果其满足以下条件，则称其为 有效子序列：

(sub[0] + sub[1]) % 2 == (sub[1] + sub[2]) % 2 == ... == (sub[x - 2] + sub[x - 1]) % 2
返回 nums 的 最长的有效子序列 的长度。

一个 子序列 指的是从原数组中删除一些元素（也可以不删除任何元素），剩余元素保持原来顺序组成的新数组。
*/
func maximumLength(nums []int) int {

	n := len(nums)

	cnt := 0
	for i := range nums {
		if nums[i]%2 == 0 {
			cnt++
		}
	}

	one := 0
	zero := 0

	for i := 0; i < n; i++ {
		if nums[i]%2 == 0 {
			zero = one + 1
		} else {
			one = zero + 1
		}
	}
	return max(one, zero, cnt, n-cnt)
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 3, 4}))
}
