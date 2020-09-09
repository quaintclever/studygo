package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	fmt.Printf("\n------- map基础定义 -------\n")
	var m1 map[string]int
	m1 = make(map[string]int)
	m1["quaint"] = 100
	m1["turkey"] = 97
	fmt.Println(m1)

	fmt.Printf("\n------- 声明的时候填充文字 -------\n")
	m2 := map[string]string{
		"name": "quaint",
		"age":  "18",
	}
	fmt.Println(m2)

	fmt.Printf("\n------- 判断某个键是否存在 -------\n")
	if v, ok := m2["name"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key 不存在")
	}

	fmt.Printf("\n------- 使用delete()函数删除键值对 -------\n")
	fmt.Println(m2)
	delete(m2, "age")
	fmt.Println(m2)

	fmt.Printf("\n------- 按照指定顺序遍历map -------\n")
	rand.Seed(time.Now().UnixNano())
	m3 := make(map[string]int)
	// 为map 复制
	for i := 1; i < 6; i++ {
		key := fmt.Sprintf("Stu%02d", i)
		val := rand.Intn(100)
		m3[key] = val
	}
	// 取出所有key
	keys := make([]string, 0, 10)
	for k := range m3 {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 遍历排序好的 keys
	for _, k := range keys {
		fmt.Println(k, m3[k])
	}

	fmt.Printf("\n------- 练习题 -------\n")
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) // [1,2,3]
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)    // 移除s[1]
	fmt.Printf("%+v\n", s)         // [1,3]
	fmt.Printf("%+v\n", m["q1mi"]) // 切片传递的是地址， 这里也会改动, 结果为 [1,3,3]

}
