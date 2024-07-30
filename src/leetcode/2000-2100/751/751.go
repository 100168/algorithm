package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
*
IP地址 是一个格式化的 32位 无符号整数，每组 8位 被打印为一个十进制数字和，点字符 '.' 起到了分组的作用。

例如，二进制数 00001111 10001000 11111111 01101011 ( 为清晰起见增加了空格)被格式化为IP地址将是 “15.136.255.107” 。
CIDR块 是一种用来表示一组特定IP地址的格式。字符串形式，由基础IP地址、斜杠和前缀长度 k 组成。
它所覆盖的地址是所有IP地址的 前 k 位 与基础IP地址相同的IP地址。

例如， “123.45.67.89/20” 是一个前缀长度为 20 的 CIDR块。
任何二进制表示形式匹配 01111011 00101101 0100xxxx xxxxxxxx 的IP地址，其中 x 可以是 0 或 1 ，都在CIDR块覆盖的集合中。
给你一个起始IP地址 ip 和我们需要覆盖的IP地址数量 n 。
你的目标是使用 尽可能少的CIDR块 来 覆盖范围 [ip, ip + n - 1] 内的所有IP地址 。不应该覆盖范围之外的其他IP地址。

返回 包含IP地址范围的 CIDR块 的 最短 列表。如果有多个答案，返回其中 任何 一个 。

示例 1：

输入：ip = "255.0.0.7", n = 10
输出：["255.0.0.7/32","255.0.0.8/29","255.0.0.16/32"]
解释：
需要覆盖的IP地址有:
- 255.0.0.7 -> 11111111 00000000 00000000 00000111
- 255.0.0.8 -> 11111111 00000000 00000000 00001000
- 255.0.0.9 -> 11111111 00000000 00000000 00001001
- 255.0.0.10 -> 11111111 00000000 00000000 00001010
- 255.0.0.11 -> 11111111 00000000 00000000 00001011
- 255.0.0.12 -> 11111111 00000000 00000000 00001100
- 255.0.0.13 -> 11111111 00000000 00000000 00001101
- 255.0.0.14 -> 11111111 00000000 00000000 00001110
- 255.0.0.15 -> 11111111 00000000 00000000 00001111
- 255.0.0.16 -> 11111111 00000000 00000000 00010000
CIDR区块“255.0.0.7/32”包含第一个地址。
CIDR区块255.0.0.8/29包含中间的8个地址(二进制格式为11111111 00000000 00000000 00001xxx)。
CIDR区块“255.0.0.16/32”包含最后一个地址。
请注意，虽然CIDR区块“255.0.0.0/28”覆盖了所有的地址，但它也包括范围之外的地址，所以我们不能使用它。
*/
func ipToCIDR(ip string, n int) []string {

	split := strings.Split(ip, ".")

	low1, _ := strconv.Atoi(split[2])
	low2, _ := strconv.Atoi(split[3])

	end := low1<<8 | low2 + n - 1

	s := low1<<8 | low2

	for s < end {
		s += s & -s
	}

	return []string{}
}

func main() {
	fmt.Println(1<<10 - 1)
}