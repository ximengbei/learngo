package main

import "fmt"

/*
  知识点1：defer的执行顺序
*/
func main() {
	//1.defer表示在函数执行结束前执行
	//2.当有多个defer时，采用栈的方式执行，先进后出
	defer fmt.Println("main end 1...")
	defer fmt.Println("main end 2...")

	fmt.Println("main:: hello go 1")
}
