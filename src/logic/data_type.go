package main

import (
	"fmt"
	"strconv"
)

// 数据类型

/*
Go 语言按类别有以下几种数据类型：
布尔型
	bool， 值是 true 或 false，默认为 false
数字类型
	整形
		uint8	无符号 8 位整型 (0 到 255)
		uint16	无符号 16 位整型 (0 到 65535)
		uint32	无符号 32 位整型 (0 到 4294967295)
		uint64	无符号 64 位整型 (0 到 18446744073709551615)
		int8	有符号 8 位整型 (-128 到 127)
		int16	有符号 16 位整型 (-32768 到 32767)
		int32	有符号 32 位整型 (-2147483648 到 2147483647)
		int64	有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
	浮点数
		float32		IEEE-754 32位浮点型数
		float64		IEEE-754 64位浮点型数
		complex64	32 位实数和虚数
		complex128	64 位实数和虚数、
	其他数字类型
		byte	类似 uint8
		rune	类似 int32
		uint	32 或 64 位
		int		与 uint 一样大小
		uintptr	无符号整型，用于存放一个指针
字符串
	string
派生类型:
	包括：
	(a) 指针类型（Pointer）
	(b) 数组类型
	(c) 结构化类型(struct)
	(d) Channel 类型
	(e) 函数类型
	(f) 切片类型
	(g) 接口类型（interface）
	(h) Map 类型
*/

var isActive bool  // 全局变量声明
var enabled, disabled = true, false  // 忽略类型的声明

func main() {
	var available bool  // 一般声明
	valid := false      // 简短声明
	available = true    // 赋值操作

	fmt.Println(isActive)
	fmt.Println(enabled)
	fmt.Println(disabled)
	fmt.Println(available)
	fmt.Println(valid)

	// 数字
	var a = 2  // 根据值自行判定变量类型。
	var b = 1.5
	var c float32 = 1.5  // 指定变量类型为float

	fmt.Println("------------------------------------------")
	fmt.Print(a)  // 输出不换行
	fmt.Println()
	fmt.Println(b)
	fmt.Println(c)


	// 类型转换, 注意：int 不能直接赋值给float
	fmt.Println("------------------------------------------")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)

	// ------- string、int、int64互相转换 ------
	fmt.Println("------------------------------------------")
	str := "123"

	//string到int
	int,err:=strconv.Atoi(str)
	fmt.Println(int, err)

	//string到int64
	int64,err := strconv.ParseInt(str, 10, 64)
	fmt.Println(int64, err)

	//int到string
	str = strconv.Itoa(int)
	fmt.Println(str, err)

	//int64到string
	str = strconv.FormatInt(int64,10)
	fmt.Println(str, err)

	//string到float32
	str = "123.456"
	float32,err := strconv.ParseFloat(str,32)
	fmt.Println(float32, err)

	//string到float64
	float64,err := strconv.ParseFloat(str,64)
	fmt.Println(float64, err)

	//float到string
	str = strconv.FormatFloat(float32, 'E', -1, 32)
	fmt.Println(str, err)

	str = strconv.FormatFloat(float64, 'E', -1, 64)
	fmt.Println(str, err)
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)


}