package main

/*
*
给你两个数字字符串 num1 和 num2 ，以及两个整数 max_sum 和 min_sum 。如果一个整数 x 满足以下条件，我们称它是一个好整数：

num1 <= x <= num2
min_sum <= digit_sum(x) <= max_sum.
请你返回好整数的数目。答案可能很大，请返回答案对 109 + 7 取余后的结果。

注意，digit_sum(x) 表示 x 各位数字之和。
*/
func count(num1 string, num2 string, min_sum int, max_sum int) int {

	mod := int(1e9 + 7)

	n := len(num2)
	diff := n - len(num1)

	memo := make([]map[int]int, n)

	for i := range memo {
		memo[i] = make(map[int]int)
	}
	var dfs func(int, bool, bool, int) int

	dfs = func(i int, limitLow bool, limitUp bool, sum int) int {
		if i == n {
			if min_sum <= sum && sum <= max_sum {
				return 1
			}
			return 0
		}

		if !limitLow && !limitUp {

			if _, ok := memo[i][sum]; ok {
				return memo[i][sum]
			}
		}

		cur := 0
		low := 0
		up := 9

		if limitLow && i >= diff {
			low = int(num1[i-diff] - '0')
		}

		if limitUp {
			up = int(num2[i] - '0')
		}

		for d := low; d <= up; d++ {
			cur += dfs(i+1, limitLow && d == low, limitUp && d == up, sum+d)
		}

		cur %= mod
		if !limitLow && !limitUp {

			memo[i][sum] = cur

		}

		return cur
	}

	return dfs(0, true, true, 0)
}
