package main

import (
	"fmt"
	"sort"
)

/*
*
给你两个正整数数组 nums 和 target ，两个数组长度相等。

在一次操作中，你可以选择两个 不同 的下标 i 和 j ，其中 0 <= i, j < nums.length ，并且：

令 nums[i] = nums[i] + 2 且
令 nums[j] = nums[j] - 2 。
如果两个数组中每个元素出现的频率相等，我们称两个数组是 相似 的。

请你返回将 nums 变得与 target 相似的最少操作次数。测试数据保证 nums 一定能变得与 target 相似。

示例 1：

输入：nums = [8,12,6], target = [2,14,10]

[6,8,14]
[2,10,14]
输出：2
解释：可以用两步操作将 nums 变得与 target 相似：
- 选择 i = 0 和 j = 2 ，nums = [10,12,4] 。
- 选择 i = 1 和 j = 2 ，nums = [10,14,2] 。
2 次操作是最少需要的操作次数。
思路：贪心
1.先从小到大排序
2.从左到右匹配
3.奇数跟奇数匹配，偶数跟偶数匹配
4.使用一个数组来存奇数跟偶数下一个匹配下标
5.只计算》target的值
*/
func makeSimilar(nums []int, target []int) int64 {

	sort.Ints(nums)
	sort.Ints(target)
	index := []int{0, 0}
	ans := 0
	for _, v := range nums {
		p := index[v&1]
		for target[p]&1 != v&1 {
			p++
		}
		index[v&1] = p + 1
		if v >= target[p] {
			ans += (v - target[p]) / 2
		}
	}
	return int64(ans)
}

func f(a []int) {
	for i, x := range a {
		if x&1 > 0 {
			a[i] = -x // 由于元素都是正数，把奇数变成相反数，这样排序后奇偶就自动分开了
		}
	}
	sort.Ints(a)
}

func makeSimilar2(nums, target []int) (ans int64) {
	f(nums)
	f(target)
	for i, x := range nums {
		ans += int64(abs(x - target[i]))
	}
	return ans / 4
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// {1, 2, 5}
// {1, 3, 4}
func main() {
	fmt.Println(makeSimilar([]int{8, 12, 6}, []int{2, 14, 10}))
	fmt.Println(makeSimilar([]int{1, 2, 5}, []int{4, 1, 3}))
}
