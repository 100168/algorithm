package main

func kthSmallest(matrix [][]int, k int) int {

	m := len(matrix)
	n := len(matrix)
	l, r := matrix[0][0], matrix[m-1][n-1]
	for l <= r {
		m := (r + l) / 2
		if check(m, matrix) <= k {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r

}

func check(t int, matrix [][]int) int {

	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := n - 1

	cnt := 1

	for i < m && j >= 0 {
		if matrix[i][j] >= t {
			j--
		} else {
			cnt += j + 1
			i++
		}
	}

	return cnt
}

func main() {
	println(kthSmallest([][]int{{1, 2}, {3, 3}}, 1))
}
