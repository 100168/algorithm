package main

import "fmt"

/*
*
给定一个下标从 0 开始的正整数数组 nums。由三个 不同 索引 (i, j, k) 组成的三元组，
如果 nums[i] + nums[j] + nums[k] 能被 nums[i]、nums[j] 或 nums[k] 中的 一个 整除，则称为 nums 的 单因数三元组。

返回 nums 的单因数三元组。

三元组索引 (0, 3, 4), (0, 4, 3), (3, 0, 4), (3, 4, 0), (4, 0, 3), 和 (4, 3, 0) 的值为 [4, 3, 2] (或者说排列为 [4, 3, 2]).
4 + 3 + 2 = 9 只能被 3 整除，所以所有的三元组都是单因数三元组。
三元组索引 (0, 2, 3), (0, 3, 2), (2, 0, 3), (2, 3, 0), (3, 0, 2), 和 (3, 2, 0) 的值为 [4, 7, 3]  (或者说排列为 [4, 7, 3]).
4 + 7 + 3 = 14 只能被 7 整除，所以所有的三元组都是单因数三元组。
一共有 12 个单因数三元组。
*/
func singleDivisorTriplet(nums []int) int64 {
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}

	ans := 0
	for k1, v1 := range cnt {
		for k2, v2 := range cnt {
			if k1 == k2 {
				v2 -= 1
				if v2 == 0 {
					continue
				}
			}
			for k3, v3 := range cnt {
				s := k1 + k2 + k3
				if k3 == k2 {
					v3 -= 1
				}
				if k3 == k1 {
					v3 -= 1
				}
				if v3 == 0 {
					continue
				}
				v3 = max(v3, 0)
				if s%k1 == 0 && s%k2 != 0 && s%k3 != 0 || s%k2 == 0 && s%k1 != 0 && s%k3 != 0 || s%k3 == 0 && s%k2 != 0 && s%k1 != 0 {
					ans += v1 * v2 * v3
				}
			}
		}
	}
	return int64(ans)

}
func main() {
	fmt.Println(singleDivisorTriplet([]int{4, 6, 7, 3, 2, 2}))
}
