package main

import "fmt"

func longestWPI(hours []int) int {

	//
	sumMap := make(map[int]int)
	sum := 0
	ans := 0
	sumMap[0] = -1
	for i, hour := range hours {

		if hour > 8 {
			sum += 1
		} else {
			sum += -1
		}
		if sum > 0 {
			ans = i + 1
		}
		if v, ok := sumMap[sum-1]; ok {
			ans = max(ans, i-v)
		}
		if _, ok := sumMap[sum]; !ok {
			sumMap[sum] = i
		}

	}
	return ans
}

func main() {
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
}
