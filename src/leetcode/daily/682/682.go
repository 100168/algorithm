package main

import "strconv"

func calPoints(operations []string) int {
	var st []int
	for _, v := range operations {
		switch v {

		case "+":
			st = append(st, st[len(st)-1]+st[len(st)-2])
			break
		case "D":
			st = append(st, st[len(st)-1]*2)
			break
		case "C":
			st = st[:len(st)-1]
			break
		default:
			num, _ := strconv.Atoi(v)
			st = append(st, num)
		}

	}

	ans := 0

	for _, v := range st {
		ans += v
	}
	return ans

}
