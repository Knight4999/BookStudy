// 使用文件操作相关知识，模拟实现linux平台cat命令的功能
package io

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// cat命令实现
func Cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') //注意
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func Du() {
	//解析命令行参数
	flag.Parse()
	if flag.NArg() == 0 {
		// 如果没有参数默认从标准输入读取内容
		Cat(bufio.NewReader(os.Stdin))
	}

	// 依次读取每个指定文件的内容并打印到终端
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		Cat(bufio.NewReader(f))
	}

}
