package main

func maxTaxiEarnings(n int, rides [][]int) int64 {

	type pair struct{ start, gold int }
	groups := make([][]pair, n+1)
	for _, offer := range rides {
		start, end, gold := offer[0], offer[1], offer[2]
		groups[end] = append(groups[end], pair{start, end - start + gold})
	}

	f := make([]int, n+2)
	for end, g := range groups {
		f[end+1] = f[end]
		for _, p := range g {
			f[end+1] = max(f[end+1], f[p.start+1]+p.gold)
		}
	}
	return int64(f[n+1])
}
