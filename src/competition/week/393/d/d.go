package main

import "math"

/*
*
给你两个数组 nums 和 andValues，长度分别为 n 和 m。

数组的 值 等于该数组的 最后一个 元素。

你需要将 nums 划分为 m 个 不相交的连续子数组
，对于第 ith 个子数组 [li, ri]，子数组元素的按位 AND 运算结果等于 andValues[i]，换句话说，
对所有的 1 <= i <= m，nums[li] & nums[li + 1] & ... & nums[ri] == andValues[i] ，其中 & 表示按位 AND 运算符。

返回将 nums 划分为 m 个子数组所能得到的可能的 最小 子数组 值 之和。如果无法完成这样的划分，则返回 -1 。

示例 1：

输入： nums = [1,4,3,3,2], andValues = [0,3,3,2]

输出： 12
解释：

唯一可能的划分方法为：

[1,4] 因为 1 & 4 == 0
[3] 因为单元素子数组的按位 AND 结果就是该元素本身
[3] 因为单元素子数组的按位 AND 结果就是该元素本身
[2] 因为单元素子数组的按位 AND 结果就是该元素本身
这些子数组的值之和为 4 + 3 + 3 + 2 = 12
*/
func minimumValueSum(nums []int, andValues []int) int {
	n := len(nums)
	m := len(andValues)

	cache := make([][]map[int]int, n+1)
	for i := range cache {
		cache[i] = make([]map[int]int, m+1)
		for j := 0; j <= m; j++ {
			cache[i][j] = make(map[int]int)
		}
	}

	var dfs func(int, int, int) int

	dfs = func(i int, addi int, value int) int {

		if i >= n {
			if addi != m {
				return math.MaxInt / 2
			} else {
				return 0
			}
		}

		if _, ok := cache[i][addi][value]; ok {
			return cache[i][addi][value]
		}
		nd := value & nums[i]

		if addi >= m {
			return math.MaxInt / 2
		}

		if nd < andValues[addi] {
			return math.MaxInt / 2
		}

		cur := min(dfs(i+1, addi, nd))
		if nd == andValues[addi] {
			cur = min(cur, dfs(i+1, addi+1, -1)+nums[i])
		}
		cache[i][addi][value] = cur
		return cur
	}
	res := dfs(0, 0, -1)
	if res == math.MaxInt/2 {
		return -1
	}
	return res
}

func main() {
	//println(minimumValueSum([]int{1, 4, 3, 3, 2}, []int{0, 3, 3, 2}))
	println(minimumValueSum([]int{1, 3, 2, 4, 7, 5, 3}, []int{0, 5, 3}))
}
