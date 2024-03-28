package main

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {

	l, r := 1, 1_000_000_000*3
	for l <= r {
		m := (r-l)/2 + l

		if check(m, divisor1, divisor2, uniqueCnt1, uniqueCnt2) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
func check(t, d1, d2, uq1, uq2 int) bool {
	m1 := t / d1
	m2 := t / d2
	x := t / lcm(d1, d2)

	left1 := max(uq1-m2+x, 0)
	left2 := max(uq2-m1+x, 0)
	rest := t - m1 - m2 + x
	return rest >= left1+left2
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {

	//println(minimizeSet(2, 7, 1, 3))
	println(minimizeSet(2, 4, 8, 2))
}
