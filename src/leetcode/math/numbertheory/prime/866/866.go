package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
*
给你一个整数 n ，返回大于或等于 n 的最小 回文质数。
一个整数如果恰好有两个除数：1 和它本身，那么它是 质数 。注意，1 不是质数。

例如，2、3、5、7、11 和 13 都是质数。
一个整数如果从左向右读和从右向左读是相同的，那么它是 回文数 。

例如，101 和 12321 都是回文数。
测试用例保证答案总是存在，并且在 [2, 2 * 108] 范围内。

输入：n = 6
输出：7
示例 2：

输入：n = 8
输出：11
示例 3：

输入：n = 13
输出：101
*/
func primePalindrome(n int) int {
	for l := 1; l <= 5; l++ {
		for i := 1; i < int(math.Pow10(l)); i++ {
			leftS := strconv.Itoa(i)
			for j := len(leftS) - 2; j >= 0; j-- {
				leftS += string(leftS[j])
			}
			curV, _ := strconv.Atoi(leftS)
			if curV >= n && isPrime(curV) {
				return curV
			}
		}
		for i := 1; i < int(math.Pow10(l)); i++ {
			leftS := strconv.Itoa(i)
			for j := len(leftS) - 1; j >= 0; j-- {
				leftS += string(leftS[j])
			}
			curV, _ := strconv.Atoi(leftS)
			if curV >= n && isPrime(curV) {
				return curV
			}
		}
	}
	return n
}
func isPrime(n int) bool {

	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/*
	class Solution {
	    public int primePalindrome(int N) {
	        while (true) {
	            if (N == reverse(N) && isPrime(N))
	                return N;
	            N++;
	            if (10_000_000 < N && N < 100_000_000)
	                N = 100_000_000;
	        }
	    }

	    public boolean isPrime(int N) {
	        if (N < 2) return false;
	        int R = (int) Math.sqrt(N);
	        for (int d = 2; d <= R; ++d)
	            if (N % d == 0) return false;
	        return true;
	    }

	    public int reverse(int N) {
	        int ans = 0;
	        while (N > 0) {
	            ans = 10 * ans + (N % 10);
	            N /= 10;
	        }
	        return ans;
	    }
	}
*/
func main() {
	fmt.Println(primePalindrome(80283903))
	fmt.Println(primePalindrome(13))
	fmt.Println(primePalindrome(6))
	fmt.Println(primePalindrome(8))
	fmt.Println(primePalindrome(1))
}
