package main

import (
	"fmt"
)

/**

给你一个长度为 n 的整数数组 nums ，n 是 偶数 ，同时给你一个整数 k 。

你可以对数组进行一些操作。每次操作中，你可以将数组中 任一 元素替换为 0 到 k 之间的 任一 整数。

执行完所有操作以后，你需要确保最后得到的数组满足以下条件：

存在一个整数 X ，满足对于所有的 (0 <= i < n) 都有 abs(a[i] - a[n - i - 1]) = X 。
请你返回满足以上条件 最少 修改次数。


示例 1：

输入：nums = [1,0,1,2,4,3], k = 4

//1,0,1
//3 4 2
输出：2

解释：
我们可以执行以下操作：

将 nums[1] 变为 2 ，结果数组为 nums = [1,2,1,2,4,3] 。
将 nums[3] 变为 3 ，结果数组为 nums = [1,2,1,3,4,3] 。
整数 X 为 2 。

输入：nums = [0,1,2,3,3,6,5,4], k = 6

输出：2

解释：
我们可以执行以下操作：

0,1,2,3,
4 5 6 3


将 nums[3] 变为 0 ，结果数组为 nums = [0,1,2,0,3,6,5,4] 。
将 nums[4] 变为 4 ，结果数组为 nums = [0,1,2,0,4,6,5,4] 。
整数 X 为 4 。

思路：
1.先确定当前位置只改一次可以获取的最大区间up,如果要改成[up+1,k]范围则需要修改两次
2.用树状数组来统计[up+1,k] 修改两次的个数
3.计算diff数组并统计出现次数
4.遍历统计次数map并计算最小值，即 cur = n/2-v+bt.query(k)
*/

func minChanges(nums []int, k int) int {

	n := len(nums)
	diff := make([]int, n/2)
	cntMap := make(map[int]int)
	bt := new(bitTree)
	bt.sum = make([]int, k+2)
	bt.l = k + 1
	//nums1 := nums[0 : n/2]
	//nums2 := make([]int, n/2)
	//copy(nums2, nums[n/2:])
	//slices.Reverse(nums2)
	//fmt.Println(nums1)
	//fmt.Println(nums2)
	for i := 0; i < n/2; i++ {
		if nums[i] > nums[n-i-1] {
			nums[n-i-1], nums[i] = nums[i], nums[n-i-1]
		}
		// 3  [0,3]
		// 3  [0,3]
		// [0,nums2[n-i-1]]
		// [0,k-nums[i] ]

		// [max(nums2[n-i-1]+1,k-nums[i]), k]
		// empty =
		upper := max(k-nums[i], nums[n-i-1]) + 1
		diff[i] = abs(nums[i] - nums[n-i-1])
		if nums[i] != 0 {
			bt.update(upper, 1)
		}
		cntMap[diff[i]]++
	}

	//[9 2 7 7 8 9 1]
	//[7 4 9 4 9 1 5]
	//[2 2 2 3 1 8 4]

	ans := n
	for t, v := range cntMap {
		low := bt.query(t)
		curL := n/2 - v + low
		ans = min(ans, curL)
	}

	return ans
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
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(minChanges([]int{0, 1, 2, 3, 3, 6, 5, 4}, 6))
	fmt.Println(minChanges([]int{1, 0, 1, 2, 4, 3}, 4))
	fmt.Println(minChanges([]int{1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0}, 1))
	fmt.Println(minChanges([]int{0, 7, 10, 16, 7, 2, 6, 14, 6, 16, 4, 18, 9, 8}, 19))
	fmt.Println(minChanges([]int{9, 2, 7, 7, 8, 9, 1, 5, 1, 9, 4, 9, 4, 7}, 9))
}
