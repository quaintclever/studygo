package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var x int64
var lock sync.Mutex

func main() {
	fmt.Printf("\n------- 并发安全和锁 -------\n")
	wg.Add(2)
	go Add()
	go Add()
	wg.Wait()
	fmt.Println(x)

}

func Add() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	wg.Done()
}
