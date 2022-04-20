package main

import "fmt"

//cityMap是引用传递
func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key=", key, ",value=", value)
	}
	cityMap["广东"] = "广州"
}

func main() {
	//1.定义map
	cityMap := make(map[string]string)
	fmt.Println(cityMap)

	//2.添加元素
	cityMap["广东"] = "深圳"
	cityMap["四川"] = "成都"
	cityMap["陕西"] = "西安"
	fmt.Println(cityMap)

	//遍历
	for key, value := range cityMap {
		fmt.Println("key=", key, ",value=", value)
	}

	//删除元素
	delete(cityMap, "四川")
	fmt.Println(cityMap)

	//传参
	printMap(cityMap)
	fmt.Println(cityMap)
}
