package main

import (
	"fmt"
)

/*
*
有一个骰子模拟器会每次投掷的时候生成一个 1 到 6 的随机数。

不过我们在使用它时有个约束，就是使得投掷骰子时，连续 掷出数字 i 的次数不能超过 rollMax[i]（i 从 1 开始编号）。

现在，给你一个整数数组 rollMax 和一个整数 n，请你来计算掷 n 次骰子可得到的不同点数序列的数量。

假如两个序列中至少存在一个元素不同，就认为这两个序列是不同的。由于答案可能很大，所以请返回 模 10^9 + 7 之后的结果。

输入：n = 2, rollMax = [1,1,2,2,2,3]
输出：34
解释：我们掷 2 次骰子，如果没有约束的话，共有 6 * 6 = 36 种可能的组合。但是根据 rollMax 数组，数字 1 和 2 最多连续出现一次，所以不会出现序列 (1,1) 和 (2,2)。因此，最终答案是 36-2 = 34。
示例 2：

输入：n = 2, rollMax = [1,1,1,1,1,1]
输出：30
示例 3：

输入：n = 3, rollMax = [1,1,1,2,2,3]
输出：181
*/
func dieSimulator(n int, rollMax []int) int {

	mod := int(1e9 + 7)

	f := make([][]int, n)

	for i := range f {
		f[i] = make([]int, 6)

		for j := range f[i] {
			f[0][j] = 1
		}
	}
	s := make([]int, n)

	s[0] = 6

	//d[i][j]

	for i := 1; i < n; i++ {
		t := 0
		for j, v := range rollMax {
			res := s[i-1]
			pre := i - v
			if pre > 0 {
				res = res - s[pre-1] + f[pre-1][j]
			} else if pre == 0 {
				res--
			}
			f[i][j] = (res%mod + mod) % mod
			t += f[i][j]
		}
		s[i] = t % mod
	}
	return s[n-1]
}

func dieSimulator3(n int, rollMax []int) (ans int) {
	const mod int = 1e9 + 7
	const m = 6
	cache := make([][m][]int, n)
	for i := range cache {
		for j := range cache[i] {
			cache[i][j] = make([]int, rollMax[j])
			for k := range cache[i][j] {
				cache[i][j][k] = -1 // -1 表示没有访问过
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, last, left int) (res int) {
		if i == 0 {
			return 1
		}
		c := &cache[i][last][left]
		if *c != -1 {
			return *c
		}
		for j, mx := range rollMax {
			if j != last {
				res += dfs(i-1, j, mx-1)
			} else if left > 0 {
				res += dfs(i-1, j, left-1)
			}
		}
		res %= mod
		*c = res
		return
	}
	for j, mx := range rollMax {
		ans += dfs(n-1, j, mx-1)
	}
	return ans % mod
}

func main() {
	fmt.Println(dieSimulator(4, []int{1, 1, 1, 1, 1, 1}))
	fmt.Println(6 * 6 * 6 * 6)
}
