package dira07

import (
	"fmt"
	"testing"
)

// map在并发执行中会直接报错
func TestNxx(t *testing.T) {
	m := make(map[int]int)
	go func() { //开一个协程写map
		for i := 0; i < 10000; i++ {
			m[i] = i
		}
	}()

	go func() { //开一个协程读map
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
	}()

	//time.Sleep(time.Second * 20)
	for {

	}
}
