package main

import "fmt"

// 学生类型
type Student struct {
	id    int
	name  string
	class string
}

//Student 构造函数
func NewStudent(id int, name, class string) *Student {
	return &Student{
		id:    id,
		name:  name,
		class: class,
	}
}

//用于存储学生信息的类型
type StudentManager struct {
	allStudents []*Student
}

//1.添加学生
func (s *StudentManager) AddStudent(newStu *Student) {
	s.allStudents = append(s.allStudents, newStu)
}

//2.编辑学生
func (s *StudentManager) ModifyStudent(newStu *Student) {
	for i, v := range s.allStudents { //遍历切片
		if v.id == newStu.id {
			s.allStudents[i] = newStu //通过比较id是否相等来确认学生是否存在，
		}
	}
	fmt.Println("该学生不存在，请输入正确的学号")
}

//3.展示所有学生信息
func (s *StudentManager) ShowStudents() {
	for _, stu := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", stu.id, stu.name, stu.class)
	}
}
