package main

func minAnagramLength(s string) int {

	n := len(s)
	for i := 0; i < n; i++ {
		if n%(i+1) != 0 {
			continue
		}
		if check(s, i+1) {
			return i + 1
		}
	}
	return -1
}

func check(s string, t int) bool {

	n := len(s)

	cntMap := make(map[uint8]int)

	for i := 0; i < t; i++ {
		cntMap[s[i]]++
	}

	for i := t; i < n; i += t {
		curMap := make(map[uint8]int)

		for j := i; j < i+t; j++ {
			curMap[s[j]]++
		}

		for k, v := range curMap {
			if cntMap[k] != v {
				return false
			}
		}

	}
	return true
}
