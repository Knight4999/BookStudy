// template包的作用，主要用于模板渲染，即将后端数据传输至前端网页对应的位置上
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP SERVER failed,err:", err)
		return
	}
}

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("E:/GoProject/BookStudy/cmd/stdlib/template/hello2.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	/* // 利用给定数据渲染模板, 并将结果写入w
	tmpl.Execute(w, "小明") */
	user := UserInfo{
		"李白",
		"唐朝",
		56,
	}
	tmpl.Execute(w, user)
}
