package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTick(t *testing.T) {
	t1 := time.NewTicker(1 * time.Second)
	tend := time.NewTicker(10 * time.Second)

ENDFOR:
	for {
		select {
		case <-t1.C:
			fmt.Println("ticker work here")
		case <-tend.C:
			break ENDFOR
		}
	}

}
