package main

import "sort"

func findOriginalArray(changed []int) []int {

	n := len(changed)

	if n%2 == 1 {
		return []int{}
	}
	ans := make([]int, 0)

	sort.Ints(changed)

	cntMap := make(map[int]int)

	for i := 0; i < n; i++ {
		cntMap[changed[i]]++
	}
	for _, v := range changed {

		if cntMap[v] == 0 {
			continue
		}
		if cntMap[v] > 0 {
			ans = append(ans, v)
		}
		if cntMap[v*2] == 0 {
			return []int{}
		}
		cntMap[v]--
		cntMap[v*2]--
	}
	return ans
}

func findOriginalArray2(changed []int) []int {
	cnt := map[int]int{}
	for _, x := range changed {
		cnt[x]++
	}

	// 单独处理 0
	cnt0 := cnt[0]
	if cnt0%2 == 1 {
		return nil
	}
	delete(cnt, 0)
	ans := make([]int, cnt0/2, len(changed)/2)

	done := map[int]bool{}
	for x := range cnt {
		// 如果 x 已处理完毕，或者 x/2 在 cnt 中，则跳过
		if done[x] || x%2 == 0 && cnt[x/2] > 0 {
			continue
		}
		// 把 x, 2x, 4x, 8x, ... 全部配对
		for cnt[x] > 0 {
			// 每次循环，把 cntX 个 x 和 cntX 个 2x 配对
			cntX := cnt[x]
			// 无法配对，至少要有 cntX 个 2x
			if cntX > cnt[x*2] {
				return nil
			}
			for i := 0; i < cntX; i++ {
				ans = append(ans, x)
			}
			// x 配对完成
			done[x] = true
			if cntX < cnt[x*2] {
				// 还剩下一些 2x
				cnt[x*2] -= cntX
				x *= 2
			} else {
				// 2x 配对完成
				done[x*2] = true
				x *= 4
			}
		}
	}
	return ans
}
