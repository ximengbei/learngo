package main

import (
	"fmt"
	"learngo/5-init/lib1"
	"learngo/5-init/lib2"
)

func init()  {
	fmt.Println("main init()...")
}

func main() {
	lib1.Lib1test()
	lib2.Lib2test()
}
