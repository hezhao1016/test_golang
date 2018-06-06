package main

import "fmt"

// 变量

// 全局变量
var msg string = "我是全局变量"
var totalCount = 0

// 多变量声明
var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {
	// 变量声明
	// 第一种，指定变量类型，声明后若不赋值，使用默认值。
	var name string
	name = "张三"

	//第二种，根据值自行判定变量类型。
	var age = 23

	// 第三种，:=为简短声明，编译器根据初始化的值自动推断相应的类型，但是不能定义在函数外部使用。
	// 这是使用变量的首选形式, 使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。
	// 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。
	sex := true

	// Tips：声明但未使用的局部变量会在编译阶段报错。但是全局变量是允许声明但不使用。
	// var n = 2

	fmt.Println(name, age, sex)


	// 多变量声明
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)

}