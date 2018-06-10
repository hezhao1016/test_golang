package main

import "fmt"

// 指针
/*
Go 具有指针。 指针保存了变量的内存地址。

类型 *T 是指向类型 T 的值的指针。其零值是 `nil`。
var p *int

& 符号会生成一个指向其作用对象的指针。
i := 42
p = &i

* 符号表示指针指向的底层的值。
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i

这也就是通常所说的“间接引用”或“非直接引用”。

与 C 不同，Go 没有指针运算。
*/

func main() {
	i, j := 42, 2701

	// *符号 指定变量是作为一个指针
	var p *int

	// &符号 返回相应变量的内存地址
	p = &i

	fmt.Printf("a 变量的地址是: %x\n", &i  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", p )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *p )

	// 修改
	*p = 21
	fmt.Printf("a 变量修改后的值是: %d\n", i)


	// 现在指针指向别的值
	p = &j
	*p /= 37
	fmt.Println(j)


	// 空指针 nil
	var ptr *int
	fmt.Println(ptr)

	// 空指针判断
	if ptr != nil {
		fmt.Println("ptr 不是空指针")
	} else {
		fmt.Println("ptr 是空指针")
	}


	// 指针数组
	fmt.Println("--------------------指针数组--------------------")
	a := []int{10,100,200}
	var index int
	const MAX int = 3  // 注意：数组初始化时，长度必须是常量
	var ptrArray [MAX]*int

	for  index = 0; index < MAX; index++ {
		ptrArray[index] = &a[index] /* 整数地址赋值给指针数组 */
	}

	for  i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i,*ptrArray[i] )
	}


	fmt.Println("--------------------指向指针的指针--------------------")
	test1()
}

// 指向指针的指针
func test1() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a )
	fmt.Printf("指针变量 *ptr = %d\n", *ptr )
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}