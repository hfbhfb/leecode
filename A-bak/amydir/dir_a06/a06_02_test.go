package dira06

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// with cancel Context
// 返回closed channel,值bool为false
func TestYxx(t *testing.T) {
	c, close := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Second)
		close()
	}()

	select {
	case v, f := <-c.Done():
		fmt.Println(v)
		fmt.Println(f)
	}
	fmt.Println("end")

}
