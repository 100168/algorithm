package main

func modifiedMatrix(matrix [][]int) [][]int {

	m := len(matrix)
	n := len(matrix[0])
	maxC := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			maxC[j] = max(maxC[j], matrix[i][j])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if matrix[i][j] == -1 {
				matrix[i][j] = maxC[j]
			}
		}
	}

	return matrix
}
