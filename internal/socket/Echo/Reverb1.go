package echo

//回声服务器
import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func Echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func HandleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go Echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
