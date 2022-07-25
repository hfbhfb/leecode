package main

import (
	"fmt"
	"log"
	"runtime"
	"testing"
	"time"
)

// cpu 100%测试
func TestCPU100(t *testing.T) {
	fmt.Println("111")

	ticker := time.NewTicker(300 * time.Second)

	fmt.Println(runtime.NumCPU())
	a := 0
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				a++
			}
		}()
	}

	select {
	case <-ticker.C:
		log.Println("Next")
	}

}
func TestCPUhello(t *testing.T) {
	fmt.Println("111")
}
