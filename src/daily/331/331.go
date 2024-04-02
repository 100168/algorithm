package main

import "strings"

func isValidSerialization(preorder string) bool {

	split := strings.Split(preorder, ",")

	n := len(split)

	cost := 1
	for i := 0; i < n; i++ {
		if cost == 0 {
			return false
		}
		if split[i] == "#" {
			cost--
		} else {
			cost++
		}
	}
	return cost == 0
}

func isValidSerialization2(preorder string) bool {
	st := make([]string, 0)

	split := strings.Split(preorder, ",")
	n := len(split)
	for i := 0; i < n; i++ {
		if split[i] != "#" {
			st = append(st, split[i])
		} else {
			for len(st) > 0 && st[len(st)-1] == "#" {
				st = st[:len(st)-1]
				if len(st) == 0 {
					return false
				}
				st = st[:len(st)-1]
			}
			st = append(st, split[i])
		}
	}
	return len(st) == 1 && st[0] == "#"
}

func isValidSerialization3(preorder string) bool {

	split := strings.Split(preorder, ",")
	n := len(split)

	var dfs func(int) (int, bool)

	dfs = func(i int) (int, bool) {

		if i < n {
			if split[i] == "#" {
				return i + 1, true
			}

			right, leftOk := dfs(i + 1)

			next, rightOk := dfs(right)
			return next, leftOk && rightOk

		}
		return i, false
	}
	visited, ok := dfs(0)
	return ok && visited >= n
}

func main() {

	println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	println(isValidSerialization2("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	println(isValidSerialization3("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}
