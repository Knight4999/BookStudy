package main

//回声服务器主程序
import (
	. "BookStudy/internal/socket/Echo"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:23500") //监听端口
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept() //阻塞，等待连接
		if err != nil {
			log.Println(err)
			continue //如果失败，跳过本次循环，等待下次连接
		}
		go HandleConn(conn)
	}
}
