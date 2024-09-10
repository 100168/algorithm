package main

import "fmt"

/*
*
给你一个长度为 n 下标从 0 开始的整数数组 nums ，它包含 1 到 n 的所有数字，请你返回上升四元组的数目。

如果一个四元组 (i, j, k, l) 满足以下条件，我们称它是上升的：

0 <= i < j < k < l < n 且
nums[i] < nums[k] < nums[j] < nums[l] 。

示例 1：

输入：nums = [1,3,2,4,5]
输出：2
解释：
- 当 i = 0 ，j = 1 ，k = 2 且 l = 3 时，有 nums[i] < nums[k] < nums[j] < nums[l] 。
- 当 i = 0 ，j = 1 ，k = 2 且 l = 4 时，有 nums[i] < nums[k] < nums[j] < nums[l] 。
没有其他的四元组，所以我们返回 2 。
示例 2：

输入：nums = [1,2,3,4]
输出：0
解释：只存在一个四元组 i = 0 ，j = 1 ，k = 2 ，l = 3 ，但是 nums[j] < nums[k] ，所以我们返回 0 。

提示：

4 <= nums.length <= 4000
1 <= nums[i] <= nums.length
nums 中所有数字 互不相同 ，nums 是一个排列。

我是废物
*/
func countQuadruplets(nums []int) int64 {

	n := len(nums)

	bt := new(bitTree)
	bt.sum = make([]int, n+1)
	ans := 0
	bt.l = n + 1

	for j, v := range nums {
		bt.update(v, 1)
		cnt := 0
		for k := n - 1; k > j; k-- {
			if nums[k] < v {
				less := bt.query(nums[k])
				ans += less * cnt
			} else {
				cnt++
			}
		}
	}
	return int64(ans)

}

type bitTree struct {
	sum []int
	l   int
}

func lowBit(index int) int {
	return index & -index
}

func (bt bitTree) query(index int) int {

	ans := 0
	for index > 0 {
		ans += bt.sum[index]
		index -= lowBit(index)
	}
	return ans
}
func (bt bitTree) update(index int, val int) {
	for index < bt.l {
		bt.sum[index] += val
		index += lowBit(index)
	}
}

func main() {
	fmt.Println(countQuadruplets([]int{1, 3, 2, 4, 5}))
	fmt.Println(countQuadruplets([]int{1, 2, 3, 4}))
}
