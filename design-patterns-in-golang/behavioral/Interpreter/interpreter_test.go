package Interpreter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//18 P18 解释器模式介绍
func TestInterpreter(t *testing.T) {
	expression := "w x z +"
	sentence := NewEvaluator(expression)
	variables := make(map[string]Expression)
	variables["w"] = &Integer{6}
	variables["x"] = &Integer{10}
	variables["z"] = &Integer{41}
	result := sentence.Interpret(variables)
	assert.Equal(t, 51,result)
}
