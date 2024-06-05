package main

import (
	"slices"
	"sort"
)

/*
*
给你一个下标从 1 开始、长度为 n 的整数数组 nums 。

现定义函数 greaterCount ，使得 greaterCount(arr, val) 返回数组 arr 中 严格大于 val 的元素数量。

你需要使用 n 次操作，将 nums 的所有元素分配到两个数组 arr1 和 arr2 中。在第一次操作中，将 nums[1] 追加到 arr1 。
在第二次操作中，将 nums[2] 追加到 arr2 。之后，在第 i 次操作中：

如果 greaterCount(arr1, nums[i]) > greaterCount(arr2, nums[i]) ，将 nums[i] 追加到 arr1 。
如果 greaterCount(arr1, nums[i]) < greaterCount(arr2, nums[i]) ，将 nums[i] 追加到 arr2 。
如果 greaterCount(arr1, nums[i]) == greaterCount(arr2, nums[i]) ，将 nums[i] 追加到元素数量较少的数组中。
如果仍然相等，那么将 nums[i] 追加到 arr1 。
连接数组 arr1 和 arr2 形成数组 result 。例如，如果 arr1 == [1,2,3] 且 arr2 == [4,5,6] ，那么 result = [1,2,3,4,5,6] 。

返回整数数组 result 。
*/
func resultArray(nums []int) []int {

	newNums := slices.Clone(nums)
	sort.Ints(newNums)

	newNums = slices.Compact(newNums)

	rk := make(map[int]int)

	for i, v := range newNums {
		rk[v] = i + 1
	}

	rkL := len(rk)

	var arr1 []int
	var arr2 []int
	bt1 := bt{sum: make([]int, len(rk)+1), n: len(rk) + 1}
	bt2 := bt{sum: make([]int, len(rk)+1), n: len(rk) + 1}

	for i := range nums {
		rank := rk[nums[i]]
		if i == 0 {
			arr1 = append(arr1, nums[i])
			bt1.update(rank)
		} else if i == 1 {
			arr2 = append(arr2, nums[i])
			bt2.update(rank)
		} else {
			greater1 := bt1.query(rkL) - bt1.query(rank)
			greater2 := bt2.query(rkL) - bt2.query(rank)
			if greater1 > greater2 {
				arr1 = append(arr1, nums[i])
				bt1.update(rank)
			} else if greater1 < greater2 {
				arr2 = append(arr2, nums[i])
				bt2.update(rank)
			} else {
				if len(arr1) <= len(arr2) {
					arr1 = append(arr1, nums[i])
					bt1.update(rank)
				} else {
					arr2 = append(arr2, nums[i])
					bt2.update(rank)
				}
			}
		}

	}
	return append(arr1, arr2...)

}

type bt struct {
	sum []int
	n   int
}

func lowBit(index int) int {
	return index & -index
}

func (b bt) query(index int) int {

	ans := 0

	for index > 0 {
		ans += b.sum[index]
		index -= lowBit(index)
	}
	return ans
}

func (b bt) update(index int) {
	for index < b.n {
		b.sum[index] += 1
		index += lowBit(index)
	}
}
