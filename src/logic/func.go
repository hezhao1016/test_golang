package main

import (
	"fmt"
	"math"
)

/*
Go 语言函数

Go 语言函数定义格式如下：

func function_name( [parameter list] ) [return_types] {
	函数体
}

特点：
- 函数可以有多个返回值
- 可以使用可变参数

不支持：
- 不支持默认参数
- 方法不支持重载
*/

// 无参函数， 无返回值
func hello() {
	fmt.Println("Hello, World!")
}

// 带参函数， 无返回值
func sayHello(name string) {
	fmt.Println("Hello,", name, "!")
}

// 带参函数， 有返回值
// 如果有多个相同Type的参数或返回值，可以省略重复Type声明，比如(num1 int,num2 int) =>(num1,num2 int)
func max(num1, num2 int) int {
	var result int

	/* 函数返回两个数的最大值 */
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 多个返回值
func sum(num1 int, num2 float32) (int, float32) {
	// 返回平方数组
	return num1 * num1, num2 * num2
}

// 可变参数
func maxNumber(first int, others ...int) int {
	// 返回一组数值中的最大值
	max := first
	for _, value := range others {
		if value > max {
			max = value
		}
	}
	return max
}


// ----------> 函数参数
// 默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
// 值传递
func swap(x, y int) {
	// 交换值
	x, y = y, x
}

// 引用传递, *int 指针，指向整型
func swapRef(x *int, y *int) {
	// 交换值
	x, y = y, x
}

// ----------> 函数用法
// 函数作为值
/* 声明函数变量, 匿名函数 */
var getSquareRoot = func(x float64) float64 {
	return math.Sqrt(x)
}

// 闭包
// Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}
//带参数的闭包函数
func add(x1, x2 int) func()(int, int)  {
	i := 0
	return func() (int, int){
		i++
		return i, x1+x2
	}
}

// 多层次参数
func add2(x1, x2 int) func(x3 int, x4 int)(int, int, int)  {
	i := 0
	return func(x3 int, x4 int) (int, int, int){
		i++
		return i, x1+x2, x3+x4
	}
}


// 方法
// Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。

/* 定义结构体 */
type Circle struct {
	radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}





// 例子：杨辉三角
func showYangHuiTsriangle() {
	//行数
	const LINES int = 10

	nums := []int{}
	for i := 0; i < LINES; i++ {
		//补空白
		for j := 0; j < (LINES - i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (i + 1); j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			fmt.Print(value, " ")
		}
		fmt.Println("")
	}
}


func main() {
	// 无参
	hello()

	// 带参
	sayHello("Horace")

	// 带返回值
	fmt.Printf( "最大值是 : %d\n", max(100, 200))

	// 返回多个值
	fmt.Println(sum(2, 3))

	// 可变参数
	fmt.Println(maxNumber(1, 9, 6, 12, 34, 5))

	// 值传递
	a := 100
	b := 200
	fmt.Printf("值传递交换前 a, b 的值为 : %d, %d\n", a, b )
	swap(a, b)
	fmt.Printf("值传递交换后 a, b 的值为 : %d, %d\n", a, b )

	// 引用传递
	fmt.Printf("引用传递交换前 a, b 的值为 : %d, %d\n", a, b )
	// &a 指向 a 指针，a 变量的地址
	swapRef(&a, &b)
	fmt.Printf("引用传递交换后 a, b 的值为 : %d, %d\n", a, b )

	// 函数作为值
	fmt.Println(getSquareRoot(9))

	// 闭包
	fmt.Println("------------------闭包---------------------")
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	// 带参数的闭包函数调用
	fmt.Println("---------------------------------------")
	add_func := add(1,2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())

	fmt.Println("---------------------------------------")
	add_func2 := add2(1,2)
	fmt.Println(add_func2(1, 1))
	fmt.Println(add_func2(0, 0))
	fmt.Println(add_func2(2, 2))


	// 使用方法
	fmt.Println("-------------------方法--------------------")
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())


	// 杨辉三角
	fmt.Println("--------- 杨辉三角 ---------")
	showYangHuiTsriangle()
	fmt.Println("--------- 杨辉三角 ---------")

}