package main

import "fmt"

/*静态类型*/
func main() {
	var a string
	//pair<statictype:string,value:"abcde">
	a = "abcde"

	//pair<type:string,value:"abcde">
	var allType interface{}
	//a赋值给任何变量，里面的类型和值是不变的
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
