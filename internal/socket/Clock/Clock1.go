package clock

//顺序时钟服务器
import (
	"io"
	"net"
	"time"
)

// 处理连接函数
func HandleConn(c net.Conn) {
	defer c.Close() //关闭连接流
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\t"))
		if err != nil {
			return //连接断开
		}
		time.Sleep(1 * time.Second)
	}
}
