package dira10

import (
	"fmt"
	"testing"
)

func withPanic() {
	fmt.Println("1111")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
		// fmt.Println("222")
		// recover()
	}()
	m := 9
	n := 0
	a := m / n
	_ = a
	fmt.Println("3333")
}

func beforePanic() {
	fmt.Println("before")
}

func afterPanic() {
	fmt.Println("after")
}

func TestYxx(t *testing.T) {
	beforePanic()
	withPanic()
	afterPanic()
}
