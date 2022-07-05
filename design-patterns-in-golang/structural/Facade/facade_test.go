package Facade

import "testing"

//03  P3 外观模式
func TestCarFacade_CreateCompleteCar(t *testing.T) {
	facade := NewCarFacade()
	facade.CreateCompleteCar()
}
