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



	// 字符串去除空格和换行符
	str = "这里是 www\n.runoob\n.com"
	fmt.Println("-------- 原字符串 ----------")
	fmt.Println(str)
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println("-------- 去除空格与换行后 ----------")
	fmt.Println(str)



	fmt.Println("---------------常用方法------------------")

	// 遍历字符串
	str = "好好学习天天向上"
	for _, v:= range str {  // 用_声明的变量可以不使用
		fmt.Printf("%c", v)
	}
	fmt.Printf("\n")


	// 判断字符串是否包含某个字符
	flag := strings.Contains(str,"好")
	fmt.Println(flag)

	// 判断字符串是否相等
	flag = str == "好好学习天天向上"
	fmt.Println(flag)
	flag = strings.EqualFold(str, "好好学习天天向上")
	fmt.Println(flag)

	// 字符串拆分成数组
	array := strings.Split("a,b,c", ",")
	fmt.Println(array)

	// 数组转字符串
	fmt.Println(strings.Join(array,","))

	// 查找下标, 一个汉字占3个长度
	index := strings.Index(str, "天")
	fmt.Println(index)
	lastIndex := strings.LastIndex(str, "天")
	fmt.Println(lastIndex)

	// 截取字符串， string()类似于toString()
	fmt.Println(string(str[0:3]))
	fmt.Println(string(str[:12]))
	fmt.Println(string(str[12:]))
	fmt.Println(string(str[:]))

	// Substring
	fmt.Println(Substring("好好学习天天向上", 0,3))
	fmt.Println(Substring("abcdefgh", 0,3))

}

// 字符串截取函数，一个汉字算一个长度
func Substring(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}