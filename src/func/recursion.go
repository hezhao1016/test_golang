package main

import "fmt"

// 递归函数


// 递归函数实现阶乘
func Factorial(n uint64)(result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// 递归函数实现斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}


func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))


	var j int
	fmt.Println("0到10的斐波那契数列：")
	for j = 0; j < 10; j++ {
		fmt.Printf("%d\t", fibonacci(j))
	}
}