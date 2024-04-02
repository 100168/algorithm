package main

func numTrees(n int) int {

	cache := make([]int, n+1)

	for i := range cache {
		cache[i] = -1
	}
	var dfs func(int) int

	dfs = func(cnt int) int {
		if cnt <= 1 {
			return 1
		}
		if cache[cnt] != -1 {
			return cache[cnt]
		}
		ans := 0
		for i := 0; i <= cnt-1; i++ {
			ans += dfs(i) * dfs(cnt-1-i)
		}
		cache[cnt] = ans
		return ans
	}
	return dfs(n)
}

func main() {
	println(numTrees(10))
}
