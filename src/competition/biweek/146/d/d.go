package main

import (
	"fmt"
	"sort"
)

func subsequencesWithMiddleMode(nums []int) int {

	cnt := make(map[int]int)

	mx := 1
	for _, v := range nums {

		cnt[v]++
		mx = max(mx, cnt[v])
	}

	mod := int(1e9 + 7)

	comb := make([][]int, mx+1)
	for i := range comb {
		comb[i] = make([]int, mx+1)
	}
	comb[0][0] = 1
	for i := 1; i <= mx; i++ {
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % mod
		}
	}

	freq := make([]int, 0)

	for _, v := range cnt {
		freq = append(freq, v)
	}

	sort.Ints(freq)

	var dfs func(int, int, int) int

	f := make([][]map[int]int, len(freq))

	for i := range f {
		f[i] = make([]map[int]int, 6)

		for j := range f[i] {
			f[i][j] = make(map[int]int)
		}
	}

	dfs = func(i int, rest int, maxV int) int {

		if rest == 0 {
			return 1
		}

		if i < 0 {
			return 0
		}

		if v, ok := f[i][rest][maxV]; ok {
			return v
		}

		cur := 0
		for j := 0; j <= min(maxV-1, freq[i], rest); j++ {
			cur = (cur + comb[freq[i]][j]*dfs(i-1, rest-j, maxV)) % mod
		}
		f[i][rest][maxV] = cur
		return cur
	}
	ans := 0

	for i := len(freq) - 1; i >= 0; i-- {
		freq[i], freq[len(freq)-1] = freq[len(freq)-1], freq[i]
		for maxV := min(freq[i], 5); maxV >= 2; maxV-- {
			ans += comb[freq[i]][maxV] * dfs(len(freq)-1, 5-maxV, maxV)
			ans %= mod
		}
		freq[i], freq[len(freq)-1] = freq[len(freq)-1], freq[i]
	}
	return ans

}

func main() {
	//fmt.Println(subsequencesWithMiddleMode([]int{1, 1, 1, 1, 1, 1}))
	fmt.Println(subsequencesWithMiddleMode([]int{1, 2, 2, 3, 3, 4}))
}
