package Abstract_Factory

import "testing"


// 02 抽象工厂模式
func TestNewSimpleShapeFactory(t *testing.T) {
	factory := NewSimpleShapeFactory()
	food := factory.CreateFood()
	food.Cook()

	vegetable := factory.CreateVegetable()
	vegetable.Cook()
}
