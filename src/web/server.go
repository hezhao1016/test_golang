package main

import (
	"net/http"
	"fmt"
	"strings"
	"reflect"
)

/*
Web 服务器
包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求：

package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

访问 http://localhost:8001/ 会看到来自程序的问候。
*/

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

type Handlers struct {
}

func (h *Handlers) ResAction(w http.ResponseWriter, req *http.Request) {
	fmt.Println("res")
	w.Write([]byte("res"))
}

func say(w http.ResponseWriter, req *http.Request) {
	pathInfo := strings.Trim(req.URL.Path, "/")
	fmt.Println("pathInfo:", pathInfo)

	parts := strings.Split(pathInfo, "/")
	fmt.Println("parts:", parts)

	var action = "ResAction"
	fmt.Println(strings.Join(parts, "|"))
	if len(parts) > 1 {
		fmt.Println("22222222")
		action = strings.Title(parts[1]) + "Action"
	}
	fmt.Println("action:", action)
	handle := &Handlers{}
	controller := reflect.ValueOf(handle)
	method := controller.MethodByName(action)
	r := reflect.ValueOf(req)
	wr := reflect.ValueOf(w)
	method.Call([]reflect.Value{wr, r})
}

func main() {
	fmt.Println("server started on localhost:8001 ...")

	http.HandleFunc("/hello", hello)
	http.Handle("/handle/", http.HandlerFunc(say))
	http.ListenAndServe("localhost:8001", nil)
	//select {} //阻塞进程
}