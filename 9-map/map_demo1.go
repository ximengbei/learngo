package main

import "fmt"

/*
	map的声明方式
*/
func main() {
	//1.第一种声明方式：定义map,key是string，value是string
	var map1 map[string]string
	fmt.Println(map1)
	//使用map前需要分配数据空间
	map1 = make(map[string]string, 10)
	map1["1"] = "java"
	map1["2"] = "go"
	map1["3"] = "c"
	fmt.Println(map1)

	//2.第二种声明方式
	map2 := make(map[string]string)
	map2["1"] = "a"
	map2["2"] = "b"
	map2["3"] = "c"
	fmt.Println(map2)

	//3.第三种声明方式
	map3 := map[string]string{
		"1": "hello",
		"2": "go",
		"3": "!",
	}
	fmt.Println(map3)
}
