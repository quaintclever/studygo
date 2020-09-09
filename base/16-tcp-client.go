package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("与服务器建立连接失败!", err)
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		s, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(s, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		// 发送数据到服务端
		_, e2 := conn.Write([]byte(inputInfo))
		if e2 != nil {
			fmt.Println("发送消息失败")
			return
		}
		var buf [128]byte
		n, e3 := conn.Read(buf[:])
		if e3 != nil {
			fmt.Println("接收服务器回复消息失败")
			return
		}
		fmt.Println("收到服务器的消息:", string(buf[:n]))
	}

}
