package dira04

import (
	"fmt"
	"testing"
	"time"
)

// 如果非要从消费端关闭channel
func TestCxx(t *testing.T) {

	ch := make(chan time.Time, 200)
	closed := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-closed:
				t.Log("sender quit")
				return
			default:
				ch <- time.Now()
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			now := <-ch
			fmt.Println(now.Format("2006-01-01 15:00:00"))
			if i == 50 {
				closed <- true
				close(ch)
				return
			}
		}
	}()

	time.Sleep(time.Second)
	time.Sleep(time.Second)
	// <-make(chan bool)
}
