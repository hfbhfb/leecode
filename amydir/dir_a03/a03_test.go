package dira03

import (
	"fmt"
	"testing"
)

// slice 扩容引发的问题
func TestA03(t *testing.T) {

	x := []int{1, 2, 3}
	y := x[:2]
	fmt.Println(cap(y))
	y = append(y, 50)
	y = append(y, 60)
	fmt.Printf("x=%v\n", x)

	y[0] = 20
	fmt.Printf("x=%v\n", x)

}
