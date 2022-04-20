package main

import (
	"encoding/json"
	"fmt"
)

type Moive struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

/*
	json解析
*/
func main() {
	moive := Moive{"喜剧之王", 2000, 10, []string{"张三", "李四", "王麻子"}}

	//1.结构体转json
	jsonStr, err := json.Marshal(moive)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}
	fmt.Printf("jsonstr = %s \n", jsonStr) //{"title":"喜剧之王","year":2000,"price":10,"actors":["张三","李四","王麻子"]}

	//2.json转结构体
	mymoive := Moive{}
	err = json.Unmarshal(jsonStr, &mymoive)
	if err != nil {
		fmt.Println("json Unmarshal error")
		return
	}
	fmt.Printf("%v \n", mymoive)
}
