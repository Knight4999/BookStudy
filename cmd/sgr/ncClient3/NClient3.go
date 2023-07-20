package main

// netcat
import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:23450") //Dial 拨号的意思
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	/* go mustCopy(os.Stdout, conn) 旧写法
	mustCopy(conn, os.Stdin) */
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} //指示主 goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done //等待后台goroutine完成
}

// 将从流中读取的数据，输出到控制台
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
