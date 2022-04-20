package main

import "fmt"

//打印数组，传参为动态数组，动态数组是引用传递
func printArray(myArray []int) {
	//_表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value=", value)
	}
	myArray[0] = 100
}

/*
	动态数组slice长度不固定
*/
func main() {
	//1.动态数组，切片slice
	myArray1 := []int{1, 2, 3, 4}
	fmt.Printf("myArray1 types is %T\n", myArray1)

	//2.动态数组中的值在传递给的函数中可以改变
	fmt.Println("修改参数前：")
	printArray(myArray1)
	fmt.Println("修改参数后：")
	for _, value := range myArray1 {
		fmt.Println("value=", value)
	}

}
