package main

func canMakeSubsequence(str1 string, str2 string) bool {

	m := len(str1)
	n := len(str2)

	p2 := 0
	for i := 0; i < m && p2 < n; i++ {
		diff := (int(str2[p2]) - int(str1[i]) + 26) % 26
		if diff <= 1 {
			p2++
		}
	}

	return p2 == n
}

func main() {
	println(canMakeSubsequence("zc", "ad"))
}
