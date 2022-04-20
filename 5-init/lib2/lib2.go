package lib2

import "fmt"

//函数名首字母大写表示函数是公开的，外部可以调用
func Lib2test() {
	fmt.Println("lib2test")
}

func init() {
	fmt.Println("lib2 init()...")
}
