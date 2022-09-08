package learning

import (
	"fmt"
	"testing"
)

type MsgError struct {
	dividee int
	divider int
}

//实现了错误中的error方法
func (msgError *MsgError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0`
	return fmt.Sprintf(strFormat, msgError.dividee)
}

func (msgError *MsgError) sql(a int, b int) (result int, errorMsg string) {
	if b == 0 {
		return 0, msgError.Error()
	}
	return b - a, ""
}

func TestError(t *testing.T) {

	m := MsgError{dividee: 10, divider: 0}

	if result, e := m.sql(1, -1); e == "" {
		fmt.Println(result)
	}
	if _, e := m.sql(1, 0); e != "" {
		fmt.Println(e)
	}
}
