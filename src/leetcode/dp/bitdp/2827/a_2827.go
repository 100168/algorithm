package main

import (
	"strconv"
	"strings"
)

// 给你正整数 low ，high 和 k 。
//
// 如果一个数满足以下两个条件，那么它是 美丽的 ：
//
// 偶数数位的数目与奇数数位的数目相同。
// 这个整数可以被 k 整除。
// 请你返回范围 [low, high] 中美丽整数的数目。
func numberOfBeautifulIntegers(low, high, k int) int {

	lows := strconv.Itoa(low)
	highs := strconv.Itoa(high)

	n1 := len(highs)
	n2 := len(lows)

	padding := n1 - n2
	//填0对齐
	lows = strings.Repeat("0", padding) + lows

	memo := make([][]map[int]int, n1)

	for i := range memo {
		memo[i] = make([]map[int]int, k)
		for j := range memo[i] {
			memo[i][j] = make(map[int]int)
		}
	}

	var dfs func(int, int, int, bool, bool, bool) int

	dfs = func(i int, diff int, rest int, isLimitL bool, isLimitH bool, isNum bool) int {

		if i == n1 {
			if diff == 0 && rest == 0 {
				return 1
			}
			return 0
		}

		if _, ok := memo[i][rest][diff]; ok && !isLimitL && !isLimitH && isNum {
			return memo[i][rest][diff]
		}

		//不选当前数
		res := 0
		if !isNum && i < padding {
			res += dfs(i+1, diff, rest, true, false, false)
		}

		low := 0
		up := 9

		down := 0
		if !isNum {
			down = 1
		}
		if isLimitL && i >= padding {
			low = int(lows[i] - '0')
		}
		if isLimitH {
			up = int(highs[i] - '0')
		}

		for d := max(low, down); d <= up; d++ {
			res += dfs(i+1, diff+(d%2)*2-1, (rest*10+d)%k, isLimitL && d == low, isLimitH && d == up, true)
		}

		if isNum && !isLimitL && !isLimitH {
			memo[i][rest][diff] = res
		}

		return res

		//6366
		//
	}
	return dfs(0, 0, 0, true, true, false)
}

func main() {
	println(numberOfBeautifulIntegers(0, 0, 1))
}
