package main

func longestDupSubstring(s string) string {

	N := 30010
	x := make([]uint8, N)
	y := make([]int, N)

	sa := make([]int, n)
	rk := make([]int, n)
	height := make([]int, n)
	n := len(s)

	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		x[i] = s[i]
		cur := s[i] - 'a'
		cnt[cur]++
	}
	for i := 1; i < 26; i++ {
		cnt[i] += cnt[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		cur := s[i] - 'a'
		sa[cnt[cur]] = i
		cnt[cur]--
	}

	for exp := 1; exp <= n; exp <<= 1 {
		m := 0
		cnt = make([]int, 26)
		copy(y, sa)

		//根据第二位排序
		for i := 0; i < n; i++ {
			cnt[x[y[i]+exp]-'a']++
		}
		for i := 1; i < 26; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			sa[cnt[x[y[i]+exp]]-'a'] = y[i]
			cnt[x[y[i]+exp]-'a']--
		}
		copy(y, sa)
		cnt = make([]int, 26)

		//根据第一位排序
		for i := 0; i < n; i++ {
			cnt[x[y[i]]-'a']++
		}
		for i := 1; i < 26; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			sa[cnt[x[y[i]]-'a']] = y[i]
			cnt[x[y[i]]-'a']--
		}
		t := make([]uint8, N)
		copy(t, x)
		for i := 1; i <= n; i++ {
			if t[sa[i]] == t[sa[i-1]] && t[sa[i]+exp] == t[sa[i-1]+exp] {
				x[sa[i]] = uint8(m)
			} else {
				x[sa[i]] = uint8(m + 1)
				m++
			}

		}
		if m == n {
			break
		}
	}
	for i := 1; i <= n; i++ {
		rk[sa[i]] = i
	}
	for i, k := 1, 0; i <= n; i++ {
		if rk[i] == 1 {
			continue
		}
		if k > 0 {
			k--
		}
		j := sa[rk[i]-1]

		for i+k <= n && j+k <= n && s[i+k] == s[j+k] {
			k++
		}
		height[rk[i]] = k
	}
	return s[:N]
}
