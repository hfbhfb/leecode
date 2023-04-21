package singleton

import (
	"testing"
)

type SingleObject struct {
	// 可以在这里定义 SingleObject 的属性和方法
}

// 定义 instance 变量，在第一次使用 GetInstance() 函数时创建 SingleObject 对象
var instance *SingleObject

// GetInstance() 函数返回 SingleObject 对象的唯一实例
func GetInstance() *SingleObject {
	if instance == nil {
		instance = &SingleObject{}
	}
	return instance
}

func TestGetInstance(t *testing.T) {
	// 两次获取实例应该返回同一个对象
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Errorf("GetInstance() returned different instances, expected same instance")
	}
}
