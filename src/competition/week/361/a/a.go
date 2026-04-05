package main

import "strconv"

func countSymmetricIntegers(low int, high int) int {

	ans := 0
	for i := low; i <= high; i++ {

		str := strconv.Itoa(i)
		n := len(str)
		if len(str)&1 == 0 {

			sum := 0

			for i := 0; i < n; i++ {

				if i < n>>1 {
					sum += int(str[i] - '0')
				} else {
					sum -= int(str[i] - '0')
				}
			}
			if sum == 0 {
				ans++
			}
		}

	}
	return ans
}
