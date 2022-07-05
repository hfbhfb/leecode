package Visitor

import "testing"

//09 P9 访问者模式
func TestElement_Accept(t *testing.T) {
	e:=new(Element)
	e.Accept(new(WeiBoVisitor))
	e.Accept(new(IQIYIVisitor))
}
