package main

import "fmt"

/*
*
桌子上有 n 个球，每个球的颜色不是黑色，就是白色。

给你一个长度为 n 、下标从 0 开始的二进制字符串 s，其中 1 和 0 分别代表黑色和白色的球。

在每一步中，你可以选择两个相邻的球并交换它们。

返回「将所有黑色球都移到右侧，所有白色球都移到左侧所需的 最小步数」。

11  000

输入：s = "100"
输出：2
解释：我们可以按以下方式将所有黑色球移到右侧：
- 交换 s[0] 和 s[1]，s = "010"。
- 交换 s[1] 和 s[2]，s = "001"。
可以证明所需的最小步数为 2 。

思路：
1.从左往右遍历并记录1出现次数（cnt），当0出现时则需要往右移cnt位
*/
func minimumSteps(s string) int64 {

	ans := 0
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cnt++
		} else {
			ans += cnt
		}
	}
	return int64(ans)

}

func main() {
	fmt.Println(minimumSteps("1100"))
}
