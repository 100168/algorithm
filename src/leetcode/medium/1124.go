package main

import "fmt"

/*
*
给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。

我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。

所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。

请你返回「表现良好时间段」的最大长度。


*/

func longestWPI(hours []int) int {
	n := len(hours)
	st := []int{0}
	sum := make([]int, n+1)
	ans := 0
	for i, v := range hours {
		sum[i+1] = sum[i]
		if v <= 8 {
			sum[i+1] -= 1
		} else {
			sum[i+1] += 1
		}
		if sum[i+1] < sum[st[len(st)-1]] {
			st = append(st, i+1)
		}
	}

	for i := n; i > 0; i-- {
		for len(st) > 0 && sum[i] > sum[st[len(st)-1]] {
			ans = max(ans, i-st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return ans
}

func longestWPI2(hours []int) int {

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
	fmt.Println(longestWPI([]int{6, 6, 9}))
}
