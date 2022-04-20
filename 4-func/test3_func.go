package main

import "fmt"


//1、返回单个值
func foo1(a string , b int ) int {
	fmt.Println("a=",a,",b=",b)
	return 0
}
//2、匿名返回多个参数
func foo2(a string , b int ) (int,int) {
	fmt.Println("a=",a,",b=",b)
	return 0,1
}
//3、返回有名称的多个参数
func foo3(a string , b int ) ( r1 int, r2 int) {
	fmt.Println("a=",a,",b=",b)

	//r1,r2属于foo3的形参，初始默认0，作用空间在函数体内
	fmt.Println("r1=",r1)
	fmt.Println("r2=",r2)

	r1=10
	r2=20
	return
}

//4、返回有名称的多个参数，返回值写法变化
func foo4(a string , b int ) ( r1 ,r2 int) {
	fmt.Println("a=",a,",b=",b)
	r1=30
	r2=40
	return
}

/*
*/
func main()  {

	//1、单个返回值
	result := foo1("hello",1)	
	fmt.Println("result=",result)	

	//2、匿名多个返回值
	result1,result2 := foo2("hello",1)	
	fmt.Println("result1=",result1,",result2=",result2)	

	//3、有名称多个返回值
	result3,result4 := foo3("hello",1)	
	fmt.Println("result3=",result3,",result4=",result4)	

	//4、有名称多个返回值
	result5,result6 := foo4("hello",1)	
	fmt.Println("result5=",result5,",result6=",result6)	
}