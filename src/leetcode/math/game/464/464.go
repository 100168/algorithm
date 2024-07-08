package main

import (
	"fmt"
)

/*
*
在 "100 game" 这个游戏中，两名玩家轮流选择从 1 到 10 的任意整数，累计整数和，先使得累计整数和 达到或超过  100 的玩家，即为胜者。

如果我们将游戏规则改为 “玩家 不能 重复使用整数” 呢？

例如，两个玩家可以轮流从公共整数池中抽取从 1 到 15 的整数（不放回），直到累计整数和 >= 100。

给定两个整数 maxChoosableInteger （整数池中可选择的最大数）和 desiredTotal（累计和），若先出手的玩家能稳赢则返回 true ，
否则返回 false 。假设两位玩家游戏时都表现 最佳 。
*/
func canIWin(maxChoosableInteger int, desiredTotal int) bool {

	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	var dfs func(i int) bool

	f := make([]int, 1<<maxChoosableInteger)

	for i := range f {
		f[i] = -1

	}

	s := make([]int, 1<<maxChoosableInteger)

	for i := 0; i < maxChoosableInteger; i++ {
		for j, k := 0, 1<<i; j < k; j++ {
			s[j|k] = s[j] + i + 1
		}
	}

	dfs = func(i int) bool {
		if f[i] != -1 {
			return f[i] == 1
		}

		cur := false
		for j := 0; j < maxChoosableInteger; j++ {
			if 1<<j&i == 0 && (s[i|1<<j] >= desiredTotal || !dfs(1<<j|i)) {
				cur = true
				break
			}
		}
		if cur {
			f[i] = 1
		} else {
			f[i] = 0
		}
		return cur
	}
	ans := dfs(0)
	return ans
}

func canIWin2(maxChoosableInteger, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	dp := make([]int8, 1<<maxChoosableInteger)
	for i := range dp {
		dp[i] = -1
	}

	s := make([]int, 1<<maxChoosableInteger)

	for i := 0; i < maxChoosableInteger; i++ {
		for j, k := 0, 1<<i; j < k; j++ {
			s[j|k] = s[j] + i + 1
		}
	}
	var dfs func(int) int8
	dfs = func(usedNum int) (res int8) {
		dv := &dp[usedNum]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()

		for i := 0; i < maxChoosableInteger; i++ {
			if usedNum>>i&1 == 0 && (s[usedNum|1<<i] >= desiredTotal || dfs(usedNum|1<<i) == 0) {
				return 1
			}
		}
		return
	}
	return dfs(0) == 1
}

func main() {
	fmt.Println(canIWin(10, 11))
	fmt.Println(canIWin2(10, 11))
}
