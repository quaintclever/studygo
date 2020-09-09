package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Printf("\n------- 反射初识 -------\n")
	var a = 12.3
	typeOf := reflect.TypeOf(a)
	valueOf := reflect.ValueOf(a)
	fmt.Println(typeOf, typeOf.Name(), typeOf.Kind())
	fmt.Println(valueOf)

	fmt.Printf("\n------- 结构体反射初识 -------\n")
	q := quaint{
		"quaint",
		18,
	}
	typeOf1 := reflect.TypeOf(q)
	fmt.Printf("%v, %v, %v \n", typeOf1, typeOf1.Kind(), typeOf1.Name())

	por := &a
	typeOf2 := reflect.TypeOf(por)
	fmt.Printf("%v, %v, %v \n", typeOf2, typeOf2.Kind(), typeOf2.Name())

	fmt.Printf("\n------- 通过反射设置变量的值 -------\n")
	porVal := reflect.ValueOf(por)
	if porVal.Elem().Kind() == reflect.Float64 {
		porVal.Elem().SetFloat(3.14)
	}
	fmt.Println(a, *por)

	valueOf.IsValid()

}

type quaint struct {
	name string
	age  int
}
