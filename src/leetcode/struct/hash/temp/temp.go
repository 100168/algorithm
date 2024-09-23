package main

import (
	"math"
	"math/rand"
	"slices"
)

// 基于上面的「更快写法」
/**

双模字符串hash*/
func minimumCost(target string, words []string, costs []int) int {
	n := len(target)
	const mod1 = 1_070_777_777
	const mod2 = 1_000_000_007
	base1 := 9e8 - rand.Intn(1e8)
	base2 := 9e8 - rand.Intn(1e8)

	type hPair struct{ h1, h2 int }
	powBase := make([]hPair, n+1)
	preHash := make([]hPair, n+1)
	powBase[0] = hPair{1, 1}
	for i, b := range target {
		powBase[i+1] = hPair{powBase[i].h1 * base1 % mod1, powBase[i].h2 * base2 % mod2}
		preHash[i+1] = hPair{(preHash[i].h1*base1 + int(b)) % mod1, (preHash[i].h2*base2 + int(b)) % mod2}
	}

	// 计算子串 target[l:r] 的哈希值
	// 空串的哈希值为 0
	subHash := func(l, r int) hPair {
		h1 := ((preHash[r].h1-preHash[l].h1*powBase[r-l].h1)%mod1 + mod1) % mod1
		h2 := ((preHash[r].h2-preHash[l].h2*powBase[r-l].h2)%mod2 + mod2) % mod2
		return hPair{h1, h2}
	}

	calcHash := func(t string) (p hPair) {
		for _, b := range t {
			p.h1 = (p.h1*base1 + int(b)) % mod1
			p.h2 = (p.h2*base2 + int(b)) % mod2
		}
		return
	}

	minCost := make([]map[hPair]int, n+1) // [words[i] 的长度][words[i] 的哈希值] -> 最小成本
	lens := map[int]struct{}{}            // 所有 words[i] 的长度集合
	for i, w := range words {
		m := len(w)
		lens[m] = struct{}{}
		h := calcHash(w)
		if minCost[m] == nil {
			minCost[m] = map[hPair]int{}
		}
		if minCost[m][h] == 0 {
			minCost[m][h] = costs[i]
		} else {
			minCost[m][h] = min(minCost[m][h], costs[i])
		}
	}

	// 有 O(√L) 个不同的长度
	sortedLens := make([]int, 0, len(lens))
	for l := range lens {
		sortedLens = append(sortedLens, l)
	}
	slices.Sort(sortedLens)

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, sz := range sortedLens {
			if sz > i {
				break
			}
			if cost, ok := minCost[sz][subHash(i-sz, i)]; ok {
				f[i] = min(f[i], f[i-sz]+cost)
			}
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}
