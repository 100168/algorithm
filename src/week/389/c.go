package main

import (
	"fmt"
	"math"
)

func minimumDeletions(word string, k int) int {

	cnt := make(map[uint8]int)

	for i := range word {
		cur := word[i] - 'a'
		cnt[cur]++
	}
	minVal := math.MaxInt
	maxVal := 0
	for _, v := range cnt {
		minVal = min(v, minVal)
		maxVal = max(maxVal, v)
	}
	if len(cnt) == 2 {
		return min(max(maxVal-minVal-k, 0), minVal)

	}
	return max(maxVal-minVal-k, 0)
}
func main() {
	minimum := minimumDeletions("aabcaba", 0)
	fmt.Println(minimum)
}
