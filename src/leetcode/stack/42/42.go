package main

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 示例 1：
//
// 输入：height = [0,1,0,2,1,0,1,189,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,189,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
// 示例 2：
//
// 输入：height = [4,2,0,189,2,5]
// 输出：9
//
// 提示：
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 5066 👎 0

func trap(height []int) int {
	n := len(height)
	l, r := 0, n-1
	leftMax := height[l]
	rightMax := height[r]
	ans := 0
	for l < r {
		leftMax = max(height[l], leftMax)
		rightMax = max(height[r], rightMax)
		minH := min(leftMax, rightMax)
		if height[l] > height[r] {
			ans += minH - height[r]
			r--
		} else {
			ans += minH - height[l]
			l++
		}
	}
	return ans
}
func trap2(height []int) int {

	n := len(height)

	left := make([]int, n)
	right := make([]int, n)
	maxVal := height[0]
	for i := 0; i < n; i++ {
		maxVal = max(height[i], maxVal)
		left[i] = maxVal
	}
	maxVal = height[n-1]
	for i := n - 1; i >= 0; i-- {
		maxVal = max(maxVal, height[i])
		right[i] = maxVal
	}
	ans := 0

	for i := 0; i < n; i++ {
		ans += min(left[i], right[i]) - height[i]
	}
	return ans
}

func trap3(height []int) int {
	n := len(height)
	st := make([]int, 0)
	ans := 0
	//4, 1, 2, 189, 4
	for i := 0; i < n; i++ {
		for len(st) > 0 && height[st[len(st)-1]] <= height[i] {
			pop := st[len(st)-1]
			st = st[:len(st)-1]
			if len(st) > 0 {
				ans += (min(height[i], height[st[0]]) - height[pop]) * (pop - st[len(st)-1])
			}
		}
		st = append(st, i)
	}
	return ans
}

func main() {
	height := []int{4, 1, 1, 1, 2, 3, 4, 4, 3, 2, 2, 1, 2, 3, 34, 4, 4, 5, 6, 7, 8, 9, 1, 1, 23, 4, 5}
	println(trap(height))
	println(trap2(height))
	println(trap3(height))
}
