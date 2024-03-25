package main

func hIndex(citations []int) int {
	//最大值
	n := len(citations)
	l, r := 1, n
	for l < r {
		m := (r + l) / 2
		//大于等于h的文章个数
		if citations[n-m] >= m {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}
