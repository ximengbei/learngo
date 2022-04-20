package main

import "fmt"

/*
	channel的使用，无缓冲的
	1.channel会使管道的两边保持同步，通过阻塞的方式保证同步
*/
func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("guroutine正在运行...")

		c <- 666 //将666发送给c
	}()

	num := <-c //从c中接受数据，赋值给num
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}
