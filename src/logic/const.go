package main

import (
	"fmt"
	"unsafe"
)

// 常量

// 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
const CNAME string = "我是全局常量，变量名一般大写"
const SUCCESS_CODE = "200"

// 多常量声明：
const c_name1, c_number2 = "a", 1

//常量还可以用作枚举
const (
	A = "abc"
	B = len(A)
	C = unsafe.Sizeof(A)
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"  // 多重赋值

	fmt.Println(CNAME)
	println(a, b, c)
	println()
	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	fmt.Println()

	fmt.Println(A, B, C)


	//iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	//在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
	//iota 可以被用作枚举值：
	const (
		d = iota   //0
		e          //1
		f          //2
		g = "ha"   //独立值，iota += 1
		h          //"ha"   iota += 1
		i = 100    //iota +=1
		j          //100  iota +=1
		k = iota   //7,恢复计数
		l          //8
	)
	fmt.Println(d,e,f,g,h,i,j,k,l)
}