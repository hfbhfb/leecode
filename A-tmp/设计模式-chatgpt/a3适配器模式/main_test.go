package adapter

// 测试代码
import (
	"fmt"
	"testing"
)

// Adaptee 是待适配的接口
type Adaptee interface {
	SpecificRequest() string
}

// AdapteeImpl 是实现 Adaptee 接口的结构体
type AdapteeImpl struct{}

func (AdapteeImpl) SpecificRequest() string {
	return "adaptee method called"
}

// Target 是目标接口，适配器将 Adaptee 转换为 Target 接口
type Target interface {
	Request() string
}

// Adapter 是适配器结构体，实现 Target 接口
type Adapter struct {
	adaptee Adaptee
}

func (a Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func TestAdapter(t *testing.T) {
	adaptee := AdapteeImpl{}
	adapter := Adapter{adaptee}
	expected := "adaptee method called"
	res := adapter.Request()

	if res != expected {
		t.Errorf("got %s, expected %s", res, expected)
	}
}

// 运行测试
// go test

func TestMyFunction(t *testing.T) {
	// 取消输出缓冲
	// https://stackoverflow.com/a/19576600/11689125
	t.Parallel()
	t.Run("test", func(t *testing.T) {

		// 打印串口信息
		fmt.Println("串口打印输出")

		// 这里执行需要测试的代码
	})
}
