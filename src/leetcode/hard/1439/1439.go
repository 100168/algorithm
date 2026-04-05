package main

type Info struct {
	m   int
	n   int
	cnt int
	k   int
	mat [][]int
}

func kthSmallest(mat [][]int, k int) int {

	m := len(mat)
	n := len(mat[0])
	l := 0
	r := 0

	for i := 0; i < m; i++ {
		l += mat[i][0]
		r += mat[i][n-1]
	}
	var info *Info
	info = new(Info)
	info.k = k
	info.mat = mat
	info.n = n
	info.m = m
	for l < r {

		mid := (r + l) >> 1
		info.cnt = 1
		dfs(info, mid, 0, l)

		if info.cnt >= k {
			r = mid
		} else {
			l = mid + 1
		}

	}
	return l
}

func dfs(info *Info, mid int, row int, sum int) {

	if row == info.m || info.cnt >= info.k {
		return
	}
	dfs(info, mid, row+1, sum)
	for i := 1; i < info.n; i++ {
		val := sum + info.mat[row][i] - info.mat[row][i-1]
		if val > mid {
			break
		}
		info.cnt++
		dfs(info, mid, row+1, val)

	}
}
