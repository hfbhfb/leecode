package dira04

import (
	"log"
	"testing"
	"time"
)

// 随机 select
func TestNxx(t *testing.T) {

	stopCh := make(chan struct{})
	ticker := time.NewTicker(100 * time.Microsecond)

	go func() {
		time.Sleep(2 * time.Second)
		close(stopCh)
	}()

	for {
		select {
		case <-stopCh:
			log.Println("Exit")
			return
		default:
		}

		time.Sleep(time.Second)
		log.Println("Working hard")

		select {
		case <-ticker.C:
			log.Println("Next")
		case <-stopCh:
			log.Println("Exit")
			return
		}

	}

}
