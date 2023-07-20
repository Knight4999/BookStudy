package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

// 含空指针的非空接口
func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) //启用输出收集
	}
	fmt.Println(buf == nil)
	f(buf)
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
