package main

import "slices"

func minimumTime(time []int, totalTrips int) int64 {

	l, r := int64(1), int64(slices.Max(time)*totalTrips)

	//循环不变量  r+1>=totalTrips l-1< totalTrips => ans ==r+1 ==>l

	for l <= r {
		m := (l + r) / 2
		if check(time, m) >= int64(totalTrips) {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return l

}

func check(time []int, target int64) int64 {

	sum := int64(0)

	for _, v := range time {
		sum += target / int64(v)
	}
	return sum
}
