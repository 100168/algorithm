package main

func findKthNumber(m int, n int, k int) int {
	l, r := 1, m*n

	for l <= r {
		t := (r-l)/2 + l

		if check(t, m, n) <= k {
			l = t + 1
		} else {
			r = t - 1
		}
	}
	return r
}

func check(t, m, n int) int {

	i := 1
	j := n

	cnt := 1
	for i <= m && j >= 1 {
		if i*j >= t {
			j--
		} else {
			cnt += j
			i++
		}
	}
	return cnt
}

func main() {
	println(findKthNumber(3, 3, 5))
}
