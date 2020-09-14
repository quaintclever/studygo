package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("listen fail", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		// 接收数据
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp error", err)
			continue
		}
		fmt.Printf("data:%v addr:%v,count:%v \n", string(data[:n]), addr, n)
		// 发送数据
		_, err = listen.WriteToUDP([]byte("服务端收到了你的请求."), addr)
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}

}
