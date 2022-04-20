package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		//监听多个channel的状态
		select {
		case c <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacii(c, quit)
}
