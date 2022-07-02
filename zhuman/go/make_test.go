package main

import (
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}

func TestMakeNoBufferChan(t *testing.T) {
	nobufferchan := make(chan interface{})
	go func() {
		for {
			time.Sleep(1 * time.Second)
			log.Println("push something to nobufferchan")
			nobufferchan <- 1
		}
	}()
	time.Sleep(8 * time.Second)
	log.Println("start wait nobufferchan")
	<-nobufferchan
	log.Println("end wait nobufferchan")
}

func TestMakeBufferChan(t *testing.T) {
	nobufferchan := make(chan interface{}, 100)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			log.Println("push something to nobufferchan")
			nobufferchan <- 1
		}
	}()
	time.Sleep(8 * time.Second)
	log.Println("start wait nobufferchan")
	<-nobufferchan
	log.Println("end wait nobufferchan")
}

func TestMakeSlicePanic(t *testing.T) {
	myslice := make([]int, 100)
	log.Println(cap(myslice))
	myslice[200] = 1
	log.Println("end wait nobufferchan")
}

func TestMakeSliceMax(t *testing.T) {
	myslice := make([]int, 100)
	myslice[len(myslice)-1] = 1
	log.Println("end wait nobufferchan")
}

func TestMakeSliceWichCapPanic(t *testing.T) {
	myslice := make([]int, 100, 600)
	log.Println(cap(myslice))
	myslice[200] = 1
	log.Println("end wait nobufferchan")
}
