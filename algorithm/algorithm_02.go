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
	//fmt.Println(cuttingRope(10))
	// 剑指 Offer 14- II. 剪绳子 II
	//fmt.Println(cuttingRope2(10))

	// 寻找两个正序数组的中位数
	// [0,0,0,0,0]
	//[-1,0,0,0,0,0,1]
	//var nums1 = []int{1, 2}
	//var nums2 = []int{1, 2}
	//fmt.Println(findMedianSortedArrays(nums1, nums2))

	// 842. 将数组拆分成斐波那契序列
	//s:= "123456579"
	//fmt.Printf("结果: %v",splitIntoFibonacci(s))

	// 649. Dota2 参议院
	//fmt.Println(predictPartyVictory("DDRRR"))

	// 738. 单调递增的数字
	fmt.Println(monotoneIncreasingDigits(113056))

}

// 738. 单调递增的数字
func monotoneIncreasingDigits(N int) int {
	ans := 0
	for true {
		num := getOnceNum(N)
		if ans == num {
			break
		}
		ans = num
		N = ans
	}
	return ans
}

func getOnceNum(N int) int {
	var arr []int
	for N != 0 {
		arr = append(arr, N%10)
		N /= 10
	}
	// 123456 => 654321
	ans := 0
	for i := len(arr) - 1; i >= 0; i-- {
		// 654321
		if i == 0 {
			ans += arr[i]
			return ans
		}
		if arr[i] <= arr[i-1] {
			ans += arr[i] * int(math.Pow(10, float64(i)))
			continue
		} else {
			ans += arr[i]*int(math.Pow(10, float64(i))) - 1
			return ans
		}
	}
	return ans
}

// 649. Dota2 参议院
func predictPartyVictory(senate string) string {
	// 优先干掉 排在自己后面以为 的 敌方
	// 如果后面 没有敌方， 优先干掉排在最前面的 敌方
	//（把前面存活的追加到后面，认为前面的全部被禁掉） 遍历到最后就知道结果了
	// 被禁掉的数量
	r, d := 0, 0
	length := len(senate)
	limitLen := length
	for i := 0; i < limitLen*2 && i < length; i++ {
		if senate[i] == 'R' {
			if r > 0 {
				// 本方死亡。
				r--
			} else {
				// 对方被 禁 一个， 本方存活， 排到最后
				d++
				senate = senate + "R"
				if i == length-1 {
					break
				}
				length++
			}
		} else if senate[i] == 'D' {
			if d > 0 {
				d--
			} else {
				r++
				senate = senate + "D"
				if i == length-1 {
					break
				}
				length++
			}
		}
	}
	if r > d {
		return "Dire"
	} else {
		return "Radiant"
	}
}

// 842. 将数组拆分成斐波那契序列
func splitIntoFibonacci2(S string) (arr []int) {
	n := len(S)
	var backtrack func(idx int, sum int, prev int) bool
	backtrack = func(idx int, sum int, prev int) bool {
		// 如果遍历到最后一个， 结束递归
		if n == idx {
			// 返回数组， 和是否能够成 斐波那契
			return len(arr) >= 3
		}
		var curr = 0
		for i := idx; i < n; i++ {
			if i > idx && S[idx] == '0' {
				break
			}
			curr = curr*10 + (int)(S[i]-'0')
			if curr > math.MaxInt32 {
				break
			}
			if len(arr) >= 2 {
				if sum > curr {
					continue
				}
				if sum < curr {
					break
				}
			}
			arr = append(arr, curr)
			// 如果 遍历可以 组成
			if backtrack(i+1, curr+prev, curr) {
				return true
			}
			fmt.Printf("==>> %v \n", arr)
			arr = arr[:len(arr)-1]
			fmt.Printf("<<== %v \n", arr)
		}
		return false
	}
	backtrack(0, 0, 0)
	return arr
}

// 842. 将数组拆分成斐波那契序列
func splitIntoFibonacci(S string) (arr []int) {
	arr, _ = backtrack(arr, S, len(S), 0, 0, 0, 0)
	return
}

/**
 * idx, 当前位置
 * sum, 前两个数的和
 * prev, 前一个数字
 */
func backtrack(arr []int, s string, length int, idx int, sum int, prev int, level int) ([]int, bool) {
	// 如果遍历到最后一个， 结束递归
	if length == idx {
		// 返回数组， 和是否能够成 斐波那契
		return arr, len(arr) >= 3
	}
	var curr = 0
	for i := idx; i < length; i++ {
		if i > idx && s[idx] == '0' {
			break
		}
		curr = curr*10 + (int)(s[i]-'0')
		if curr > math.MaxInt32 {
			break
		}
		if len(arr) >= 2 {
			if sum > curr {
				continue
			}
			if sum < curr {
				break
			}
		}
		fmt.Printf("==> level:%d, add之前地址: %p, 内容: %v \n", level, arr, arr)
		arr = append(arr, curr)
		fmt.Printf("<== level:%d, add之后地址: %p, 内容: %v \n\n", level, arr, arr)
		// 如果 遍历可以 组成
		arr, result := backtrack(arr, s, length, i+1, curr+prev, curr, level+1)
		fmt.Printf("<== level:%d, cal返回地址: %p, 内容: %v \n\n", level, arr, arr)
		if result {
			return arr, result
		}
		fmt.Printf("==> level:%d, del之前地址: %p, 内容: %v \n", level, arr, arr)
		arr = arr[:len(arr)-1]
		fmt.Printf("<== level:%d, del之后地址: %p, 内容: %v \n\n", level, arr, arr)
	}
	return arr, false
}

// 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var p1 = len(nums1)
	var p2 = len(nums2)

	if p1+p2 == 0 {
		return 0
	}

	var mid = (p1 + p2 - 1) / 2
	var flag = (p1 + p2 - 1) % 2

	var ans = 0.0
	for i := 0; i <= mid; i++ {
		if p1 == 0 || (p2 > 0 && nums1[p1-1] <= nums2[p2-1]) {
			ans = float64(nums2[p2-1])
			p2--
			continue
		}
		if p2 == 0 || (p1 > 0 && nums1[p1-1] >= nums2[p2-1]) {
			ans = float64(nums1[p1-1])
			p1--
			continue
		}
	}

	if flag != 0 {
		if p1 == 0 || (p2 > 0 && nums1[p1-1] <= nums2[p2-1]) {
			ans = (float64(nums2[p2-1]) + ans) / 2
		} else if p2 == 0 || (p1 > 0 && nums1[p1-1] >= nums2[p2-1]) {
			ans = (float64(nums1[p1-1]) + ans) / 2
		}
	}
	return ans
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
