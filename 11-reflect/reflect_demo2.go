package main

import (
	"fmt"
	"io"
	"os"
)

/*具体的类型*/
func main() {

	//tty:pair<type:*os.File,value:"/dev/tty">
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error ", err)
	}

	//r:pair<type:,value:>
	var r io.Reader
	//r:pair<type:*os.File,value:"/dev/tty">
	r = tty

	//w:pair<type:,value:>
	var w io.Writer
	//w:pair<type:*os.File,value:"/dev/tty">
	w = r.(io.Writer) //强制转换

	w.Write([]byte("hello this is a test!\n"))
}
