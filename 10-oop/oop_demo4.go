package main

import "fmt"

//1.定义接口
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}

//2.定义类，实现接口
type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("cat is Sleep()...")
}
func (this *Cat) GetColor() string {
	fmt.Println("cat is GetColor()...")
	return this.color
}
func (this *Cat) GetType() string {
	fmt.Println("cat is GetType()...")
	return "Cat"
}

//3.定义类，实现接口
type Dog struct {
	color string //猫的颜色
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep()...")
}
func (this *Dog) GetColor() string {
	fmt.Println("Dog is GetColor()...")
	return this.color
}
func (this *Dog) GetType() string {
	fmt.Println("Dog is GetType()...")
	return "Dog"
}

//4.
func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color=", animal.GetColor())
	fmt.Println("type=", animal.GetType())
}

/*
	多态
*/
func main() {

	//1.定义接口对象
	var animal AnimalIF

	//2.指向Cat的实现，调用的是子Cat的sleep方法
	animal = &Cat{"Greent"}
	animal.Sleep()

	fmt.Println("-------------------------")

	//3.指向Dog的实现，调用的是子Dog的sleep方法
	animal = &Dog{"Yellow"}
	animal.Sleep()

	fmt.Println("-------------------------")

	//4.通过调用方法打印
	cat := Cat{"red"}
	dog := Dog{"black"}
	showAnimal(&cat)
	fmt.Println("-------------------------")
	showAnimal(&dog)

}
