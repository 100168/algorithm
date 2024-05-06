package main

import "fmt"

/*
交换 定义为选中一个数组中的两个 互不相同 的位置并交换二者的值。

环形 数组是一个数组，可以认为 第一个 元素和 最后一个 元素 相邻 。

给你一个 二进制环形 数组 nums ，返回在 任意位置 将数组中的所有 1 聚集在一起需要的最少交换次数。
*/
func minSwaps(nums []int) int {

	cnt := 0
	for _, v := range nums {
		if v == 1 {
			cnt++
		}
	}
	n := len(nums)

	l := 0
	curCnt := 0
	maxCnt := 0
	for r := 0; r < 2*n; r++ {
		if nums[r%n] == 1 {
			curCnt++
		}
		for r-l >= cnt {
			if nums[l%n] == 1 {
				curCnt--
			}
			l++
		}
		maxCnt = max(maxCnt, curCnt)
	}

	return cnt - maxCnt
}

func main() {
	fmt.Println(minSwaps([]int{0, 1, 0, 1, 1, 0, 0}))
}
