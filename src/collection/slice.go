package main

import "fmt"

// Go 语言切片是对数组的抽象。
// Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，
// Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

func main() {
	// 声明一个未指定大小的数组来定义切片
	var slice []int

	// 或使用make()函数来创建切片
	slice2 := make([]int,0, 100)

	// len() 方法获取长度, cap() 可以测量切片最长可以达到多少
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice2), cap(slice2), slice2)

	// slice 的零值是 `nil`。 一个 nil 的 slice 的长度和容量是 0。
	if slice == nil {
		fmt.Println("nil!")
	}


	// 直接初始化切片
	s := []int {1,2,3 }
	fmt.Println(s)

	// 定义数组
	arr := [5]string{"a", "b", "c", "d", "e"}

	// 通过 数组arr初始化切片s1
	s1 := arr[:]
	s2 := arr[1:3]
	// 通过切片s1初始化切片s3
	s3 := s1[1:4]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)


	fmt.Println("--------------------切片截取--------------------")
	test1()

	fmt.Println("--------------------添加元素--------------------")
	testAppend()

}

// 切片截取
func test1() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 省略下标代表从 0 开始 */
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 省略上标代表到 len(s) 结束 */
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}

// append() 和 copy() 函数
func testAppend() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}