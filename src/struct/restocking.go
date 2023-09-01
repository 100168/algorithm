package _struct

// Store /**商店id和商店补货日*/
type Store struct {
	id string
	//到货提前日
	daysInAdvanceOfArrival int
	//补货日列表
	restockDays []RestockDay
}

type OrderDay struct {
	//下单日期
	days []int
	//状态
	status OrderStatus
}
type RestockDay struct {
	day int
	//提前期
	daysInAdvance int
}

type OrderStatus int

const (
	Open OrderStatus = iota
	Closed
)

func calculateRestockDay(stores Store, orderDay OrderDay) {

}
