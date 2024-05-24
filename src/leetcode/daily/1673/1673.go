package main

import (
	"fmt"
)

/*
*
给你一个整数数组 nums 和一个正整数 k ，返回长度为 k 且最具 竞争力 的 nums 子序列。

数组的子序列是从数组中删除一些元素（可能不删除元素）得到的序列。

在子序列 a 和子序列 b 第一个不相同的位置上，如果 a 中的数字小于 b 中对应的数字，
那么我们称子序列 a 比子序列 b（相同长度下）更具 竞争力 。
例如，[1,3,4] 比 [1,3,5] 更具竞争力，在第一个不相同的位置，也就是最后一个位置上， 4 小于 5 。

[3,5,2,6], k = 2
*/

func mostCompetitive(nums []int, k int) []int {
	//长度为k的最小递增子序列
	var ans []int
	n := len(nums)
	for i, v := range nums {
		for len(ans) > 0 && len(ans)+(n-i) > k && ans[len(ans)-1] > v {
			ans = ans[:len(ans)-1]
		}
		ans = append(ans, v)
	}
	return ans[:k]
}

// 直接二分
func mostCompetitive1(nums []int, k int) []int {
	//长度为k的最小递增子序列
	var ans []int
	n := len(nums)
	for i, v := range nums {
		// l = max(0,k-(n-i))
		l, r := max(0, k-(n-i)), len(ans)-1
		for l <= r {
			mid := (l + r) / 2
			if ans[mid] <= v {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		ans = ans[:l]
		ans = append(ans, v)
	}
	return ans[:k]
}

func main() {
	//fmt.Println(mostCompetitive([]int{3, 5, 2, 6}, 2))
	//fmt.Println(mostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4))
	//fmt.Println(mostCompetitive([]int{71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2}, 2))
	fmt.Println(mostCompetitive1([]int{84, 10, 71, 23, 66, 61, 62, 64, 34, 41, 80, 25, 91, 43, 4, 75, 65, 13, 37, 41, 46, 90, 55, 8, 85, 61, 95, 71}, 24))
	fmt.Println(mostCompetitive([]int{84, 10, 71, 23, 66, 61, 62, 64, 34, 41, 80, 25, 91, 43, 4, 75, 65, 13, 37, 41, 46, 90, 55, 8, 85, 61, 95, 71}, 24))
}
