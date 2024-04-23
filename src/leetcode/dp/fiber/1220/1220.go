package main

import "fmt"

/*
	给你一个整数n，请你帮忙统计一下我们可以按下述规则形成多少个长度为n的字符串：

字符串中的每个字符都应当是小写元音字母（'a', 'e', 'i', 'o', 'u'）
每个元音'a'后面都只能跟着'e'
每个元音'e'后面只能跟着'a'或者是'i'
每个元音'i'后面不能 再跟着另一个'i'
每个元音'o'后面只能跟着'i'或者是'u'
每个元音'u'后面只能跟着'a'
由于答案可能会很大，所以请你返回 模10^9 + 7之后的结果。
*/
func countVowelPermutation(n int) int {

	f := []int{1, 1, 1, 1, 1}

	mod := 1_000_000_007
	for i := 2; i <= n; i++ {
		a := f[0]
		e := f[1]
		i := f[2]
		//e->a   u->a i->a
		f[0] = (f[1] + f[4] + i) % mod
		//a->e i->e
		f[1] = (a + i) % mod
		//e->i  o->i
		f[2] = (e + f[3]) % mod
		f[4] = (i + f[3]) % mod
		f[3] = i

	}

	return (f[0] + f[1] + f[2] + f[3] + f[4]) % mod
}

func main() {
	fmt.Println(countVowelPermutation(2))
}
