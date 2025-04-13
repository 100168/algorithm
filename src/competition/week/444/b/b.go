package main

type Router struct {
	n int
	p []pair

	pm map[pair]bool
}

type pair struct {
	s    int
	d    int
	time int
}

func Constructor(memoryLimit int) Router {

	r := Router{n: memoryLimit, p: make([]pair, 0), pm: make(map[pair]bool)}
	return r
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {

	p := pair{source, destination, timestamp}

	if len(this.p) == 0 || !this.pm[p] {
		this.p = append(this.p, p)
		this.pm[p] = true
		if len(this.p) > this.n {
			pp := this.p[0]
			this.p = this.p[1:len(this.p)]
			delete(this.pm, pp)

		}
		return true
	}
	return false

}

func (this *Router) ForwardPacket() []int {

	if len(this.p) == 0 {
		return []int{}
	}
	p := this.p[0]
	this.p = this.p[1:]

	return []int{p.s, p.d, p.time}

}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {

	p := this.p
	l := this.find(startTime)
	r := this.find(endTime+1) - 1

	if r < l {
		return 0
	}

	ans := 0
	for _, v := range p[l : r+1] {
		if v.d == destination {
			ans++
		}
	}
	return ans

}

func (this *Router) find(t int) int {

	l, r := 0, len(this.p)-1
	for l <= r {
		mid := (l + r) / 2

		if this.p[mid].time < t {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
func main() {
	router := Constructor(3)
	router.AddPacket(5, 2, 110)

	println(router.GetCount(5, 110, 110))

}
