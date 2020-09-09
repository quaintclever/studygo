package main

import "fmt"

// 流程控制练习
func main() {

	fmt.Printf("\n------- if else练习 -------\n")
	score := 65
	if score >= 90 {
		fmt.Printf("恭喜你, 优秀\n")
	} else if score >= 60 {
		fmt.Println("恭喜你, 及格了!")
	} else {
		fmt.Println("不及格!")
	}

	if score := 90; score >= 90 {
		fmt.Printf("恭喜你, 优秀\n")
	} else if score >= 60 {
		fmt.Println("恭喜你, 及格了!")
	} else {
		fmt.Println("不及格!")
	}
	fmt.Println(score)

	fmt.Printf("\n------- for 循环练习 -------\n")
	fmt.Printf("\n------- switch case 循环练习 -------\n")
	// fallthrough ==> case 穿透

	fmt.Printf("\n------- goto 练习 -------\n")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
breakTag:
	fmt.Println("come here!")

	fmt.Printf("\n------- break 跳出循环 -------\n")
	fmt.Printf("\n------- continue 继续下次循环 -------\n")

}
