package main

/*
*
给你一个整数 n 。如果两个整数 x 和 y 满足下述条件，则认为二者形成一个质数对：

1 <= x <= y <= n
x + y == n
x 和 y 都是质数
请你以二维有序列表的形式返回符合题目要求的所有 [xi, yi] ，列表需要按 xi 的 非递减顺序 排序。如果不存在符合要求的质数对，则返回一个空数组。

注意：质数是大于 1 的自然数，并且只有两个因子，即它本身和 1 。
*/
func findPrimePairs(n int) [][]int {

	isPrime := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	ans := make([][]int, 0)
	for i := 2; i*2 <= n; i++ {

		if isPrime[i] && isPrime[n-i] {
			ans = append(ans, []int{i, n - i})
		}
	}
	return ans
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
