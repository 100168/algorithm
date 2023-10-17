package main

func minOperations(s1 string, s2 string, x int) int {

	n := len(s1)
	newChar := make([]uint8, n)
	flag := uint8(0)
	for i := 0; i < n; i++ {
		newChar[i] = s1[i] ^ s2[i]
		flag = newChar[i] ^ flag
	}
	if flag > 0 {
		return -1
	}
	ans := 0
	return ans
}

func min(a, b int) int {
	if a > b {

		return b
	}
	return a
}
func main() {
	println(minOperations("101101", "000000", 6))
}
