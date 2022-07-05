package Factory

import "testing"

// 01 这是工厂模式
func TestNewRestaurant(t *testing.T) {
	NewRestaurant("d").GetFood()
	NewRestaurant("q").GetFood()
}
