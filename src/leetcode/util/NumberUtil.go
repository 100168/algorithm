package main

import (
	"fmt"
	"strings"
)

func brace(s string) string {
	s = strings.Replace(s, "[", "{", len(s))
	s = strings.Replace(s, "]", "}", len(s))
	return s
}

func join(s string) string {

	split := strings.Split(s, ",")

	for i := range split {

		split[i] = "'" + split[i] + "'"
	}

	return strings.Join(split, ",")

}

func diff(s1, s2 string) string {

	sp1 := strings.Split(s1, ",")
	sp2 := strings.Split(s2, ",")

	m1 := make(map[string]bool)
	for _, v := range sp1 {

		m1[v] = true
	}

	set := make(map[string]bool)
	ans := make([]string, 0)

	for _, v := range sp2 {

		if set[v] {
			continue
		}
		set[v] = true
		if !m1[v] {
			ans = append(ans, v)
		}
	}

	return strings.Join(ans, ",")

}
func main() {

	nums := []int{1, 1, 1, 1}

	n := len(nums)

	f := make([]int, 1<<n)

	for i := 0; i < n; i++ {

		mask := (1<<n - 1) ^ 1<<i
		f[1<<i] = nums[i]

		for sub := mask; sub > 0; sub = mask & (sub - 1) {
			f[sub|1<<i] = f[sub] + nums[i]
		}
	}
	fmt.Println(f)
}
