package main

import (
	"fmt"
	"time"
)

//从goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d \n", i)
		time.Sleep(1 * time.Second)
	}
}

//主goroutine
func main() {
	//1.创建一个go程，去执行newTask()流程
	go newTask()

	//主方法的go程
	i := 0
	for {
		i++
		fmt.Printf("main Goroutine : i = %d \n", i)
		time.Sleep(1 * time.Second)
	}
}
