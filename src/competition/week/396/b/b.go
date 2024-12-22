package main

import "fmt"

func minimumOperationsToMakeKPeriodic(word string, k int) int {

	n := len(word)

	cntMap := make(map[string]int)

	maxCnt := 0

	for i := n; i >= k; i -= k {
		cur := word[i-k : i]
		cntMap[cur]++
		maxCnt = max(maxCnt, cntMap[cur])
	}

	return n/k - maxCnt
}

func main() {
	fmt.Println(minimumOperationsToMakeKPeriodic("rry", 1))
}
