package main

import (
	"fmt"
)

// 全局变量
var all = "全局变量"

// 变量定义， 格式化字符串输出
func main() {
	fmt.Println("\n==========================")
	// 定义变量 标准声明
	var name string
	var age int
	var isBoy bool
	// 赋初始值
	name = "quaint"
	age = 24
	isBoy = true
	fmt.Printf("hello %s, this year your is %d, you is a boy? %v ! ", name, age, isBoy)
	fmt.Println("\n------------------------")
	// 批量定义变量
	type double float64
	var (
		a string  = "quaint"
		b float32 = 3.14
		c double  = 3.1415926
	)
	fmt.Printf(a, b, c)
	fmt.Println("\n------------------------")
	// 类型推导
	var lalala = "lalala"
	var isOk = true
	// 临时变量
	temp := "tempStr"
	fmt.Printf(lalala, isOk, temp)

	fmt.Println("\n------------------------")
	// 匿名变量
	_, str := foo()
	num, _ := foo()
	fmt.Println(str, num)

	fmt.Println("\n------------------------")
	// 常量定义 , 批量常量定义 类似 var
	const PI = 3.1415926
	fmt.Println(PI)

	fmt.Println("\n------------------------")
	//iota
	const (
		n1 = iota
		n2
		n3
		n4
	)
	// 0 1 2 3
	fmt.Println(n1, n2, n3, n4)
	const (
		a1 = iota
		a2 = 77
		a3
		a4
	)
	// 0 77 77 77 77
	fmt.Println(a1, a2, a3, a4)
	const (
		b1 = iota
		b2 = 77
		b3 = iota
		b4
	)
	// 0 77 2 3
	fmt.Println(b1, b2, b3, b4)

}

// 前面括号为入参, 后面括号为出参
func foo() (int, string) {
	return 1, "foo"
}
