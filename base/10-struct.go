package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Printf("\n------- 类型定义和类型别名的区别 -------\n")
	type newInt int
	var a newInt = 1
	type myInt = int
	var b myInt = 2
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)

	fmt.Printf("\n------- 结构体实例化 -> 基本实例化 -------\n")
	var p people
	p.name = "quaint"
	p.age = 18
	fmt.Println(p)
	fmt.Printf("%+v\n", p)

	fmt.Printf("\n------- 匿名结构体 -------\n")
	var noName struct {
		Name string
		Age  int
	}
	noName.Name = "noName"
	noName.Age = 11
	fmt.Printf("%+v\n", noName)

	fmt.Printf("\n------- 创建指针类型结构体 -------\n")
	pp := new(people)
	pp.age = 12
	pp.name = "test"
	fmt.Printf("%+v %T \n", pp, pp)
	fmt.Printf("%s %d %T \n", pp.name, pp.age, *pp)

	fmt.Printf("\n------- 面试题 -------\n")
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	fmt.Printf("\n------- 结构体的“继承” -------\n")
	dog := new(Dog)
	dog.Animal = new(Animal)
	dog.Name = "哈士奇"
	dog.Skill = "拆家"
	dog.GetName()
	dog.GetSkill()

	fmt.Printf("\n------- 结构体与JSON序列化 -------\n")
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)

	fmt.Printf("\n------- 结构体与JSON序列化2 -------\n")
	dogTest := &Dog{
		Animal: new(Animal),
		Skill:  "拆家二号",
	}
	dogTest.Name = "二哈二号"
	data2, err2 := json.Marshal(dogTest)
	if err2 != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data2)

	dog4Json := &Dog{}
	err = json.Unmarshal(data2, dog4Json)
	if err != nil {
		fmt.Println("json 反序列化失败")
	}
	fmt.Printf("%+v %p \n", dogTest, dogTest)
	fmt.Printf("%+v %p \n", dog4Json, dog4Json)

}

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

type Animal struct {
	Name string
}

func (a Animal) GetName() {
	fmt.Println("我的名字是:" + a.Name)
}

type Dog struct {
	*Animal
	Skill string
}

func (d Dog) GetSkill() {
	fmt.Println("我可以" + d.Skill)
}

type student struct {
	name string
	age  int
}

type people struct {
	name string
	age  int
}
