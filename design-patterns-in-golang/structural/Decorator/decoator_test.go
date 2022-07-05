package Decorator

import (
	"fmt"
	"log"
	"os"
	"testing"
)

//22 P22 装饰器模式  (示例很好, 封装了一下日志的重定向和计算时间 )
func TestPi(t *testing.T) {
	fmt.Println(Pi(5000))
	fmt.Println(Pi(10000))
	fmt.Println(Pi(50000))

}

func TestWrapLogger(t *testing.T) {
	f:= WrapLogger(Pi,log.New(os.Stdout,"test ",1))
	f(100000)
}