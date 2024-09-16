package main

import (
	"fmt"
	"math"
)

// Aho-Corasick 自动机节点
type ACNode struct {
	edges  map[rune]int
	fail   int
	output []int
}

func newACNode() *ACNode {
	return &ACNode{edges: make(map[rune]int), output: []int{}}
}

// Aho-Corasick 自动机结构体
type AhoCorasick struct {
	nodes []*ACNode
	fail  []int
}

func newAhoCorasick() *AhoCorasick {
	root := newACNode()
	nodes := []*ACNode{root}
	fail := []int{0}
	return &AhoCorasick{nodes: nodes, fail: fail}
}

// 插入模式串到 Aho-Corasick 自动机
func (ac *AhoCorasick) insert(word string, index int) {
	node := 0
	for _, char := range word {
		if _, exists := ac.nodes[node].edges[char]; !exists {
			ac.nodes[node].edges[char] = len(ac.nodes)
			ac.nodes = append(ac.nodes, newACNode())
			ac.fail = append(ac.fail, 0)
		}
		node = ac.nodes[node].edges[char]
	}
	ac.nodes[node].output = append(ac.nodes[node].output, index)
}

// 构建 Aho-Corasick 自动机的失败链接
func (ac *AhoCorasick) build() {
	queue := []int{0}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for char, nextNode := range ac.nodes[node].edges {
			fail := ac.fail[node]
			for _, exists := ac.nodes[fail].edges[char]; !exists && fail != 0; fail = ac.fail[fail] {
			}
			if _, exists := ac.nodes[fail].edges[char]; exists {
				ac.fail[nextNode] = ac.nodes[fail].edges[char]
			} else {
				ac.fail[nextNode] = 0
			}
			ac.nodes[nextNode].output = append(ac.nodes[nextNode].output, ac.nodes[ac.fail[nextNode]].output...)
			queue = append(queue, nextNode)
		}
	}
}

// 查找最少有效字符串数量
func minValidStrings(words []string, target string) int {
	n := len(target)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	ac := newAhoCorasick()
	for i, word := range words {
		ac.insert(word, i)
	}
	ac.build()

	for i := 0; i < n; i++ {
		if dp[i] == math.MaxInt {
			continue
		}

		node := 0
		for j := i; j < n; j++ {
			char := rune(target[j])
			for _, exists := ac.nodes[node].edges[char]; !exists && node != 0; node = ac.fail[node] {
			}
			if _, exists := ac.nodes[node].edges[char]; exists {
				node = ac.nodes[node].edges[char]
			} else {
				node = 0
			}

			for _, wordIndex := range ac.nodes[node].output {
				wordLength := len(words[wordIndex])
				if i+wordLength <= n {
					dp[i+wordLength] = min(dp[i+wordLength], dp[i]+1)
				}
			}
		}
	}

	if dp[n] == math.MaxInt {
		return -1
	}
	return dp[n]
}

func main() {
	words := []string{"ab", "abc", "bc", "bca"}
	target := "abcabc"
	fmt.Println(minValidStrings(words, target)) // Output should be 3
}
