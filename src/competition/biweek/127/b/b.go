package main

func minimumLevels(possible []int) int {

	n := len(possible)
	sum := 0
	for _, v := range possible {
		if v == 1 {
			sum += 1
		} else {
			sum -= 1
		}
	}

	sumA := 0
	for i, v := range possible[:n-1] {
		if v == 1 {
			sumA += 1
		} else {
			sumA -= 1
		}
		diff := sum - sumA
		if sumA > diff {
			return i + 1
		}

	}
	return -1

}
