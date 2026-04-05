package main

func findTheLongestSubstring(s string) int {
	sumMap := make(map[int]int)
	sum := 0
	ans := 0
	charMap := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	for i, v := range s {

		curV := int(1 << (v - 'a'))
		if charMap[string(v)] {
			sum = sum ^ curV
		}
		if sum != 0 {
			if v, ok := sumMap[sum]; ok {
				ans = max(ans, i-v)
				continue
			}
			sumMap[sum] = i
		} else {
			ans = i + 1
		}
	}

	return ans

}

func main() {
	println(findTheLongestSubstring("cbc"))
}
