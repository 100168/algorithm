package main

import (
	"fmt"
	"slices"
	"sort"
)

func maximumTotalDamage(power []int) int64 {

	sort.Ints(power)

	cnt := make(map[int]int)
	for i := range power {
		cnt[power[i]]++
	}
	power = slices.Compact(power)
	n := len(power)
	yes := make([]int, n+1)
	no := make([]int, n+1)
	ans := 0
	for i := range power {
		yes[i] = power[i] * cnt[power[i]]
		maxVal := 0
		if i > 0 {
			maxVal = no[i-1]
			if power[i-1] < power[i]-2 {
				maxVal = max(maxVal, yes[i-1])
			}
		}
		if i > 1 {
			maxVal = max(maxVal, no[i-2])
			if power[i-2] < power[i]-2 {
				maxVal = max(maxVal, yes[i-2])
			}
		}
		if i > 2 {
			maxVal = max(yes[i-3], maxVal)
		}
		yes[i] += maxVal
		no[i] = maxVal
		ans = max(ans, yes[i], no[i])
	}
	return int64(ans)

}

func main() {
	fmt.Println(maximumTotalDamage([]int{7, 1, 6, 3}))
}
