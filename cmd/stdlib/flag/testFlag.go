// flag 包用于定义命令行参数
package main

import (
	"flag"
	"fmt"
	"time"
)

// os.Args demo
func main() {
	//os.Args是一个[]string,如果你只是简单的想要获取命令行参数，可以使用os.Agrs获取命令行参数
	//os.Args是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称。
	/* if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	} */
	flagDemo()
}

func flagDemo() {
	//定义命令行参数
	//flag.Type
	/* name := flag.String("name", "娜迦海妖", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("结婚", false, "是否结婚")
	delay := flag.Duration("d", 0, "时间间隔") */

	//flag.TypeVar,flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
	var name string
	var age int
	var married bool
	var delay time.Duration

	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
