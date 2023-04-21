package Adapter

import "testing"

//20 P20 适配器模式
func TestAdaptee_SpecificExecute(t *testing.T) {
	adapter := Adapter{}
	adapter.Execute()
}