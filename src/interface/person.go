package main

import "fmt"

// Person接口, 带返回值的方法
type Person interface {
	name() string
	age() int
	sayHello()
}

// 男人
type Man struct {
}

// 必须实现接口中的所有方法
func (man Man) name() string {
	return "夏雨"
}
func (man Man) age() int {
	return 39
}
func (man Man) sayHello() {
	fmt.Printf("我是男人：%s，我的年龄是：%d\n", man.name(), man.age())
}

// 女人
type Woman struct {
	// 定义一些属性
	yanzhi string
}

func (woman Woman) name() string {
	return "袁泉"
}
func (woman Woman) age() int {
	return 36
}
func (woman Woman) sayHello() {
	fmt.Printf("我是女人：%s，我的年龄是：%d，颜值：%s\n", woman.name(), woman.age(), woman.yanzhi)
}

func main(){
	var person Person

	person = new(Man)
	fmt.Println(person.name())
	fmt.Println(person.age())
	person.sayHello()

	// 将一个Woman对象赋值给person
	var woman = Woman{yanzhi:"99分"}
	person = woman
	fmt.Println(person.name())
	fmt.Println(person.age())
	person.sayHello()

}