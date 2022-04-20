package main

import (
	"fmt"
	"reflect"
)

//定义类
type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", this)
}

//获取对象的属性和方法
func ShowFiledAndMethod(input interface{}) {
	//获取input的类型
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is :", inputType.Name())
	fmt.Println("----------------------------------")

	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value is :", inputValue)
	fmt.Println("----------------------------------")

	//通过type获取对象里面的字段
	//1.获取到interface的reflect.Type,通过Type得到NumField，进行遍历
	//2.得到每个field，数据类型
	//3.通过field有一个interface（）方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)              //获取字段类型
		value := inputValue.Field(i).Interface() //获取字段的值
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	fmt.Println("----------------------------------")

	//通过type获取里面的方法(无法打印出来值，待调试)
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v \n", m.Name, m.Type)
	}
	fmt.Println("----------------------------------")
}

/*
	类对象反射，通过反射获取对象的字段和方法
*/
func main() {
	user := User{1, "张三", 30}
	ShowFiledAndMethod(user)
}
