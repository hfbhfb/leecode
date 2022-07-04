package main

import (
	"fmt"
	"testing"
)

func TestTypefun(t *testing.T) {

	two(22, wraptypefun(callback))
}

//需要传递函数
func callback(i int, j int) {
	// fmt.Println("i am callBack")
	// fmt.Println(i)
	// fmt.Println(j)
	fmt.Println(i + j)
}

/*
//main中调用的函数
func one(i int, f func(int, int)) {
	two(i, wraptypefun(f))
}
*/

//one()中调用的函数
func two(i int, c CallInterface) {
	c.call(i)
}

//定义的type函数
type wraptypefun func(int, int)

//fun实现的Call接口的call()函数
func (fworparam wraptypefun) call(i int) {
	fworparam(i, i+1)
}

//接口
type CallInterface interface {
	call(int)
}
