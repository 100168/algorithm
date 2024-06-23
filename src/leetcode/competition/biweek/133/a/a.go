package main

import "strings"

func clearDigits(s string) string {

	st := make([]rune, 0)
	for _, v := range s {
		if strings.ContainsRune("0123456789", v) {
			st = st[:len(st)-1]
		} else {
			st = append(st, v)
		}
	}
	return string(st)
}
