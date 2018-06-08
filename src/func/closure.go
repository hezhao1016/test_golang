package _func

import "fmt"

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

// 例如，函数 adder 返回一个闭包。每个闭包都被绑定到其各自的 sum 变量上。
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 练习：斐波纳契闭包
func fibonacci() func() int {
	var a, b, n = 0, 1, 0
	return func() int {
		n++
		if n > 1 {
			a, b = b, a+b
		}
		return a
	}
}


func main() {
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

	// 多层次参数
	fmt.Println("---------------------------------------")
	add_func2 := add2(1,2)
	fmt.Println(add_func2(1, 1))
	fmt.Println(add_func2(0, 0))
	fmt.Println(add_func2(2, 2))


	// 例子
	fmt.Println("---------------------------------------")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// 斐波纳契函数
	fmt.Println("---------------------------------------")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}