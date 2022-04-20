package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))   //获取type
	fmt.Println("value: ", reflect.ValueOf(arg)) //获取具体的值
}

/*
	静态类型反射
*/
func main() {
	//1.定义变量
	var num float64 = 1.2345
	//2.反射
	reflectNum(num)
}
