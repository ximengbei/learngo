package main

import "fmt"

//定义reader接口
type Reader interface {
	ReaderBook()
}

//定义writer接口
type Writer interface {
	WriterBook()
}

//定义类并实现两个接口
type Book struct {
}

func (this *Book) ReaderBook() {
	fmt.Println("Reader a Book")
}
func (this *Book) WriterBook() {
	fmt.Println("Writer a Book")
}

/*具体类型示例*/
func main() {
	//b:pair<type:Book,value:book{}的地址>
	book := &Book{}

	var r Reader
	//r:pair<type:Book,value:book{}的地址>
	r = book

	r.ReaderBook()

	var w Writer
	//r:pair<type:Book,value:book{}的地址>
	w = r.(Writer)

	w.WriterBook()
}
