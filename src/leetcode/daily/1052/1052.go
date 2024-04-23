package main

func maxSatisfied(customers []int, grumpy []int, minutes int) int {

	n := len(customers)

	user := make([]int, n+1)
	unUser := make([]int, n+1)

	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + customers[i-1]
	}
	for i := 1; i <= n; i++ {
		if grumpy[i-1] == 1 {
			unUser[i] = unUser[i-1]
			user[i] = max(unUser[max(i-minutes, 0)]+s[i]-s[max(i-minutes, 0)], user[i-1])
		} else {
			unUser[i] = unUser[i-1] + customers[i-1]
			user[i] = max(unUser[max(i-minutes, 0)]+s[i]-s[max(i-minutes, 0)], user[i-1]+customers[i-1])
		}
	}
	return max(unUser[n], user[n])
}

func main() {
	println(maxSatisfied([]int{9, 10, 4, 5}, []int{1, 0, 1, 1}, 1))

	//10000    100   100 100 100 100 100 100
}
