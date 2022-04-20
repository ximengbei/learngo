package main

import "fmt"

/*
	切片的追加
	1）len长度是指左指针到右指针之间的距离
	2）cap容量表示左指针到底层数组末尾的额距离
	3）当容量满时，自动追加一倍
*/
func main() {

	//1.定义切片并赋值，3是长度，5是容量
	var numbers = make([]int, 3, 5)
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//2.向切片中追加值
	numbers = append(numbers, 100)
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 200)
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//3.向容量满的slice追加元素，会自动分配初始化时给出的容量
	numbers = append(numbers, 300)
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

}
