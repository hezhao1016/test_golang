// 每个golang源代码文件开头都拥有一个package声明，表示该golang代码所属的package。
// 如果是为了将代码编译成一个可执行程序，那么package必须是main
// 如果是为了将代码编译成库，那么package则没有限制
package main

import "fmt"

// 注意：在一个可执行程序只有一个main函数
func main() {
	// 输出语句
	fmt.Println("Hello, World!")
}