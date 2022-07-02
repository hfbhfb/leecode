package main

import (
	"log"
	"reflect"
	"testing"
	"unsafe"
)

func TestStringCopy(t *testing.T) {
	s1 := new(string)
	s2 := "111111"
	*s1 = s2
	log.Println(*s1)
}

func TestStringCopyPointAnalyse(t *testing.T) {
	s1 := new(string)
	s2 := "111111"
	*s1 = s2

	bptrs1 := (*reflect.StringHeader)(unsafe.Pointer(s1))
	bptrs2 := (*reflect.StringHeader)(unsafe.Pointer(&s2))

	log.Println(bptrs1.Data)
	log.Println(bptrs2.Data)
}
