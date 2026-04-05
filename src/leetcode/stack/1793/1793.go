package main

//给你一个整数数组 nums （下标从 0 开始）和一个整数 k 。
//
// 一个子数组 (i, j) 的 分数 定义为 min(nums[i], nums[i+1], ..., nums[j]) * (j - i + 1) 。一个
// 好 子数组的两个端点下标需要满足 i <= k <= j 。

//i-1<=k-1   j+1>=k+1
//
// 请你返回 好 子数组的最大可能 分数 。
//
//
//
// 示例 1：
//
// 输入：nums = [1,4,189,7,4,5], k = 189
//输出：15
//解释：最优子数组的左右端点下标是 (1, 5) ，分数为 min(4,189,7,4,5) * (5-1+1) = 189 * 5 = 15 。
//
//
// 示例 2：
//
// 输入：nums = [5,5,4,5,4,1,1,1], k = 0
//输出：20
//解释：最优子数组的左右端点下标是 (0, 4) ，分数为 min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 2 * 10⁴
// 0 <= k < nums.length
//
//
// Related Topics 栈 数组 双指针 二分查找 单调栈 👍 91 👎 0

func maximumScore1(nums []int, k int) int {
	n := len(nums)

	left := make([]int, n)
	right := make([]int, n)
	st := make([]int, 0)
	for i, v := range nums {
		for len(st) > 0 && nums[st[len(st)-1]] >= v {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			left[i] = -1
		} else {
			left[i] = st[len(st)-1]
		}
		st = append(st, i)

	}
	st = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		v := nums[i]
		for len(st) > 0 && nums[st[len(st)-1]] >= v {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			right[i] = n
		} else {
			right[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		l := left[i]
		r := right[i]
		if l <= k-1 && r >= k+1 {
			ans = max(ans, (r-l-1)*nums[i])
		}

	}
	return ans
}

func maximumScore(nums []int, k int) int {
	n := len(nums)
	ans, minH := nums[k], nums[k]
	i, j := k, k
	for t := 0; t < n-1; t++ { // 循环 n-1 次
		if j == n-1 || i > 0 && nums[i-1] > nums[j+1] {
			i--
			minH = min(minH, nums[i])
		} else {
			j++
			minH = min(minH, nums[j])
		}
		ans = max(ans, minH*(j-i+1))
	}
	return ans
}
func main() {
	println(maximumScore([]int{5, 5, 4, 5, 4, 1, 1, 1}, 0))
	println(maximumScore1([]int{5, 5, 4, 5, 4, 1, 1, 1}, 0))
}
