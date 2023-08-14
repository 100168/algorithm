package main

import (
	"fmt"
	"go-leetcode/src/util"
	"math"
)

func minSubarray(nums []int, p int) int {

	parent := make([]int, 0)

	point := 1

	max := int(math.Pow10(1))

	for point <= max && point > 0 {
		parent = append(parent, point)
		point *= p
	}

	//n := len(nums)
	last := 0
	//for i := 0; i < n; i++ {
	//	cur := last + nums[i]
	//	l := 0
	//	r := len(parent)
	//	for cur > p {
	//		for l < r {
	//			mid := (l + r) >> 1
	//			if parent[mid] > parent {
	//
	//			}
	//		}
	//
	//	}
	//
	//}

	exist := make(map[int]int)
	for i := 1; i < p; i++ {
		exist[i] = -1
	}
	cur := 0

	ans := math.MaxInt
	for i, num := range nums {
		cur = (cur + num) % p
		key := ((cur - last) + p) % p
		if exist[key] != -1 {
			ans = util.Min(ans, i+1-exist[key])
		}
		exist[cur] = i + 1

	}
	return ans
}

func main() {
	nums := []int{6, 3, 5, 2}

	subarray := minSubarray(nums, 9)
	fmt.Println(subarray)
}
