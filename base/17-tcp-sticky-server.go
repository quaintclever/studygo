package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	listener, e := net.Listen("tcp", ":8080")
	if e != nil {
		fmt.Println(e)
		return
	}
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go process2(conn)
	}
}

func process2(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//var buf [1024] byte
	for {
		//n, err := reader.Read(buf[:])
		msg, err := Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("err is ", err)
			break
		}
		//recStr := string(buf[:n])
		//fmt.Println("收到客户端发来的消息: ",recStr)
		fmt.Println("收到客户端发来的消息: ", msg)
	}

}

func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewReader(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// buffered 返回缓冲中 现有 的 可读取的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	// 读取真正的数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
