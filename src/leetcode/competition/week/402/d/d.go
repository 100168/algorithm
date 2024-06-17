package main

import "fmt"

/*
*
数组 arr 中 大于 前面和后面相邻元素的元素被称为 峰值 元素。

给你一个整数数组 nums 和一个二维整数数组 queries 。

你需要处理以下两种类型的操作：

queries[i] = [1, li, ri] ，求出子数组 nums[li..ri] 中 峰值 元素的数目。
queries[i] = [2, indexi, vali] ，将 nums[indexi] 变为 vali 。
请你返回一个数组 answer ，它依次包含每一个第一种操作的答案。

注意：

子数组中 第一个 和 最后一个 元素都 不是 峰值元素。
*/
func countOfPeaks(nums []int, queries [][]int) []int {

	bt := new(bitTree)

	bt.sum = make([]int, len(nums)+1)
	bt.n = len(nums) + 1

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			bt.update(i, 1)
		}
	}
	n := len(nums)
	ans := make([]int, 0)
	for i := range queries {
		op, l, r := queries[i][0], queries[i][1], queries[i][2]
		if op == 1 {
			cur := bt.query(r-1) - bt.query(l)
			ans = append(ans, max(cur, 0))
		} else {
			for j := max(l-1, 1); j <= min(l+1, n-2); j++ {
				if nums[j] > nums[j-1] && nums[j] > nums[j+1] {
					bt.update(j, -1)
				}
			}
			nums[l] = r
			for j := max(l-1, 1); j <= min(l+1, n-2); j++ {
				if nums[j] > nums[j-1] && nums[j] > nums[j+1] {
					bt.update(j, 1)
				}
			}
		}
	}
	return ans
}

type bitTree struct {
	sum []int
	n   int
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

func (bt bitTree) update(index int, value int) {
	for index < bt.n {
		bt.sum[index] += value
		index += lowBit(index)
	}
}

func main() {
	fmt.Println(countOfPeaks([]int{8, 7, 10}, [][]int{{2, 2, 4}, {2, 1, 9}, {1, 0, 2}}))
}
