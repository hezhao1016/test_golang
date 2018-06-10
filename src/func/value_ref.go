package main

import "fmt"

// 典型的引用传递类型：slice，map和channel
// 典型的值传递类型：数组、struct

type Test struct {
	A int64
	B string
}

func main(){
	// 引用类型
	fmt.Println("--------------map-------------------")
	map_param := make(map[string]string,0)
	fmt.Println(map_param)
	mapfunc(map_param)
	fmt.Println(map_param)

	fmt.Println("---------------slice------------------")
	slice_param := make([]int64,0)
	slice_param = append(slice_param,999)
	fmt.Println(slice_param)
	slicefunc(slice_param)
	fmt.Println(slice_param)

	// 值类型
	fmt.Println("--------------array-------------------")
	arr_param := [4]int64{1,2,3,4}
	fmt.Println(arr_param) // 1,2,3,4
	arrfunc(arr_param)
	fmt.Println(arr_param) // 1,2,3,4

	fmt.Println("----------------struct-----------------")
	test := Test{A:1,B:"ooo",}
	fmt.Println(test) // 1,ooo
	testfunc(test)
	fmt.Println(test) // 1,ooo
}

func mapfunc(param map[string]string){
	param["Name"] = "Jack"
}

func slicefunc(param []int64){
	param[0] = 10000
}

func arrfunc(param [4]int64){
	param[0] = 10000
}

func testfunc(a Test){
	a.A = 9999
	a.B = "hah"
}