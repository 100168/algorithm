package main

type subArray struct {
	//index = 排名第几，value = 字符
	sa []int
	//index = 字符 value 排名第几
	rk []int
}

func sufferArray(s string) {

	n := len(s)

	sb := new(subArray)
	sb.sa = make([]int, 10000)
	sb.rk = make([]int, 10000)

	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		cur := s[i] - 'a'
		cnt[cur]++
	}
	for i := 1; i < 26; i++ {
		cnt[i] += cnt[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		cur := s[i] - 'a'
		sb.sa[cnt[cur]-1] = i
		sb.rk[i] = cnt[cur-1]
		cnt[cur]--
	}

	for exp := 1; 1<<exp < n; exp <<= 1 {
		x := make([]int, 10000)
		y := make([]int, 10000)
		copy(sb.sa, x)
		copy(sb.rk, y)
		cnt = make([]int, 26)
		for i := 0; i < n; i++ {
			cnt[y[i]+exp]++
		}
		for i := 1; i < 26; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			//当前字符
			cur := s[i+1<<exp] - 'a'
			//当前字符排名
			//
			sb.sa[cnt[cur]-1] = i
			sb.rk[i] = cnt[cur] - 1
			cnt[cur]--
		}
	}
}
