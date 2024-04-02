package main

func substringXorQueries(s string, queries [][]int) [][]int {

	m := len(s)
	n := len(queries)
	ans := make([][]int, n)

	for i, v := range queries {
		cur := v[0] ^ v[1]
		curS := ""
		for j := 0; 1<<j <= cur; j++ {
			if 1<<j&cur != 0 {
				curS = "1" + curS
			} else {
				curS = "0" + curS
			}
		}
		l := len(curS)
		for j := l; j <= m; j++ {
			if s[j-l:j] == curS {
				ans[i] = []int{j - l, j - 1}
				break
			}
		}
		if ans[i] == nil {
			ans[i] = []int{-1, -1}
		}

	}
	return ans
}

func main() {

	println(substringXorQueries("101101", [][]int{{0, 5}, {1, 2}}))
}
