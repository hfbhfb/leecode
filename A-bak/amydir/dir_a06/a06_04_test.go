package dira06

import (
	"context"
	"fmt"
	"testing"
)

//
func TestMxx(t *testing.T) {
	c := context.WithValue(context.Background(), "key1", "value1")

	k1 := c.Value("key1")
	k2 := c.Value("key2")
	fmt.Println(k1)
	fmt.Println(k2)

	fmt.Println("end")

}
