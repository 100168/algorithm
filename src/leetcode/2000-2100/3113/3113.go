package main

import (
	"fmt"
	"math"
)

/**
给你一个 正 整数数组 nums 。

请你求出 nums 中有多少个子数组，满足子数组中 第一个 和 最后一个 元素都是这个子数组中的 最大 值。

[1,4,3,3,2]
*/

func numberOfSubarrays(nums []int) int64 {

	ans := 0
	type pair struct {
		cnt int
		v   int
	}
	var st []pair

	st = append(st, pair{0, math.MaxInt})

	for _, v := range nums {
		for st[len(st)-1].v < v {
			st = st[:len(st)-1]
		}
		if st[len(st)-1].v == v {
			ans += st[len(st)-1].cnt
			st[len(st)-1].cnt++
		} else {
			st = append(st, pair{1, v})
		}
	}
	return int64(ans + len(nums))
}

func main() {
	//fmt.Println(numberOfSubarrays([]int{1, 4, 3, 3, 2}))
	fmt.Println(numberOfSubarrays([]int{3, 3, 3}))
}
