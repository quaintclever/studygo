package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("建立tcp 监听 失败")
		return
	}
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("接收连接失败")
			continue
		}
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("读取连接中的数据失败", err)
		}
		recvStr := string(buf[:n])
		fmt.Println("收到客户端发来的数据:", recvStr)
		_, err2 := conn.Write([]byte("您发送的消息已收到"))
		if err2 != nil {
			fmt.Println("发送回复消息失败")
		}
	}
}
