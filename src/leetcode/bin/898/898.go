package main

func subarrayBitwiseORs(arr []int) int {

	ans := make([]int, 0)
	for _, v := range arr {
		ans = append(ans, v)
		newAns := make([]int, 0)
		for _, e := range ans {
			cur := e | v
			if len(newAns) > 0 && newAns[len(newAns)-1] == cur {
				continue
			}
			newAns = append(newAns, cur)
		}
		ans = newAns
	}
	return len(ans)
}

func subarrayBitwiseORs2(arr []int) int {

	ans := make([]int, 0)
	ansMap := make(map[int]bool)
	for _, v := range arr {
		ans = append(ans, v)
		k := 0
		for _, e := range ans {
			cur := e | v
			if ans[k] == cur {
				continue
			}
			k++
			ans[k] = cur
			ansMap[cur] = true
		}
		ans = ans[:k+1]
	}
	return len(ansMap)
}

func main() {
	println(subarrayBitwiseORs([]int{1, 2, 4}))
	println(subarrayBitwiseORs2([]int{1, 2, 4}))
}
