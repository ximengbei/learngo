package lib1

import "fmt"

//函数名首字母大写表示函数是公开的，外部可以调用
func Lib1test() {
	fmt.Println("lib1test")
}

func init() {
	fmt.Println("lib1 init()...")
}
