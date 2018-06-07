package main

import "fmt"

// 循环

func main() {
	// for循环是Go语言唯一的循环结构, 有4种形式：
	// for condition { }
	// for init; condition; increment { }
	// for { }

	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
	// for key, value := range oldMap {
	// 	newMap[key] = value
	// }

	numbers := [6]int{1, 2, 3, 5}

	/* 1.类似于Java for 循环 */
	fmt.Println("-----------------------类似于Java for 循环------------------------")
	for i := 0; i < 10; i++ {
		fmt.Printf("i 的值为: %d\n", i)
	}

	/* 2.类似于Java while 循环 */
	fmt.Println("-----------------------类似于Java while 循环------------------------")
	n := 9
	for n < 15 {
		n++
		fmt.Printf("n 的值为: %d\n", n)
	}

	/* 3.类似于Python for 循环, 用于遍历集合 */
	fmt.Println("-----------------------for range------------------------")
	for i, x:= range numbers {
		fmt.Printf("a - 第 %d 位 x 的值 = %d\n", i, x)
	}

	/* 这样也能遍历集合 */
	fmt.Println("-----------------------遍历集合------------------------")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("b - 第 %d 位 x 的值 = %d\n", i, numbers[i])
	}

	// 无条件的for循环
	fmt.Println("-----------------------无条件的for循环------------------------")
	j := 0
	for true {
		j++

		if j % 2 == 1 {
			/* 跳过此次循环 */
			continue
		}

		fmt.Printf("无条件循环[%d]...\n", j)

		if j >= 10 {
			/* 使用 break 语句跳出循环 */
			break
		}
	}


	// 循环嵌套, 输出 2 到 100 间的素数
	fmt.Println("-----------------------循环嵌套------------------------")
	var i, k int

	for i=2; i < 100; i++ {
		for k=2; k <= (i/k); k++ {
			if i % k == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if k > i/k {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}
