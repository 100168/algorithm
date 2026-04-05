package main

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {

	diffY := fy - sy
	diffX := fx - sx

	if diffX < 0 {
		diffX = -diffX
	}
	if diffY < 0 {
		diffY = -diffY
	}

	if diffX > diffY {
		diffX -= diffY
	} else {
		diffY -= diffX
	}
	diffSum := diffY + diffX

	if diffSum == 0 && t == 1 || t < diffSum {
		return false
	}
	return true
}
