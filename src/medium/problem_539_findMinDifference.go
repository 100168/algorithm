package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

/*
给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。*/
func getMinutes(t string) int {
	return (int(t[0]-'0')*10+int(t[1]-'0'))*60 + int(t[3]-'0')*10 + int(t[4]-'0')
}

func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	ans := math.MaxInt32
	t0Minutes := getMinutes(timePoints[0])
	preMinutes := t0Minutes
	for _, t := range timePoints[1:] {
		minutes := getMinutes(t)
		ans = min(ans, minutes-preMinutes) // 相邻时间的时间差
		preMinutes = minutes
	}
	ans = min(ans, t0Minutes+1440-preMinutes) // 首尾时间的时间差
	return ans
}

func main() {
	fmt.Println(time.Now())
}
