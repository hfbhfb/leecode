package Proxy

import (
	"testing"
)

//21 P21 代理模式
func TestAgentTask_RentHouse(t *testing.T) {
	agent := NewAgentTask()
	agent.RentHouse("北京",8000)
}