package main

func maxBottlesDrunk(numBottles int, numExchange int) int {

	ans := numBottles

	for numBottles > 0 && numBottles >= numExchange {
		numBottles = numBottles - numExchange
		numExchange++
		numBottles++
		ans++
	}
	return ans
}
