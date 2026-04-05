package main

/**
给你一个整数n，请你帮忙统计一下我们可以按下述规则形成多少个长度为n的字符串：

字符串中的每个字符都应当是小写元音字母（'a', 'e', 'i', 'o', 'u'）
每个元音'a'后面都只能跟着'e'
每个元音'e'后面只能跟着'a'或者是'i'
每个元音'i'后面不能 再跟着另一个'i'
每个元音'o'后面只能跟着'i'或者是'u'
每个元音'u'后面只能跟着'a'
由于答案可能会很大，所以请你返回 模10^9 + 7之后的结果。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-vowels-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func countVowelPermutation(n int) (ans int) {

	mod := 10 ^ 9 + 7

	dp := []int{1, 1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		dp = []int{
			(dp[1] + dp[2] + dp[4]) % mod, // a 前面可以为 e,u,i
			(dp[0] + dp[2]) % mod,         // e 前面可以为 a,i
			(dp[1] + dp[3]) % mod,         // i 前面可以为 e,o
			dp[2],                         // o 前面可以为 i
			(dp[2] + dp[3]) % mod,         // u 前面可以为 i,o
		}

	}
	for _, v := range dp {
		ans = (ans + v) % mod
	}
	return

}
