package strconv

import (
	"fmt"

	"strconv"
)

func StoI() {

	//Atoi()函数用于将字符串类型的整数转换为int类型,func Atoi(s string) (i int, err error)
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1)
	}
}

func ItoS() {
	//Itoa()函数用于将int类型数据转换为对应的字符串表示，func Itoa(i int) string
	i1 := 10086
	fmt.Println(strconv.Itoa(i1))

}

// Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
func ParseTest() {
	// ParseBool(),返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	str := []string{"1", "0", "t", "f", "T", "F", "False", "True", "FALSE", "TRUE", "nihao"}
	for _, s := range str {
		fmt.Println(strconv.ParseBool(s))
	}

	//PaserInt(), func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	sint := "-1345123"
	i1, err := strconv.ParseInt(sint, 16, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i1)

	//PaserUnit(),func ParseUint(s string, base int, bitSize int) (n uint64, err error),与PasInt类似，不支持正负号，用于无符号整形
	sint2 := "1323"
	i2, err := strconv.ParseUint(sint2, 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i2)

	//PaserFloat(),解析一个表示浮点数的字符串并返回其值。func ParseFloat(s string, bitSize int) (f float64, err error)
	//如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。
	//bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；
	//返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。
	sint3 := "13235"
	i3, err := strconv.ParseFloat(sint3, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i3)

}

// Format系列函数实现了将给定类型数据格式化为string类型数据的功能。
func FormatTest() {
	//FormatBool()
	s := strconv.FormatBool(true)
	fmt.Println(s)

	//FormatInt(),返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。s
	v := int64(-42)
	s1 := strconv.FormatInt(v, 10)
	fmt.Println(s1)
	s16 := strconv.FormatInt(v, 16)
	fmt.Println(s16)

	//FormatUint(),是FormatInt的无符号版本
	t := uint64(42)
	t1 := strconv.FormatUint(t, 10)
	fmt.Println(t1)
	t16 := strconv.FormatUint(t, 16)
	fmt.Println(t16)

	//FormatFloat(),函数将浮点数表示为字符串并返回。func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	//bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	//fmt表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。
	//prec控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	f := 3.1415926535
	f1 := strconv.FormatFloat(f, 'E', -1, 64)
	fmt.Println(f1)

}
func IsPrint() {
	//func IsPrint(r rune) bool 判断某个字符是否可以打印
	s := 'a'
	b := strconv.IsPrint(s)
	fmt.Println(b)

	//func CanBackquote(s string) bool ,返回字符串s是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串。
	str := "asyve;asbvuqw"
	b2 := strconv.CanBackquote(str)
	fmt.Println(b2)
}

func AppendTest() {
	buf := []byte("bool:")
	buf = strconv.AppendBool(buf, true)
	fmt.Println(string(buf))
}
