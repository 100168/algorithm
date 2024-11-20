package main

func isPossibleToSplit(nums []int) bool {

	cnt := make(map[int]int)

	for _, v := range nums {
		cnt[v]++
	}

	maxCnt := 0

	for _, v := range cnt {
		maxCnt = max(maxCnt, v)
	}

	return maxCnt <= 2

}
