package main

import (
	"math"
)

func closestToTarget(arr []int, target int) int {

	ans := math.MaxInt

	pres := make([]int, 0)
	for _, v := range arr {

		pres = append(pres, v)

		newPres := make([]int, 0)

		for _, e := range pres {

			cur := e & v
			if len(newPres) > 0 && newPres[len(newPres)-1] == cur {
				continue
			}
			newPres = append(newPres, cur)
			ans = min(ans, abs(cur-target))
		}
		pres = newPres
	}
	return ans
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a

}
