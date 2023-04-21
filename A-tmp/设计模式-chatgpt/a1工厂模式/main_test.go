package main

import (
	"testing"
)

import "fmt"

type Animal interface {
	Eat()
}

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("Cat is eating")
}

type Dog struct{}

func (d *Dog) Eat() {
	fmt.Println("Dog is eating")
}

type AnimalFactory struct{}

func (af *AnimalFactory) CreateAnimal(animalType string) Animal {
	switch animalType {
	case "cat":
		return &Cat{}
	case "dog":
		return &Dog{}
	default:
		return nil
	}
}

func main() {
	factory := AnimalFactory{}
	cat := factory.CreateAnimal("cat")
	dog := factory.CreateAnimal("dog")

	cat.Eat() // output: Cat is eating
	dog.Eat() // output: Dog is eating
}

func TestGetInstance(t *testing.T) {
	// t.Parallel()
	// 两次获取实例应该返回同一个对象
	main()
}
