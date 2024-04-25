package main

func distanceTraveled(mainTank int, additionalTank int) int {

	ans := 0
	for mainTank/5 > 0 {
		ans += (mainTank - mainTank%5) * 10
		cost := min(mainTank/5, additionalTank)
		additionalTank -= cost
		mainTank = mainTank%5 + cost
	}
	return ans + mainTank*10

}
func distanceTraveled2(mainTank int, additionalTank int) int {
	return 10 * (mainTank + min((mainTank-1)/4, additionalTank))
}
