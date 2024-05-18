package main

/*
*
给你一个下标从 0 开始的整数数组 nums ，你可以在一些下标之间遍历。对于两个下标 i 和 j（i != j），
当且仅当 gcd(nums[i], nums[j]) > 1 时，我们可以在两个下标之间通行，其中 gcd 是两个数的 最大公约数 。

你需要判断 nums 数组中 任意 两个满足 i < j 的下标 i 和 j ，是否存在若干次通行可以从 i 遍历到 j 。

如果任意满足条件的下标对都可以遍历，那么返回 true ，否则返回 false 。
*/
func canTraverseAllPairs(nums []int) bool {

	n := len(nums)

	uf := new(unionFind)
	uf.parent = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	primeMap := make(map[int]int)

	for i, v := range nums {

		for j := 2; j*j <= v; j++ {
			if v%j == 0 {
				if _, ok := primeMap[j]; ok {
					uf.union(i, primeMap[j])
				}
				primeMap[j] = i

				for v%j == 0 {
					v /= j
				}
			}
		}
		if v > 1 {
			if _, ok := primeMap[v]; ok {
				uf.union(i, primeMap[v])
			}
			primeMap[v] = i
		}
	}

	for i := 1; i < n; i++ {
		if uf.find(0) != uf.find(i) {
			return false
		}
	}
	return true
}

type unionFind struct {
	parent []int
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf unionFind) union(a, b int) {
	fa := uf.find(a)
	fb := uf.find(b)
	if fa == fb {
		return
	}
	uf.parent[fa] = fb
}
