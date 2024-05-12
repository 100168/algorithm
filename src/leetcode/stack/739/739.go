package main

/*
*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，
其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

1,2,3,4,5
*/
func dailyTemperatures(temperatures []int) []int {

	st := make([]int, 0)
	n := len(temperatures)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		for len(st) > 0 && temperatures[st[len(st)-1]] < temperatures[i] {
			cur := st[len(st)-1]
			ans[st[len(st)-1]] = i - cur
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
