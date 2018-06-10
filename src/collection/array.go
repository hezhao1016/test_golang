package main

import "fmt"

// 数组

/*
数组元素类型要一致
长度固定
*/

func main() {
	// 声明数组
	var balance [10]float32
	
	// 赋值
	// 通过大括号来初始化
	balance = [10]float32{1000.0}
	fmt.Println(balance)

	// 初始化数组
	// 初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
	var nums1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var nums2 = [5]float32{1000.0, 2.0, 3.4, 7.0}
	fmt.Println(nums1)
	fmt.Println(nums2)

	// 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
	var nums3 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(nums3)

	// 根据下标赋值
	nums3[1] = 999

	// 访问数组元素
	var salary = nums3[1]
	fmt.Println(salary)


	var n [10]int
	fmt.Println("---------------遍历数组------------------")

	/* 为数组 n 初始化元素 */
	for i := 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j := 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}

	fmt.Println("---------------数组长度------------------")
	// 数组长度
	length := len(n)
	fmt.Println(length)


	// 多维数组
	fmt.Println("---------------多维数组------------------")

	var twoArr = [5][2]int{
		{0,0},
		{1,2},
		{2,4},
		{3,6},
		{4,8}}

	/* 输出数组元素 */
	for  i := 0; i < len(twoArr); i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i,j, twoArr[i][j] )
		}
	}

	fmt.Println("---------------函数参数------------------")
	/* 数组长度为 5 */
	var balance2 = []int {1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage( balance2)

	/* 输出返回的平均值 */
	fmt.Printf( "平均值为: %f ", avg )
}

// 向函数传递数组参数, 可规定数组大小, arr [5]int
func getAverage(arr []int) float32 {
	var i,sum int
	var avg float32
	var size = len(arr)

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}