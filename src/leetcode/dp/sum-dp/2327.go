package main

/*
*
在第 1 天，有一个人发现了一个秘密。

给你一个整数 delay ，表示每个人会在发现秘密后的 delay 天之后，
每天 给一个新的人 分享 秘密。同时给你一个整数 forget ，表示每个人在发现秘密 forget 天之后会 忘记 这个秘密。
一个人 不能 在忘记秘密那一天及之后的日子里分享秘密。

给你一个整数 n ，请你返回在第 n 天结束时，知道秘密的人数。由于答案可能会很大，请你将结果对 109 + 7 取余 后返回。
*/

/*
 */
func peopleAwareOfSecret(n int, delay int, forget int) int {
	// 1,1,1,2,1,

	//a:可以造谣
	//b:目前不能造谣
	//c:已经不能造谣
	mod := 1000000007
	sum := make([]int, n+1)
	for i := 1; i <= delay; i++ {
		sum[i] = 1
	}
	//  s[i-delay] - s[i-forget]
	for i := delay + 1; i <= n; i++ {
		//f 表示新增的利息
		f := (sum[i-delay] - sum[max(i-forget, 0)]) % mod
		//当前的总数等于昨天的数量加上新增的数量
		sum[i] = (sum[i-1] + f) % mod
	}
	//答案等于单前数量-忘记的数量
	return ((sum[n]-sum[max(n-forget, 0)])%mod + mod) % mod
}

func peopleAwareOfSecret3(n, delay, forget int) int {
	const mod int = 1e9 + 7
	sum := make([]int, n+1)
	sum[1] = 1
	for i := 2; i <= n; i++ {
		f := (sum[max(i-delay, 0)] - sum[max(i-forget, 0)]) % mod
		sum[i] = (sum[i-1] + f) % mod
	}
	return ((sum[n]-sum[max(n-forget, 0)])%mod + mod) % mod // 防止结果为负数
}

func main() {
	println(peopleAwareOfSecret(6, 2, 4))
}
