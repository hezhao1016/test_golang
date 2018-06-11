package dm

func Max(num1,num2 int) int {
	var result int

	/* 函数返回两个数的最大值 */
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}