package main

import "fmt"

// 结构体指针
// 结构体字段可以通过结构体指针来访问。
// 通过指针间接的访问是透明的。

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// 创建指针
	p := &v
	// 通过指针修改
	p.X = 1e9
	fmt.Println(v)

	// 重新指向一个对象
	*p = Vertex{10, 20}
	fmt.Println(v)


	fmt.Println("-------------------------------")
	// 调用值传递方法，不生效
	calc(v)
	fmt.Println(v)

	// 调用引用传递方法
	calcRef(&v)
	fmt.Println(v)
}

func calc(vertex Vertex){
	// 注意，与Java不同，直接修改并不会改变vertex的值
	vertex.X = 11
}

func calcRef(vertex *Vertex){
	// 给字段重新赋值
	//vertex.X = 11

	// 或者，重新指向一个对象， 这种情况要使用 *
	*vertex = Vertex{11, 22}
}
