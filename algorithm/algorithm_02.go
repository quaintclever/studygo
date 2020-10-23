package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("======= =======")

	// 剑指 Offer 12. 矩阵中的路径
	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//word := "ABCCED"
	//fmt.Println(exist(board, word))

	// 剑指 Offer 13. 机器人的运动范围
	//count := movingCount(1, 2, 1)
	//fmt.Println(count)

	// 剑指 Offer 14- I. 剪绳子
	fmt.Println(cuttingRope(10))
	// 剑指 Offer 14- II. 剪绳子 II
	fmt.Println(cuttingRope2(10))

}

// 剑指 Offer 14- II. 剪绳子 II
func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return myPow(3, a)
	}
	if b == 1 {
		return myPow(3, a-1) * 4 % 1000000007
	}
	return myPow(3, a) * 2 % 1000000007
}

// % 1000000007
func myPow(a, b int) int {
	var ans = 1
	if b == 0 {
		return ans
	}
	for i := 0; i < b; i++ {
		ans = ans * a % 1000000007
	}
	return ans
}

// 剑指 Offer 14- I. 剪绳子
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	}
	if b == 1 {
		return int(math.Pow(3, float64(a-1)) * 4)
	}
	return int(math.Pow(3, float64(a)) * 2)
}

// 剑指 Offer 13. 机器人的运动范围
func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis := make([]bool, n)
		for j := 0; j < n; j++ {
			vis = append(vis, false)
		}
		visited[i] = vis
	}
	return dfs(0, 0, m, n, k, visited)
}

func dfs(i int, j int, m int, n int, k int, visited [][]bool) int {
	if i < 0 || i >= m || j < 0 || j >= n || sumNum(i)+sumNum(j) > k || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	return 1 + dfs(i+1, j, m, n, k, visited) + dfs(i, j+1, m, n, k, visited)
}

func sumNum(n int) int {
	res := 0
	for n != 0 {
		res += n % 10
		n = n / 10
	}
	return res
}

/*
===========================================
	剑指 Offer 12. 矩阵中的路径
===========================================
*/
func exist(board [][]byte, word string) bool {

	if len(word) == 0 {
		return false
	}

	wordByte := []byte(word)
	for i, b1 := range board {
		for j, b2 := range b1 {
			// 找到了第一个 字母
			if b2 == wordByte[0] {
				// 复制一份数组
				copyBoard := copyArray(board)
				if canContinue(copyBoard, i, j, wordByte, 1) {
					return true
				}
			}
		}
	}
	return false
}

// 搜索从该点出发，是否可行
func canContinue(board [][]byte, i int, j int, w []byte, idx int) bool {
	// 如果 全部找到了， 返回true
	if len(w) <= idx {
		return true
	}
	board[i][j] = '-'
	// 寻找上下左右 是否有满足的条件
	var ans = false
	if i > 0 {
		ans = ans || board[i-1][j] == w[idx]
		if ans {
			ans = canContinue(copyArray(board), i-1, j, w, idx+1)
			if ans {
				return true
			}
		}
	}
	if j > 0 {
		ans = ans || board[i][j-1] == w[idx]
		if ans {
			ans = canContinue(copyArray(board), i, j-1, w, idx+1)
			if ans {
				return true
			}
		}
	}
	if i < len(board)-1 {
		ans = ans || board[i+1][j] == w[idx]
		if ans {
			ans = canContinue(copyArray(board), i+1, j, w, idx+1)
			if ans {
				return true
			}
		}
	}
	if j < len(board[0])-1 {
		ans = ans || board[i][j+1] == w[idx]
		if ans {
			ans = canContinue(copyArray(board), i, j+1, w, idx+1)
			if ans {
				return true
			}
		}
	}
	return ans
}

// 复制数组
func copyArray(src [][]byte) [][]byte {
	target := make([][]byte, len(src))
	for i, v := range src {
		target[i] = make([]byte, len(v))
		copy(target[i], v)
	}
	return target
}
