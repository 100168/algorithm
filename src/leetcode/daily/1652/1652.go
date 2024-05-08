package main

func decrypt(code []int, k int) []int {

	n := len(code)

	s := make([]int, n+1)

	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + code[i-1]
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if k > 0 {
			if i+1+k <= n {
				ans[i] = s[i+1+k] - s[i+1]
			} else {
				ans[i] = s[n] - s[i+1] + s[k+i-n+1]

			}
		} else if k == 0 {
			ans[i] = 0
		} else {
			if i+k >= 0 {
				ans[i] = s[i] - s[i+k]
			} else {
				ans[i] = s[i] + s[n] - s[n-(-k-i)]
			}
		}

	}
	return ans
}
