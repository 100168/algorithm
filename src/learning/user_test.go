package learning

import (
	"fmt"
	_ "testing"
)

type User struct {
	name string
	age  int8
}

func (u *User) setUser(name string, age int8) {
	u.name = name
	u.age = age
}

//指针调用，共享值
func (u *User) getUser() {
	fmt.Println("age={}", u.age)
	fmt.Println("name={}", u.name)
}

//值调用，创建一个临时副本
func (u User) setUser2(name string, a int8) {
	u.name = name
	u.age = a

	fmt.Println("user2", u.name)
	fmt.Println("user2-age", u.age)

}

func main() {
	a := User{
		name: " ",
		age:  10,
	}

	b := &a
	//a.setUser("a", 20)
	//a.getUser()

	//可以使用指针调用值方法
	fmt.Println("------")
	b.setUser2("b", 10)
	b.getUser()
}
