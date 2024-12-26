package main

import "slices"

/*
*
给你一个整数数组 nums 。

开始时，选择一个满足 nums[curr] == 0 的起始位置 curr ，并选择一个移动 方向 ：向左或者向右。

此后，你需要重复下面的过程：

如果 curr 超过范围 [0, n - 1] ，过程结束。
如果 nums[curr] == 0 ，沿当前方向继续移动：如果向右移，则 递增 curr ；如果向左移，则 递减 curr 。
如果 nums[curr] > 0:
将 nums[curr] 减 1 。
反转 移动方向（向左变向右，反之亦然）。
沿新方向移动一步。
如果在结束整个过程后，nums 中的所有元素都变为 0 ，则认为选出的初始位置和移动方向 有效 。

返回可能的有效选择方案数目。
*/
func countValidSelections(nums []int) int {

	n := len(nums)
	ans := 0
	for i, v := range nums {

		if v == 0 {
			nes := slices.Clone(nums)
			var dfs func(int, bool) bool
			dfs = func(i int, dir bool) bool {

				if i < 0 || i > n-1 {

					for _, t := range nes {
						if t != 0 {
							return false
						}
					}
					return true
				}
				cur := false

				if nes[i] == 0 {
					if dir {
						cur = dfs(i-1, dir)
					} else {
						cur = dfs(i+1, dir)
					}
				} else {
					nes[i]--
					if dir {
						cur = dfs(i+1, !dir)
					} else {
						cur = dfs(i-1, !dir)
					}
				}

				return cur
			}

			if dfs(i, true) {
				ans++
			}

			nes = slices.Clone(nums)
			if dfs(i, false) {
				ans++
			}

		}
	}
	return ans
}
