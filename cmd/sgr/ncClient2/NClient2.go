package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:23500") //Dial 拨号的意思
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() //关闭连接流
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)

}

// 将从流中读取的数据，输出到控制台
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
