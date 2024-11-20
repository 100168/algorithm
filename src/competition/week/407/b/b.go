package main

import "strings"

/*
小红和小明在玩一个字符串元音游戏。

给你一个字符串 s，小红和小明将轮流参与游戏，小红 先 开始：

在小红的回合，她必须移除 s 中包含 奇数 个元音的任意 非空
子字符串
。
在小明的回合，他必须移除 s 中包含 偶数 个元音的任意 非空
子字符串
。
第一个无法在其回合内进行移除操作的玩家输掉游戏。假设小红和小明都采取 最优策略 。

如果小红赢得游戏，返回 true，否则返回 false。

英文元音字母包括：a, e, i, o, 和 u。*/

func doesAliceWin(s string) bool {

	pattern := "aeiou"
	cnt := 0

	for _, v := range s {
		if strings.ContainsRune(pattern, v) {
			cnt++
		}
	}

	if cnt == 0 {
		return false
	}
	return true
}
