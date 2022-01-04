package main

import "fmt"

/*给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。*/
//双指针
func trap1(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

//动态规划
func trap2(height []int) (ans int) {
	n := len(height)
	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	maxLeft[0] = height[0]

	maxRight[n-1] = height[n-1]
	for i := 1; i < n; i++ {

		//为什么是height[i]而不是height[i-1]
		//因为可以自己减自己就为0
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	}

	for i := range height {
		minH := min(maxLeft[i], maxRight[i])

		ans += minH - height[i]
	}
	return ans
}

//单调栈
func trap3(height []int) (ans int) {
	var stack []int
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWith := i - left - 1
			curHeight := min(h, height[left]) - height[top]

			ans += curWith * curHeight
		}

		stack = append(stack, i)
	}

	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	ans := trap3(a)

	fmt.Println(ans)
}
