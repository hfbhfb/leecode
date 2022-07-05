package Memento

import "testing"

//10 P10 备忘录模式
func TestMemento(t *testing.T) {
	n:= NewNUmber(10)
	n.Double()
	memento:=n.CreateMemento()
	n.Half()
	n.ReinstateMemento(memento)
}