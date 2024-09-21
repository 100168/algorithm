package main

type trie struct {
	root *node
}

type node struct {
	son  [26]*node
	len  int
	fail *node
	last *node
}

func (t *trie) insert(word string) {
	root := t.root
	cur := root
	for _, v := range word {
		c := v - 'a'
		if cur.son[c] == nil {
			cur.son[c] = &node{}
		}
		cur = cur.son[c]
	}
	cur.len = len(word)
}

func (t *trie) buildFail() {

	t.root.fail = t.root
	t.root.last = t.root
	var q []*node
	for i, son := range t.root.son[:] {
		if son == nil {
			t.root.son[i] = t.root
			continue
		}
		son.fail = t.root
		son.last = t.root
		q = append(q, son)
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i, son := range cur.son[:] {
			if son == nil {
				cur.son[i] = cur.fail.son[i]
				continue
			}
			son.fail = cur.fail.son[i]
			//为什么需要记录last，因为fail节点不一定是单词节点
			if son.fail.len > 0 {
				son.last = son.fail
			} else {
				son.last = son.fail.last
			}
			q = append(q, cur.son[i])
		}
	}
}

func main() {

}
