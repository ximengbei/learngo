package main

import "fmt"

/*
	interface是万能数据类型
*/
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called ...")
	fmt.Println(arg)

	//interface{}通过”类型断言“的机制，区分底层数据类型
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type,value=", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}
	myFunc(book)
	fmt.Println("-------------------------")
	myFunc(100)
	fmt.Println("-------------------------")
	myFunc("hello")
	fmt.Println("-------------------------")
	myFunc(false)
}
