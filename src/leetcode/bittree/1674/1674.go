package main

import (
	"fmt"
)

/**
给你一个长度为 偶数 n 的整数数组 nums 和一个整数 limit 。每一次操作，你可以将 nums 中的任何整数替换为 1 到 limit 之间的另一个整数。

如果对于所有下标 i（下标从 0 开始），nums[i] + nums[n - 1 - i] 都等于同一个数，则数组 nums 是 互补的 。例如，数组 [1,2,3,4] 是互补的，因为对于所有下标 i ，nums[i] + nums[n - 1 - i] = 5 。

返回使数组 互补 的 最少 操作次数。
*/

func minMoves(nums []int, limit int) int {
	return minChanges(nums, limit)
}

func minChanges(nums []int, k int) int {

	n := len(nums)
	diff := make([]int, n/2)
	cntMap := make(map[int]int)
	bt := new(bitTree)
	bt.sum = make([]int, k*2+2)
	bt.l = k*2 + 2
	for i := 0; i < n/2; i++ {
		if nums[i] > nums[n-i-1] {
			nums[n-i-1], nums[i] = nums[i], nums[n-i-1]
		}
		// [1+nums[n-1-i],nums[n-1-i]+k]
		// [nums[i],nums[i]+k]

		upper := nums[n-1-i] + k + 1
		low := nums[i]
		diff[i] = abs(nums[i] + nums[n-i-1])
		bt.update(upper, 1)
		if low > 1 {
			bt.updateRange(2, low, 1)
		}
		cntMap[diff[i]]++
	}

	//limit 22891
	//20744, 7642, 19090, 9992,
	//15721, 3458, 16848, 2457,
	//36465, 11100,35938 ,12499

	ans := n
	for t, v := range cntMap {
		low := bt.query(t)
		curL := n/2 - v + low
		ans = min(ans, curL)
	}
	return min(ans, n/2)
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

func (bt bitTree) updateRange(l, r int, val int) {
	bt.update(l, val)
	bt.update(r+1, -val)
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
	//fmt.Println(minMoves([]int{1, 2, 2, 1}, 2))
	fmt.Println(minMoves([]int{20744, 7642, 19090, 9992, 2457, 16848, 3458, 15721}, 22891))
}
