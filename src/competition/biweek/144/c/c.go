package main

import (
	"fmt"
	"sort"
)

/*
*
 */
func maxRemoval(nums []int, queries [][]int) int {
	n := len(nums)

	type even struct {
		pos, delta int
	}
	var events []even

	for _, query := range queries {
		li, ri := query[0], query[1]
		events = append(events, even{pos: li, delta: 1})
		events = append(events, even{pos: ri + 1, delta: -1})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].pos == events[j].pos {
			return events[i].delta < events[j].delta
		}
		return events[i].pos < events[j].pos
	})

	coverage := make([]int, n)
	currCoverage := 0
	eventIdx := 0

	for i := 0; i < n; i++ {
		for eventIdx < len(events) && events[eventIdx].pos == i {
			currCoverage += events[eventIdx].delta
			eventIdx++
		}
		coverage[i] = currCoverage
	}

	for i := 0; i < n; i++ {
		if coverage[i] < nums[i] {
			return -1
		}
	}

	type QueryInfo struct {
		L, R, Index int
	}
	queryInfo := make([]QueryInfo, len(queries))
	for i, query := range queries {
		queryInfo[i] = QueryInfo{L: query[0], R: query[1], Index: i}
	}

	sort.Slice(queryInfo, func(i, j int) bool {
		return queryInfo[i].R-queryInfo[i].L < queryInfo[j].R-queryInfo[j].L
	})

	countDeleted := 0
	used := make([]bool, len(queries))
	for _, q := range queryInfo {
		li, ri := q.L, q.R

		for i := li; i <= ri; i++ {
			coverage[i]--
		}

		valid := true
		for i := 0; i < n; i++ {
			if coverage[i] < nums[i] {
				valid = false
				break
			}
		}

		if !valid {
			for i := li; i <= ri; i++ {
				coverage[i]++
			}
		} else {
			used[q.Index] = true
			countDeleted++
		}
	}

	return countDeleted
}
func main() {
	fmt.Println(maxRemoval([]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}}))
}
