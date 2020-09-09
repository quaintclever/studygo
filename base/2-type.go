package main

import (
	"fmt"
	"math"
)

// 基本数据类型练习
func main() {

	fmt.Printf("\n------- 整形 -------\n")
	var a int = 100
	// 10
	fmt.Printf("%d \n", a)
	// 2
	fmt.Printf("%b \n", a)
	// 8
	fmt.Printf("%o \n", a)
	// 16
	fmt.Printf("%x \n", a)

	fmt.Printf("\n------- 浮点型 -------\n")
	var b float32 = 77.77
	// 默认6位  77.769997
	fmt.Printf("%f \n", b)
	// 77.77
	fmt.Printf("%.2f \n", b)
	// 3.141593
	fmt.Printf("%f\n", math.Pi)
	// 3.14
	fmt.Printf("%.2f\n", math.Pi)

	fmt.Printf("\n------- 复数 -------\n")
	var c complex64 = 1 + 2i
	var d complex128 = 2 + 3i
	// (1+2i)
	fmt.Printf("%v \n", c)
	// (2+3i)
	fmt.Printf("%v \n", d)

	fmt.Printf("\n------- 多行字符串 -------\n")
	var e = `quaint 
test
the mul line
`
	fmt.Print(e)

	fmt.Printf("\n------- byte和rune类型 -------\n")
	var f = '中'
	var g = 'x'
	fmt.Printf("%v(%c), type = %T \n", f, f, f)
	fmt.Printf("%v(%c), type = %T \n", g, g, g)

	str := "hello 字符串!"
	for i := 0; i < len(str); i++ { // byte
		fmt.Printf("%v(%c) ", str[i], str[i])
	}
	fmt.Printf("\n")
	for _, s := range str { // rune
		fmt.Printf("%v(%c) ", s, s)
	}

	fmt.Printf("\n------- 修改字符串 -------\n")
	s1 := "big"
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Printf("%v, %s\n", byteS1, byteS1)
	s2 := "行秋"
	byteS2 := []rune(s2)
	byteS2[0] = '去'
	fmt.Printf("%v, %s", byteS2, string(byteS2))

}
