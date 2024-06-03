package main

import (
	"math"
	"sort"
)

/*
*

写法太low
*/

func smallestFactorization(num int) int {

	ans := 0

	if num < 10 {
		return num
	}
	p := 1
	for i := 9; i > 1; i-- {
		for num%i == 0 {
			num /= i
			ans += p * i
			p *= 10
		}
	}
	if num > 1 {
		return 0
	}
	if ans > 0 && ans < math.MaxInt32 {
		return ans
	}
	return 0

}
func smallestFactorization2(num int) int {

	if num < 10 {
		return num
	}
	c := num
	primes := make(map[int]int)

	for i := 2; i*i <= c; i++ {

		if i >= 8 {
			return 0
		}
		cnt := 0
		for c%i == 0 {
			cnt++
			c /= i
		}
		primes[i] = cnt
	}
	if c > 1 {
		if c > 7 {
			return 0
		} else {
			primes[c] = 1
		}
	}

	nums := make([]int, 0)
	for k, v := range primes {
		if k > 3 {
			for v > 0 {
				nums = append(nums, k)
				v--
			}
		}
		if k == 3 {
			for v >= 2 {
				nums = append(nums, 9)
				v -= 2
			}
		}
		if k == 2 {
			for v >= 3 {
				nums = append(nums, 8)
				v -= 3
			}
		}
		primes[k] = v
		if v == 0 {
			delete(primes, k)
		}
	}

	if primes[2] > 0 && primes[3] > 0 {
		nums = append(nums, 6)
		primes[2]--
		primes[3]--
		if primes[2] > 0 {
			nums = append(nums, 2)
			primes[2]--
		}
	}
	if primes[2] > 0 {
		nums = append(nums, 1<<primes[2])
	}
	if primes[3] > 0 {
		nums = append(nums, 3)
	}

	sort.Ints(nums)

	ans := 0

	for _, v := range nums {
		ans = ans*10 + v
	}

	if ans > 0 && ans < math.MaxInt32 {
		return ans
	}
	return 0
}
