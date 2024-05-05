package main

import "slices"

func largestComponentSize(nums []int) int {

	n := len(nums)
	parent := make([]int, n)
	size := make([]int, n)

	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	find := func(a int) int {
		for parent[a] != a {
			parent[a] = parent[parent[a]]
			a = parent[a]
		}
		return a
	}

	union := func(a, b int) {
		fa := find(a)
		fb := find(b)
		if fa == fb {
			return
		}

		if size[fa] < size[fb] {
			parent[fa] = fb
			size[fb] += size[fa]
			size[fa] = 0
		} else {
			parent[fb] = fa
			size[fa] += size[fb]
			size[fb] = 0
		}

	}

	maxVal := slices.Max(nums)
	indexNums := make([]int, maxVal+1)

	for i := range indexNums {
		indexNums[i] = -1
	}

	for j, v := range nums {
		for i := 2; i*i <= v; i++ {
			if v%i != 0 {
				continue
			}
			for v%i == 0 {
				v /= i
			}
			if indexNums[i] != -1 {
				union(indexNums[i], j)
			}
			indexNums[i] = j
		}
		if v > 1 {
			if indexNums[v] != -1 {
				union(indexNums[v], j)
			}
			indexNums[v] = j
		}

	}

	return slices.Max(size)
}
