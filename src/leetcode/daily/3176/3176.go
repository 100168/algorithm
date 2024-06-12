package main

import "fmt"

/*
*
给你一个整数数组 nums 和一个 非负 整数 k 。如果一个整数序列 seq 满足在范围下标范围 [0, seq.length - 2] 中存在 不超过 k 个下标 i 满足 seq[i] != seq[i + 1] ，那么我们称这个整数序列为 好 序列。

请你返回 nums 中 好
子序列

	的最长长度

思路：
1。子序列问题
2.相邻相关
3.相邻无关
4.选或不选。枚举选哪个，枚举选哪个空间复杂度更低
*/
func maximumLength(nums []int, k int) int {

	n := len(nums)

	memo := make([][]int, n+1)

	for i := range memo {
		memo[i] = make([]int, k+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int

	dfs = func(i, rest int) int {
		if i < 0 {
			return 0
		}
		if memo[i][rest] != -1 {
			return memo[i][rest]
		}

		cur := 0
		for j := 0; j < i; j++ {
			if i == n || nums[i] == nums[j] {
				cur = max(cur, dfs(j, rest)+1)
			} else if rest > 0 {
				cur = max(cur, dfs(j, rest-1)+1)
			}
		}
		memo[i][rest] = cur
		return cur
	}
	return dfs(n, k)
}

/*
换一种定义方式
f[x][j]以x结尾最多j对不一样最长子序列长度
转移：
1.f[x][j] = f[x][j]+1   nums[i]==x
2.f[x][j] = max(f[y][j-1]+1)
*/
func maximumLength2(nums []int, k int) int {
	fs := map[int][]int{}
	records := make([]struct{ mx, mx2, num int }, k+1)
	for _, x := range nums {
		if fs[x] == nil {
			fs[x] = make([]int, k+1)
		}
		f := fs[x]
		for j := k; j >= 0; j-- {
			f[j]++
			if j > 0 {
				p := records[j-1]
				m := p.mx
				if x == p.num {
					m = p.mx2
				}
				f[j] = max(f[j], m+1)
			}

			// records[j] 维护 fs[.][j] 的 mx,mx2,num
			v := f[j]
			p := &records[j]
			if v > p.mx {
				if x != p.num {
					p.num = x
					p.mx2 = p.mx
				}
				p.mx = v
			} else if x != p.num && v > p.mx2 {
				p.mx2 = v
			}
		}
	}
	return records[k].mx
}

func maximumLength3(nums []int, k int) int {
	fs := map[int][]int{}
	mx := make([]int, k+2)
	for _, x := range nums {
		if fs[x] == nil {
			fs[x] = make([]int, k+1)
		}
		f := fs[x]
		for j := k; j >= 0; j-- {
			f[j] = max(f[j], mx[j]) + 1
			mx[j+1] = max(mx[j+1], f[j])
		}
	}
	return mx[k+1]
}
func main() {
	fmt.Println(maximumLength([]int{1, 2, 1, 1, 3}, 2))
	fmt.Println(maximumLength([]int{1, 2, 2, 1, 1, 2}, 0))
}
