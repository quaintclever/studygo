package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Printf("\n------- 切片定义 -------\n")
	var a []int
	var b = []int{}
	var c = []string{"a", "b"}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)
	// 不能直接比较
	//fmt.Println(c==b)

	fmt.Printf("\n------- 从数组中提取切片 -------\n")
	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	slice1 := array[:10]
	slice2 := array[3:7]
	slice3 := array[6:]
	fmt.Printf("%v,%d,%d \n", slice1, len(slice1), cap(slice1))
	fmt.Printf("%v,%d,%d \n", slice2, len(slice2), cap(slice2))
	// 这里会报错
	// fmt.Printf("%d \n",slice2[6])
	fmt.Printf("%v,%d,%d \n", slice3, len(slice3), cap(slice3))

	fmt.Printf("\n------- 完整切片表达式 -------\n")
	var array2 = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl := array2[0:3:5]
	fmt.Printf("%v,%d,%d \n", sl, len(sl), cap(sl))

	fmt.Printf("\n------- 使用make()函数构造切片 -------\n")
	var maArray = make([]int, 0, 10)
	fmt.Println(maArray, len(maArray), cap(maArray))
	// 添加
	maArray = append(maArray, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(maArray, len(maArray), cap(maArray))
	// 删除
	maArray = append(maArray[:6], maArray[7:]...)
	fmt.Println(maArray, len(maArray), cap(maArray))

	fmt.Printf("\n------- 练习题 -------\n")
	var aa = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		aa = append(aa, fmt.Sprintf("%v", i))
	}
	var sortArr = [...]int{3, 7, 8, 9, 1}
	sort.Ints(sortArr[:])
	fmt.Println(sortArr)

}
