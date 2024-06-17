package main

import "fmt"

func findWinningPlayer(skills []int, k int) int {

	cnt := 0
	index := 1

	type pair struct{ x, index int }

	newSkills := make([]pair, len(skills))
	for i := range skills {
		newSkills[i] = pair{skills[i], i}
	}
	for {
		if newSkills[0].x > newSkills[index].x {
			cnt++
			newSkills = append(newSkills, newSkills[index])
		} else {
			newSkills[0] = newSkills[index]
			cnt = 1
		}
		if cnt == k {
			return newSkills[0].index
		}
		index++
	}
}

func main() {
	fmt.Println(findWinningPlayer([]int{4, 2, 6, 3, 9}, 2))
}
