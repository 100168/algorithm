package main

import (
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	sameCnt := 0
	var polarDegrees []float64
	for _, p := range points {
		if p[0] == location[0] && p[1] == location[1] {
			sameCnt++
		} else {
			polarDegrees = append(polarDegrees, math.Atan2(float64(p[1]-location[1]), float64(p[0]-location[0])))
		}
	}
	sort.Float64s(polarDegrees)

	n := len(polarDegrees)
	for _, deg := range polarDegrees {
		polarDegrees = append(polarDegrees, deg+2*math.Pi)
	}

	maxCnt := 0
	degree := float64(angle) * math.Pi / 180

	l := 0
	for r := 0; r < 2*n; r++ {
		for polarDegrees[r]-polarDegrees[l] > degree {
			l++
		}
		maxCnt = max(maxCnt, r-l+1)
	}

	return sameCnt + maxCnt
}
