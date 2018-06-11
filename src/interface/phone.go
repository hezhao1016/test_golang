package main

// Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

import "fmt"

// 定义Phone接口
type Phone interface {
	call()
}

// 定义诺基亚Phone结构体
type NokiaPhone struct {
}

// 实现call()方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

// 定义iPhone结构体
type IPhone struct {
}

// 实现call()方法
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	// 创建一个NokiaPhone实例，指向Phone接口对象
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}