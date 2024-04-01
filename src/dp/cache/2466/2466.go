package main

import "fmt"

func countGoodStrings(low int, high int, zero int, one int) int {

	cache := make([]int, high+1)

	mod := 1_000_000_007

	for i := range cache {
		cache[i] = -1
	}
	var dfs func(int) int

	dfs = func(rest int) int {
		if rest < 0 {
			return 0
		}
		if rest == 0 {
			return 1
		}
		if cache[rest] != -1 {
			return cache[rest]
		}
		cache[rest] = 1 + (dfs(rest-zero)+dfs(rest-one))%mod
		return cache[rest]
	}
	return (dfs(high) - dfs(low-1) + mod) % mod

}

func main() {
	fmt.Println(countGoodStrings(1, 1, 1, 1))
}
