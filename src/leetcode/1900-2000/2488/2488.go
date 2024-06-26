package main

import "fmt"

/*
*
给你一个长度为 n 的数组 nums ，该数组由从 1 到 n 的 不同 整数组成。另给你一个正整数 k 。

统计并返回 nums 中的 中位数 等于 k 的非空子数组的数目。

注意：

数组的中位数是按 递增 顺序排列后位于 中间 的那个元素，如果数组长度为偶数，则中位数是位于中间靠 左 的那个元素。
例如，[2,3,1,4] 的中位数是 2 ，[8,4,3,5,1] 的中位数是 4 。
子数组是数组中的一个连续部分。

示例 1：

输入：nums = [3,2,1,4,5], k = 4
输出：3
解释：中位数等于 4 的子数组有：[4]、[4,5] 和 [1,4,5] 。
示例 2：

输入：nums = [2,3,1], k = 3
输出：1
解释：[3] 是唯一一个中位数等于 3 的子数组。

nums 中的整数互不相同
*/
func countSubarrays(nums []int, k int) int {

	for i := range nums {
		if nums[i] < k {
			nums[i] = 2
		} else if nums[i] > k {
			nums[i] = -2
		} else {
			nums[i] = 0
		}
	}
	sum := make(map[int]int)

	sum[0] = 1

	var ans = 0
	s := 0
	for _, num := range nums {
		s += num
		ans += sum[s]
		sum[s]++
	}
	return ans
}

func countSubarrays2(nums []int, k int) int {
	pos := 0
	for nums[pos] != k {
		pos++
	}

	// i=pos 的时候 x 是 0，直接记到 cnt 中，这样下面不是大于 k 就是小于 k
	cnt, x := map[int]int{0: 1}, 0
	for i := pos - 1; i >= 0; i-- { // 从 pos-1 开始累加 x
		if nums[i] < k {
			x++
		} else {
			x--
		}
		cnt[x]++
	}

	// i=pos 的时候 x 是 0，直接加到答案中，这样下面不是大于 k 就是小于 k
	ans, x := cnt[0]+cnt[-1], 0
	for _, v := range nums[pos+1:] { // 从 pos+1 开始累加 x
		if v > k {
			x++
		} else {
			x--
		}
		ans += cnt[x] + cnt[x-1]
	}
	return ans
}

func main() {
	fmt.Println(countSubarrays([]int{3, 2, 1, 4, 5}, 4))
	//fmt.Println(countSubarrays([]int{2, 3, 1}, 3))
}
