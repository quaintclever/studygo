package main

import "fmt"

func main() {

	// 剑指 Offer 12. 矩阵中的路径
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
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
