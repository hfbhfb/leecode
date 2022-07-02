package main

import (
	"log"
	"reflect"
	"testing"
	"unsafe"
)

func TestSliceCopy(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	s2 := s1[:5]
	s3 := s1[0:0]
	s4 := s1[0:1]
	log.Println(s1)
	log.Println(s2)
	log.Println(s3)
	log.Println(s4)
}

func TestSliceCopyPointAnalyse(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	s2 := s1[:5]
	s3 := s1[4:5]
	s4 := s1[0:1]
	log.Println(s1)
	log.Println(s2)
	log.Println(s3)
	log.Println(s4)

	bptrs1 := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	bptrs2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
	bptrs3 := (*reflect.SliceHeader)(unsafe.Pointer(&s3))
	bptrs4 := (*reflect.SliceHeader)(unsafe.Pointer(&s4))
	log.Println(bptrs1.Data)
	log.Println(bptrs2.Data)
	log.Println(bptrs3.Data)
	log.Println(bptrs4.Data)

}
