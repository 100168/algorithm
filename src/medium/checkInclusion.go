package main

import "fmt"

/*给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，s1 的排列之一是 s2 的 子串 。*/
func checkInclusion(s1 string, s2 string) bool {

	n := len(s1)
	m := len(s2)
	ans := [26]int{}
	for i := 0; i < n; i++ {
		index := s1[i] - 'a'
		ans[index]--
	}

	l := 0

	for r := 0; r < m; r++ {
		index := s2[r] - 'a'
		ans[index]++
		if ans[index] > 0 {
			ans[index]--
			//用来表示有多少个字符是s1中没有的
			l++
		}
		if r-l+1 == n {
			return true
		}
	}
	return false
}

func main() {

	inclusion := checkInclusion("abc", "ddddnajbjljkjc")
	fmt.Println(inclusion)
}
