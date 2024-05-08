package main

/*
给你一个字符串 word 和一个字符串数组 forbidden 。

如果一个字符串不包含 forbidden 中的任何字符串，我们称这个字符串是 合法 的。

请你返回字符串 word 的一个 最长合法子字符串 的长度。

子字符串 指的是一个字符串中一段连续的字符，它可以为空。
*/
func longestValidSubstring(word string, forbidden []string) int {

	n := len(word)

	ans := 0

	t := new(trie)
	t.root = new(node)

	maxLen := 0
	for _, v := range forbidden {
		t.insert(v)
		maxLen = max(maxLen, len(v))
	}

	//最长就10个，优秀

	l := 0
	for r := 0; r < n; r++ {
	next:
		for {
			for i := max(r-maxLen+1, l); i <= r; i++ {
				if t.find(word[i : r+1]) {
					l++
					continue next
				}
			}
			break
		}
		ans = max(ans, r-l+1)

	}
	return ans
}

type trie struct {
	root *node
}

type node struct {
	next [26]*node
	path int
	end  int
}

func (t *trie) insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := int(word[i] - 'a')
		if cur.next[c] == nil {
			cur.next[c] = &node{}
		}
		cur = cur.next[c]
		cur.path++
	}
	cur.end++
}
func (t *trie) find(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := int(word[i] - 'a')
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return cur.end > 0
}
