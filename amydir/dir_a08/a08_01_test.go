package dira08

import (
	"strconv"
	"testing"
)

// ## 一个示例
//   * 在上面的代码中，golang 的参数传递过程是：
//     * 1,分配一块内存 p， 并且将对象 b 的内容拷贝到 p 中；
//     * 2,创建 iface 对象 i，将 i.tab 赋值为 itab<Stringer, Binary>。将 i.data 赋值为 p；
//     * 3,使用 i 作为参数调用 test 函数。

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

type Stringer interface {
	String() string
}

func test(s Stringer) {
	s.String()
}

func TestXxx(t *testing.T) {
	b := Binary(0x123)
	test(b)
}
