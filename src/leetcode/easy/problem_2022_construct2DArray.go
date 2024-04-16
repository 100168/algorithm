package main

func construct2DArray(original []int, m int, n int) [][]int {
	length := len(original)
	if m*n != length {
		return nil
	}

	ans := make([][]int, m)

	index := 0
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if index >= length {
				return nil
			}
			ans[i][j] = original[index]
			index++

		}
	}
	return ans
}
