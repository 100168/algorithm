package main

import (
	"fmt"
	"strings"
)

/*给出第一个词first 和第二个词second，考虑在某些文本text中可能以 "first second third"
形式出现的情况，其中second紧随first出现，third紧随second出现。

对于每种这样的情况，将第三个词 "third" 添加到答案中，并返回答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/occurrences-after-bigram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func findOcurrences(text string, first string, second string) []string {

	split := strings.Split(text, " ")

	var ans []string

	for i := range split {
		if i+2 < len(split) {

			if split[i] == first && second == split[i+1] {
				ans = append(ans, split[i+2])
			}
		}
	}
	return ans
}

func main() {
	text := "alice is a good girl she is a good student"

	ocurrences := findOcurrences(text, "a", "good")

	fmt.Println(ocurrences)
}
