package main

import (
	"fmt"
	"sort"
)

/*
给出一个字符串数组words组成的一本英语词典。从中找出最长的一个单词，该单词是由words词典中其他单词逐步添加一个字母组成。若其中有多个可行的答案，则返回答案中字典序最小的单词。

若无答案，则返回空字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-word-in-dictionary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func longestWord(words []string) string {

	sort.Strings(words)
	//myList := make([]string, 0)
	//
	//for k =  {
	//
	//}
	//dict := map[string]bool{}
	//for _, k := range words {
	//	dict[k] = true
	//	if len(k) > len(longest) {
	//		longest = k
	//	}
	//}
	fmt.Println(words)
	return ""
}
func main() {
	words := []string{"w", "wo", "wor", "world", "worl", "z"}
	longestWord(words)
}
