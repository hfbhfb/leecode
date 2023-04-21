package Mediator

import "testing"

//12  P12 中介模式
func TestMediator(t *testing.T) {
	mediator:=NewMediator()
	mediator.Ted.Talk()

}