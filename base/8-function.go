package main

import "fmt"

func main() {

	fmt.Printf("\n------- 函数作为参数 -------\n")
	fmt.Println(calc(1, 2, add))
	fmt.Printf("\n------- 函数作为返回值 -------\n")
	ans := do("+")(3, 4)
	fmt.Println(ans)

	fmt.Printf("\n------- 函数defer -------\n")
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5

	fmt.Printf("\n------- 面试题 -------\n")
	//x := 1
	//y := 2
	//defer calc_t("AA", x, calc_t("A", x, y))
	//x = 10
	//defer calc_t("BB", x, calc_t("B", x, y))
	//y = 20

	fmt.Printf("\n------- 分金币 -------\n")
	/*
		你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
		分配规则如下：
		a. 名字中每包含1个'e'或'E'分1枚金币
		b. 名字中每包含1个'i'或'I'分2枚金币
		c. 名字中每包含1个'o'或'O'分3枚金币
		d: 名字中每包含1个'u'或'U'分4枚金币
		写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
		程序结构如下，请实现 ‘dispatchCoin’ 函数
	*/
	var (
		coins = 50
		users = []string{
			"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		}
		distribution = make(map[string]int, len(users))
	)
	left := dispatchCoin(coins, users, distribution)
	fmt.Println("剩下：", left)
	fmt.Println(distribution)

}

func dispatchCoin(coins int, users []string, distribution map[string]int) (left int) {

	for _, name := range users {
		coin := 0
		for _, s := range name {
			switch s {
			case 'e', 'E':
				coin += 1
			case 'i', 'I':
				coin += 2
			case 'o', 'O':
				coin += 3
			case 'u', 'U':
				coin += 4
			}
		}
		distribution[name] = coin
		coins -= coin
	}
	return coins
}

func calc_t(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func add(a, b int) int {
	return a + b
}

func calc(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func do(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return func(i int, i2 int) int {
			return i - i2
		}
	default:
		return nil
	}
}
