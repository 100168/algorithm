package main

func removeStars(s string) string {

	var st []byte

	for _, v := range s {
		if v == '*' {
			st = st[:len(st)-1]
			continue
		}
		st = append(st, byte(v))
	}
	return string(st)
}
