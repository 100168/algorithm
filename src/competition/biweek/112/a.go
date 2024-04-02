package main

func canBeEqual(s1 string, s2 string) bool {

	bS1 := []int32(s1)
	bS2 := []int32(s2)
	if len(bS1) != len(bS2) {
		return false
	}

	for i := 0; i < len(s2); i++ {
		if bS2[i] != bS1[i] {
			if i+2 < len(s2) {
				bS1[i], bS1[i+2] = bS1[i+2], bS1[i]
				if bS1[i] != bS2[i] {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}
