package main

// leetcode submit region begin(Prohibit modification and deletion)
func smallestNumber(num int64) int64 {
	flag := num < 0
	cnt := make([]int, 10)
	num = abs(num)
	cur := num
	for cur > 0 {
		mod := cur % 10
		cnt[mod]++
		cur /= 10
	}
	ans := int64(0)

	if !flag {
		for i := 1; i < 10; i++ {
			if cnt[i] > 0 {
				ans *= 10
				ans += int64(i)
				cnt[i]--
			}
			for cnt[0] > 0 && ans != 0 {
				ans *= 10
				cnt[0]--
			}
			for cnt[i] > 0 {
				ans *= 10
				ans += int64(i)
				cnt[i]--
			}
		}
	} else {
		for i := 9; i >= 0; i-- {
			for cnt[i] > 0 {
				ans *= 10
				ans += int64(i)
				cnt[i]--
			}
		}
	}
	if flag {
		return -ans
	}
	return ans

}
func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}
func main() {
	println(smallestNumber(-103))
	println(smallestNumber(3001))
}
