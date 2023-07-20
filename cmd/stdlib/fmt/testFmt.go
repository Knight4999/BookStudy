package test6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FmtDemo() {
	/* fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")
	fmt.Println("---------------------------------")
	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)

	s := "枯藤"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s) */

	/* var (
		name    string
		age     int
		married bool
	) */
	/* fmt.Scan(&name, &age, &married)//
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married) */
	/* fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married) */
	/* fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married) */

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text) //去除收尾的空格
	fmt.Println(text)

	fmt.Sscan("无可奈何花落去")
}
