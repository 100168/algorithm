package main

func findLatestTime(s string) string {

	b := []byte(s)

	if s[4] == '?' {
		b[4] = '9'
	}
	if s[3] == '?' {
		b[3] = '5'
	}

	if s[0] == '?' {
		if s[1] == '?' || s[1] < '2' {
			b[0] = '1'
		} else {
			b[0] = '0'
		}
	}
	if s[1] == '?' {
		if b[0] == '1' {
			b[1] = '1'
		} else {
			b[1] = '9'
		}
	}

	return string(b)

}
