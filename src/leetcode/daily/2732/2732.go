package main

/*
*
给你一个下标从 0 开始大小为 m x n 的二进制矩阵 grid 。

从原矩阵中选出若干行构成一个行的 非空 子集，如果子集中任何一列的和至多为子集大小的一半，那么我们称这个子集是 好子集。

更正式的，如果选出来的行子集大小（即行的数量）为 k，那么每一列的和至多为 floor(k / 2) 。

请你返回一个整数数组，它包含好子集的行下标，请你将子集中的元素 升序 返回。

如果有多个好子集，你可以返回任意一个。如果没有好子集，请你返回一个空数组。

一个矩阵 grid 的行 子集 ，是删除 grid 中某些（也可能不删除）行后，剩余行构成的元素集合。

输入：grid =
[[0,1,1,0]
,[0,0,0,1]
,[1,1,1,1]]
输出：[0,1]
解释：我们可以选择第 0 和第 1 行构成一个好子集。
选出来的子集大小为 2 。
- 第 0 列的和为 0 + 0 = 0 ，小于等于子集大小的一半。
- 第 1 列的和为 1 + 0 = 1 ，小于等于子集大小的一半。
- 第 2 列的和为 1 + 0 = 1 ，小于等于子集大小的一半。
- 第 3 列的和为 0 + 1 = 1 ，小于等于子集大小的一半。
*/
func goodSubsetofBinaryMatrix(grid [][]int) []int {
	maskToIdx := map[int]int{}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 {
			return []int{i}
		}
		maskToIdx[mask] = i
	}

	for x, i := range maskToIdx {
		for y, j := range maskToIdx {
			if x&y == 0 {
				return []int{min(i, j), max(i, j)}
			}
		}
	}
	return nil
}

/*
*SOSDP
 */
func goodSubsetofBinaryMatrix2(grid [][]int) []int {
	n := len(grid[0])
	f := make([]int, 1<<n)
	for i := range f {
		f[i] = -1
	}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 {
			return []int{i}
		}
		f[mask] = i
	}

	u := 1<<n - 1
	for s := 1; s < u; s++ {
		for b := 0; b < n; b++ {
			if 1<<b&s == 0 {
				continue
			}
			f[s] = max(f[s], f[s^1<<b])
			i := f[s]
			if i < 0 {
				continue
			}
			j := f[u^s]
			if j >= 0 {
				return []int{min(i, j), max(i, j)}
			}
		}
	}
	return nil
}
