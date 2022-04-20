package main

import "fmt"

//定义结构体
type Book struct {
	name string
	auth string
}

//传递的是book的副本，无法修改
func changeBook(book Book) {
	book.auth = "李四"
}

//指针传递可以修改
func changeBook2(book *Book) {
	book.auth = "王麻子"
}

func main() {

	//1.使用结构体
	var book1 Book
	book1.name = "java编程思想"
	book1.auth = "张三"
	fmt.Println(book1)

	//2.值传递无法修改
	changeBook(book1)
	fmt.Println(book1)

	//3.指针传递值可以修改
	changeBook2(&book1)
	fmt.Println(book1)
}
