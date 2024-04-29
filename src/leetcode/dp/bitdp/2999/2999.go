package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
*
给你三个整数 start ，finish 和 limit 。同时给你一个下标从 0 开始的字符串 s ，表示一个 正 整数。

如果一个 正 整数 x 末尾部分是 s （换句话说，s 是 x 的 后缀），且 x 中的每个数位至多是 limit ，那么我们称 x 是 强大的 。

请你返回区间 [start..finish] 内强大整数的 总数目 。

如果一个字符串 x 是 y 中某个下标开始（包括 0 ），到下标为 y.length - 1 结束的子字符串，
那么我们称 x 是 y 的一个后缀。比方说，25 是 5125 的一个后缀，但不是 512 的后缀
*/
func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {

	lows := strconv.Itoa(int(start))
	ups := strconv.Itoa(int(finish))
	ls := len(lows)
	lf := len(ups)

	lows = strings.Repeat("0", lf-ls) + lows

	diff := lf - len(s)

	memo := make([]int, lf)

	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int, bool, bool) int

	dfs = func(i int, isLimitLow, isLimitUp bool) int {
		if i == lf {

			return 1
		}

		if !isLimitLow && !isLimitUp && memo[i] != -1 {
			return memo[i]
		}

		low := 0
		up := 9
		if isLimitLow {
			low = int(lows[i] - '0')
		}

		if isLimitUp {
			up = int(ups[i] - '0')
		}

		res := 0

		if i < diff {
			for d := low; d <= min(up, limit); d++ {
				res += dfs(i+1, isLimitLow && d == low, isLimitUp && d == up)
			}
		} else if i >= diff {
			cur := int(s[i-diff] - '0')
			if cur >= low && cur <= up {
				res += dfs(i+1, isLimitLow && cur == low, isLimitUp && cur == up)
			}

		}

		if !isLimitLow && !isLimitUp {
			memo[i] = res
		}

		return res
	}

	return int64(dfs(0, true, true))
}

func main() {
	fmt.Println(numberOfPowerfulInt(15, 215, 6, "10"))
}
