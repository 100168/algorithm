package main

/*
累加数 是一个字符串，组成它的数字可以形成累加序列。

一个有效的 累加序列 必须 至少 包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。

给你一个只包含数字'0'-'9'的字符串，编写一个算法来判断给定输入是否是 累加数 。如果是，返回 true ；否则，返回 false 。

说明：累加序列里的数 不会 以 0 开头，所以不会出现1, 2, 03 或者1, 02, 3的情况。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/additive-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func stringAdd(x, y string) string {
	res := []byte{}
	carry, cur := 0, 0
	for x != "" || y != "" || carry != 0 {
		cur = carry
		if x != "" {
			cur += int(x[len(x)-1] - '0')
			x = x[:len(x)-1]
		}
		if y != "" {
			cur += int(y[len(y)-1] - '0')
			y = y[:len(y)-1]
		}
		carry = cur / 10
		cur %= 10
		res = append(res, byte(cur)+'0')
	}
	for i, n := 0, len(res); i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return string(res)
}

func valid(num string, secondStart, secondEnd int) bool {
	n := len(num)
	firstStart, firstEnd := 0, secondStart-1
	for secondEnd <= n-1 {
		third := stringAdd(num[firstStart:firstEnd+1], num[secondStart:secondEnd+1])
		thirdStart := secondEnd + 1
		thirdEnd := secondEnd + len(third)
		if thirdEnd >= n || num[thirdStart:thirdEnd+1] != third {
			break
		}
		if thirdEnd == n-1 {
			return true
		}
		firstStart, firstEnd = secondStart, secondEnd
		secondStart, secondEnd = thirdStart, thirdEnd
	}
	return false
}

func isAdditiveNumber(num string) bool {
	n := len(num)
	for secondStart := 1; secondStart < n-1; secondStart++ {
		if num[0] == '0' && secondStart != 1 {
			break
		}
		for secondEnd := secondStart; secondEnd < n-1; secondEnd++ {
			if num[secondStart] == '0' && secondStart != secondEnd {
				break
			}
			if valid(num, secondStart, secondEnd) {
				return true
			}
		}
	}
	return false
}
