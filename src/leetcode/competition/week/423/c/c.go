package main

import "fmt"

/*
*
给你一个整数数组 nums。好子序列 的定义是：子序列中任意 两个 连续元素的绝对差 恰好 为 1。

子序列 是指可以通过删除某个数组的部分元素（或不删除）得到的数组，并且不改变剩余元素的顺序。

返回 nums 中所有 可能存在的 好子序列的 元素之和。

因为答案可能非常大，返回结果需要对 109 + 7 取余。

注意，长度为 1 的子序列默认为好子序列。

示例 1：

输入：nums = [1,2,1]

输出：14

解释：

好子序列包括：[1], [2], [1], [1,2], [2,1], [1,2,1]。
这些子序列的元素之和为 14。
*/

func sumOfGoodSubsequences(nums []int) (ans int) {
	const mod = 1_000_000_007
	f := map[int]int{}
	cnt := map[int]int{}
	for _, x := range nums {
		c := cnt[x-1] + cnt[x+1] + 1
		f[x] = (f[x] + f[x-1] + f[x+1] + x*c) % mod
		cnt[x] = (cnt[x] + c) % mod
	}

	for _, s := range f {
		ans += s
	}
	return ans % mod
}

func main() {
	fmt.Println(sumOfGoodSubsequences([]int{1, 2, 1}))
}
