package main

import "fmt"

func answerString(word string, numFriends int) string {

	n := len(word)

	mx := 0

	if numFriends == 1 {
		return word
	}
	for i := range word {

		if int(word[i]-'a') > mx {
			mx = int(word[i] - 'a')
		}
	}

	ans := ""
	for i := range word {
		if int(word[i]-'a') != mx {
			continue
		}

		cur := ""

		end := i
		for j := i + 1; j < n; j++ {
			if i+(n-j)+1 > numFriends {
				end++
			}
		}
		cur = word[i : end+1]
		if ans < cur {
			ans = cur
		}
	}
	return ans

}
func main() {
	//fmt.Println(answerString("dbca", 2))
	fmt.Println(answerString("gh", 2))
}
