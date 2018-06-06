package main

import "fmt"

// Go 语言基础语法

func main()  {
	// Go 标记
	// Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。
	fmt.Println("Hello, World!")

	// 行分隔符
	// 在 Go 程序中，一行代表一个语句结束。每个语句不需要像Java一样以分号 ; 结尾，因为这些工作都将由 Go 编译器自动完成。
	// 如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法。
	fmt.Println("Hello");fmt.Println("Golang!")

	// 单行注释
	/* 我是多行注释 */

	// import 导包
	/*
	import "fmt"最常用的一种形式

	import "./test"导入同一目录下test包中的内容

	import f "fmt"导入fmt，并给他启别名f

	import . "fmt"，将fmt启用别名"."，这样就可以直接使用其内容，而不用再添加ｆｍｔ，如fmt.Println可以直接写成Println

	import  _ "fmt" 表示不使用该包，而是只是使用该包的init函数，并不显示的使用该包的其他内容。注意：这种形式的import，当import时就执行了fmt包中的init函数，而不能够使用该包的其他函数。
	*/

	// const 关键字来进行常量的定义。

	// 在函数体外部使用 var 关键字来进行全局变量的声明和赋值。

	// type 关键字来进行结构(struct)和接口(interface)的声明。

	// func 关键字来进行函数的声明。

	// 可见性规则
	// Go语言中，使用大小写来决定该常量、变量、类型、接口、结构或函数是否可以被外部包所调用。
	// 函数名首字母小写即为 private
	// 函数名首字母大写即为 public
}