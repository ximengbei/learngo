package main

import "fmt"

//全局变量
var f int      //方法1
var g int = 5  //方法2
var h = 6      //方法3
//i := 7       //方法4，无法声明

/*
	四种变量的声明方式
	1）方法1、方法2、方法3可以声明全局变量
	2）方法4用来声明局部变量，只能用于函数体内
*/
func main(){
	//方法1：声明一个变量，默认值是0
	var a int;
	fmt.Println("方法1：a=",a);
	fmt.Printf("数据类型： a = %T \n",a)

	//方法2：声明一个变量，初始化一个值
	var b int = 1
	fmt.Println("方法2：b=",b)
	fmt.Printf("数据类型： b = %T \n",b)

	var bb string = "hello world"
	fmt.Println("方法2：bb=",bb)
	fmt.Printf("数据类型： bb = %T \n",bb)

	//方法3：在初始化的时候，可以省去数据类型，go自动匹配数据类型
	var c = 2;
	fmt.Println("方法3：c=",c)
	fmt.Printf("数据类型： c = %T \n",c)

	//方法4：省去var关键字，自动匹配数据类型，:=
	d := 3
	fmt.Println("方法4：d=",d)
	fmt.Printf("数据类型： d = %T \n",d)

	e := "who are you"
	fmt.Println("方法4：e=",e)
	fmt.Printf("数据类型： e = %T \n",e)

	//全局变量
	fmt.Println("全局变量：f=",f)
	fmt.Println("全局变量：g=",g)
	fmt.Println("全局变量：h=",h)

	//声明多个变量
	var xx, yy int = 11,22
	fmt.Println("xx=",xx,", yy=",yy)

	var oo, qq = 33, "what is you name?"
	fmt.Println("oo=",oo,", qq=",qq)

	//多行声明变量
	var (
		mm int = 50
		nn bool = true
	)
	fmt.Println("mm=",mm,", nn=",nn)
}