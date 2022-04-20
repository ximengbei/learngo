package main

import "fmt"

//defer在return之后调用，后执行
func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

//return先执行
func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

/*
	知识点2：defer和return谁先执行
*/
func main() {
	returnAndDefer()
}
