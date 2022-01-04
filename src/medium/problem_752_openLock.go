package main

import "fmt"

/*你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
每个拨轮可以自由旋转：例如把 '9' 变为'0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。

锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。

列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。

字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/open-the-lock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func openLock(deadends []string, target string) int {
	const start = "0000"
	if target == start {
		return 0
	}

	dead := map[string]bool{}
	for _, s := range deadends {
		dead[s] = true
	}
	if dead[start] {
		return -1
	}

	// 枚举 status 通过一次旋转得到的数字
	get := func(status string) (ret []string) {
		s := []byte(status)
		for i, b := range s {
			s[i] = b - 1
			if s[i] < '0' {
				s[i] = '9'
			}
			ret = append(ret, string(s))
			s[i] = b + 1
			if s[i] > '9' {
				s[i] = '0'
			}
			ret = append(ret, string(s))
			s[i] = b
		}
		return
	}

	type pair struct {
		status string
		step   int
	}
	q := []pair{{start, 0}}
	seen := map[string]bool{start: true}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range get(p.status) {
			if !seen[nxt] && !dead[nxt] {
				if nxt == target {
					return p.step + 1
				}
				seen[nxt] = true
				q = append(q, pair{nxt, p.step + 1})
			}
		}
	}
	return -1
}

func openLock2(deadends []string, target string) int {
	const start = "0000"
	if target == start {
		return 0
	}

	ans := 0
	dead := map[string]bool{}
	for _, s := range deadends {
		dead[s] = true
	}
	if dead[start] {
		return -1
	}
	hashGo := make(map[string]bool)

	var que []string

	que = append(que, start)
	for len(que) > 0 {

		n := len(que)

		for i := 0; i < n; i++ {
			key := que[0]
			que = que[1:]

			if !dead[key] && !hashGo[key] {
				if key == target {
					return ans
				}
				hashGo[key] = true
				for i := 0; i < 3; i++ {
					que = append(que, add(key, i))
					que = append(que, decrease(key, i))
				}

			}
		}
		ans++

	}

	return -1
}
func add(target string, index int) string {
	s := []byte(target)

	if s[index] == '9' {
		s[index] = '0'
	} else {
		s[index] = s[index] + 1
	}
	target = string(s)
	return target
}
func decrease(target string, index int) string {
	s := []byte(target)

	if s[index] == '0' {
		s[index] = '9'
	} else {
		s[index] = s[index] - 1
	}
	target = string(s)
	return target
}
func main() {
	s := "9"
	fmt.Println(s[0])
	fmt.Println((1 << 63) - 1)
}
