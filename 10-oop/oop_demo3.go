package main

import "fmt"

/*1.定义父类*/
type Human struct {
	name string
	sex  string
}
func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

/*2.定义子类*/
type SuperMan struct {
	Human
	level int
}
//重写父类方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}
//重写父类方法
func (this *SuperMan) Walk() {
	fmt.Println("SuperMan.Walk()...")
}
//定义子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}
func (this *SuperMan) Show() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {

	//1.定义父类对象
	human := Human{"张三", "female"}
	human.Eat()
	human.Walk()

	//2.定义子类对象
	//superman := SuperMan{Human{"李四","female"},99}
	var superman SuperMan
	superman.name = "王麻子"
	superman.sex = "male"
	superman.level = 100
	superman.Eat()
	superman.Walk()
	superman.Fly()
	superman.Show()
}
