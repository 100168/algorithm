package main

import "fmt"

/*

给你一个整数数组 nums 。

如果数组 nums 的一个分割满足以下条件，我们称它是一个 美丽 分割：

数组 nums 分为三段
非空子数组
：nums1 ，nums2 和 nums3 ，三个数组 nums1 ，nums2 和 nums3 按顺序连接可以得到 nums 。
子数组 nums1 是子数组 nums2 的
前缀
 或者 nums2 是 nums3 的
前缀
。
请你返回满足以上条件的分割 数目 。

子数组 指的是一个数组里一段连续 非空 的元素。

前缀 指的是一个数组从头开始到中间某个元素结束的子数组。*/

func beautifulSplits(nums []int) int {

	ans := 0

	z0 := zfunc(nums[:])

	n := len(nums)
	for i := 1; i < n-1; i++ {

		z := zfunc(nums[i:])

		for j := i + 1; j < n; j++ {

			if i <= j-i && z0[i] >= i || z[j-i] >= j-i {
				ans++
			}
		}

	}
	return ans

}

func zfunc(nums []int) []int {

	n := len(nums)
	p := make([]int, len(nums))

	p[0] = n

	c, r := 1, 1

	for i := 1; i < n; i++ {
		l := min(p[i-c], max(r-i, 0))
		for i+l < n && nums[i+l] == nums[l] {
			l++
		}
		p[i] = l
		if l > r {
			c = i
			r = l + i
		}

	}
	return p
}

func main() {
	//fmt.Println(beautifulSplits([]int{1, 1, 2, 1}))
	//fmt.Println(beautifulSplits([]int{1, 1, 0, 1, 3, 2, 2, 2}))
	//fmt.Println(beautifulSplits([]int{1, 1, 0, 1, 0}))
	//fmt.Println(beautifulSplits([]int{2, 3, 2, 2, 1}))
	//fmt.Println(beautifulSplits([]int{3, 3, 3, 1, 3}))
	//fmt.Println(beautifulSplits([]int{0, 2, 0, 2, 1, 3, 1, 0}))
	fmt.Println(beautifulSplits([]int{0, 0, 0, 0, 0}))
}
