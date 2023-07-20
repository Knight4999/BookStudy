package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserID   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var DB *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:wzh123456@tcp(127.0.0.1:3306)/test1")
	if err != nil {
		defer DB.Close()
		fmt.Println("连接服务器失败", err)
		return
	}
	DB = database
}
func main() {
	// insert()
	// update()
	// delete()
	// selectmethod()
	alter()
}

func insert() {
	r, err := DB.Exec("insert into person (username,sex,email) Values(?,?,?);", "stu01", "m", "123@gmail.com")
	if err != nil {
		fmt.Println("解析SQL语句失败", err)
		return
	}
	id, err := r.LastInsertId() //返回插入的主键ID
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("插入成功:", id)
}

func update() {
	r, err := DB.Exec("update person set username = ? where user_id = ?", "stu03", 3)
	if err != nil {
		fmt.Println("解析SQL语句失败", err)
		return
	}
	row, err := r.RowsAffected() //返回受影响的行数
	if err != nil {
		fmt.Println("解析SQL语句失败", err)
		return
	}
	fmt.Println("update succ:", row)
}

func delete() {
	r, err := DB.Exec("delete from person where user_id=?", 5)
	if err != nil {
		fmt.Println("解析失败", err)
		return
	}
	id, err := r.RowsAffected()
	if err != nil {
		fmt.Println("删除失败", err)
	}
	fmt.Println("删除成功", id)
}

func selectmethod() {
	var person []Person
	err := DB.Select(&person, "select user_id,username,sex,email from person where user_id = ?", 2)
	if err != nil {
		fmt.Println("解析失败", err)
		return
	}
	fmt.Println(person)
}

func alter() {
	r, err := DB.Exec("alter table person add ? after sex", "salary DOUBLE")
	if err != nil {
		fmt.Println("解析失败", err)
		return
	}
	row, err := r.RowsAffected()
	if err != nil {
		fmt.Println("添加字段失败", err)
	}
	fmt.Println("添加字段成功", row)
}
