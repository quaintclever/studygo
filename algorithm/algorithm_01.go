package main

import (
	"fmt"
	"strings"
)

func main() {

	// 5. 构建二叉树
	preorder := make([]int, 0)
	preorder = append(preorder, 1, 2, 3)
	inorder := make([]int, 0)
	inorder = append(inorder, 2, 1, 3)
	tree := buildTree(preorder, inorder)
	fmt.Printf("%v", tree)

	// 4. 写一个程序，统计一个字符串中每个单词出现的次数。
	// 比如：”how do you do”中how=1 do=2 you=1。
	//printNum("how do you do")

	// 3. 求数组 2个数 和 为 8 的 坐标
	//arrs := [5]int{1, 3, 5, 7, 8}
	//get2NumIs8(arrs)

	// 2. 打印99乘法表
	//muTable()
	// 1. 寻找出现一次的数字
	//findOnceNum()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	tree := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	if len(preorder) == 1 {
		return tree
	}
	var idx int
	for i, v := range inorder {
		if v == tree.Val {
			idx = i
		}
	}
	if idx == 0 {
		tree.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	} else if idx == len(inorder) {
		tree.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	} else {
		tree.Left = buildTree(preorder[1:idx+1], inorder[:idx])
		tree.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	}
	return tree
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
