package main

import (
	"fmt"
	"strings"
)

func main() {

	// 4. 写一个程序，统计一个字符串中每个单词出现的次数。
	// 比如：”how do you do”中how=1 do=2 you=1。
	printNum("how do you do")

	// 3. 求数组 2个数 和 为 8 的 坐标
	//arrs := [5]int{1, 3, 5, 7, 8}
	//get2NumIs8(arrs)

	// 2. 打印99乘法表
	//muTable()
	// 1. 寻找出现一次的数字
	//findOnceNum()
}

// 4. 写一个程序，统计一个字符串中每个单词出现的次数。
// 比如：”how do you do”中how=1 do=2 you=1。
func printNum(str string) {
	arrStr := strings.Split(str, " ")
	m := make(map[string]int)
	for _, s := range arrStr {
		if v, ok := m[s]; ok {
			m[s] = v + 1
		} else {
			m[s] = 1
		}
	}
	fmt.Printf("%+v\n", m)
}

// 3. 求数组 2个数 和 为 8 的 坐标
func get2NumIs8(arrs [5]int) {
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if arrs[i]+arrs[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

// 2. 打印99乘法表
func muTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

}

// 1. 寻找出现一次的数字
func findOnceNum() {
	var arr = [...]int{1, 1, 2, 2, 3, 3, 4, 4, 5}
	var ans = 0
	for _, a := range arr {
		ans ^= a
	}
	fmt.Println(ans)
}
