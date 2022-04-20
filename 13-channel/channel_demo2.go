package main

import (
	"fmt"
	"time"
)

//有缓冲
func main() {
	c := make(chan int, 3)

	fmt.Println("len(c)=", len(c), "cap(c)=", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在运行,发送的元素=", i, " len(c)=", len(c), "cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}

	fmt.Println("main go程结束")
}
