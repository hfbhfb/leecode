package State

import "testing"

//14 P14 状态模式
func TestState(t *testing.T) {
	machine := NewMachine()
	machine.Off()
	machine.On()
	machine.On()
	machine.Off()
}