package main

import "fmt"

func main() {

	fmt.Printf("\n------- 面试题 -------\n")
	var peo People1 = &Student1{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

}

type People1 interface {
	Speak(string) string
}
type Student1 struct{}

func (stu Student1) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}
