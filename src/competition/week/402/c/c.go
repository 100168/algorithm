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
		for j := max(i-3, 0); j < i; j++ {
			maxVal = max(maxVal, no[j])
			if power[j] < power[i]-2 {
				maxVal = max(maxVal, yes[j])
			}
		}
		yes[i] += maxVal
		no[i] = maxVal
		ans = max(ans, yes[i], no[i])
	}
	return int64(ans)

}

func maximumTotalDamage2(power []int) int64 {
	cnt := map[int]int{}
	for _, x := range power {
		cnt[x]++
	}

	sort.Ints(power)
	power = slices.Compact(power)
	n := len(power)
	f := make([]int, n+1)
	j := 0
	for i, x := range power {
		for power[j] < x-2 {
			j++
		}
		f[i+1] = max(f[i], f[j]+x*cnt[x])
	}
	return int64(f[n])
}
func main() {
	fmt.Println(maximumTotalDamage([]int{7, 1, 6, 3}))
}
