package main

import "fmt"

/*
	切片截取和copy
*/
func main() {

	//1.定义切片
	slice1 := []int{1, 2, 3} //leng=3,cap=3
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(slice1), cap(slice1), slice1)

	//2.截取切片
	slice1_1 := slice1[0:2]
	slice1_1[0] = 100
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(slice1_1), cap(slice1_1), slice1_1)

	//3.copy可以将底层数组的slice一起拷贝
	slice1_2 := make([]int, 3)
	copy(slice1_2, slice1)
	fmt.Printf("length = %d, cap = %d, slice = %v\n", len(slice1_2), cap(slice1_2), slice1_2)
}
