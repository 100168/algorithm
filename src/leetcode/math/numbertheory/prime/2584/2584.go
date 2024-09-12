package main

/*
给你一个长度为 n 的整数数组 nums ，下标从 0 开始。

如果在下标 i 处 分割 数组，其中 0 <= i <= n - 2 ，使前 i + 1 个元素的乘积和剩余元素的乘积互质，则认为该分割 有效 。

例如，如果 nums = [2, 3, 3] ，那么在下标 i = 0 处的分割有效，因为 2 和 9 互质，而在下标 i = 1 处的分割无效，因为 6 和 3 不互质。在下标 i = 2 处的分割也无效，因为 i == n - 1 。
返回可以有效分割数组的最小下标 i ，如果不存在有效分割，则返回 -1 。

当且仅当 gcd(val1, val2) == 1 成立时，val1 和 val2 这两个值才是互质的，其中 gcd(val1, val2) 表示 val1 和 val2 的最大公约数。
*/
func findValidSplit(nums []int) int {

	n := len(nums)

	gcdMap := make(map[int]int)

	//质因子分解
	for i := 0; i < n; i++ {
		v := nums[i]
		for j := 2; j*j <= v; j++ {
			if v%j == 0 {
				gcdMap[j]++
				for v%j == 0 {
					v /= j
				}
			}
		}
		if v > 1 {
			gcdMap[v]++
		}

	}

	cur := make(map[int]bool)
	for i := 0; i < n-1; i++ {
		v := nums[i]
		for j := 2; j*j <= v; j++ {
			if v%j == 0 {
				cur[j] = true
				if gcdMap[j]--; gcdMap[j] == 0 {
					delete(gcdMap, j)
					delete(cur, j)
				}
				for v%j == 0 {
					v /= j
				}
			}
		}
		if v > 1 {
			cur[v] = true
			if gcdMap[v]--; gcdMap[v] == 0 {
				delete(gcdMap, v)
				delete(cur, v)
			}
		}
		if len(cur) == 0 {
			return i
		}

	}
	return -1

}

func findValidSplit2(nums []int) int {
	left := map[int]int{}           // left[p] 表示质数 p 首次出现的下标
	right := make([]int, len(nums)) // right[i] 表示左端点为 i 的区间的右端点的最大值
	f := func(p, i int) {
		if l, ok := left[p]; ok {
			right[l] = i // 记录左端点 l 对应的右端点的最大值
		} else {
			left[p] = i // 第一次遇到质数 p
		}
	}

	for i, x := range nums {
		for d := 2; d*d <= x; d++ { // 分解质因数
			if x%d == 0 {
				f(d, i)
				for x /= d; x%d == 0; x /= d {
				}
			}
		}
		if x > 1 {
			f(x, i)
		}
	}

	maxR := 0
	for l, r := range right {
		if l > maxR { // 最远可以遇到 maxR
			return maxR // 也可以写 l-1
		}
		maxR = max(maxR, r)
	}
	return -1
}

func findValidSplit3(nums []int) int {

	indexMap := make(map[int]int)

	for i, v := range nums {
		for j := 2; j*j <= v; j++ {
			for ; v%j == 0; v /= j {
				indexMap[j] = i
			}
		}
		if v > 1 {
			indexMap[v] = i
		}
	}

	r := 0

	for i := 0; i <= r; i++ {
		v := nums[i]
		for j := 2; j*j <= v; j++ {
			for ; v%j == 0; v /= j {
				r = max(r, indexMap[j])
			}
		}
		if v > 1 {
			r = max(r, indexMap[v])
		}
	}

	if r == len(nums)-1 {
		return -1
	}
	return r

}
