package main

import (
	"fmt"
	"runtime"
	"time"
)

//var wg sync.WaitGroup

func main() {

	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go hello(i)
	//}
	//wg.Wait()

	fmt.Printf("\n------- 设置最大cpu 执行数字 -------\n")
	runtime.GOMAXPROCS(2)
	go hello1("A")
	go hello1("B")

	time.Sleep(time.Second)
	fmt.Println("hello, main... over..")
}

func hello1(s string) {
	for i := 0; i < 100; i++ {
		fmt.Println("hello->", s)
	}
}

func hello(num int) {
	//defer wg.Done()
	fmt.Println("hello goroutine! -->", num)
}
