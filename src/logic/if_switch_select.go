package main

import "fmt"

// 选择结构

func main() {

	// ------> if-else
	fmt.Println("-----------------------if-else------------------------")

	var a int = 100
	var b int = 200

	/* 判断布尔表达式 */
	if a < 20 {
		fmt.Printf("a 小于 20\n" )

		// 注意 else、else if语句必须与右花括号 } 处在一行
	} else if a < 50 {
		fmt.Printf("a 小于 50\n" )
	} else {
		fmt.Printf("a 不小于 20\n" )

		if b == 200 {
			fmt.Printf("b 的值为 : %d\n", b)
		}
	}
	fmt.Printf("a 的值为 : %d\n", a)



	// ------> switch
	// 类型不被局限于常量或整数，但必须是相同的类型
	// 单个case中，可以出现多个结果选项
	// 与Java的规则相反，Go语言不需要用break来明确退出一个case
	// 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case (就和Java的case后面没加break的用法是一样的)
	// 可以不设定switch之后的条件表达式，在此种情况下，整个switch结构与多个if...else...的逻辑作用等同。
	fmt.Println("-----------------------switch------------------------")

	var grade string = "B"
	var marks int = 90

	// case 语句可以在一行代码写完
	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70 : grade = "C"
		default: grade = "D"
	}

	switch {
		case grade == "A" :
			fmt.Printf("优秀!\n" )
		case grade == "B", grade == "C" :
			fmt.Printf("良好\n" )
		case grade == "D" :
			fmt.Printf("及格\n" )
		case grade == "F":
			fmt.Printf("不及格\n" )
		default:
			fmt.Printf("差\n" );
	}
	fmt.Printf("你的等级是 %s\n", grade );


	// 不设定switch之后的条件表达式
	var age = 10
	switch {
		case 0 <= age && age <= 10:
			fmt.Printf("小朋友\n")

			// 加了fallthrough后，会直接运行【紧跟的后一个】case或default语句，不论条件是否满足都会执行
			fallthrough
		case 11 <= age && age <= 20:
			fmt.Printf("大朋友\n")
		default:
			fmt.Printf("%d岁的朋友\n", age)
	}


	// Type Switch
	// switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	var x interface{}

	switch i := x.(type) {
		case nil:
			fmt.Printf("x 的类型 :%T",i)
		case int:
			fmt.Printf("x 是 int 型")
		case float64:
			fmt.Printf("x 是 float64 型")
		case func(int) float64:
			fmt.Printf("x 是 func(int) 型")
		case bool, string:
			fmt.Printf("x 是 bool 或 string 型" )
		default:
			fmt.Printf("未知类型")
	}
	fmt.Println()



	// ------> select
	// select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。
	// select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。

	//以下描述了 select 语句的语法：
	//	每个case都必须是一个通信
	//	所有channel表达式都会被求值
	//	所有被发送的表达式都会被求值
	//	如果任意某个通信可以进行，它就执行；其他被忽略。
	//	如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
	//	否则：
	//	如果有default子句，则执行该语句。
	//	如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

	fmt.Println("-----------------------select------------------------")

	var c1, c2, c3 chan int
	var i1, i2 int
	select {
		case i1 = <-c1:
			fmt.Printf("received ", i1, " from c1\n")
		case c2 <- i2:
			fmt.Printf("sent ", i2, " to c2\n")
		case i3, ok := (<-c3):  // same as: i3, ok := <-c3
			if ok {
				fmt.Printf("received ", i3, " from c3\n")
			} else {
				fmt.Printf("c3 is closed\n")
			}
		default:
			fmt.Printf("no communication\n")
	}

}