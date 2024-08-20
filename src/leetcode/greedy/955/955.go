package main

import "fmt"

/*
给定由 n 个字符串组成的数组 strs，其中每个字符串长度相等。

选取一个删除索引序列，对于 strs 中的每个字符串，删除对应每个索引处的字符。

比如，有 strs = ["abcdef", "uvwxyz"]，删除索引序列 {0, 2, 3}，删除后 strs 为["bef", "vyz"]。

假设，我们选择了一组删除索引 answer，那么在执行删除操作之后，最终得到的数组的元素是按 字典序（strs[0] <= strs[1] <= strs[2] ... <= strs[n - 1]）排列的，然后请你返回 answer.length 的最小可能值。

示例 1：

输入：strs = ["ca","bb","ac"]
输出：1
解释：
删除第一列后，strs = ["a", "b", "c"]。
现在 strs 中元素是按字典排列的 (即，strs[0] <= strs[1] <= strs[2])。
我们至少需要进行 1 次删除，因为最初 strs 不是按字典序排列的，所以答案是 1。
*/
func minDeletionSize(strs []string) int {
	m, n := len(strs), len(strs[0])
	buf := make([][]byte, m) // 记录每一行可以保留的字符
	for i := 0; i < m; i++ {
		buf[i] = []byte{}
	}
	for j := 0; j < n; j++ {
		buf[0] = append(buf[0], strs[0][j])
		i := 1
		for i < m {
			buf[i] = append(buf[i], strs[i][j])
			if string(buf[i]) < string(buf[i-1]) {
				break
			}
			i++
		}
		if i < m { //说明遇到非字典序的情况，要删掉这一列
			for k := 0; k <= i; k++ {
				buf[k] = buf[k][:len(buf[k])-1]
			}
		}
	}

	return n - len(buf[0])
}

//ebodtwc",
//neoytml",
//neuyzlu",
//pnmtivg"]

func main() {
	//fmt.Println(minDeletionSize([]string{"xga", "xfb", "yfa"}))
	//fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
	//fmt.Println(minDeletionSize([]string{"bbjwefkpb", "axmksfchw"}))
	fmt.Println(minDeletionSize([]string{"jsebodtwc", "cnneoytml", "eeneuyzlu", "ewpnmtivg"}))
}
