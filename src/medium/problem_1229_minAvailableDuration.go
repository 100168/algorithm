package main

import "sort"

/*你是一名行政助理，手里有两位客户的空闲时间表：slots1 和 slots2，以及会议的预计持续时间duration，请你为他们安排合适的会议时间。

「会议时间」是两位客户都有空参加，并且持续时间能够满足预计时间duration 的 最早的时间间隔。

如果没有满足要求的会议时间，就请返回一个 空数组。



「空闲时间」的格式是[start, end]，由开始时间start和结束时间end组成，表示从start开始，到 end结束。

题目保证数据有效：同一个人的空闲时间不会出现交叠的情况，也就是说，对于同一个人的两个空闲时间[start1, end1]和[start2, end2]，要么start1 > end2，要么start2 > end1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/meeting-scheduler
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {

	point1 := 0
	point2 := 0

	l1 := len(slots1)
	l2 := len(slots2)

	sort.Slice(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0] //按照每行的第一个元素排序
	})
	sort.Slice(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0] //按照每行的第一个元素排序
	})

	for point1 < l1 && point2 < l2 {

		ans := inter(point1, point2, slots1, slots2, duration)
		if ans != nil {
			return ans
		}
		if slots1[point1][1] < slots2[point2][1] {
			point1++
		} else {
			point2++
		}
	}

	return nil
}

func inter(point1 int, point2 int, slots1 [][]int, slots2 [][]int, k int) []int {

	if slots1[point1][0] <= slots2[point2][0] &&
		slots1[point1][1] >= slots2[point2][0]+k &&
		slots2[point2][1] >= slots2[point2][0]+k {
		return []int{slots2[point2][0], slots2[point2][0] + k}
	}

	if slots2[point2][0] <= slots1[point1][0] &&
		slots2[point2][1] >= slots1[point1][0]+k &&
		slots1[point1][1] >= slots1[point1][0]+k {
		return []int{slots1[point1][0], slots1[point1][0] + k}
	}

	return nil

}
