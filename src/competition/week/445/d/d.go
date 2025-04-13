package main

import (
	"math/big"
	"strings"
)

const mod = 1e9 + 7

func countNonDecreasing(l, r string, b int) int {
	lNum, _ := new(big.Int).SetString(l, 10)
	rNum, _ := new(big.Int).SetString(r, 10)

	rStr := convertToBase(rNum, b)
	maxLen := len(rStr)

	countR := countInBase(rNum, b, maxLen)
	lMinus1 := new(big.Int).Sub(lNum, big.NewInt(1))
	var countL int
	if lMinus1.Sign() < 0 {
		countL = 0
	} else {
		countL = countInBase(lMinus1, b, maxLen)
	}

	return (countR - countL + mod) % mod
}

func countInBase(n *big.Int, base, maxLen int) int {
	if n.Sign() < 0 {
		return 0
	}

	s := convertToBaseWithMaxLen(n, base, maxLen)
	digits := make([]int, maxLen)
	for i, c := range s {
		digits[i] = int(c - '0')
	}

	memo := make([][]int, maxLen)
	for i := range memo {
		memo[i] = make([]int, base)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(pos, pre int, isLimit, isNum bool) int
	dp = func(pos, pre int, isLimit, isNum bool) int {
		if pos == maxLen {
			if isNum {
				return 1
			}
			return 0
		}
		if !isLimit && isNum && memo[pos][pre] != -1 {
			return memo[pos][pre]
		}

		res := 0
		up := digits[pos]
		if !isLimit {
			up = base - 1
		}

		if !isNum {
			res += dp(pos+1, 0, isLimit && 0 == up, true)
			res %= mod

			for d := 1; d <= up; d++ {
				res += dp(pos+1, d, isLimit && d == up, true)
				res %= mod
			}
		} else {
			start := pre
			if start < 0 {
				start = 0
			}
			for d := start; d <= up; d++ {
				res += dp(pos+1, d, isLimit && d == up, true)
				res %= mod
			}
		}

		if isNum && !isLimit {
			memo[pos][pre] = res % mod
		}

		return res % mod
	}

	return dp(0, 0, true, false)
}

func convertToBase(n *big.Int, base int) string {
	if n.Sign() == 0 {
		return "0"
	}
	var digits []rune
	baseBig := big.NewInt(int64(base))
	zero := big.NewInt(0)
	num := new(big.Int).Set(n)
	for num.Cmp(zero) > 0 {
		mod := new(big.Int)
		num.DivMod(num, baseBig, mod)
		digits = append(digits, rune(mod.Int64()+'0'))
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return string(digits)
}

func convertToBaseWithMaxLen(n *big.Int, base, maxLen int) string {
	s := convertToBase(n, base)
	if len(s) < maxLen {
		s = strings.Repeat("0", maxLen-len(s)) + s
	}
	return s
}
