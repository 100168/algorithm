package main

import (
	"fmt"
	"slices"
	"sort"
)

/**
给你一个下标从 0 开始的字符串 s 、字符串 a 、字符串 b 和一个整数 k 。

如果下标 i 满足以下条件，则认为它是一个 美丽下标 ：

0 <= i <= s.length - a.length
s[i..(i + a.length - 1)] == a
存在下标 j 使得：
0 <= j <= s.length - b.length
s[j..(j + b.length - 1)] == b
|j - i| <= k
以数组形式按 从小到大排序 返回美丽下标。
*/

func beautifulIndices(s string, a string, b string, k int) []int {

	ls := len(s)
	la := len(a)
	lb := len(b)
	findNext := func(pattern string) []int {
		n := len(pattern)
		next := make([]int, n)
		for i := 1; i < n; i++ {
			j := next[i-1]

			for j > 0 && pattern[j] != pattern[i] {
				j = next[j-1]
			}

			if pattern[j] == pattern[i] {
				j++
			}
			next[i] = j
		}
		return next
	}

	nexta := findNext(a)
	nextb := findNext(b)

	var findA []int
	var findB []int

	var ans []int
	ja := 0
	jb := 0
	for i := 0; i < ls; i++ {
		for ja != 0 && a[ja] != s[i] {
			ja = nexta[ja-1]
		}
		for jb != 0 && b[jb] != s[i] {
			jb = nextb[jb-1]
		}
		if a[ja] == s[i] {
			ja++
		}
		if b[jb] == s[i] {
			jb++
		}
		if ja == la {
			findA = append(findA, i-ja+1)
			ja = nexta[ja-1]
		}
		if jb == lb {
			findB = append(findB, i-jb+1)
			jb = nextb[jb-1]
		}
	}

	find := func(t int, nums []int) int {
		l, r := 0, len(nums)-1
		for l <= r {
			m := (l + r) / 2
			if nums[m] >= t {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		return l
	}

	for _, v := range findA {
		//
		t := find(v, findB)

		if t < len(findB) && abs(findB[t]-v) <= k || t-1 >= 0 && abs(findB[t-1]-v) <= k {
			ans = append(ans, v)
		}
	}
	//for _, v := range findB {
	//	t := find(v, findA)
	//
	//	if t < len(findA) && abs(findA[t]-v) <= k || t-1 >= 0 && abs(findA[t-1]-v) <= k {
	//		ans = append(ans, v)
	//	}
	//}
	sort.Ints(ans)
	return slices.Compact(ans)
}

func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

func main() {
	//fmt.Println(beautifulIndices("isawsquirrelnearmysquirrelhouseohmy", "my", "squirrel", 15))
	//fmt.Println(beautifulIndices("ithhi", "t", "hhi", 1))
	//fmt.Println(beautifulIndices("jajrfw", "rf", "j", 3))
	fmt.Println(beautifulIndices("lagopphhnl", "gopph", "hnl", 8))
}
