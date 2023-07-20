package main

import "fmt"

func main() {
	//创建一个管道
	ch := make(chan int)
	fmt.Printf("ch: %T\n", ch)

	//chan有发送（send）和接受（receive）两个主要操作，都是通信行为
	x := 10
	ch <- x  //发送语句,将值x发送到管道ch中
	x = <-ch //赋值语句中的接受表达式，从管道ch中接受值，并赋值给Y
	<-ch     //接受语句，舍弃结果

	close(ch) //关闭通道

	//无缓冲通道
	ch = make(chan int)    //无缓冲通道
	ch = make(chan int, 0) //无缓冲通道

	ch2 := make(chan string, 3) //容量为3的缓冲通道
	ch2 <- "A"
	ch2 <- "B"
	ch2 <- "C"
	/* ch2 <- "D" */
	fmt.Println(<-ch2)

}

func mirrorQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("asia.gopl.io")
	}()
	go func() {
		responses <- request("europe.gopl.io")
	}()
	go func() {
		responses <- request("americas.gopl.io")
	}()
	return <-responses
}

func request(hostname string) (responses string) { return responses }
