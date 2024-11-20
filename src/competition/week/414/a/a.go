package main

import (
	"strconv"
	"strings"
)

/*
*
给你一个字符串 date，它的格式为 yyyy-mm-dd，表示一个公历日期。

date 可以重写为二进制表示，只需要将年、月、日分别转换为对应的二进制表示（不带前导零）并遵循 year-month-day 的格式。

返回 date 的 二进制 表示。

示例 1：

输入： date = "2080-02-29"

输出： "100000100000-10-11101"

解释：

100000100000, 10 和 11101 分别是 2080, 02 和 29 的二进制表示。
*/
func convertDateToBinary(date string) string {

	split := strings.Split(date, "-")

	ans := ""
	for _, v := range split {

		atoi, _ := strconv.Atoi(v)
		formatInt := strconv.FormatInt(int64(atoi), 2)
		ans += formatInt
		ans += "-"
	}

	return ans[:len(ans)-1]

}

func main() {

}
