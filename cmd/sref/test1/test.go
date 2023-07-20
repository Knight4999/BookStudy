package test1

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func Stest() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	fmt.Printf("%T\n", w)

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	//调用value的type
	t1 := v.Type()
	fmt.Println(t1.String())

	//reflect.Value.Interface 方法. 它返回一个 interface{} 类型表示 reflect.Value 对应类型的具体值
	v1 := reflect.ValueOf(3)
	x := v1.Interface() //返回v1对应的接口类型具体值
	i := x.(int)
	fmt.Printf("%d\n", i)

	//使用reflect.Value 设置值
	j := 2
	a := reflect.ValueOf(2)  //副本值，无地址值
	b := reflect.ValueOf(j)  //函数副本，无地址值
	c := reflect.ValueOf(&j) //指针副本，无地址值
	//通过Elem函数来获取任意变量c可寻址的value值。reflect.ValueOf(&x).Elem()
	d := c.Elem()
	//调用CanAddr函数判断reflect.Value 是否可以取地址值
	fmt.Println(a.CanAddr())
	fmt.Println(b.CanAddr())
	fmt.Println(c.CanAddr())
	fmt.Println(d.CanAddr())

	//从一个可以寻址的reflect.Value()获取变量的方式，并修改变量值
	x1 := 2
	y1 := reflect.ValueOf(&x1).Elem()
	//1.addr()函数，返回一个Value,里面保存了指向变量的指针
	//2.在这个Value上调用Interface(),返回一个包含改指针的interface{}值
	//3.最后，如果我们知道变量的类型，我们可以使用类型的断言机制将得到的interface{}类型的接口强制环为普通的类型指针
	px := y1.Addr().Interface().(*int)
	*px = 10
	fmt.Println(x1)

	//通过可寻址的reflect.Value修改变量值方式，调用Set()函数
	y1.Set(reflect.ValueOf(1998)) //此处赋值的类型必须与元数据类型一致。
	fmt.Println(y1.CanSet())      //canSet()函数，用于判断可寻址变量是否可更改
	fmt.Println(x1)
}
