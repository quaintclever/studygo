package main

import "fmt"

func main() {

	isomorphic := isIsomorphic("paper", "title")
	fmt.Println(isomorphic)
}

/*
示例 3:

输入: s = "paper", t = "title"
p -> t
a -> i
p    t

输出: true
*/
func isIsomorphic(s string, t string) bool {

	checkKey := make(map[uint8]uint8)
	checkVel := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		if v, ok := checkKey[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else if v, ok := checkVel[t[i]]; ok {
			if v != s[i] {
				return false
			}
		} else {
			checkKey[s[i]] = t[i]
			checkVel[t[i]] = s[i]
		}
	}
	return true
}
