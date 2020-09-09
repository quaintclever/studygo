package main

import "fmt"

func main() {

	fmt.Printf("\n------- channel操作 -------\n")
	ch := make(chan int, 0)
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- 1

	fmt.Printf("\n------- channel 2 -------\n")
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for {
			if i, ok := <-ch1; ok {
				ch2 <- i * i
			} else {
				break
			}
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}

}
