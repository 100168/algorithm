package main

func maximumSumOfHeights2(maxHeights []int) int64 {
	n := len(maxHeights)
	stack := make([]int, 0)
	stack = append(stack, n)
	sum := make([]int64, n+1)
	curSum := int64(0)
	ans := int64(0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 1 && maxHeights[stack[len(stack)-1]] >= maxHeights[i] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curSum -= int64(maxHeights[pop]) * int64(stack[len(stack)-1]-pop)
		}
		curSum += int64(maxHeights[i]) * int64(stack[len(stack)-1]-i)
		stack = append(stack, i)
		sum[i] = curSum
	}
	stack = stack[:0]
	stack = append(stack, -1)
	curSum = 0
	for i := 0; i < n; i++ {
		for len(stack) > 1 && maxHeights[stack[len(stack)-1]] >= maxHeights[i] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curSum -= int64(maxHeights[pop]) * int64(pop-stack[len(stack)-1])
		}
		curSum += int64(maxHeights[i]) * int64(i-stack[len(stack)-1])
		stack = append(stack, i)
		cur := curSum + sum[i+1]
		if cur > ans {
			ans = cur
		}
	}
	return ans
}

func main() {

	println(maximumSumOfHeights2([]int{5, 3, 4, 1, 1}))
}
