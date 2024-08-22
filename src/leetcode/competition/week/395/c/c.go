package main

import (
	"fmt"
)

/*
给你两个整数 n 和 x 。你需要构造一个长度为 n 的 正整数 数组 nums ，
对于所有 0 <= i < n - 1 ，满足 nums[i + 1] 大于 nums[i] ，并且数组 nums 中所有元素的按位 AND 运算结果为 x 。

返回 nums[n - 1] 可能的 最小 值。

秒啊
*/
func minEnd(n int, x int) int64 {

	n--

	low := ^x
	j := 0
	for n>>j > 0 {
		//最低位
		l := low & -low
		//0 or 1
		lb := n >> j & 1
		x |= l * lb
		//去掉最低位
		low ^= l
		j++
	}
	return int64(x)
}

func main() {

	fmt.Println(minEnd(1, 1))
}
