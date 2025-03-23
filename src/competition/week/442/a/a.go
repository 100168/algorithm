package main

func maxContainers(n int, w int, maxWeight int) int {

	return min(maxWeight/w, n*n) * w
}
