package dira1

import (
	"fmt"
	"testing"
	"unsafe"
)

// golang 结构体占用字节数 打印

func TestA02(t *testing.T) {
	var arune rune
	ints := []int{1, 2, 3, 3, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1}
	ret := NumSet{
		MapInfo: make(map[rune]int),
	}

	fmt.Printf("map 的类型 %T map占中的字节数是 %d\n\n", ret, unsafe.Sizeof(ret))
	fmt.Printf("map 的类型 %T map占中的字节数是 %d\n\n", ret.MapInfo, unsafe.Sizeof(ret.MapInfo))
	fmt.Printf("ints 的类型 %T ints 占中的字节数是 %d\n\n", ints, unsafe.Sizeof(ints))
	fmt.Printf("arune 的类型 %T arune 占中的字节数是 %d\n\n", arune, unsafe.Sizeof(arune))

}

type NumSet struct {
	MapInfo map[rune]int
}
