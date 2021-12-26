package learning

import (
	"fmt"
	"testing"
)

type Animal interface {
	MyName()
}
type Cat struct {
}

type Dog struct {
}

func (cat Cat) MyName() {
	fmt.Println("my name is cat")
}

func (dog Dog) MyName() {
	fmt.Println("my name is dog")
}

func TestName(t *testing.T) {
	dog := Dog{}
	cat := Cat{}

	dog.MyName()
	cat.MyName()
}
