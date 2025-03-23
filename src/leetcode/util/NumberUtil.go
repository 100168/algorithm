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

	fmt.Println(brace("[[1,2],[1,1],[3,4],[4,5],[5,6],[7,7]]"))
}
