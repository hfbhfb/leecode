package dira09

import "runtime"

func M07() []int {
	m := make(chan int, 1)
	m <- 1
	_ = <-m
	runtime.CPUProfile()
	return nil
}
