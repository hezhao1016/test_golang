package main

import "fmt"

// 一个结构体（`struct`）就是一个字段的集合。

// 定义结构体
type Books struct {
	bookId int
	// 相同类型可在一行声明
	title, author, subject string
}

func main() {
	var book1 Books
	book1.title = "Go 语言"
	book1.author = "www.runoob.com"
	book1.subject = "Go 语言教程"
	book1.bookId = 6495407

	book2 := Books{6495700, "Python 教程", "www.runoob.com", "Python 语言教程"}

	book3 := Books{6495800, "Ruby 教程", "www.runoob.com", "Ruby 语言教程"}

	// 打印 Book1 信息
	printBook(book1)
	printBook(book2)
	printBook(book3)

}

func printBook(book Books) {
	fmt.Println("-----------------------------------")
	fmt.Printf("Book bookId : %d\n", book.bookId)
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
}