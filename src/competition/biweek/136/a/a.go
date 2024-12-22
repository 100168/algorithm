package main

func winningPlayerCount(n int, pick [][]int) int {

	ans := make(map[int]bool)

	cntMap := make(map[int]map[int]int)

	for _, v := range pick {

		x, y := v[0], v[1]

		if cntMap[x] == nil {
			cntMap[x] = make(map[int]int)
		}
		cntMap[x][y]++
		if cntMap[x][y] > x {
			ans[x] = true
		}
	}
	return len(ans)
}
