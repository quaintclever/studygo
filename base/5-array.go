package main

import "fmt"

func main() {

	fmt.Printf("\n------- 数组的初始化 -------\n")
	var a = [3]int{1, 2, 3}
	var b = [2]string{"quaint", "test"}
	var c = [...]byte{'a', 'b'}
	var d = [...]int{1: 1, 3: 5}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Printf("\n------- 数组的遍历 -------\n")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for i, s := range b {
		fmt.Println(i, s)
	}

	fmt.Printf("\n------- 多维数组定义 -------\n")
	var arr = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr)
	fmt.Printf("\n------- 多维数组遍历 -------\n")
	for _, a := range arr {
		for _, i := range a {
			fmt.Println(i)
		}
	}

	fmt.Printf("\n------- 练习题1 -------\n")
	fmt.Println(sumArr(a))

}

func sumArr(arr [3]int) (sum int) {
	sum = 0
	for _, i := range arr {
		sum += i
	}
	return
}
