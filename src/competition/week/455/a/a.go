package main

func checkPrimeFrequency(nums []int) bool {

	cnt := make(map[int]int)

	for _, v := range nums {
		cnt[v]++
	}

	for _, v := range cnt {

		if v == 1 {
			continue
		}

		flag := true
		for j := 2; j*j <= v; j++ {
			if v%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}
