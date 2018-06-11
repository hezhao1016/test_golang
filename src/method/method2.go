package main

import (
	"math"
	"fmt"
)

// Go 没有类。然而，仍然可以在结构体类型上定义方法。
// 方法接收者 出现在 func 关键字和方法名之间的参数中。

type Vertex struct {
	X, Y float64
}

// 方法可以与命名类型或命名类型的指针关联。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v, v.Abs())

	// 将会改变v的属性值
	v.Scale(5)
	fmt.Println(v, v.Abs())
}