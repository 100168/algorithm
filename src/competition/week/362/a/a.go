package main

import (
	"sort"
)

func numberOfPoints(nums [][]int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	ansMap := make(map[int]bool)
	for _, num := range nums {

		for j := num[0]; j <= num[1]; j++ {
			ansMap[j] = true
		}
	}
	return len(ansMap)
}

func main() {

	nums := [][]int{{1, 3}, {5, 8}}
	numberOfPoints(nums)
}
