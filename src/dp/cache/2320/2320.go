package main

func countHousePlacements(n int) int {

	mod := 1_000_000_007
	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, 2)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) int

	dfs = func(i int, up int) int {
		if i < 0 {
			return 1
		}
		if cache[i][up] != -1 {
			return cache[i][up]
		}
		cur := 0
		cur = (dfs(i-1, up^1) + dfs(i-2, up)) % mod
		cache[i][up] = cur
		return cur
	}
	return dfs(n, 1) + dfs(n, 0)
}

func main() {
	println(countHousePlacements(1))

}
