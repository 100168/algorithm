package main

import "fmt"

/*
*
给定一个整数 n ，返回 n! 结果中尾随零的数量。

提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1

示例 1：

输入：n = 3
输出：0
解释：3! = 6 ，不含尾随 0
示例 2：

输入：n = 5
输出：1
解释：5! = 120 ，有一个尾随 0
示例 3：

输入：n = 0
输出：0

	class Solution {
	    public int trailingZeroes(int n) {
	        int ans = 0;
	        while (n != 0) {
	            n /= 5;
	            ans += n;
	        }
	        return ans;
	    }
	}
*/
func trailingZeroes(n int) int {
	ans := 0
	for n > 0 {
		n /= 5
		ans += n
	}
	return ans
}

func trailingZeroes2(n int) int {
	ans := 0
	for i := 5; i <= n; i += 5 {
		for x := i; x%5 == 0; x /= 5 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(trailingZeroes(30))
	fmt.Println(trailingZeroes2(30))
}
