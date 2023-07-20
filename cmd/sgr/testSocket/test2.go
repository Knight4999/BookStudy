package main

import (
	. "BookStudy/internal/socket/Clock"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:23500")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConn(conn) //并发处理连接
	}
}
