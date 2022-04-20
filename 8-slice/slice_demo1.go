package main

import "fmt"

//打印数组，注意数组传参的长度需要和定义的一样
//是值传递，方法内不能改变数组中的值
func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index=", index, ",value=", value)
	}
}

/*
	固定数组的长度是固定的
*/
func main() {

	//1.固定长度数组
	var myArray1 [10]int
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	//2.固定长度数组，部分赋值
	myArray2 := [10]int{1, 2, 3, 4}
	for index, value := range myArray2 {
		fmt.Println("index=", index, ",value=", value)
	}

	//3.数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)

	//4.传递参数，必须保证数组长度一致
	myArray3 := [4]int{1, 2, 3, 4}
	printArray(myArray3)
}
