package dira04

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

var (
	wg       sync.WaitGroup
	execTime time.Duration = time.Second
)

func finishReq(timeout time.Duration) int {
	ch := make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(execTime)
		ch <- 42
	}()

	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		fmt.Println("Timeout")
		return -1
	}

}

func TestMxx(t *testing.T) {
	timeout := 50 * time.Microsecond

	log.Printf("Result:%d\n", finishReq(timeout))
	wg.Wait()

}
