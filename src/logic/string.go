package main

import (
	"fmt"
	"strings"
	"bytes"
	"unsafe"
)

// 字符串

func main() {
	// ------> 字符串拼接
	// 1.直接+=进行拼接
	var str = "hello"
	str += ",world"
	fmt.Println(str)

	// 2.利用Sprintf
	str = "hello"
	str = fmt.Sprintf("%s%s", str, ",golang")
	fmt.Println(str)

	// 3.对于[]string利用strings.Join进行拼接
	str = "hello"
	str = strings.Join([]string{str, ",man"},"")
	fmt.Println(str)

	// 4.利用[]string和slice的append的性质进行拼接
	var strs []string
	strs = append(strs,"abc", "def")
	str = strings.Join(strs,"")
	fmt.Println(str)

	// 5.使用bytes.Buffer
	var buf bytes.Buffer
	buf.WriteString("12312")  // WriteString()参数一定是string，不能是byte
	buf.WriteString("werwer")
	str = buf.String()
	fmt.Println(str)

	// 优缺点：
	// 在已有字符串数组的场合，使用 strings.Join() 能有比较好的性能
	// 在一些性能要求较高的场合，尽量使用 buffer.WriteString() 以获得更好的性能
	// 较少字符串连接的场景下性能最好，而且代码更简短清晰，可读性更好
	// 如果需要拼接的不仅仅是字符串，还有数字之类的其他需求的话，可以考虑 fmt.Sprintf



	a := "hello"
	fmt.Println(unsafe.Sizeof(a))
	// 输出结果为：16
	// 字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
}

