package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	logfile io.Writer
	err     error
)

func init() {

	logfile, err = os.OpenFile("E:/rz.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed,err:", err)
		return
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
func main() {
	/* log.Println("这是一条普通的日志")
	v := "很普通的"
	log.Printf("这是一条%s日志", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。") */

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate) // 设置日志输出格式SetFlags()
	log.Println("这是一条普通的日志")
	//获取日志前缀
	prx := log.Prefix()
	fmt.Println(prx)
	//配置日志前缀
	log.SetPrefix("[wzh]")
	log.Println("莫名其妙的帅气")
	//配置日志输出位置     func SetOutput(w io.Writer)
	/* logfile, err := os.OpenFile("E:/rz.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed,err:", err)
		return
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate) */
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")

	own := log.New(logfile, "[ONE]", log.LstdFlags|log.Llongfile)
	own.Println("这是自定义的logger记录的日志。")
}
