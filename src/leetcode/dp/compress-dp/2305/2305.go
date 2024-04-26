package main

import (
	"fmt"
	"math"
)

/*
*

给你一个整数数组 cookies ，其中 cookies[i] 表示在第 i 个零食包中的饼干数量。
另给你一个整数 k 表示等待分发零食包的孩子数量，所有 零食包都需要分发。在同一个零食包中的所有饼干都必须分发给同一个孩子，不能分开。

分发的 不公平程度 定义为单个孩子在分发过程中能够获得饼干的最大总数。

返回所有分发的最小不公平程度。
*/
func distributeCookies(cookies []int, k int) int {

	n := len(cookies)

	memo := make([][]int, 1<<n)

	for i := range memo {
		memo[i] = make([]int, k+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(mask, index int) int {
		if mask == 1<<n-1 {
			return 0
		}

		if index == 0 {
			return math.MaxInt
		}

		if memo[mask][index] != -1 {
			return memo[mask][index]
		}
		nextMask := 0
		for i := 0; i < n; i++ {
			if 1<<i&mask == 0 {
				nextMask |= 1 << i
			}
		}

		ans := math.MaxInt
		next := nextMask
		//枚举子集的子集
		for next > 0 {
			cur := 0
			for i := 0; i < n; i++ {
				if 1<<i&next > 0 {
					cur += cookies[i]
				}
			}
			ans = min(ans, max(cur, dfs(mask^next, index-1)))
			next = (next - 1) & nextMask
		}

		memo[mask][index] = ans
		return ans
	}
	return dfs(0, k)
}

func distributeCookies2(cookies []int, k int) int {

	n := len(cookies)

	sum := make([]int, 1<<n)
	for i, v := range cookies {
		for j, k := 0, 1<<i; j < k; j++ {
			sum[j|k] = sum[j] + v
		}
	}

	f := make([]int, 1<<n)
	for i := range f {
		f[i] = sum[i]
	}
	for i := 1; i < k; i++ {
		for mask := 1<<n - 1; mask >= 0; mask-- {
			for sub := mask; sub > 0; sub = (sub - 1) & mask {
				f[mask] = min(f[mask], max(f[mask^sub], sum[sub]))
			}
		}
	}
	return f[1<<n-1]
}

func distributeCookies3(a []int, k int) int {
	n := 1 << len(a)
	sum := make([]int, n)
	for i, v := range a {
		for j, bit := 0, 1<<i; j < bit; j++ {
			sum[bit|j] = sum[j] + v
		}
	}

	f := append([]int{}, sum...)
	for i := 1; i < k; i++ {
		for j := n - 1; j > 0; j-- {
			for s := j; s > 0; s = (s - 1) & j {
				f[j] = min(f[j], max(f[j^s], sum[s]))
			}
		}
	}
	return f[n-1]
}

func main() {
	fmt.Println(distributeCookies([]int{8, 15, 10, 20, 8}, 2))
	fmt.Println(distributeCookies2([]int{8, 15, 10, 20, 8}, 2))
}
