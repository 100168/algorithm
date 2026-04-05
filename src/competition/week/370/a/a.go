package main

func findChampion(grid [][]int) int {

	n := len(grid)
	for i := 0; i < n; i++ {
		flag := true
		for j := 0; j < n; j++ {
			if i != j && grid[i][j] == 0 {
				flag = false
			}
		}
		if flag {
			return i
		}
	}
	return -1
}
