// package可以不和文件夹名相同
package dm

import "fmt"

// 全局变量
var Name = "demo"
var remark = "备注，小写字母开头的变量和函数是不可导出的，是private的。"

// 常量
const VERSION = "v0.0.1"
const msg = "该常量私有"

// 小写开头的函数是private的，外部无法访问
func sayHello() {
	fmt.Println("hello")
}

func SayHello()  {
	fmt.Println("Hello")
}