package main

import "fmt"

func main() {

	fmt.Printf("\n------- 指针定义 -------\n")
	//a:=10
	//b:=&a
	//fmt.Printf("%d,%p \n",a,&a)
	//fmt.Printf("%d,%p \n",b,b)
	//fmt.Printf("%p \n",&b)

	fmt.Printf("\n------- 创建指针 -------\n")
	var a = new(int)
	*a = 100
	fmt.Printf("%d,%p", *a, a)

}
