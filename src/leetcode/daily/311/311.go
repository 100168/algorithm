package main

/*
*
给定两个 稀疏矩阵 ：大小为 m x k 的稀疏矩阵 mat1 和大小为 k x n 的稀疏矩阵 mat2 ，
返回 mat1 x mat2 的结果。你可以假设乘法总是可能的。

c[i][j] =
*/
func multiply(mat1 [][]int, mat2 [][]int) [][]int {

	m := len(mat1)
	n := len(mat2[0])

	ans := make([][]int, m)

	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i, r := range mat1 {
		//把k想成每一列
		// 每一列*每一行，只需要知道当前列需要的每个值
		for k, v := range r {
			if v == 0 {
				continue
			}
			for j, x := range mat2[k] {
				ans[i][j] += v * x
			}
		}
	}
	return ans
}
