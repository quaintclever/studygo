package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("连接服务端失败:", err)
	}
	defer socket.Close()

	sendDate := []byte("Hello server")
	_, err = socket.Write(sendDate)
	if err != nil {
		fmt.Println("发送数据失败", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("接收数据失败,err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v \n", string(data[:n]), remoteAddr, n)

}
