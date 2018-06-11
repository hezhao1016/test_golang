package main

import (
	"./demo"  // 相对路径导入
	"fmt"
)

func main() {
	// dm才是真正的包名
	fmt.Println(dm.Name)
	fmt.Println(dm.VERSION)
	dm.SayHello()

	fmt.Println(dm.Max(1, 2))
}
