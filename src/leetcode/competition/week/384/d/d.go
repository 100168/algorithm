package main

import "cmp"

// 字符串编码
func countMatchingSubarrays(nums []int, pattern []int) int {

	n := len(nums)
	m := len(pattern)

	base := 131313

	ps := make([]int, n)
	h := make([]int, n)
	ps[0] = 1
	ans := 0
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			h[i] = h[i-1]*base + 1
			ps[i] = ps[i-1] * base
		} else if nums[i] == nums[i-1] {
			h[i] = h[i-1]*base + 2
			ps[i] = ps[i-1] * base
		} else if nums[i] < nums[i-1] {
			h[i] = h[i-1]*base + 3
			ps[i] = ps[i-1] * base
		}
	}

	hh := ps[m]
	pp := 0
	for _, v := range pattern {
		if v == 1 {
			pp = pp*base + 1
		}
		if v == 0 {
			pp = pp*base + 2
		}
		if v == -1 {
			pp = pp*base + 3
		}
	}

	for i := m; i < len(h); i++ {
		cur := h[i] - h[i-m]*hh
		if cur == pp {
			ans++
		}
	}
	return ans
}

// kmp
func countMatchingSubarrays2(nums []int, pattern []int) int {

	m := len(pattern)

	//多少个前缀跟后缀是一样的
	//j 表示下一个要匹配的位置
	//
	next := make([]int, m)
	for i := 1; i < m; i++ {
		j := next[i-1]
		for j > 0 && pattern[j] != pattern[i] {
			//向前跳，看看前j个数有多少个前后缀相同
			j = next[j-1]
		}

		if pattern[j] == pattern[i] {
			j++
		}
		next[i] = j
	}

	ans := 0

	j := 0
	for i := 1; i < len(nums); i++ {
		v := cmp.Compare(nums[i], nums[i-1])
		for j > 0 && pattern[j] != v {
			j = next[j-1]
		}
		if pattern[j] == v {
			j++
		}
		if j == m {
			ans++
			j = next[j-1]
		}
	}
	return ans
}

func main() {
	println(countMatchingSubarrays([]int{1, 4, 4, 1, 3, 5, 5, 3}, []int{1, 0, -1}))
	println(countMatchingSubarrays2([]int{1, 4, 4, 1, 3, 5, 5, 3}, []int{1, 0, -1}))
}
