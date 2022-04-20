package main

import (
	"fmt"
	"runtime"
	"time"
)

/*

 */
func main() {
	//用go创建承载一个形参为空，返回值为空的函数
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			//退出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}() //表示调用了函数

		fmt.Println("A")
	}() //表示调用了函数

	//用go创建承载一个有参函数
	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
