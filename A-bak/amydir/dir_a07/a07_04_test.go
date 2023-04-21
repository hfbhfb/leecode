package dira07

import (
	"fmt"
	"testing"
)

// 7053 : 9919
// 7054 : 9920
// 7055 : 9921
// --- PASS: TestMxx (0.03s)
// 没有到9999,说明有数据丢失

var s []int

func appendValue(i int) {
	s = append(s, i)
}

func TestMxx(t *testing.T) {
	for i := 0; i < 10000; i++ { //10000个协程同时添加切片
		go appendValue(i)
	}

	for i, v := range s { //同时打印索引和值
		fmt.Println(i, ":", v)
	}
}
