package main

func minSpeedOnTime(dist []int, hour float64) int {

	l, r := 1, 10_000_000
	for l <= r {
		m := (l + r) / 2
		if check(dist, m) <= hour {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func check(dist []int, speed int) float64 {

	sum := float64(0)
	for _, v := range dist[:len(dist)-1] {
		sum += float64((v-1)/speed + 1)
	}

	return sum + float64(dist[len(dist)-1])/float64(speed)
}
func main() {
	println(minSpeedOnTime([]int{1, 3, 2}, 2.7))
}
