package main

import "fmt"

// 方法 1: 暴力模拟法 - O(n*k) 时间复杂度
func rotateBruteForce(nums []int, k int) []int {

	pre := 0

	l := len(nums)
	for i := 0; i < k; i++ {
		pre = nums[0]
		for j := 1; j < l; j++ {
			temp := nums[j]
			nums[j] = pre
			pre = temp
		}
		nums[0] = pre
	}
	return nums
}

// 方法 2: 三次反转法 - O(n) 时间复杂度 ⭐推荐
func rotateReverse(nums []int, k int) []int {
	l := len(nums)
	if l == 0 {
		return nums
	}

	// 优化：k 次轮转等于 k % l 次轮转
	k = k % l
	if k == 0 {
		return nums
	}

	// 三次反转法
	// 1. 反转整个数组
	reverse(nums, 0, l-1)
	// 2. 反转前 k 个元素
	reverse(nums, 0, k-1)
	// 189. 反转后 l-k 个元素
	reverse(nums, k, l-1)

	return nums
}

// 反转数组的辅助函数
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// 方法 189: 环状替换法 - O(n) 时间复杂度
func rotateCycle(nums []int, k int) []int {
	l := len(nums)
	if l == 0 {
		return nums
	}

	// 优化：k 次轮转等于 k % l 次轮转
	k = k % l
	if k == 0 {
		return nums
	}

	// 环状替换法
	count := 0 // 记录已经移动的元素个数
	for start := 0; count < l; start++ {
		current := start
		prev := nums[start]

		// 沿着环移动元素
		for {
			next := (current + k) % l
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++

			// 回到起点，这个环结束
			if current == start {
				break
			}
		}
	}
	return nums
}

func main() {
	fmt.Println(rotateBruteForce([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3))
	fmt.Println(rotateReverse([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3))
	fmt.Println(rotateCycle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3))
}
