package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//关闭channel,如果不关闭，则会报死锁的错误
		close(c)
	}()

	for {
		//ok为true表示channel没有关闭，如果为false表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main Finished...")
}
