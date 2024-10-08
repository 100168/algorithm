package main

/*
*
假设你是一个专业的狗仔，参加了一个 n 人派对，其中每个人被从 0 到 n - 1 标号。
在这个派对人群当中可能存在一位 “名人”。所谓 “名人” 的定义是：其他所有 n - 1 个人都认识他/她，而他/她并不认识其他任何人。

现在你想要确认这个 “名人” 是谁，或者确定这里没有 “名人”。而你唯一能做的就是问诸如 “A 你好呀，请问你认不认识 B呀？” 的问题，以确定 A 是否认识 B。你需要在（渐近意义上）尽可能少的问题内来确定这位 “名人” 是谁（或者确定这里没有 “名人”）。

在本题中，你可以使用辅助函数 bool knows(a, b) 获取到 A 是否认识 B。请你来实现一个函数 int findCelebrity(n)。

派对最多只会有一个 “名人” 参加。若 “名人” 存在，请返回他/她的编号；若 “名人” 不存在，请返回 -1。

进阶：如果允许调用 API knows 的最大次数为 3 * n ，你可以设计一个不超过最大调用次数的解决方案吗？

名人:所有人都认识他，他不认识任何人

[[1,1],
[1,1]]

[[1,0],[0,1]]
*/
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {

		celebrity := make([]bool, n)
		for i := range celebrity {
			celebrity[i] = true
		}
		cnt := 0
		for i := 0; i < n; i++ {
			if !celebrity[i] {
				continue
			}
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				if !knows(j, i) {
					celebrity[i] = false
					break
				} else {
					celebrity[j] = false
				}
			}

			if celebrity[i] {
				cnt++
			}
			if cnt > 1 {
				return -1
			}
		}
		for i := range celebrity {
			if celebrity[i] {
				for j := 0; j < n; j++ {
					if j == i {
						continue
					}
					if knows(i, j) {
						return -1
					}
				}
				return i
			}
		}
		return -1
	}
}

func solution2(knows func(a int, b int) bool) func(n int) int {
	isCelebrity := func(i int, n int) bool {
		for j := 0; j < n; j++ {
			if i == j {
				continue // 他们认识自己就不用询问。
			}
			if knows(i, j) || !knows(j, i) {
				return false
			}
		}
		return true
	}

	return func(n int) int {
		celebrityCandidate := 0
		for i := 0; i < n; i++ {
			if knows(celebrityCandidate, i) {
				celebrityCandidate = i
			}
		}
		if isCelebrity(celebrityCandidate, n) {
			return celebrityCandidate
		}
		return -1
	}
}
