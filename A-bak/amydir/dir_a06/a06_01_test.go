package dira06

import (
	"context"
	"fmt"
	"testing"
)

func DoneEnpty() <-chan struct{} {
	return nil
}

// 等一个nil channel永远不会得到数据
func TestXxx(t *testing.T) {
	b := context.Background()

	select {
	case <-b.Done():
		fmt.Println("done background")
	case <-DoneEnpty():
		fmt.Println("done background")
	}

	fmt.Println("end")
}
