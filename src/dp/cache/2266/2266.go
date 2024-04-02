package main

func countTexts(pressedKeys string) int {

	mod := 1_000_000_007
	n := len(pressedKeys)

	cache := make([]int, n)
	for i := range cache {
		cache[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 1
		}
		if cache[i] != -1 {
			return cache[i]
		}
		cur := 0
		maxLen := 3
		if pressedKeys[i] == '7' || pressedKeys[i] == '9' {
			maxLen = 4
		}
		for j := i; j >= 0 && i-j+1 <= maxLen; j-- {
			if pressedKeys[j] == pressedKeys[i] {
				cur += dfs(j - 1)
				cur %= mod
			} else {
				break
			}
		}
		cache[i] = cur
		return cur

	}
	return dfs(n - 1)
}

func countTexts2(pressedKeys string) int {

	mod := 1_000_000_007
	n := len(pressedKeys)

	dp := make([]int, n+1)

	dp[0] = 1

	for i := 1; i <= n; i++ {
		maxLen := 3
		if pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9' {
			maxLen = 4
		}
		for j := i; j > 0 && i-j+1 <= maxLen; j-- {
			if pressedKeys[j] == pressedKeys[i] {
				dp[i] += dp[j-1]
				dp[i] %= mod
			} else {
				break
			}
		}
	}

	return dp[n]
}

func main() {
	println(countTexts("22233"))
	println(countTexts2("22233"))
}
