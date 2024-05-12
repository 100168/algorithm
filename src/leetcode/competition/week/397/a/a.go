package main

func findPermutationDifference(s string, t string) int {

	ans := 0

	index := make(map[uint8]int)
	for i := range t {

		index[t[i]] = i

	}
	for i := range s {

		ans += abs(index[s[i]] - index[t[i]])
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
