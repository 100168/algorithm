package main

func getMaxLen(nums []int) int {

	ne := 0
	po := 0
	ans := 0
	for _, v := range nums {
		if v == 0 {
			ne = 0
			po = 0
			continue
		}
		if v > 0 {
			po = po + 1
			if ne > 0 {
				ne = ne + 1
			}
		} else {
			temp := po
			po = ne
			if ne > 0 {
				po++
			}
			ne = temp + 1
		}

		ans = max(ans, po)
	}
	return ans
}
