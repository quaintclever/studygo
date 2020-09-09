package main

import (
	"fmt"
)

// 1. 使用“面向对象”的思维方式编写一个学生信息管理系统。
// 		a. 学生有id、姓名、年龄、分数等信息
//		b. 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能

var students = make([]*Student, 0)

func main() {
	showMenu()
}

func showMenu() {
	fmt.Printf("----------------\n")
	fmt.Println("1.学生列表")
	fmt.Println("2.添加学生")
	fmt.Println("3.编辑学生信息")
	fmt.Println("4.删除学生")
	fmt.Println("other.退出")
	fmt.Printf("----------------\n")
	var a string
	_, err := fmt.Scanf("%s", &a)
	defer errRecover()
	errorProcess(err)

	switch a {
	case "1":
		getList()
	case "2":
		addStudent()
	case "3":
		updateStudent()
	case "4":
		delStudent()
	}
}

func getList() {
	fmt.Printf("\n------- 查询学生列表 -------\n")
	fmt.Printf("[\n")
	for _, v := range students {
		fmt.Printf("\t%+v \n", *v)
	}
	fmt.Printf("]\n")
	showMenu()
}

func addStudent() {
	fmt.Printf("\n------- 添加学生开始 -------\n")
	var add = new(Student)
	fmt.Println("请依次输入 姓名 年龄 成绩 \n Tips ==> 以空格隔开")
	_, err := fmt.Scanf("%s %d %f", &add.name, &add.age, &add.score)
	defer errRecover()
	errorProcess(err)

	if l := len(students); l > 0 {
		add.id = students[l-1].id + 1
	} else {
		add.id = 1
	}
	students = append(students, add)
	fmt.Println("添加成功!")
	showMenu()
}

func updateStudent() {
	fmt.Printf("\n------- 修改学生开始 -------\n")

	fmt.Println("请输入要修改学生的id")
	var id int
	_, err := fmt.Scanf("%d", &id)
	defer errRecover()
	errorProcess(err)

	byId := getStudentById(id)
	if byId != nil {
		fmt.Println("请依次输入修改后的 姓名 年龄 成绩 \n Tips ==> 以空格隔开")
		_, err2 := fmt.Scanf("%s %d %f", &byId.name, &byId.age, &byId.score)
		errorProcess(err2)
		fmt.Println("修改成功!")
	} else {
		fmt.Println("修改失败, 用户不存在!")
	}
	showMenu()
}

func delStudent() {
	fmt.Printf("\n------- 删除学生开始 -------\n")

	fmt.Println("请输入要修改学生的id")
	var id int
	_, err := fmt.Scanf("%d", &id)
	defer errRecover()
	errorProcess(err)

	for i, v := range students {
		if v.id == id {
			students = append(students[:i], students[i+1:]...)
			fmt.Println("删除成功!")
			showMenu()
			return
		}
	}
	fmt.Println("删除失败,用户不存在")
	showMenu()

}

func getStudentById(id int) *Student {
	for _, v := range students {
		if v.id == id {
			return v
		}
	}
	return nil
}

func errRecover() {
	msg := recover()
	if msg != nil {
		println(msg)
		showMenu()
	}

}

func errorProcess(err error) {
	if err != nil {
		fmt.Printf("遇到未知错误! 操作失败! %v \n", err)
		panic("遇到未知错误! 操作失败!")
	}
}

type Student struct {
	id    int
	name  string
	age   int
	score float32
}
