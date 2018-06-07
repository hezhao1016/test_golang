package main

import "fmt"

//运算符

/*
Go 语言内置的运算符有：
- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
- 赋值运算符
- 其他运算符

算术运算符
下表列出了所有Go语言的算术运算符。假定 A 值为 10，B 值为 20。
运算符	描述	实例
+	相加	A + B 输出结果 	30
-	相减	A - B 输出结果 	-10
*	相乘	A * B 输出结果 	200
/	相除	B / A 输出结果 	2
%	求余	B % A 输出结果	0
++	自增	A++ 输出结果	 	11
--	自减	A-- 输出结果 	9

关系运算符
下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20。
运算符	描述	实例
==	检查两个值是否相等，如果相等返回 True 否则返回 False。			(A == B) 为 False
!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。		(A != B) 为 True
>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。		(A > B) 为 False
<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。		(A < B) 为 True
>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。	(A >= B) 为 False
<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。	(A <= B) 为 True

逻辑运算符
下表列出了所有Go语言的逻辑运算符。假定 A 值为 True，B 值为 False。
运算符	描述	实例
&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。	(A && B) 为 False
||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。	(A || B) 为 True
!	逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。		!(A && B) 为 True

赋值运算符
下表列出了所有Go语言的赋值运算符。
运算符	描述	实例
=	简单的赋值运算符，将一个表达式的值赋给一个左值		C = A + B 将 A + B 表达式结果赋值给 C
+=	相加后再赋值										C += A 等于 C = C + A
-=	相减后再赋值										C -= A 等于 C = C - A
*=	相乘后再赋值										C *= A 等于 C = C * A
/=	相除后再赋值										C /= A 等于 C = C / A
%=	求余后再赋值										C %= A 等于 C = C % A
<<=	左移后赋值										C <<= 2 等于 C = C << 2
>>=	右移后赋值										C >>= 2 等于 C = C >> 2
&=	按位与后赋值										C &= 2 等于 C = C & 2
^=	按位异或后赋值									C ^= 2 等于 C = C ^ 2
|=	按位或后赋值										C |= 2 等于 C = C | 2

其他运算符
下表列出了Go语言的其他运算符。
运算符	描述	实例
&	返回变量存储地址	&a 	将给出变量的实际地址。
*	指针变量。	*a			是一个指针变量

运算符优先级
有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：
优先级	运算符
7	^ !
6	* / % << >> & &^
5	+ - | ^
4	== != < <= >= >
3	<-
2	&&
1	||
*/

func main() {
	fmt.Println("---------------------------------------------")
	mathOperator()

	fmt.Println("---------------------------------------------")
	gxOperator()

	fmt.Println("---------------------------------------------")
	logicOperator()

	fmt.Println("---------------------------------------------")
	logicOperator()

	fmt.Println("---------------------------------------------")
	fzOperator()

	fmt.Println("---------------------------------------------")
	highOperator()
}

//算术运算符
func mathOperator() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c )
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c )
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c )
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c )
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c )
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a )
	a=21   // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a )
}

//关系运算符
func gxOperator() {
	var a int = 21
	var b int = 10

	if a == b  {
		fmt.Printf("第一行 - a 等于 b\n" )
	} else {
		fmt.Printf("第一行 - a 不等于 b\n" )
	}
	if a < b {
		fmt.Printf("第二行 - a 小于 b\n" )
	} else {
		fmt.Printf("第二行 - a 不小于 b\n" )
	}

	if a > b {
		fmt.Printf("第三行 - a 大于 b\n" )
	} else {
		fmt.Printf("第三行 - a 不大于 b\n" )
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第四行 - a 小于等于 b\n" )
	}
	if b >= a {
		fmt.Printf("第五行 - b 大于等于 a\n" )
	}
}

//逻辑运算符
func logicOperator() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("第一行 - 条件为 true\n" )
	}
	if a || b {
		fmt.Printf("第二行 - 条件为 true\n" )
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n" )
	} else {
		fmt.Printf("第三行 - 条件为 false\n" )
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n" )
	}
}

//赋值运算符
func fzOperator() {
	var a int = 21
	var c int

	c =  a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )

	c +=  a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )

	c -=  a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )

	c *=  a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )

	c /=  a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )

	c  = 200

	c <<=  2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c )

	c >>=  2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )

	c &=  2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )

	c ^=  2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )

	c |=  2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )

	// 交换两个参数的值
	x := 100
	y := 200
	fmt.Printf("交换前：x=%d, y=%d\n", x, y)
	x, y = y, x
	fmt.Printf("交换后：x=%d, y=%d\n", x, y)
}

//其他运算符
func otherOperator()  {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a )
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b )
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c )

	/*  & 和 * 运算符实例 */
	ptr = &a    /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}

//运算符优先级
func highOperator() {
	var a int = 20
	var b int = 10
	var c int = 15
	var d int = 5
	var e int

	e = (a + b) * c / d      // ( 30 * 15 ) / 5
	fmt.Printf("(a + b) * c / d 的值为 : %d\n",  e )

	e = ((a + b) * c) / d    // (30 * 15 ) / 5
	fmt.Printf("((a + b) * c) / d 的值为  : %d\n" ,  e )

	e = (a + b) * (c / d)   // (30) * (15/5)
	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n",  e )

	e = a + (b * c) / d     //  20 + (150/5)
	fmt.Printf("a + (b * c) / d 的值为  : %d\n" ,  e )
}