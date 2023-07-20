package io

import (
	"fmt"
	"log"
	"os"
)

func A() {
	/* var buf [16]byte

	os.Stdin.Read(buf[:])
	os.Stdin.WriteString(string(buf[:])) */

	//Create()  根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666
	file, err := os.Create("E:/my.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Write()
	if _, err := file.Write([]byte("Michael Jackson")); err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//NewFile() 并不会真实创建函数
	f := os.NewFile(0, "E:/newfile.ini")
	defer f.Close()
	/* //Open() 打开一个只读文件。
	_, err = os.Open("e:/newfile.ini")
	if err != nil {
		fmt.Println(err)
	} */

	//OpenFile() 打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限
	fs, err := os.OpenFile("e:/rockstart.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("请输入写到文件中的数据：")
	var text [100]byte
	os.Stdin.Read(text[:])
	if _, err = fs.WriteString(string(text[:])); err != nil {
		fmt.Println(err)
		fs.Close()
	}
}
