package main

import "fmt"

/*
*
给你一个只包含小写英文字母的字符串 s 。

每一次 操作 ，你可以选择 s 中两个 相邻 的字符，并将它们交换。

请你返回将 s 变成回文串的 最少操作次数 。

注意 ，输入数据会确保 s 一定能变成一个回文串。

示例 1：

输入：s = "aabb"
输出：2
解释：
我们可以将 s 变成 2 个回文串，"abba" 和 "baab" 。
- 我们可以通过 2 次操作得到 "abba" ："aabb" -> "abab" -> "abba" 。
- 我们可以通过 2 次操作得到 "baab" ："aabb" -> "abab" -> "baab" 。
因此，得到回文串的最少总操作次数为 2 。

思路:贪心

1.奇数情况可以变成偶数情况，碰到奇数不处理，先建偶数部分处理完成再将奇数情况加进去
*/
func minMovesToMakePalindrome(s string) int {

	n := len(s)

	bytes := []byte(s)

	ans := 0

	l, r := 0, n-1
	for l < r {
		right := r
		for right > l && bytes[right] != bytes[l] {
			right--
		}
		if right > l {
			temp := bytes[right]
			ans += r - right
			swap(right, r, bytes, 1)
			bytes[r] = temp
			r--
		} else {
			ans += n/2 - l
		}
		l++

	}
	fmt.Println(string(bytes))
	return ans
}

func swap(start int, end int, bytes []byte, odd int) {
	for i := start; i != end; i += odd {
		bytes[i] = bytes[i+odd]
	}
}

func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

func minMovesToMakePalindrome2(s string) int {
	n := len(s)
	sBytes := []byte(s)
	l := 0
	r := n - 1
	var res int
	for l < r {
		ptrR := r
		for sBytes[l] != sBytes[ptrR] && l < ptrR {
			ptrR--
		}
		if l < ptrR {
			for i := ptrR; i < r; i++ {
				sBytes[i], sBytes[i+1] = sBytes[i+1], sBytes[i]
				res++
			}
			r--
		} else {
			res += n/2 - l
		}
		l++
	}
	return res
}

// letctel
func main() {
	//fmt.Println(minMovesToMakePalindrome("aabb"))
	//fmt.Println(minMovesToMakePalindrome("aca"))
	//fmt.Println(minMovesToMakePalindrome("leteltc"))
	fmt.Println(minMovesToMakePalindrome("skwhhaaunskegmdtutlgtteunmuuludii"))
}
