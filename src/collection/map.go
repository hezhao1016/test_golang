package main

import (
	"fmt"
)

// Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
// Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

// 定义 Map
// 可以使用内建函数 make 也可以使用 map 关键字来定义 Map
// 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

func main() {
	// 创建集合
	var countryCapitalMap map[string]string
	// 使用 make 函数
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap [ "France" ] = "Paris"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	// 使用键输出Map值
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	// 查看元素在集合中是否存在
	captial, ok := countryCapitalMap [ "Japan" ]  // 如果确定是真实的,则存在,否则不存在
	if ok {
		fmt.Println("日本的首都是", captial)
	} else {
		fmt.Println("日本的首都不存在")
	}

	fmt.Println("------------------------testUpdate------------------------------")
	testUpdate()

	fmt.Println("------------------------testDelete------------------------------")
	testDelete()

	fmt.Println("------------------------testWithStruct------------------------------")
	testWithStruct()

	fmt.Println("------------------------WordCount------------------------------")
	mwords := WordCount("hello horace, haha!")
	fmt.Println(mwords)
}


// 修改 map
func testUpdate()  {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 删除元素
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 通过双赋值检测某个键存在
	// 如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key
func testDelete()  {
	// 创建map
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	// 打印地图
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}

	//删除元素
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	// 打印地图
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}
}


// --------------------------------------------------------------------------------

type Vertex struct {
	Lat, Long float64
}

var m1 map[string]Vertex

// map 的文法跟结构体文法相似，不过必须有键名。
var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// 如果顶级的类型只有类型名的话，可以在文法的元素中省略键名。
var m3 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func testWithStruct() {
	// map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。
	m1 = make(map[string]Vertex)

	m2["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m1["Bell Labs"])
	fmt.Println(m2)
	fmt.Println(m3)
}


// 练习：map
// 实现 `WordCount`。它应当返回一个含有 s 中每个 “词” 个数的 map。函数 wc.Test 针对这个函数执行一个测试用例，并输出成功还是失败。
func WordCount(s string) map[string]int {
	mwords := make(map[string]int)
	var word string

	for _, c := range s{
		word = string(c)
		count, exists := mwords[word]
		if exists {
			mwords[word] = count + 1
		} else {
			mwords[word] = 1
		}
	}

	return mwords
}