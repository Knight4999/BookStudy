package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //读取单个字节
		abort <- struct{}{}
	}()
	fmt.Println("火箭发射倒计时，中断请按空格键！")
	select {
	case <-time.After(10 * time.Second): //返回一个channal，10秒后会向通道中发送一个值
	case <-abort:
		fmt.Println("火箭发射中止！")
		return
	}
	launch()
}

func launch() {
	fmt.Println("发送成功!")
}
