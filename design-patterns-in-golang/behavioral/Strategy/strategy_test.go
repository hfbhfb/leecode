package Strategy

import "testing"

//15  P15 策略模式
func TestContext_Execute(t *testing.T) {
	strategyA:= NewStrategyA()
	c:=NewContext()
	c.SetStrategy(strategyA)
	c.Execute()

	strategyB := NewStrategyB()
	c.SetStrategy(strategyB)
	c.Execute()
}
