package main

import "fmt"

// 变量作用域

/*
Go 语言中变量可以在三个地方声明：

函数内定义的变量称为局部变量
函数外定义的变量称为全局变量
函数定义中的变量称为形式参数,即函数形参

初始值:
int		0
float32	0
pointer	nil
*/

// 在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。
/* 声明全局变量 */
var g int


func main() {
	// 在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。
	/* 声明局部变量 */
	var a, b int

	/* 初始化参数 */
	a = 10
	b = 20
	g = a + b

	fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)


	// Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。
	var g int = 10
	fmt.Printf ("结果： g = %d\n",  g)


	c := sum( a, b)
	fmt.Printf("main()函数中 c = %d\n",  c)
}


/* 函数形参 */
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n",  a)
	fmt.Printf("sum() 函数中 b = %d\n",  b)
	return a + b
}