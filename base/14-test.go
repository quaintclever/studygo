package main

import (
	"fmt"
	"math/rand"
)

// 1. 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
//a. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
//b. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//c. 主goroutine从resultChan取出结果并打印到终端输出
func main() {

	jobChan := make(chan int64, 24)
	resultChan := make(chan int64, 24)

	go func() {
		for i := 0; i < 24; i++ {
			ran := rand.Int63()
			jobChan <- ran
			fmt.Println("随机数:", ran)
		}
		close(jobChan)
	}()

	for i := 0; i < 24; i++ {
		go func() {
			i := <-jobChan
			var res int64 = 0
			for ; i > 0; i /= 10 {
				res += i % 10
			}
			resultChan <- res
		}()
	}

	for i := 0; i < 24; i++ {
		fmt.Println(<-resultChan)
	}
	close(resultChan)
}
