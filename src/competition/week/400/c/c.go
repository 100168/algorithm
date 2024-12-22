package main

import (
	"container/list"
	"fmt"
)

/*
*
给你一个字符串 s 。它可能包含任意数量的 '*' 字符。你的任务是删除所有的 '*' 字符。

当字符串还存在至少一个 '*' 字符时，你可以执行以下操作：

删除最左边的 '*' 字符，同时删除该星号字符左边一个字典序 最小 的字符。如果有多个字典序最小的字符，你可以删除它们中的任意一个。
请你返回删除所有 '*' 字符以后，剩余字符连接而成的
字典序最小

	的字符串。
*/
func clearStars1(s string) string {

	cnt := make([][]int, 26)

	indexs := make([]int, 0)

	for i := range cnt {
		cnt[i] = make([]int, 0)
	}
	delIndex := make(map[int]bool)
	for i := range s {
		if s[i] == '*' {
			indexs = append(indexs, i)
		} else {
			cur := int(s[i] - 'a')
			cnt[cur] = append(cnt[cur], i)
		}
	}

	for _, c := range indexs {
		for _, v := range cnt {
			if len(v) == 0 || v[0] > c {
				continue
			}
			l, r := 0, len(v)-1
			for l <= r {
				m := (r + l) / 2
				if v[m] < c {
					l = m + 1
				} else {
					r = m - 1
				}
			}
			delIndex[v[r]] = true
			break
		}
	}

	ans := ""

	for i := range s {
		if delIndex[i] || s[i] == '*' {
			continue
		}
		ans += string(s[i])
	}
	return ans

}

// 使用list.List来模拟栈
func clearStars(s string) string {
	stack := list.New() // 使用内置的list.List来作为栈
	foundStar := false

	for _, c := range s {
		if c != '*' {
			// 如果不是'*'，并且栈不为空且栈顶元素小于等于当前字符
			// 或者栈顶元素为'*'且已经处理过'*'（foundStar为true）
			for stack.Len() > 0 &&
				(foundStar || stack.Back().Value.(rune) <= c) {
				stack.Remove(stack.Back()) // 弹出栈顶元素
			}
			stack.PushBack(c) // 将当前字符压入栈中
			foundStar = false // 重置foundStar
		} else {
			foundStar = true
			if stack.Len() > 0 {
				stack.Remove(stack.Back()) // 弹出栈顶元素（'*'及其左边字典序最小的字符）
			}
		}
	}

	// 将栈中剩余元素转为字符串（逆序）
	result := make([]rune, 0, stack.Len())
	for e := stack.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(rune))
	}

	// 因为栈是逆序的，所以需要反转结果字符串
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func main() {

	fmt.Println(clearStars("d*o*"))
}
