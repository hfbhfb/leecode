package dira10

import (
	"fmt"
	"math"
	"testing"
	"unsafe"

	"github.com/imroc/biu"
)

func TestZxx(t *testing.T) {
	var f32 float32 = 1.21
	var i32 int32 = 121

	fmt.Println(math.Pi)

	// fmt.Println(biu.ToBinaryString(f32))  // panic: data type is unsupported
	fmt.Println(biu.ToBinaryString(i32))

	fmt.Println(unsafe.Sizeof(f32))
	fmt.Println(unsafe.Sizeof(i32))
}
