package main

import "fmt"

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
// 值传递

// 引用传递

// ----------> 函数用法

// 函数作为值

// 闭包

// 方法




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

	// 杨辉三角
	fmt.Println("--------- 杨辉三角 ---------")
	showYangHuiTsriangle()
	fmt.Println("--------- 杨辉三角 ---------")

}