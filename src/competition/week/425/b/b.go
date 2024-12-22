package main

func isPossibleToRearrange(s string, t string, k int) bool {

	cnt := make(map[string]int)

	k = len(s) / k
	for i := k; i <= len(s); i += k {
		cnt[s[i-k:i]]++
	}
	for i := k; i <= len(s); i += k {
		cur := t[i-k : i]
		if _, ok := cnt[cur]; !ok {
			return false
		}
		cnt[cur]--
		if cnt[cur] == 0 {
			delete(cnt, cur)
		}
	}
	return true

}
func main() {
}
