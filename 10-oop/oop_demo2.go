package main

import "fmt"

/*
	定义类
	1）如果类名首字母大写，则表示该类能够访问
	2）如果属性名称大写，则该属性可以对外访问，否则只能该类内部访问
*/
type Hero struct {
	Name  string
	Ad    int
	Level int
}

/*
说明：此种写法是值传递的

func (this Hero) Show()  {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this Hero) GetName() {
	fmt.Println("Name = ", this.Name)
}

func (this Hero) setName(newName string) {
	//this是调用该方法的对象的一个副本，修改对对象无效
	this.Name = newName
}

*/

func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this *Hero) GetName() {
	fmt.Println("Name = ", this.Name)
}

func (this *Hero) SetName(newName string) {
	//this是调用该方法的对象的一个副本，修改对对象无效
	this.Name = newName
}

func main() {
	//1.创建对象
	hero := Hero{Name: "张三", Ad: 20, Level: 1}
	hero.Show()

	//2.修改对象的值,通过指针，引用传递，可以修改
	hero.SetName("李四")
	hero.Show()
}
