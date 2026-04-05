package main

import "fmt"

func wonderfulSubstrings(word string) int64 {

	n := len(word)
	pos := make([]int64, 1<<10)
	pos[0] = 1
	status := 0
	ans := int64(0)

	for i := 0; i < n; i++ {
		cur := word[i] - 'a'
		status ^= 1 << cur
		ans += pos[status]
		pos[status]++
		for j := 0; j < 10; j++ {
			ans += pos[status^(1<<j)]
		}

	}
	return ans

}
func main() {
	substrings := wonderfulSubstrings("aabb")
	fmt.Println(substrings)
}
