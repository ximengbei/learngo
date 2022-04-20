package main

import "fmt"

//const定位枚举类型，iota只能出现在const关键词里面，只能配合const使用
const (
	//iota,默认值是0，每行累加1
	sz = iota * 10
	cd
	xa
)

//iota每行加1，中途可以改变公式
const (
	a, b = iota + 1, iota + 2 //iota=0,a=1,b=2
	c, d                      //iota=1,a=2,b=3
	e, f                      //iota=2,a=3,b=4
	g, h = iota * 2, iota * 3 //iota=3,a=6,b=9
	j, k                      //iota=4,a=8,b=12
)

/*
	常量
*/
func main() {
	//1、常量是只读的
	const height int = 100
	fmt.Println("height=", height)
	//height = 200 //常量不能修改，报错：cannot assign to height (constant 100 of type int)

	//2、枚举
	fmt.Println("sz=", sz)
	fmt.Println("cd=", cd)
	fmt.Println("xa=", xa)

	fmt.Println("a=", a, ",b=", b)
	fmt.Println("c=", c, ",d=", d)
	fmt.Println("e=", e, ",f=", f)
	fmt.Println("g=", g, ",h=", h)
	fmt.Println("j=", j, ",k=", k)
}
