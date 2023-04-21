package dira06

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//
func TestZxx(t *testing.T) {
	c, cancel := context.WithTimeout(context.Background(), time.Second)
	_ = cancel

	select {
	case v, f := <-c.Done():
		fmt.Println(v)
		fmt.Println(f)
	}
	fmt.Println("end")

}
