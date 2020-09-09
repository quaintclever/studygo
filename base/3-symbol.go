package main

import "fmt"

// 运算符练习
func main() {

	// 找到出现一次的数字, 异或运算练习
	var arr = [...]int{1, 1, 2, 2, 3, 3, 4, 4, 5}
	var ans = 0
	for _, a := range arr {
		ans ^= a
	}
	fmt.Println(ans)

}
