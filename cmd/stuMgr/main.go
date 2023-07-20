package main

import (
	"fmt"
	"os"
)

//简易学生管理系统

//1.添加学生信息
//2.修改学生信息
//3.展示学生信息

// 主菜单
func starMenu() {
	fmt.Println("欢迎来到学生信息管理系统！")
	fmt.Println("1.添加学生")
	fmt.Println("2.编辑学生")
	fmt.Println("3.查看所有学生信息")
}

// 获取用户输入，生成学生信息
func GetInfo() *Student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请输入学生ID：")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入学生姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学生班级：")
	fmt.Scanf("%s\n", &class)
	stu := NewStudent(id, name, class)
	return stu
}

func main() {
	sm := StudentManager{make([]*Student, 0, 100)} //初始化，允许存储最大100个学生
	for {
		starMenu()
		var id int
		fmt.Println("请输入功能编号：")
		fmt.Scanf("%d\n", &id)
		switch id {
		case 1:
			//1.添加学生信息
			stu := GetInfo()
			sm.AddStudent(stu)
		case 2:
			//2.修改学生信息
			stu := GetInfo()
			sm.ModifyStudent(stu)
		case 3:
			//3.查看所有学生信息
			sm.ShowStudents()
		case 4:
			os.Exit(0)
		}
	}
	fmt.Println("asysflhuvsav")
}
