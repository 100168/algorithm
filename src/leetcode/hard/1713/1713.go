package main

func minOperations(target []int, arr []int) int {

	indexMap := make(map[int]int)

	for i, v := range target {

		indexMap[v] = i

	}
	se := make([]int, 0)

	for _, v := range arr {

		a, ok := indexMap[v]
		if ok {
			if len(se) == 0 || se[len(se)-1] < a {
				se = append(se, a)
			} else {
				index := find(se, a)
				se[index] = a
			}
		}
	}

	return len(target) - len(se)

}

func find(arr []int, target int) int {

	l := 0
	r := len(arr)

	for l < r {

		m := l + r>>1
		if arr[m] >= target {

			r = m
		} else {
			l = m + 1
		}
	}

	return l

}
