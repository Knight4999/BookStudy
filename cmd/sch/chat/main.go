package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:23450")
	if err != nil {
		log.Fatal(err)
	}
	go brodcaster()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

type clients chan<- string //对外发送消息通道

var (
	entering = make(chan clients)
	leaving  = make(chan clients)
	messages = make(chan string) // 所有接受的客户端消息
)

func brodcaster() {
	clients := make(map[clients]bool)
	for {
		select {
		case msg := <-messages:
			//把所有接收的消息广播给所有的客户
			//发送消息通道
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(c net.Conn) {
	ch := make(chan string) //对外发送客户消息通道
	go clientWriter(c, ch)
	who := c.RemoteAddr().String()
	ch <- "You are" + who
	messages <- who + "has arrived"
	entering <- ch

	input := bufio.NewScanner(c)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + "has left"
	c.Close()
}

func clientWriter(c net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(c, msg)
	}
}
