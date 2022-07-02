package main

import (
	"log"
	"reflect"
	"testing"
)

type StructHHH struct {
	Name string
	Age  int
}

func (a StructHHH) SetAAName(newName string) {
	a.Name = newName
	log.Println(a.Name)
}

func TestReflectHelloWorld(t *testing.T) {
	a := &StructHHH{Name: "hfb", Age: 18}
	t1 := reflect.TypeOf(a)
	log.Println(t1.Name(), t1.NumMethod())
}

func Invoke(any interface{}, name string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(inputs)
}

func TestReflectOptMethod(t *testing.T) {
	a := &StructHHH{Name: "hfb", Age: 18}

	func(any interface{}, name string, args ...interface{}) {
		inputs := make([]reflect.Value, len(args))
		for i, _ := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		reflect.ValueOf(any).MethodByName(name).Call(inputs)
	}(a, "SetAAName", "aaa")

	log.Println(a.Name)
}
