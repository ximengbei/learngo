package main

import "fmt"

func main()  {
	//1.声明切片，并初始化赋值
	slice1 := []int{1,2,3}
	fmt.Println("length=",len(slice1),",values=",slice1)

	//2.声明切片，然后分配空间
	var slice2 []int
	slice2 = make([]int, 5)
	fmt.Println("length=",len(slice2),",values=",slice2)

	//3.声明切片，同时分配空间
	var slice3 []int = make([]int, 10)
	fmt.Println("length=",len(slice3),",values=",slice3)

	//4.声明切片，同事分配空间，:=推导出是切片类型
	slice4 := make([]int,20)
	fmt.Println("length=",len(slice4),",values=",slice4)

	//5.判断一个slice是否为0
	var slice5 []int
	if(slice5==nil){
		fmt.Println("slice5 是一个空切片")
	}else{
		fmt.Println("slice5 是有空间的")
	}
}