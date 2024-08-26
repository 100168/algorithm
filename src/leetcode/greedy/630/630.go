package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/**
这里有 n 门不同的在线课程，按从 1 到 n 编号。给你一个数组 courses ，
其中 courses[i] = [durationi, lastDayi] 表示第 i 门课将会 持续 上 durationi 天课，并且必须在不晚于 lastDayi 的时候完成。

你的学期从第 1 天开始。且不能同时修读两门及两门以上的课程。

返回你最多可以修读的课程数目。



示例 1：

输入：courses = [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
输出：3
解释：
这里一共有 4 门课程，但是你最多可以修 3 门：
首先，修第 1 门课，耗费 100 天，在第 100 天完成，在第 101 天开始下门课。
第二，修第 3 门课，耗费 1000 天，在第 1100 天完成，在第 1101 天开始下门课程。
第三，修第 2 门课，耗时 200 天，在第 1300 天完成。
第 4 门课现在不能修，因为将会在第 3300 天完成它，这已经超出了关闭日期。


1.思路反悔贪心
*/

func scheduleCourse(courses [][]int) int {

	//先按结束时间从小到大排序
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	s := 0

	h := new(myHeap)

	//返回贪心
	for _, v := range courses {

		//当前结束时间
		s += v[0]
		heap.Push(h, v[0])
		//当前结束时间超过截止时间
		if s > v[1] {
			//找到持续时间最久的课程
			pop := heap.Pop(h).(int)
			//减去这个课程
			s -= pop
		}
	}
	return h.Len()
}

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func main() {

	fmt.Println(scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}))
	fmt.Println(scheduleCourse([][]int{{1, 2}, {2, 3}}))
}
