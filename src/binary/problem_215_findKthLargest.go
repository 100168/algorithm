package main

import (
	"fmt"
)

//找第k大数
func findKthLargest(nums []int, k int) int {

	return quicklySort(nums, 0, len(nums)-1, k-1)

}

func partition(nums []int, left int, right int) int {

	if left > right {
		return -1
	}
	if left == right {
		return left
	}
	flag := nums[right]
	index := left - 1

	for i := left; i < right; i++ {
		if nums[i] >= flag {
			index++
			swap(index, i, nums)
		}
	}
	index++
	swap(index, right, nums)
	return index

}

func quicklySort(nums []int, l, r, k int) int {

	m := partition(nums, l, r)

	ans := 0
	if m == k {
		ans = nums[m]
	} else if k > m {
		ans = quicklySort(nums, m+1, r, k)
	} else {
		ans = quicklySort(nums, l, m-1, k)

	}
	return ans
}

func swap(a, b int, nums []int) {
	temp := nums[a]
	nums[a] = nums[b]
	nums[b] = temp
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 2, 24, 1, 3, 343, 21, 34, 4242}

	sort := quicklySort(nums, 0, len(nums)-1, 10)

	fmt.Println(nums)
	fmt.Println(sort)

}
